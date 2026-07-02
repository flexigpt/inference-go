package integration

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

func Example_googleGenerateContent_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	pp, mp, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderGoogleGemini,
		modelpreset.PresetGemini25Flash,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Google Gemini preset provider:", err)
		return
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "GEMINI_API_KEY/GOOGLE_API_KEY not set; skipping live Google Gemini call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google Gemini API key:", err)
		return
	}

	modelParam := mp.ModelParam
	modelParam.Stream = false
	modelParam.MaxPromptLength = min(modelParam.MaxPromptLength, 4096)
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 2048)
	modelParam.SystemPrompt = "You are a concise, helpful assistant."

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}

	resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{
			newUserTextInput("Say hello from Google Gemini in one short sentence."),
		},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FetchCompletion error:", err)
		if resp != nil && resp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", resp.Error.Message)
		}
		return
	}

	fmt.Fprintln(os.Stderr, "Gemini assistant:", responseText(resp))
	fmt.Println("OK")
	// Output: OK
}
