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

const (
	metaMuseSpark13ContributorToolID   = "meta-echo-tool"
	metaMuseSpark13ContributorToolName = "meta_echo"
)

func Example_meta_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	providerSet, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	provider, model, err := addCatalogModelProvider(
		ctx,
		providerSet,
		modelpreset.ProviderMeta,
		modelpreset.PresetMuseSpark13Contributor,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Meta Muse Spark 1.3 Contributor:", err)
		return
	}

	apiKey := os.Getenv("META_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "META_API_KEY not set; skipping live Meta call")
		fmt.Println("OK")
		return
	}
	if err := providerSet.SetProviderAPIKey(ctx, provider.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Meta API key:", err)
		return
	}

	summaryStyle := spec.ReasoningSummaryStyleOmitted
	modelParam := model.ModelParam
	modelParam.Stream = false
	modelParam.MaxPromptLength = min(modelParam.MaxPromptLength, 4096)
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 2048)
	modelParam.SystemPrompt = "You are a concise, helpful assistant."
	modelParam.Reasoning = &spec.ReasoningParam{
		Type:         spec.ReasoningTypeSingleWithLevels,
		Level:        spec.ReasoningLevelLow,
		SummaryStyle: &summaryStyle,
	}
	modelParam.CacheControl = &spec.CacheControl{
		Kind: spec.CacheControlKindEphemeral,
		TTL:  spec.CacheControlTTL24h,
		Key:  "meta-muse-basic-example",
	}

	options, err := presetFetchOptions(ctx, providerSet, provider, model)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}

	response, err := providerSet.FetchCompletion(ctx, provider.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{
			newUserTextInput("Explain the difference between a goroutine and an OS thread in two sentences."),
		},
	}, options)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FetchCompletion error:", err)
		if response != nil && response.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", response.Error.Message)
		}
		return
	}

	fmt.Fprintln(os.Stderr, "Meta Muse assistant:", responseText(response))
	fmt.Println("OK")
	// Output: OK
}

// Example_meta_multiParameterStreaming demonstrates:
//
// - the Meta Muse Spark 1.3 Contributor preset
// - streamed text and reasoning output
// - detailed reasoning summaries
// - top-level prompt cache controls
// - text, image, and file input
// - JSON Schema output
// - a function tool with auto tool policy
//
// Meta's preset intentionally does not support reasoning context or reasoning
// mode. Do not add those fields without a verified Meta-specific capability
// override.
func Example_meta_multiParameterStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	providerSet, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	provider, model, err := addCatalogModelProvider(
		ctx,
		providerSet,
		modelpreset.ProviderMeta,
		modelpreset.PresetMuseSpark13Contributor,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Meta Muse Spark 1.3 Contributor:", err)
		return
	}

	apiKey := os.Getenv("META_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "META_API_KEY not set; skipping live Meta call")
		fmt.Println("OK")
		return
	}
	if err := providerSet.SetProviderAPIKey(ctx, provider.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Meta API key:", err)
		return
	}

	summaryStyle := spec.ReasoningSummaryStyleDetailed
	modelParam := model.ModelParam
	modelParam.Stream = true
	modelParam.MaxPromptLength = min(modelParam.MaxPromptLength, 8192)
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 2048)
	modelParam.SystemPrompt = "Use the supplied inputs and tools when helpful. Return a concise structured response."
	modelParam.Reasoning = &spec.ReasoningParam{
		Type:         spec.ReasoningTypeSingleWithLevels,
		Level:        spec.ReasoningLevelHigh,
		SummaryStyle: &summaryStyle,
	}
	modelParam.CacheControl = &spec.CacheControl{
		Kind: spec.CacheControlKindEphemeral,
		TTL:  spec.CacheControlTTL24h,
		Key:  "meta-muse-multi-parameter-example",
	}
	modelParam.OutputParam = &spec.OutputParam{
		Format: &spec.OutputFormat{
			Kind: spec.OutputFormatKindJSONSchema,
			JSONSchemaParam: &spec.JSONSchemaParam{
				Name: "meta_muse_result",
				Schema: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"answer":    map[string]any{"type": "string"},
						"used_tool": map[string]any{"type": "boolean"},
					},
					"required":             []any{"answer", "used_tool"},
					"additionalProperties": false,
				},
				Strict: true,
			},
		},
	}

	// A 1x1 transparent PNG.
	imageData := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPhfDwAChwGA60e6kgAAAABJRU5ErkJggg=="
	tool := newEchoToolChoice(
		metaMuseSpark13ContributorToolID,
		metaMuseSpark13ContributorToolName,
	)

	options, err := presetFetchOptions(ctx, providerSet, provider, model)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}
	options.StreamHandler = func(event spec.StreamEvent) error {
		switch event.Kind {
		case spec.StreamContentKindThinking:
			if event.Thinking != nil {
				fmt.Fprintf(os.Stderr, "[thinking] %s\n", event.Thinking.Text)
			}
		case spec.StreamContentKindText:
			if event.Text != nil {
				fmt.Fprint(os.Stderr, event.Text.Text)
			}
		}
		return nil
	}

	response, err := providerSet.FetchCompletion(ctx, provider.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{
					{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: "Describe the image and file, then answer briefly."},
					},
					{
						Kind: spec.ContentItemKindImage,
						ImageItem: &spec.ContentItemImage{
							ImageData: imageData,
							ImageMIME: spec.DefaultImageDataMIME,
							ImageName: "transparent-pixel.png",
							Detail:    spec.ImageDetailLow,
						},
					},
					{
						Kind: spec.ContentItemKindFile,
						FileItem: &spec.ContentItemFile{
							FileURL:  "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
							FileMIME: "application/pdf",
							FileName: "dummy.pdf",
						},
					},
				},
			},
		}},
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeAuto,
		},
	}, options)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nFetchCompletion error:", err)
		if response != nil && response.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", response.Error.Message)
		}
		return
	}

	fmt.Fprintln(os.Stderr, "\nMeta Muse assistant:", responseText(response))
	fmt.Println("OK")
	// Output: OK
}
