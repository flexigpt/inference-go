package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

const sendFile = true

// Example_openAIResponses_toolsAndAttachments demonstrates a more advanced
// Responses call that:
//
//   - defines function and web-search tools,
//   - sends text + image + file as input content,
//   - enables streaming of both text and reasoning.
//
// The example only attempts a live call when OPENAI_API_KEY is set. The
// image/file payloads are placeholders; for a real run, replace them with
// valid data or URLs.
func Example_openAIResponses_toolsAndAttachments() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, "openai-responses-extended", &inference.AddProviderConfig{
		SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
		Origin:                   spec.DefaultOpenAIOrigin,
		ChatCompletionPathPrefix: "/v1/responses",
		APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding OpenAI Responses provider:", err)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENAI_API_KEY not set; skipping extended OpenAI Responses example")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, "openai-responses-extended", apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	// Tool: summarize_document(document: string, focus: string).
	summarizeTool := spec.ToolChoice{
		Type:        spec.ToolTypeFunction,
		ID:          "summarize-document",
		Name:        "summarize_document",
		Description: "Summarize a document with an optional focus.",
		Arguments: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"document": map[string]any{
					"type":        "string",
					"description": "Full text of the document to summarize.",
				},
				"focus": map[string]any{
					"type":        "string",
					"description": "Optional topic to focus on.",
				},
			},
			"required": []any{"document"},
		},
	}

	// Web search tool: used for retrieving fresh information when needed.
	webSearchTool := spec.ToolChoice{
		Type:        spec.ToolTypeWebSearch,
		ID:          "web-search",
		Name:        "web_search",
		Description: "Search the web for recent information.",
		WebSearchArguments: &spec.WebSearchToolChoiceItem{
			MaxUses:           2,
			SearchContextSize: "medium",
			AllowedDomains:    []string{}, // any domain
			UserLocation: &spec.WebSearchToolChoiceItemUserLocation{
				City:     "San Francisco",
				Country:  "US",
				Region:   "CA",
				Timezone: "America/Los_Angeles",
			},
		},
	}

	toolChoices := []spec.ToolChoice{summarizeTool, webSearchTool}

	// Placeholder image data (not a real image). In a real application, provide a valid base64-encoded image.
	// 1x1 transparent PNG.
	fakeImageData := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPhfDwAChwGA60e6kgAAAABJRU5ErkJggg=="

	userMessage := spec.InputOutputContent{
		Role: spec.RoleUser,
		Contents: []spec.InputOutputContentItemUnion{
			{
				Kind: spec.ContentItemKindText,
				TextItem: &spec.ContentItemText{
					Text: "This is a test. Very briefly reply with attached PDF name and data and describe the image. Use tools where appropriate. Keep the final answer short.",
				},
			},
			{
				Kind: spec.ContentItemKindImage,
				ImageItem: &spec.ContentItemImage{
					ImageMIME: spec.DefaultImageDataMIME,
					ImageData: fakeImageData,
					ImageName: "example-image",
					Detail:    spec.ImageDetailLow,
					ImageURL:  "",
					ID:        "abc",
				},
			},
		},
	}
	if sendFile {
		// Placeholder PDF URL.
		fileURL := "https://www.w3schools.com/asp/text/textfile.txt"
		userMessage.Contents = append(userMessage.Contents, spec.InputOutputContentItemUnion{
			Kind: spec.ContentItemKindFile,
			FileItem: &spec.ContentItemFile{
				FileName: "dummy",
				FileMIME: "text/plain",
				FileURL:  fileURL,
			},
		})
	}
	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "gpt-5-mini",
			Stream:          true,
			MaxPromptLength: 8192,
			MaxOutputLength: 8192,
			SystemPrompt: "You are a research assistant that first uses tools " +
				"when needed, then answers succinctly.",
			Reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelMedium,
			},
		},
		Inputs: []spec.InputUnion{
			{
				Kind:         spec.InputKindInputMessage,
				InputMessage: &userMessage,
			},
		},
		ToolChoices: toolChoices,
	}

	// Stream both text and reasoning to stdout.
	opts := &spec.FetchCompletionOptions{
		StreamHandler: func(ev spec.StreamEvent) error {
			switch ev.Kind {
			case spec.StreamContentKindText:
				if ev.Text != nil {
					fmt.Fprintln(os.Stderr, ev.Text.Text)
				}
			case spec.StreamContentKindThinking:
				if ev.Thinking != nil {
					// In a real app you might log this separately; here we
					// just prefix it.
					fmt.Fprintf(os.Stderr, "\n[thinking] %s \n", ev.Thinking.Text)
				}
			}
			return nil
		},
		StreamConfig: &spec.StreamConfig{
			// Use library defaults; override here if you want.
		},
	}

	resp, err := ps.FetchCompletion(ctx, "openai-responses-extended", req, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nFetchCompletion error:", err)
		if resp != nil && resp.Error != nil {
			fmt.Fprintln(os.Stderr, "Provider error:", resp.Error.Message)
		}
		return
	}

	fmt.Fprintln(os.Stderr, "\n\n--- normalized outputs ---")
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
