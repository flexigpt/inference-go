package integration

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

// Example_openAIChat_toolsAndJSONSchema demonstrates:
//   - streaming text
//   - JSON schema output (response_format=json_schema)
//   - function tools
func Example_openAIChat_toolsAndJSONSchema() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, "openai-chat", &inference.AddProviderConfig{
		SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
		Origin:                   spec.DefaultOpenAIOrigin,
		ChatCompletionPathPrefix: spec.DefaultOpenAIChatCompletionsPrefix,
		APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
		DefaultHeaders:           spec.OpenAIChatCompletionsDefaultHeaders,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding OpenAI Chat provider:", err)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENAI_API_KEY not set; skipping live OpenAI Chat call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, "openai-chat", apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	tools := []spec.ToolChoice{
		{
			Type:        spec.ToolTypeFunction,
			ID:          "math",
			Name:        "multiply",
			Description: "Multiply two integers.",
			Arguments: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"a": map[string]any{"type": "integer"},
					"b": map[string]any{"type": "integer"},
				},
				"required":             []any{"a", "b"},
				"additionalProperties": false,
			},
		},
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:         "gpt-4.1",
			Stream:       true,
			SystemPrompt: "answer directly",
			OutputParam: &spec.OutputParam{
				Format: &spec.OutputFormat{
					Kind: spec.OutputFormatKindJSONSchema,
					JSONSchemaParam: &spec.JSONSchemaParam{
						Name: "result",
						Schema: map[string]any{
							"type": "object",
							"properties": map[string]any{
								"answer": map[string]any{"type": "string"},
							},
							"required":             []any{"answer"},
							"additionalProperties": false,
						},
						Strict: true,
					},
				},
			},
		},
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind: spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{
						Text: "What is 6*7 and what's a recent Go release? (If unknown, say unknown.)",
					},
				}},
			},
		}},
		ToolChoices: tools,
		ToolPolicy: &spec.ToolPolicy{
			Mode:            spec.ToolPolicyModeAuto,
			DisableParallel: true,
		},
	}

	_, err = ps.FetchCompletion(ctx, "openai-chat", req, &spec.FetchCompletionOptions{
		CompletionKey: "gpt41",
		StreamHandler: func(ev spec.StreamEvent) error {
			if ev.Kind == spec.StreamContentKindText && ev.Text != nil {
				fmt.Fprint(os.Stderr, ev.Text.Text)
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
