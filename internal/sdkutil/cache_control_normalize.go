package sdkutil

import (
	"fmt"
	"slices"
	"strings"

	"github.com/flexigpt/inference-go/spec"
)

func normalizeRequestCacheControls(
	req *spec.FetchCompletionRequest,
	cacheCaps *spec.CacheCapabilities,
) []spec.Warning {
	if req == nil {
		return nil
	}

	var warnings []spec.Warning

	req.ModelParam.CacheControl, warnings = normalizeCacheControlForScope(
		req.ModelParam.CacheControl,
		"modelParam.cacheControl",
		cacheControlScope(cacheCaps, "topLevel"),
		warnings,
	)

	for i := range req.Inputs {
		in := &req.Inputs[i]
		switch in.Kind {
		case spec.InputKindInputMessage:
			if in.InputMessage != nil {
				in.InputMessage.CacheControl, warnings = normalizeCacheControlForScope(
					in.InputMessage.CacheControl,
					fmt.Sprintf("inputs[%d].inputMessage.cacheControl", i),
					cacheControlScope(cacheCaps, "inputOutputContent"),
					warnings,
				)
			}
		case spec.InputKindOutputMessage:
			if in.OutputMessage != nil {
				in.OutputMessage.CacheControl, warnings = normalizeCacheControlForScope(
					in.OutputMessage.CacheControl,
					fmt.Sprintf("inputs[%d].outputMessage.cacheControl", i),
					cacheControlScope(cacheCaps, "inputOutputContent"),
					warnings,
				)
			}
		case spec.InputKindReasoningMessage:
			if in.ReasoningMessage != nil {
				in.ReasoningMessage.CacheControl, warnings = normalizeCacheControlForScope(
					in.ReasoningMessage.CacheControl,
					fmt.Sprintf("inputs[%d].reasoningMessage.cacheControl", i),
					cacheControlScope(cacheCaps, "reasoningContent"),
					warnings,
				)
			}
		case spec.InputKindFunctionToolCall:
			if in.FunctionToolCall != nil {
				in.FunctionToolCall.CacheControl, warnings = normalizeCacheControlForScope(
					in.FunctionToolCall.CacheControl,
					fmt.Sprintf("inputs[%d].functionToolCall.cacheControl", i),
					cacheControlScope(cacheCaps, "toolCall"),
					warnings,
				)
			}
		case spec.InputKindCustomToolCall:
			if in.CustomToolCall != nil {
				in.CustomToolCall.CacheControl, warnings = normalizeCacheControlForScope(
					in.CustomToolCall.CacheControl,
					fmt.Sprintf("inputs[%d].customToolCall.cacheControl", i),
					cacheControlScope(cacheCaps, "toolCall"),
					warnings,
				)
			}
		case spec.InputKindWebSearchToolCall:
			if in.WebSearchToolCall != nil {
				in.WebSearchToolCall.CacheControl, warnings = normalizeCacheControlForScope(
					in.WebSearchToolCall.CacheControl,
					fmt.Sprintf("inputs[%d].webSearchToolCall.cacheControl", i),
					cacheControlScope(cacheCaps, "toolCall"),
					warnings,
				)
			}
		case spec.InputKindFunctionToolOutput:
			if in.FunctionToolOutput != nil {
				in.FunctionToolOutput.CacheControl, warnings = normalizeCacheControlForScope(
					in.FunctionToolOutput.CacheControl,
					fmt.Sprintf("inputs[%d].functionToolOutput.cacheControl", i),
					cacheControlScope(cacheCaps, "toolOutput"),
					warnings,
				)
			}
		case spec.InputKindCustomToolOutput:
			if in.CustomToolOutput != nil {
				in.CustomToolOutput.CacheControl, warnings = normalizeCacheControlForScope(
					in.CustomToolOutput.CacheControl,
					fmt.Sprintf("inputs[%d].customToolOutput.cacheControl", i),
					cacheControlScope(cacheCaps, "toolOutput"),
					warnings,
				)
			}
		case spec.InputKindWebSearchToolOutput:
			if in.WebSearchToolOutput != nil {
				in.WebSearchToolOutput.CacheControl, warnings = normalizeCacheControlForScope(
					in.WebSearchToolOutput.CacheControl,
					fmt.Sprintf("inputs[%d].webSearchToolOutput.cacheControl", i),
					cacheControlScope(cacheCaps, "toolOutput"),
					warnings,
				)
			}
		default:
		}
	}

	for i := range req.ToolChoices {
		req.ToolChoices[i].CacheControl, warnings = normalizeCacheControlForScope(
			req.ToolChoices[i].CacheControl,
			fmt.Sprintf("toolChoices[%d].cacheControl", i),
			cacheControlScope(cacheCaps, "toolChoice"),
			warnings,
		)
	}

	return warnings
}

func normalizeCacheControlForScope(
	cc *spec.CacheControl,
	scopePath string,
	scopeCaps *spec.CacheControlCapabilities,
	warnings []spec.Warning,
) (*spec.CacheControl, []spec.Warning) {
	if cc == nil {
		return nil, warnings
	}

	if scopeCaps == nil {
		warnings = append(warnings, spec.Warning{
			Code:    "cacheControl_dropped_unsupported",
			Message: scopePath + " was dropped because cache control is unsupported by this SDK/model.",
		})
		return nil, warnings
	}

	out := *cc
	out.Key = strings.TrimSpace(out.Key)

	if out.Kind != "" && len(scopeCaps.SupportedKinds) > 0 && !slices.Contains(scopeCaps.SupportedKinds, out.Kind) {
		warnings = append(warnings, spec.Warning{
			Code: "cacheControl_dropped_unsupported_kind",
			Message: fmt.Sprintf(
				"%s was dropped because cacheControl.kind %q is unsupported by this SDK/model.",
				scopePath,
				out.Kind,
			),
		})
		return nil, warnings
	}

	if out.TTL != "" && (!scopeCaps.SupportsTTL ||
		(len(scopeCaps.SupportedTTLs) > 0 && !slices.Contains(scopeCaps.SupportedTTLs, out.TTL))) {
		warnings = append(warnings, getCacheTTLWarning(scopePath, out.TTL))
		out.TTL = ""
	}

	if out.Key != "" && !scopeCaps.SupportsKey {
		warnings = append(warnings, spec.Warning{
			Code:    "cacheControl_key_dropped_unsupported",
			Message: scopePath + ".key was dropped because cache keys are unsupported by this SDK/model.",
		})
		out.Key = ""
	}

	if out.Kind == "" && out.TTL == "" && out.Key == "" {
		return nil, warnings
	}

	return &out, warnings
}

func getCacheTTLWarning(scopePath string, ttl spec.CacheControlTTL) spec.Warning {
	return spec.Warning{
		Code: "cacheControl_ttl_dropped_unsupported",
		Message: fmt.Sprintf(
			"%s.ttl %q was dropped because cache TTL/retention is unsupported by this SDK/model.",
			scopePath,
			ttl,
		),
	}
}

func cacheControlScope(
	cacheCaps *spec.CacheCapabilities,
	scope string,
) *spec.CacheControlCapabilities {
	if cacheCaps == nil {
		return nil
	}

	switch scope {
	case "topLevel":
		return cacheCaps.TopLevel
	case "inputOutputContent":
		return cacheCaps.InputOutputContent
	case "reasoningContent":
		return cacheCaps.ReasoningContent
	case "toolChoice":
		return cacheCaps.ToolChoice
	case "toolCall":
		return cacheCaps.ToolCall
	case "toolOutput":
		return cacheCaps.ToolOutput
	default:
		return nil
	}
}
