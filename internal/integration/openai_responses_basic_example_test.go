package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

// Example_openAIResponses_basicConversation demonstrates a minimal non-streaming
// call to OpenAI's Responses API using text-only input.
func Example_openAIResponses_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, "openai-responses", &inference.AddProviderConfig{
		SDKType: spec.ProviderSDKTypeOpenAIResponses,
		Origin:  spec.DefaultOpenAIOrigin,
		// Only used when Origin is overridden; kept here for clarity.
		ChatCompletionPathPrefix: "/v1/responses",
		APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding OpenAI Responses provider:", err)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENAI_API_KEY not set; skipping live OpenAI Responses call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, "openai-responses", apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "gpt-5-mini",
			Stream:          false,
			MaxPromptLength: 4096,
			MaxOutputLength: 256,
			SystemPrompt:    "You are a concise assistant.",
			Reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelLow,
			},
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
								Text: "Explain the difference between goroutines and OS threads in 2–3 sentences.",
							},
						},
					},
				},
			},
		},
	}

	resp, err := ps.FetchCompletion(ctx, "openai-responses", req, nil)
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
				fmt.Fprintln(os.Stderr, "OpenAI Responses assistant:", c.TextItem.Text)
			}
		}
	}

	fmt.Println("OK")
	// Output: OK
}
