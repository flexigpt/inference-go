package sdkutil

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/flexigpt/inference-go/spec"
)

var EmptyJSONArgs = map[string]any{
	"type":                 "object",
	"properties":           map[string]any{},
	"additionalProperties": false,
}

func NormalizeRequestForSDK(
	ctx context.Context,
	req *spec.FetchCompletionRequest,
	opts *spec.FetchCompletionOptions,
	sdkType spec.ProviderSDKType,
	providerCapabilities spec.ModelCapabilities,
) (cappedReq *spec.FetchCompletionRequest, warnings []spec.Warning, err error) {
	if req == nil {
		return nil, nil, errors.New("nil request")
	}

	caps := &providerCapabilities
	if opts != nil && opts.CapabilityResolver != nil {
		caps, err = opts.CapabilityResolver.ResolveModelCapabilities(ctx, spec.ResolveModelCapabilitiesRequest{
			ProviderSDKType: sdkType,
			ModelName:       req.ModelParam.Name,
			CompletionKey:   opts.CompletionKey,
		})
		if err != nil {
			return nil, nil, err
		}
		if caps == nil {
			return nil, nil, errors.New("capability resolver returned nil ModelCapabilities")
		}
	}

	nreq, err := CloneFetchCompletionRequest(req)
	if err != nil {
		return nil, nil, err
	}

	// Modalities validation (inferred from inputs).
	used := getInputModalitiesForValidation(nreq.Inputs)
	if err := requireModalities(used, caps.ModalitiesIn); err != nil {
		return nil, nil, err
	}

	// Reasoning validation / safe-dropping.
	if nreq.ModelParam.Reasoning != nil {
		if caps.ReasoningCapabilities == nil || !caps.ReasoningCapabilities.SupportsReasoningConfig ||
			!supportsReasoningType(*nreq.ModelParam.Reasoning, caps.ReasoningCapabilities) {
			warnings = append(warnings, spec.Warning{
				Code:    "reasoning_dropped_unsupported",
				Message: "Reasoning was dropped because it is not supported by the selected SDK/model.",
			})
			nreq.ModelParam.Reasoning = nil
		} else {
			// SummaryStyle: safe to drop if unsupported.
			if nreq.ModelParam.Reasoning.SummaryStyle != nil && !caps.ReasoningCapabilities.SupportsSummaryStyle {
				warnings = append(warnings, spec.Warning{
					Code:    "reasoning_summaryStyle_dropped",
					Message: "reasoning.summaryStyle is not supported and was dropped.",
				})
				rp := *nreq.ModelParam.Reasoning
				rp.SummaryStyle = nil
				nreq.ModelParam.Reasoning = &rp
			}

			// Level validation (do not “nearest-map”; drop reasoning in bestEffort).
			if nreq.ModelParam.Reasoning.Type == spec.ReasoningTypeSingleWithLevels {
				if !supportsReasoningLevel(
					nreq.ModelParam.Reasoning.Level,
					caps.ReasoningCapabilities.SupportedReasoningLevels,
				) {
					warnings = append(warnings, spec.Warning{
						Code: "reasoning_dropped_invalid_level",
						Message: fmt.Sprintf(
							"Reasoning was dropped because level %q is unsupported.",
							nreq.ModelParam.Reasoning.Level,
						),
					})
					nreq.ModelParam.Reasoning = nil
				}
			}
		}
	}

	// Enforce SDK constraints: e.g. Anthropic temperature disallowed when reasoning enabled.
	if nreq.ModelParam.Reasoning != nil &&
		caps.ReasoningCapabilities != nil &&
		caps.ReasoningCapabilities.TemperatureDisallowedWhenEnabled {
		if nreq.ModelParam.Temperature != nil {
			warnings = append(warnings, spec.Warning{
				Code:    "temperature_dropped_reasoning_enabled",
				Message: "temperature was dropped because reasoning/thinking is enabled for this SDK/model.",
			})
			nreq.ModelParam.Temperature = nil
		}
	}

	// Stop sequences.
	if len(nreq.ModelParam.StopSequences) > 0 {
		if caps.StopSequenceCapabilities == nil || !caps.StopSequenceCapabilities.IsSupported {
			warnings = append(warnings, spec.Warning{
				Code:    "stopSequences_dropped_unsupported",
				Message: "stopSequences was dropped because it is not supported by this SDK/model.",
			})
			nreq.ModelParam.StopSequences = nil
		} else {
			if caps.StopSequenceCapabilities.MaxSequences > 0 &&
				len(nreq.ModelParam.StopSequences) > caps.StopSequenceCapabilities.MaxSequences {
				warnings = append(warnings, spec.Warning{
					Code: "stopSequences_truncated",
					Message: fmt.Sprintf(
						"stopSequences was truncated to max=%d.",
						caps.StopSequenceCapabilities.MaxSequences,
					),
				})
				nreq.ModelParam.StopSequences = nreq.ModelParam.StopSequences[:caps.StopSequenceCapabilities.MaxSequences]
			}

			if caps.StopSequenceCapabilities.DisallowedWithReasoning && nreq.ModelParam.Reasoning != nil {
				warnings = append(warnings, spec.Warning{
					Code:    "stopSequences_dropped_reasoning",
					Message: "stopSequences was dropped because it is incompatible with reasoning for this SDK/model.",
				})
				nreq.ModelParam.StopSequences = nil
			}
		}
	}

	// OutputParam: format is contract-like (error if unsupported); verbosity is safe to drop.
	if nreq.ModelParam.OutputParam != nil && caps.OutputCapabilities != nil {
		op := nreq.ModelParam.OutputParam

		if op.Format != nil {
			if !supportsOutputFormat(op.Format.Kind, caps.OutputCapabilities.SupportedOutputFormats) {
				return nil, warnings, fmt.Errorf("output format %q unsupported for sdkType=%s", op.Format.Kind, sdkType)
			}
		}

		if op.Verbosity != nil && !caps.OutputCapabilities.SupportsVerbosity {
			warnings = append(warnings, spec.Warning{
				Code:    "verbosity_dropped_unsupported",
				Message: "outputParam.verbosity was dropped because it is not supported by this SDK/model.",
			})
			cop := *op
			cop.Verbosity = nil
			nreq.ModelParam.OutputParam = &cop
		}
	}

	// OutputParam: if caps.OutputCapabilities is nil, treat format as unsupported and verbosity as droppable.
	if nreq.ModelParam.OutputParam != nil && caps.OutputCapabilities == nil {
		if nreq.ModelParam.OutputParam.Format != nil {
			return nil, warnings, errors.New(
				"outputParam.format requested but output capabilities are unavailable/unsupported for this SDK/model",
			)
		}
		if nreq.ModelParam.OutputParam.Verbosity != nil {
			warnings = append(warnings, spec.Warning{
				Code:    "verbosity_dropped_unsupported",
				Message: "outputParam.verbosity was dropped because it is not supported by this SDK/model.",
			})
			cop := *nreq.ModelParam.OutputParam
			cop.Verbosity = nil
			nreq.ModelParam.OutputParam = &cop
		}
	}

	// Tools: validate ToolChoices / ToolPolicy against capabilities.
	if (len(nreq.ToolChoices) > 0 || nreq.ToolPolicy != nil) && caps.ToolCapabilities == nil {
		return nil, warnings, errors.New("tools/toolPolicy provided but tools are not supported by selected SDK/model")
	}

	if len(nreq.ToolChoices) > 0 && caps.ToolCapabilities != nil {
		filtered := make([]spec.ToolChoice, 0, len(nreq.ToolChoices))
		for _, tc := range nreq.ToolChoices {
			if supportsToolType(tc.Type, caps.ToolCapabilities.SupportedToolTypes) {
				filtered = append(filtered, tc)
				continue
			}
			warnings = append(warnings, spec.Warning{
				Code: "toolChoice_dropped_unsupported",
				Message: fmt.Sprintf(
					"toolChoice type %q was dropped because it is unsupported by this SDK/model.",
					tc.Type,
				),
			})
		}
		nreq.ToolChoices = filtered
	}

	// Tool policy minimal validation (structural; provider-specific resolution remains in adapters).
	if nreq.ToolPolicy != nil {
		if caps.ToolCapabilities != nil &&
			!supportsToolPolicyMode(nreq.ToolPolicy.Mode, caps.ToolCapabilities.SupportedToolPolicyModes) {
			return nil, warnings, fmt.Errorf(
				"toolPolicy.mode %q unsupported for sdkType=%s",
				nreq.ToolPolicy.Mode,
				sdkType,
			)
		}
		if (nreq.ToolPolicy.Mode == spec.ToolPolicyModeAny ||
			nreq.ToolPolicy.Mode == spec.ToolPolicyModeTool) &&
			len(nreq.ToolChoices) == 0 {
			return nil, warnings, fmt.Errorf("toolPolicy.mode=%s requires toolChoices", nreq.ToolPolicy.Mode)
		}

		if nreq.ToolPolicy.Mode == spec.ToolPolicyModeTool && len(nreq.ToolPolicy.AllowedTools) == 0 {
			return nil, warnings, errors.New("toolPolicy.mode=tool requires allowedTools")
		}

		// Forced tool count constraint (bestEffort: keep first N).
		if caps.ToolCapabilities != nil && caps.ToolCapabilities.MaxForcedTools > 0 &&
			nreq.ToolPolicy.Mode == spec.ToolPolicyModeTool {
			if len(nreq.ToolPolicy.AllowedTools) > caps.ToolCapabilities.MaxForcedTools {
				warnings = append(warnings, spec.Warning{
					Code: "allowedTools_truncated",
					Message: fmt.Sprintf(
						"allowedTools truncated to %d due to SDK/model limitation.",
						caps.ToolCapabilities.MaxForcedTools,
					),
				})
				cp := *nreq.ToolPolicy
				cp.AllowedTools = cp.AllowedTools[:caps.ToolCapabilities.MaxForcedTools]
				nreq.ToolPolicy = &cp
			}
		}
	}

	// Client tool outputs: normalize according to SDK/model transport capability.
	toolWarnings, err := normalizeClientToolOutputsForSDK(nreq, caps.ToolCapabilities)
	if err != nil {
		return nil, warnings, err
	}
	warnings = append(warnings, toolWarnings...)
	warnings = append(warnings, normalizeRequestCacheControls(nreq, caps.CacheCapabilities)...)

	return nreq, warnings, nil
}

