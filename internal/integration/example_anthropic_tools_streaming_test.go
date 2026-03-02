package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

// Example_anthropic_toolsAndThinkingStreaming demonstrates:
//   - streaming text + thinking
//   - function tools + anthropic server web search
//   - JSON schema output request
func Example_anthropic_toolsAndThinkingStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, "anthropic", &inference.AddProviderConfig{
		SDKType:                  spec.ProviderSDKTypeAnthropic,
		Origin:                   spec.DefaultAnthropicOrigin,
		ChatCompletionPathPrefix: spec.DefaultAnthropicChatCompletionPrefix,
		APIKeyHeaderKey:          spec.DefaultAnthropicAuthorizationHeaderKey,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Anthropic provider:", err)
		return
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "ANTHROPIC_API_KEY not set; skipping live Anthropic call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, "anthropic", apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Anthropic API key:", err)
		return
	}

	tools := []spec.ToolChoice{
		{
			Type:        spec.ToolTypeFunction,
			ID:          "extract-key-points",
			Name:        "extract_key_points",
			Description: "Extract 3 key points from the provided text.",
			Arguments: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"text": map[string]any{"type": "string"},
				},
				"required":             []any{"text"},
				"additionalProperties": false,
			},
		},
		{
			Type: spec.ToolTypeWebSearch,
			ID:   "web-search",
			Name: spec.DefaultWebSearchToolName,
			WebSearchArguments: &spec.WebSearchToolChoiceItem{
				MaxUses:           1,
				SearchContextSize: spec.WebSearchContextSizeMedium,
			},
		},
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:         "claude-sonnet-4-6",
			Stream:       true,
			SystemPrompt: "Use tools when helpful. Keep the final answer short.",
			Reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelMedium,
			},
			OutputParam: &spec.OutputParam{
				Format: &spec.OutputFormat{
					Kind: spec.OutputFormatKindJSONSchema,
					JSONSchemaParam: &spec.JSONSchemaParam{
						Name: "answer",
						Schema: map[string]any{
							"type": "object",
							"properties": map[string]any{
								"summary":     map[string]any{"type": "string"},
								"source_used": map[string]any{"type": "boolean"},
							},
							"required":             []any{"summary", "source_used"},
							"additionalProperties": false,
						},
					},
				},
			},
		},
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "What is the latest stable Go version? If unknown, say so."},
				}},
			},
		}},
		ToolChoices: tools,
		ToolPolicy: &spec.ToolPolicy{
			Mode:            spec.ToolPolicyModeAuto,
			DisableParallel: true,
		},
	}

	_, err = ps.FetchCompletion(ctx, "anthropic", req, &spec.FetchCompletionOptions{
		CompletionKey: "sonnet46",
		StreamHandler: func(ev spec.StreamEvent) error {
			switch ev.Kind {
			case spec.StreamContentKindThinking:
				if ev.Thinking != nil {
					fmt.Fprintf(os.Stderr, "[thinking] %s\n", ev.Thinking.Text)
				}
			case spec.StreamContentKindText:
				if ev.Text != nil {
					fmt.Fprint(os.Stderr, ev.Text.Text)
				}
			}
			return nil
		},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nFetchCompletion error:", err)
		return
	}

	fmt.Println("OK")
	// Output: OK
}
