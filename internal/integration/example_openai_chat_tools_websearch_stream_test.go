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

const (
	openAIChatToolsProviderName                = "openai-chat"
	openAIChatToolsFunctionToolID              = "math"
	openAIChatToolsFunctionToolName            = "multiply"
	openAIChatToolsFunctionToolDescription     = "Multiply two integers."
	openAIChatToolsJSONKeyType                 = "type"
	openAIChatToolsJSONValueObject             = "object"
	openAIChatToolsJSONKeyProperties           = "properties"
	openAIChatToolsJSONKeyRequired             = "required"
	openAIChatToolsJSONKeyAdditionalProperties = "additionalProperties"
	openAIChatToolsJSONAnswerKey               = "answer"
	openAIChatToolsSchemaName                  = "result"
	openAIChatToolsModelName                   = "gpt-4.1"
	openAIChatToolsCompletionKey               = "gpt41"
	openAIChatToolsSystemPrompt                = "answer directly"
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

	_, err = ps.AddProvider(ctx, openAIChatToolsProviderName, &inference.AddProviderConfig{
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
	if err := ps.SetProviderAPIKey(ctx, openAIChatToolsProviderName, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	tools := []spec.ToolChoice{
		{
			Type:        spec.ToolTypeFunction,
			ID:          openAIChatToolsFunctionToolID,
			Name:        openAIChatToolsFunctionToolName,
			Description: openAIChatToolsFunctionToolDescription,
			Arguments: map[string]any{
				openAIChatToolsJSONKeyType: openAIChatToolsJSONValueObject,
				openAIChatToolsJSONKeyProperties: map[string]any{
					"a": map[string]any{openAIChatToolsJSONKeyType: "integer"},
					"b": map[string]any{openAIChatToolsJSONKeyType: "integer"},
				},
				openAIChatToolsJSONKeyRequired:             []any{"a", "b"},
				openAIChatToolsJSONKeyAdditionalProperties: false,
			},
		},
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:         openAIChatToolsModelName,
			Stream:       true,
			SystemPrompt: openAIChatToolsSystemPrompt,
			OutputParam: &spec.OutputParam{
				Format: &spec.OutputFormat{
					Kind: spec.OutputFormatKindJSONSchema,
					JSONSchemaParam: &spec.JSONSchemaParam{
						Name: openAIChatToolsSchemaName,
						Schema: map[string]any{
							openAIChatToolsJSONKeyType: openAIChatToolsJSONValueObject,
							openAIChatToolsJSONKeyProperties: map[string]any{
								openAIChatToolsJSONAnswerKey: map[string]any{openAIChatToolsJSONKeyType: "string"},
							},
							openAIChatToolsJSONKeyRequired:             []any{openAIChatToolsJSONAnswerKey},
							openAIChatToolsJSONKeyAdditionalProperties: false,
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

	_, err = ps.FetchCompletion(ctx, openAIChatToolsProviderName, req, &spec.FetchCompletionOptions{
		CompletionKey: openAIChatToolsCompletionKey,
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
