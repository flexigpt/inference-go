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

func Example_openAITerraToMetaMuse_reasoningStitching() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	terraProvider, terraModel, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderOpenAIResponses,
		modelpreset.PresetGPT56Terra,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding OpenAI Terra preset provider:", err)
		return
	}

	museProvider, museModel, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderMeta,
		modelpreset.PresetMuseSpark13Contributor,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Meta Muse preset provider:", err)
		return
	}

	openAIKey := os.Getenv("OPENAI_API_KEY")
	metaKey := os.Getenv("META_API_KEY")
	if openAIKey == "" || metaKey == "" {
		fmt.Fprintln(
			os.Stderr,
			"OPENAI_API_KEY or META_API_KEY not set; skipping live reasoning stitching example",
		)
		fmt.Println("OK")
		return
	}

	if err := ps.SetProviderAPIKey(ctx, terraProvider.Name, openAIKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}
	if err := ps.SetProviderAPIKey(ctx, museProvider.Name, metaKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Meta API key:", err)
		return
	}

	terraOpts, err := presetFetchOptions(ctx, ps, terraProvider, terraModel)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating OpenAI Terra capability resolver:", err)
		return
	}
	museOpts, err := presetFetchOptions(ctx, ps, museProvider, museModel)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating Meta Muse capability resolver:", err)
		return
	}

	terraModelParam := terraModel.ModelParam
	terraModelParam.Stream = false
	terraModelParam.MaxPromptLength = min(terraModelParam.MaxPromptLength, 4096)
	terraModelParam.MaxOutputLength = min(terraModelParam.MaxOutputLength, 1024)
	terraModelParam.SystemPrompt = "You are a concise assistant."
	terraModelParam.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}

	museModelParam := museModel.ModelParam
	museModelParam.Stream = false
	museModelParam.MaxPromptLength = min(museModelParam.MaxPromptLength, 4096)
	museModelParam.MaxOutputLength = min(museModelParam.MaxOutputLength, 1024)
	museModelParam.SystemPrompt = "You are a concise assistant."
	museModelParam.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}

	firstInputs := []spec.InputUnion{
		newUserTextInput(
			"Compute (17 * 23) - 41. Think briefly, then give the result in one short sentence.",
		),
	}

	terraResp, err := ps.FetchCompletion(ctx, terraProvider.Name, &spec.FetchCompletionRequest{
		ModelParam: terraModelParam,
		Inputs:     firstInputs,
	}, terraOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "OpenAI Terra FetchCompletion error:", err)
		if terraResp != nil && terraResp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", terraResp.Error.Message)
		}
		return
	}

	// Keep the original user input, then append all normalized Terra outputs.
	museInputs := append([]spec.InputUnion(nil), firstInputs...)
	museInputs = append(museInputs, outputUnionsToInputs(terraResp.Outputs)...)
	museInputs = append(
		museInputs,
		newUserTextInput(
			"Using the previous conversation, verify the calculation and give the result in one short sentence.",
		),
	)

	museResp, err := ps.FetchCompletion(ctx, museProvider.Name, &spec.FetchCompletionRequest{
		ModelParam: museModelParam,
		Inputs:     museInputs,
	}, museOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Meta Muse FetchCompletion error:", err)
		if museResp != nil && museResp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", museResp.Error.Message)
		}
		return
	}

	finalTerraInputs := append([]spec.InputUnion(nil), museInputs...)
	finalTerraInputs = append(
		finalTerraInputs,
		outputUnionsToInputs(museResp.Outputs)...,
	)
	finalTerraInputs = append(
		finalTerraInputs,
		newUserTextInput(
			"Using the full previous conversation, give the final verified result in one short sentence.",
		),
	)

	finalTerraResp, err := ps.FetchCompletion(ctx, terraProvider.Name, &spec.FetchCompletionRequest{
		ModelParam: terraModelParam,
		Inputs:     finalTerraInputs,
	}, terraOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "OpenAI Terra final FetchCompletion error:", err)
		if finalTerraResp != nil && finalTerraResp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", finalTerraResp.Error.Message)
		}
		return
	}

	fmt.Fprintln(os.Stderr, "OpenAI Terra first assistant:", responseText(terraResp))
	fmt.Fprintln(os.Stderr, "Meta Muse assistant:", responseText(museResp))
	fmt.Fprintln(
		os.Stderr,
		"OpenAI Terra final assistant:",
		responseText(finalTerraResp),
	)
	fmt.Println("OK")
	// Output: OK
}
