//go:build !integration

package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ppipada/inference-go"
	"github.com/ppipada/inference-go/spec"
)

// Example_anthropic_basicConversation demonstrates a minimal non-streaming
// call to Anthropic's Messages API using the normalized inference-go API.
func Example_anthropic_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug()
	if err != nil {
		fmt.Println("error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, "anthropic", &inference.AddProviderConfig{
		SDKType:                  spec.ProviderSDKTypeAnthropic,
		Origin:                   spec.DefaultAnthropicOrigin,
		ChatCompletionPathPrefix: spec.DefaultAnthropicChatCompletionPrefix,
		APIKeyHeaderKey:          spec.DefaultAnthropicAuthorizationHeaderKey,
		// DefaultHeaders are optional; the official SDK sets anthropic-version.
	})
	if err != nil {
		fmt.Println("error adding Anthropic provider:", err)
		return
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Println("ANTHROPIC_API_KEY not set; skipping live Anthropic call")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, "anthropic", apiKey); err != nil {
		fmt.Println("error setting Anthropic API key:", err)
		return
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "claude-3-5-sonnet-20241022",
			Stream:          false,
			MaxPromptLength: 4096,
			MaxOutputLength: 256,
			SystemPrompt:    "You are a concise, helpful assistant.",
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
								Text: "Say hello from Anthropic in one short sentence.",
							},
						},
					},
				},
			},
		},
	}

	resp, err := ps.FetchCompletion(ctx, "anthropic", req, nil)
	if err != nil {
		fmt.Println("FetchCompletion error:", err)
		if resp != nil && resp.Error != nil {
			fmt.Println("Provider error:", resp.Error.Message)
		}
		return
	}

	for _, out := range resp.Outputs {
		if out.Kind != spec.OutputKindOutputMessage || out.OutputMessage == nil {
			continue
		}
		for _, c := range out.OutputMessage.Contents {
			if c.Kind == spec.ContentItemKindText && c.TextItem != nil {
				fmt.Println("Anthropic assistant:", c.TextItem.Text)
			}
		}
	}

	// Output:
}
