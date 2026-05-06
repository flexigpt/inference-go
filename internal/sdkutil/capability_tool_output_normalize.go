package sdkutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"github.com/flexigpt/inference-go/spec"
)

const toolOutputCollapsedToString = "toolOutput_collapsed_to_string"

func normalizeClientToolOutputsForSDK(
	req *spec.FetchCompletionRequest,
	toolCaps *spec.ToolCapabilities,
) ([]spec.Warning, error) {
	warnings := []spec.Warning{}
	if req == nil || toolCaps == nil {
		return warnings, nil
	}

	// Backward-compatible default:
	// if capability is unset, assume structured content-list transport is okay.
	if len(toolCaps.SupportedClientToolOutputFormats) == 0 {
		return warnings, nil
	}

	if slices.Contains(
		toolCaps.SupportedClientToolOutputFormats,
		spec.ToolOutputFormatKindContentItemList,
	) {
		return warnings, nil
	}

	if !slices.Contains(
		toolCaps.SupportedClientToolOutputFormats,
		spec.ToolOutputFormatKindString,
	) {
		return nil, errors.New(
			"client tool outputs are unsupported by this SDK/model: no supportedClientToolOutputFormats matched",
		)
	}

	for i := range req.Inputs {
		in := &req.Inputs[i]

		switch in.Kind {
		case spec.InputKindFunctionToolOutput:
			if in.FunctionToolOutput == nil {
				continue
			}
			changed, err := collapseToolOutputToSingleText(in.FunctionToolOutput)
			if err != nil {
				return warnings, fmt.Errorf("normalize function tool output at inputs[%d]: %w", i, err)
			}
			if changed {
				warnings = append(warnings, spec.Warning{
					Code: toolOutputCollapsedToString,
					Message: fmt.Sprintf(
						"inputs[%d].functionToolOutput was collapsed to a single string because this SDK/model only supports string client tool outputs.",
						i,
					),
				})
			}

		case spec.InputKindCustomToolOutput:
			if in.CustomToolOutput == nil {
				continue
			}
			changed, err := collapseToolOutputToSingleText(in.CustomToolOutput)
			if err != nil {
				return warnings, fmt.Errorf("normalize custom tool output at inputs[%d]: %w", i, err)
			}
			if changed {
				warnings = append(warnings, spec.Warning{
					Code: toolOutputCollapsedToString,
					Message: fmt.Sprintf(
						"inputs[%d].customToolOutput was collapsed to a single string because this SDK/model only supports string client tool outputs.",
						i,
					),
				})
			}
		default:
			// Non Tool output.
		}
	}

	return warnings, nil
}

func collapseToolOutputToSingleText(out *spec.ToolOutput) (bool, error) {
	if out == nil || len(out.Contents) == 0 {
		return false, nil
	}

	// Already representable as a string payload.
	if singleTextToolOutput(out) != nil {
		return false, nil
	}

	raw, err := json.Marshal(out.Contents)
	if err != nil {
		return false, err
	}

	out.Contents = []spec.ToolOutputItemUnion{{
		Kind: spec.ContentItemKindText,
		TextItem: &spec.ContentItemText{
			Text: string(raw),
		},
	}}
	return true, nil
}

func singleTextToolOutput(out *spec.ToolOutput) *spec.ContentItemText {
	if out == nil || len(out.Contents) != 1 {
		return nil
	}
	item := out.Contents[0]
	if item.Kind != spec.ContentItemKindText || item.TextItem == nil {
		return nil
	}
	return item.TextItem
}
