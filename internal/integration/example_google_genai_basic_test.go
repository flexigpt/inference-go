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
	googleBasicProviderName  = "google"
	googleBasicModelName     = "gemini-3-flash-preview"
	googleBasicCompletionKey = "gemini-flash"
)

// Example_googleGenerateContent_basicConversation demonstrates a minimal non-streaming
// call to Google's Generative AI API (Gemini) using the normalized inference-go API.
func Example_googleGenerateContent_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, googleBasicProviderName, &inference.AddProviderConfig{
		SDKType: spec.ProviderSDKTypeGoogleGenerateContent,
		Origin:  spec.DefaultGoogleGenerateContentOrigin,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Google GenAI provider:", err)
		return
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "GEMINI_API_KEY not set; skipping live Google GenAI call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, googleBasicProviderName, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google GenAI API key:", err)
		return
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            googleBasicModelName,
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
								Text: "Say hello from Google Gemini in one short sentence.",
							},
						},
					},
				},
			},
		},
	}

	resp, err := ps.FetchCompletion(
		ctx, googleBasicProviderName, req,
		&spec.FetchCompletionOptions{CompletionKey: googleBasicCompletionKey},
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
				fmt.Fprintln(os.Stderr, "Gemini assistant:", c.TextItem.Text)
			}
		}
	}
	fmt.Println("OK")
	// Output: OK
}
