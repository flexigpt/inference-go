package integration

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/debugclient"
	"github.com/flexigpt/inference-go/spec"
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

	result := "ECHO: " + args.Text

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
				TextItem: &spec.ContentItemText{Text: result},
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

func nonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
