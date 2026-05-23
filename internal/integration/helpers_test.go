package integration

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/debugclient"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

const (
	toolJSONKeyType                 = "type"
	toolJSONValueObject             = "object"
	toolJSONKeyProperties           = "properties"
	toolJSONKeyText                 = "text"
	toolJSONValueString             = "string"
	toolJSONKeyRequired             = "required"
	toolJSONKeyAdditionalProperties = "additionalProperties"
	toolJSONSchemaName              = "answer"
	toolJSONKeySummary              = "summary"
	toolJSONKeySourceUsed           = "source_used"
	toolJSONValueBoolean            = "boolean"
)

// newProviderSetWithDebug constructs a ProviderSetAPI with:
//
//   - a text slog.Logger writing to stdout at debug level
//   - an HTTPCompletionDebugger that logs HTTP request/response metadata
//
// The examples reuse this helper to keep them short.
func newProviderSetWithDebug(level slog.Level) (*inference.ProviderSetAPI, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)

	return inference.NewProviderSetAPI(
		inference.WithLogger(logger),
		inference.WithDebugClientBuilder(func(p spec.ProviderParam) spec.CompletionDebugger {
			cfg := &debugclient.DebugConfig{
				// Capture and log request/response metadata (headers, URLs,
				// status codes, etc.). Bodies are scrubbed by default.
				LogToSlog: true,
			}
			return debugclient.NewHTTPCompletionDebugger(cfg)
		}),
	)
}

func addCatalogModelProvider(
	ctx context.Context,
	ps *inference.ProviderSetAPI,
	providerName spec.ProviderName,
	modelID modelpreset.ModelPresetID,
) (modelpreset.ProviderPreset, modelpreset.ModelPreset, error) {
	pp, err := modelpreset.Provider(providerName)
	if err != nil {
		return modelpreset.ProviderPreset{}, modelpreset.ModelPreset{}, err
	}

	mp, err := modelpreset.Model(providerName, modelID)
	if err != nil {
		return modelpreset.ProviderPreset{}, modelpreset.ModelPreset{}, err
	}

	if _, err := ps.AddProviderFromPreset(ctx, "", pp); err != nil {
		return modelpreset.ProviderPreset{}, modelpreset.ModelPreset{}, err
	}

	return pp, mp, nil
}

func presetFetchOptions(
	ctx context.Context,
	ps *inference.ProviderSetAPI,
	pp modelpreset.ProviderPreset,
	mp modelpreset.ModelPreset,
) (*spec.FetchCompletionOptions, error) {
	completionKey := string(mp.ID)

	resolver, err := ps.NewPresetCapabilityResolver(ctx, pp.Name, pp, mp, completionKey)
	if err != nil {
		return nil, err
	}

	return &spec.FetchCompletionOptions{
		CompletionKey:      completionKey,
		CapabilityResolver: resolver,
	}, nil
}

func newUserTextInput(text string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindInputMessage,
		InputMessage: &spec.InputOutputContent{
			Role: spec.RoleUser,
			Contents: []spec.InputOutputContentItemUnion{
				{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: text},
				},
			},
		},
	}
}

func newEchoToolChoice(id, name string) spec.ToolChoice {
	return spec.ToolChoice{
		Type:        spec.ToolTypeFunction,
		ID:          id,
		Name:        name,
		Description: "Echo the provided text back in a deterministic tool result.",
		Arguments: map[string]any{
			toolJSONKeyType: toolJSONValueObject,
			toolJSONKeyProperties: map[string]any{
				toolJSONKeyText: map[string]any{toolJSONKeyType: toolJSONValueString},
			},
			toolJSONKeyRequired:             []any{toolJSONKeyText},
			toolJSONKeyAdditionalProperties: false,
		},
	}
}

func firstFunctionToolCall(resp *spec.FetchCompletionResponse) (*spec.ToolCall, error) {
	if resp == nil {
		return nil, errors.New("nil response")
	}
	for _, out := range resp.Outputs {
		if out.Kind == spec.OutputKindFunctionToolCall && out.FunctionToolCall != nil {
			return out.FunctionToolCall, nil
		}
	}
	return nil, fmt.Errorf("no function tool call found in %d outputs", len(resp.Outputs))
}