func getInputModalitiesForValidation(inputs []spec.InputUnion) []spec.Modality {
	// Treat text as always required.
	needText := true
	needImage := false
	needFile := false
	needAudio := false
	needVideo := false

	for _, in := range inputs {
		if IsInputUnionEmpty(in) {
			continue
		}

		var msg *spec.InputOutputContent
		switch in.Kind {
		case spec.InputKindInputMessage:
			msg = in.InputMessage
		case spec.InputKindOutputMessage:
			msg = in.OutputMessage
		case spec.InputKindFunctionToolOutput:
			if in.FunctionToolOutput != nil {
				for _, c := range in.FunctionToolOutput.Contents {
					switch c.Kind {
					case spec.ContentItemKindText:
						needText = true
					case spec.ContentItemKindImage:
						needImage = true
					case spec.ContentItemKindFile:
						needFile = true
					default:
						// Add new content types as needed.
					}
				}
			}
		case spec.InputKindCustomToolOutput:
			if in.CustomToolOutput != nil {
				for _, c := range in.CustomToolOutput.Contents {
					switch c.Kind {
					case spec.ContentItemKindText:
						needText = true
					case spec.ContentItemKindImage:
						needImage = true
					case spec.ContentItemKindFile:
						needFile = true
					default:
						// Add new content types as needed.
					}
				}
			}
		default:
			// Modality checks via only input and output message.
		}

		if msg == nil {
			continue
		}
		for _, c := range msg.Contents {
			switch c.Kind {
			case spec.ContentItemKindText:
				needText = true
			case spec.ContentItemKindImage:
				needImage = true
			case spec.ContentItemKindFile:
				needFile = true
			default:
				// Add new content types as needed.
				// Refusal doesn't constitute a modality requirement as of now.
			}
		}
	}

	var out []spec.Modality
	if needText {
		out = append(out, spec.ModalityTextIn)
	}
	if needImage {
		out = append(out, spec.ModalityImageIn)
	}
	if needFile {
		out = append(out, spec.ModalityFileIn)
	}
	if needAudio {
		out = append(out, spec.ModalityAudioIn)
	}
	if needVideo {
		out = append(out, spec.ModalityVideoIn)
	}
	return out
}

func requireModalities(used, supported []spec.Modality) error {
	for _, m := range used {
		if !containsModality(supported, m) {
			return fmt.Errorf("input modality %q unsupported", m)
		}
	}
	return nil
}

func containsModality(list []spec.Modality, v spec.Modality) bool {
	return slices.Contains(list, v)
}

func supportsReasoningType(r spec.ReasoningParam, caps *spec.ReasoningCapabilities) bool {
	if caps == nil {
		return false
	}
	return slices.Contains(caps.SupportedReasoningTypes, r.Type)
}

func supportsReasoningLevel(level spec.ReasoningLevel, allowed []spec.ReasoningLevel) bool {
	return slices.Contains(allowed, level)
}

func supportsOutputFormat(kind spec.OutputFormatKind, allowed []spec.OutputFormatKind) bool {
	return slices.Contains(allowed, kind)
}

func supportsToolPolicyMode(mode spec.ToolPolicyMode, allowed []spec.ToolPolicyMode) bool {
	return slices.Contains(allowed, mode)
}

func supportsToolType(t spec.ToolType, allowed []spec.ToolType) bool {
	return slices.Contains(allowed, t)
}
