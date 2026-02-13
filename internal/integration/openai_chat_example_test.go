package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

// Example_openAIChat_basicConversation demonstrates a minimal non-streaming
// call to OpenAI's Chat Completions API.
func Example_openAIChat_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug()
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

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "gpt-4.1-mini",
			Stream:          false,
			MaxPromptLength: 4096,
			MaxOutputLength: 256,
			SystemPrompt:    "You are a concise assistant.",
		},
		Inputs: []spec.InputUnion{
			{
				Kind: spec.InputKindInputMessage,
				InputMessage: &spec.InputOutputContent{
					Role: spec.RoleUser,
					Contents: []spec.InputOutputContentItemUnion{
						{
							Kind: spec.ContentItemKindText,
							TextItem: &spec.ContentItemText{
								Text: "Say hello from OpenAI Chat Completions in one short sentence.",
							},
						},
					},
				},
			},
		},
	}

	resp, err := ps.FetchCompletion(ctx, "openai-chat", req, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FetchCompletion error:", err)
		if resp != nil && resp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", resp.Error.Message)
		}
		return
	}

	for _, out := range resp.Outputs {
		if out.Kind != spec.OutputKindOutputMessage || out.OutputMessage == nil {
			continue
		}
		for _, c := range out.OutputMessage.Contents {
			if c.Kind == spec.ContentItemKindText && c.TextItem != nil {
				fmt.Fprintln(os.Stderr, "OpenAI Chat assistant:", c.TextItem.Text)
			}
		}
	}

	fmt.Println("OK")
	// Output: OK
}
