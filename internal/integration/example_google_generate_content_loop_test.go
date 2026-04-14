package integration

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

func TestGoogleGenerateContent_FunctionToolRoundTripLoop(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY/GOOGLE_API_KEY not set")
	}

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Minute)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		t.Fatal(err)
	}

	const providerName = "google-loop"
	_, err = ps.AddProvider(ctx, providerName, &inference.AddProviderConfig{
		SDKType: spec.ProviderSDKTypeGoogleGenerateContent,
		Origin:  spec.DefaultGoogleGenerateContentOrigin,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := ps.SetProviderAPIKey(ctx, providerName, apiKey); err != nil {
		t.Fatal(err)
	}

	sawSignedReasoning := false
	for i := range 3 {
		saw, err := runGoogleGenerateContentEchoRoundTrip(ctx, ps, providerName, fmt.Sprintf("google loop %d", i))
		if err != nil {
			t.Fatalf("iteration %d failed: %v", i, err)
		}
		if saw {
			sawSignedReasoning = true
		}
	}

	if !sawSignedReasoning {
		t.Fatalf("did not observe any signed reasoning across loop; thought-signature replay path was not exercised")
	}
}

func runGoogleGenerateContentEchoRoundTrip(
	ctx context.Context,
	ps *inference.ProviderSetAPI,
	providerName spec.ProviderName,
	payload string,
) (bool, error) {
	tool := spec.ToolChoice{
		Type:        spec.ToolTypeFunction,
		ID:          "echo-tool",
		Name:        "echo_text",
		Description: "Echo the provided text back in a deterministic tool result.",
		Arguments: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"text": map[string]any{"type": "string"},
			},
			"required":             []any{"text"},
			"additionalProperties": false,
		},
	}

	history := []spec.InputUnion{
		newUserTextInput(
			fmt.Sprintf(
				`Think briefly, then call echo_text with text %q. After the tool result arrives, answer in one short sentence.`,
				payload,
			),
		),
	}

	sawToolCall := false
	sawSignedReasoning := false

	for turn := range 4 {
		policy := &spec.ToolPolicy{Mode: spec.ToolPolicyModeNone}
		if !sawToolCall {
			policy = &spec.ToolPolicy{
				Mode: spec.ToolPolicyModeTool,
				AllowedTools: []spec.AllowedTool{
					{ToolChoiceID: tool.ID},
				},
				DisableParallel: true,
			}
		}

		resp, err := ps.FetchCompletion(ctx, providerName, &spec.FetchCompletionRequest{
			ModelParam: spec.ModelParam{
				Name:            "gemini-2.5-flash",
				MaxOutputLength: 256,
				SystemPrompt: "Preserve prior assistant reasoning/tool state across turns. " +
					"After the tool result is available, answer plainly and do not call the tool again.",
				Reasoning: &spec.ReasoningParam{
					Type:   spec.ReasoningTypeHybridWithTokens,
					Tokens: 1024,
				},
			},
			Inputs:      history,
			ToolChoices: []spec.ToolChoice{tool},
			ToolPolicy:  policy,
		}, &spec.FetchCompletionOptions{
			CompletionKey: "gemini-2.5-flash",
		})
		if err != nil {
			return sawSignedReasoning, fmt.Errorf("turn %d fetch: %w", turn, err)
		}

		if responseHasGoogleSignature(resp) {
			sawSignedReasoning = true
		}

		history = append(history, outputUnionsToInputs(resp.Outputs)...)

		call, _ := firstFunctionToolCall(resp)
		if call != nil && !sawToolCall {
			sawToolCall = true

			toolOutput, err := runEchoTool(call)
			if err != nil {
				return sawSignedReasoning, fmt.Errorf("turn %d tool exec: %w", turn, err)
			}

			history = append(history, spec.InputUnion{
				Kind:               spec.InputKindFunctionToolOutput,
				FunctionToolOutput: toolOutput,
			})
			continue
		}

		if !sawToolCall {
			return sawSignedReasoning, fmt.Errorf("turn %d: expected tool call before final answer", turn)
		}
		if responseText(resp) == "" {
			return sawSignedReasoning, fmt.Errorf("turn %d: expected final assistant text", turn)
		}

		return sawSignedReasoning, nil
	}

	return sawSignedReasoning, errors.New("conversation did not complete within 4 turns")
}

func responseHasGoogleSignature(resp *spec.FetchCompletionResponse) bool {
	if resp == nil {
		return false
	}
	for _, out := range resp.Outputs {
		switch out.Kind {
		case spec.OutputKindReasoningMessage:
			if out.ReasoningMessage != nil && out.ReasoningMessage.Signature != "" {
				return true
			}
		case spec.OutputKindFunctionToolCall:
			if out.FunctionToolCall != nil && out.FunctionToolCall.Signature != "" {
				return true
			}
		case spec.OutputKindCustomToolCall:
			if out.CustomToolCall != nil && out.CustomToolCall.Signature != "" {
				return true
			}
		case spec.OutputKindOutputMessage:
			if out.OutputMessage == nil {
				continue
			}
			for _, c := range out.OutputMessage.Contents {
				if c.Kind == spec.ContentItemKindText &&
					c.TextItem != nil &&
					c.TextItem.Signature != "" {
					return true
				}
			}
		default:
			// Continue.
		}
	}
	return false
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
