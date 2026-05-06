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
	openAIResponsesBasicProviderName  = "openai-responses"
	openAIResponsesBasicPathPrefix    = "/v1/responses"
	openAIResponsesBasicModelName     = "gpt-5-mini"
	openAIResponsesBasicCompletionKey = "gpt5mini"
)

// Example_openAIResponses_basicConversation demonstrates a minimal non-streaming
// call to OpenAI's Responses API using text-only input.
func Example_openAIResponses_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, openAIResponsesBasicProviderName, &inference.AddProviderConfig{
		SDKType: spec.ProviderSDKTypeOpenAIResponses,
		Origin:  spec.DefaultOpenAIOrigin,
		// Only used when Origin is overridden; kept here for clarity.
		ChatCompletionPathPrefix: openAIResponsesBasicPathPrefix,
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
	if err := ps.SetProviderAPIKey(ctx, openAIResponsesBasicProviderName, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            openAIResponsesBasicModelName,
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

	resp, err := ps.FetchCompletion(
		ctx,
		openAIResponsesBasicProviderName,
		req,
		&spec.FetchCompletionOptions{CompletionKey: openAIResponsesBasicCompletionKey},
	)
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
