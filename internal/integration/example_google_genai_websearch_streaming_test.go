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

// Example_googleGenerateContent_webSearchAndThinkingStreaming demonstrates:
//   - server-side Google web search grounding
//   - streaming text + thinking
//   - normalized webSearch outputs synthesized from grounding metadata
//
// Note: for Gemini GenerateContent, web search is not a client-side tool-output
// round trip like function tools. It is server-side grounding.
func Example_googleGenerateContent_webSearchAndThinkingStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, "google-search", &inference.AddProviderConfig{
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
		fmt.Fprintln(os.Stderr, "GEMINI_API_KEY/GOOGLE_API_KEY not set; skipping live Google GenAI web-search example")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, "google-search", apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google GenAI API key:", err)
		return
	}

	webSearchTool := spec.ToolChoice{
		Type:               spec.ToolTypeWebSearch,
		ID:                 "google-web-search",
		Name:               spec.DefaultWebSearchToolName,
		Description:        "Search the web for recent information.",
		WebSearchArguments: &spec.WebSearchToolChoiceItem{
			// Intentionally minimal here. The Gemini adapter currently exposes
			// web search primarily as an enable/disable capability.
		},
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "gemini-3-flash-preview",
			Stream:          true,
			MaxOutputLength: 512,
			SystemPrompt: "Use web search when helpful. Keep the final answer short. " +
				"If you are unsure, say so plainly.",
			Reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelLow,
			},
		},
		Inputs: []spec.InputUnion{
			newUserTextInput(
				"What is the latest stable Go release? If unknown, say unknown. then list features. Then analyze and give its benefit over last release.",
			),
		},
		ToolChoices: []spec.ToolChoice{webSearchTool},
		ToolPolicy: &spec.ToolPolicy{
			// Auto is the right mode for Gemini web search grounding.
			Mode: spec.ToolPolicyModeAuto,
		},
	}

	resp, err := ps.FetchCompletion(ctx, "google-search", req, &spec.FetchCompletionOptions{
		CompletionKey: "gemini-search",
		StreamHandler: func(ev spec.StreamEvent) error {
			switch ev.Kind {
			case spec.StreamContentKindThinking:
				if ev.Thinking != nil {
					fmt.Fprintf(os.Stderr, "\n\n#######[thinking] %s\n", ev.Thinking.Text)
				}
			case spec.StreamContentKindText:
				if ev.Text != nil {
					fmt.Fprintf(os.Stderr, "\n\n#######[text] %s\n", ev.Text.Text)
				}
			}
			return nil
		},
	})
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
		case spec.OutputKindWebSearchToolCall:
			if out.WebSearchToolCall != nil {
				fmt.Fprintf(os.Stderr, "Web search call: %+v\n", out.WebSearchToolCall.WebSearchToolCallItems)
			}
		case spec.OutputKindWebSearchToolOutput:
			if out.WebSearchToolOutput != nil {
				for _, item := range out.WebSearchToolOutput.WebSearchToolOutputItems {
					if item.Kind == spec.WebSearchToolOutputKindSearch && item.SearchItem != nil {
						fmt.Fprintf(os.Stderr, "Search result: %s (%s)\n", item.SearchItem.Title, item.SearchItem.URL)
					}
				}
			}
		case spec.OutputKindReasoningMessage:
			if out.ReasoningMessage != nil && len(out.ReasoningMessage.Thinking) > 0 {
				fmt.Fprintln(os.Stderr, "Reasoning:", out.ReasoningMessage.Thinking[0])
			}
		case spec.OutputKindOutputMessage:
			if out.OutputMessage != nil {
				for _, c := range out.OutputMessage.Contents {
					if c.Kind == spec.ContentItemKindText && c.TextItem != nil {
						fmt.Fprintln(os.Stderr, "Final answer:", c.TextItem.Text)
					}
				}
			}
		default:
		}
	}

	fmt.Println("OK")
	// Output: OK
}
