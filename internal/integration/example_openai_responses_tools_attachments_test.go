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
	openAIResponsesSummarizeToolID          = "summarize-document"
	openAIResponsesSummarizeToolName        = "summarize_document"
	openAIResponsesSummarizeToolDescription = "Summarize a document with an optional focus."

	openAIResponsesWebSearchToolID          = "web-search"
	openAIResponsesWebSearchToolName        = "web_search"
	openAIResponsesWebSearchToolDescription = "Search the web for recent information."
)

const sendOpenAIResponsesExampleFile = true

// Example_openAIResponses_toolsAndAttachments demonstrates a more advanced
// Responses call:
//
//   - catalog-based OpenAI Responses provider setup
//   - preset model defaults
//   - preset capability resolver
//   - function and web-search tools
//   - text + image + file input content
//   - streaming text and reasoning
//
// The example only attempts a live call when OPENAI_API_KEY is set.
func Example_openAIResponses_toolsAndAttachments() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
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
		modelpreset.PresetOpenAIResponsesGPT5Mini,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding OpenAI Responses preset provider:", err)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENAI_API_KEY not set; skipping extended OpenAI Responses example")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	summarizeTool := spec.ToolChoice{
		Type:        spec.ToolTypeFunction,
		ID:          openAIResponsesSummarizeToolID,
		Name:        openAIResponsesSummarizeToolName,
		Description: openAIResponsesSummarizeToolDescription,
		Arguments: map[string]any{
			toolJSONKeyType: toolJSONValueObject,
			toolJSONKeyProperties: map[string]any{
				"document": map[string]any{
					toolJSONKeyType: toolJSONValueString,
					"description":   "Full text of the document to summarize.",
				},
				"focus": map[string]any{
					toolJSONKeyType: toolJSONValueString,
					"description":   "Optional topic to focus on.",
				},
			},
			toolJSONKeyRequired:             []any{"document"},
			toolJSONKeyAdditionalProperties: false,
		},
	}

	webSearchTool := spec.ToolChoice{
		Type:        spec.ToolTypeWebSearch,
		ID:          openAIResponsesWebSearchToolID,
		Name:        openAIResponsesWebSearchToolName,
		Description: openAIResponsesWebSearchToolDescription,
		WebSearchArguments: &spec.WebSearchToolChoiceItem{
			MaxUses:           2,
			SearchContextSize: spec.WebSearchContextSizeMedium,
			UserLocation: &spec.WebSearchToolChoiceItemUserLocation{
				City:     "San Francisco",
				Country:  "US",
				Region:   "CA",
				Timezone: "America/Los_Angeles",
			},
		},
	}

	// 1x1 transparent PNG.
	fakeImageData := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPhfDwAChwGA60e6kgAAAABJRU5ErkJggg=="

	userMessage := spec.InputOutputContent{
		Role: spec.RoleUser,
		Contents: []spec.InputOutputContentItemUnion{
			{
				Kind: spec.ContentItemKindText,
				TextItem: &spec.ContentItemText{
					Text: "Briefly describe the image and attached file. Use tools where appropriate. Keep the final answer short.",
				},
			},
			{
				Kind: spec.ContentItemKindImage,
				ImageItem: &spec.ContentItemImage{
					ImageMIME: spec.DefaultImageDataMIME,
					ImageData: fakeImageData,
					ImageName: "example-image",
					Detail:    spec.ImageDetailLow,
				},
			},
		},
	}

	if sendOpenAIResponsesExampleFile {
		userMessage.Contents = append(userMessage.Contents, spec.InputOutputContentItemUnion{
			Kind: spec.ContentItemKindFile,
			FileItem: &spec.ContentItemFile{
				FileName: "example.txt",
				FileMIME: "text/plain",
				FileURL:  "https://www.w3schools.com/asp/text/textfile.txt",
			},
		})
	}

	modelParam := mp.ModelParam
	modelParam.Stream = true
	modelParam.MaxPromptLength = min(modelParam.MaxPromptLength, 8192)
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 8192)
	modelParam.SystemPrompt = "You are a research assistant that first uses tools when needed, then answers succinctly."
	modelParam.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelMedium,
	}
	modelParam.OutputParam = &spec.OutputParam{
		Verbosity: new(spec.OutputVerbosityMedium),
		Format: &spec.OutputFormat{
			Kind: spec.OutputFormatKindJSONSchema,
			JSONSchemaParam: &spec.JSONSchemaParam{
				Name: "final_answer",
				Schema: map[string]any{
					toolJSONKeyType: toolJSONValueObject,
					toolJSONKeyProperties: map[string]any{
						"image_description": map[string]any{toolJSONKeyType: toolJSONValueString},
						"file_name":         map[string]any{toolJSONKeyType: toolJSONValueString},
						"answer":            map[string]any{toolJSONKeyType: toolJSONValueString},
					},
					toolJSONKeyRequired: []any{
						"image_description",
						"file_name",
						"answer",
					},
					toolJSONKeyAdditionalProperties: false,
				},
				Strict: true,
			},
		},
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}
	opts.StreamHandler = func(ev spec.StreamEvent) error {
		switch ev.Kind {
		case spec.StreamContentKindText:
			if ev.Text != nil {
				fmt.Fprintln(os.Stderr, ev.Text.Text)
			}
		case spec.StreamContentKindThinking:
			if ev.Thinking != nil {
				fmt.Fprintf(os.Stderr, "\n[thinking] %s\n", ev.Thinking.Text)
			}
		}
		return nil
	}
	opts.StreamConfig = &spec.StreamConfig{}

	resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{
			{
				Kind:         spec.InputKindInputMessage,
				InputMessage: &userMessage,
			},
		},
		ToolChoices: []spec.ToolChoice{summarizeTool, webSearchTool},
		ToolPolicy: &spec.ToolPolicy{
			Mode:            spec.ToolPolicyModeAuto,
			DisableParallel: true,
		},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nFetchCompletion error:", err)
		if resp != nil && resp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", resp.Error.Message)
		}
		return
	}

	fmt.Fprintln(os.Stderr, "\n--- normalized outputs ---")
	for _, out := range resp.Outputs {
		switch out.Kind {
		case spec.OutputKindFunctionToolCall:
			if out.FunctionToolCall != nil {
				fmt.Fprintf(os.Stderr, "Function tool call: %s(%s)\n",
					out.FunctionToolCall.Name,
					out.FunctionToolCall.Arguments,
				)
			}

		case spec.OutputKindWebSearchToolCall:
			if out.WebSearchToolCall != nil {
				fmt.Fprintf(os.Stderr, "Web search call: %+v\n", out.WebSearchToolCall.WebSearchToolCallItems)
			}

		case spec.OutputKindOutputMessage:
			if out.OutputMessage != nil {
				for _, c := range out.OutputMessage.Contents {
					if c.Kind == spec.ContentItemKindText && c.TextItem != nil {
						fmt.Fprintln(os.Stderr, "Final answer:", c.TextItem.Text)
					}
				}
			}

		case spec.OutputKindReasoningMessage:
			if out.ReasoningMessage != nil && len(out.ReasoningMessage.Summary) > 0 {
				fmt.Fprintln(os.Stderr, "Reasoning summary:", out.ReasoningMessage.Summary[0])
			}
		default:
		}
	}

	fmt.Println("OK")
	// Output: OK
}
