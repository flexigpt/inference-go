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

func Example_openAIResponses_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	pp, mp, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderOpenAIResponses,
		modelpreset.PresetGPT5Mini,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding OpenAI Responses preset provider:", err)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENAI_API_KEY not set; skipping live OpenAI Responses call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	modelParam := mp.ModelParam
	modelParam.Stream = false
	modelParam.MaxPromptLength = min(modelParam.MaxPromptLength, 4096)
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 4096)
	modelParam.SystemPrompt = "You are a concise assistant."
	modelParam.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}

	resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{
			newUserTextInput("Explain the difference between goroutines and OS threads in 2-3 sentences."),
		},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FetchCompletion error:", err)
		if resp != nil && resp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", resp.Error.Message)
		}
		return
	}

	fmt.Fprintln(os.Stderr, "OpenAI Responses assistant:", responseText(resp))
	fmt.Println("OK")
	// Output: OK
}