func runEchoTool(call *spec.ToolCall) (*spec.ToolOutput, error) {
	if call == nil {
		return nil, errors.New("nil tool call")
	}

	callID := nonEmpty(call.CallID, call.ID)
	if callID == "" {
		return nil, errors.New("tool call missing call id")
	}

	var args struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal([]byte(call.Arguments), &args); err != nil {
		return nil, fmt.Errorf("decode tool arguments: %w", err)
	}
	if strings.TrimSpace(args.Text) == "" {
		return nil, errors.New("tool argument text is empty")
	}

	return &spec.ToolOutput{
		Type:     call.Type,
		ChoiceID: call.ChoiceID,
		ID:       callID,
		Role:     spec.RoleTool,
		CallID:   callID,
		Name:     call.Name,
		Contents: []spec.ToolOutputItemUnion{
			{
				Kind:     spec.ContentItemKindText,
				TextItem: &spec.ContentItemText{Text: "ECHO: " + args.Text},
			},
		},
	}, nil
}

func firstToolOutputText(out *spec.ToolOutput) string {
	if out == nil {
		return ""
	}
	for _, item := range out.Contents {
		if item.Kind == spec.ContentItemKindText && item.TextItem != nil {
			return item.TextItem.Text
		}
	}
	return ""
}

func responseText(resp *spec.FetchCompletionResponse) string {
	if resp == nil {
		return ""
	}

	var b strings.Builder
	for _, out := range resp.Outputs {
		if out.Kind != spec.OutputKindOutputMessage || out.OutputMessage == nil {
			continue
		}
		for _, item := range out.OutputMessage.Contents {
			if item.Kind == spec.ContentItemKindText && item.TextItem != nil {
				if b.Len() > 0 {
					b.WriteString(" ")
				}
				b.WriteString(item.TextItem.Text)
			}
		}
	}
	return strings.TrimSpace(b.String())
}

func outputUnionsToInputs(outputs []spec.OutputUnion) []spec.InputUnion {
	if len(outputs) == 0 {
		return nil
	}

	out := make([]spec.InputUnion, 0, len(outputs))
	for _, item := range outputs {
		switch item.Kind {
		case spec.OutputKindOutputMessage:
			if item.OutputMessage != nil {
				out = append(out, spec.InputUnion{
					Kind:          spec.InputKindOutputMessage,
					OutputMessage: item.OutputMessage,
				})
			}
		case spec.OutputKindReasoningMessage:
			if item.ReasoningMessage != nil {
				out = append(out, spec.InputUnion{
					Kind:             spec.InputKindReasoningMessage,
					ReasoningMessage: item.ReasoningMessage,
				})
			}
		case spec.OutputKindFunctionToolCall:
			if item.FunctionToolCall != nil {
				out = append(out, spec.InputUnion{
					Kind:             spec.InputKindFunctionToolCall,
					FunctionToolCall: item.FunctionToolCall,
				})
			}
		case spec.OutputKindCustomToolCall:
			if item.CustomToolCall != nil {
				out = append(out, spec.InputUnion{
					Kind:           spec.InputKindCustomToolCall,
					CustomToolCall: item.CustomToolCall,
				})
			}
		case spec.OutputKindWebSearchToolCall:
			if item.WebSearchToolCall != nil {
				out = append(out, spec.InputUnion{
					Kind:              spec.InputKindWebSearchToolCall,
					WebSearchToolCall: item.WebSearchToolCall,
				})
			}
		case spec.OutputKindWebSearchToolOutput:
			if item.WebSearchToolOutput != nil {
				out = append(out, spec.InputUnion{
					Kind:                spec.InputKindWebSearchToolOutput,
					WebSearchToolOutput: item.WebSearchToolOutput,
				})
			}
		}
	}

	return out
}

func nonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
