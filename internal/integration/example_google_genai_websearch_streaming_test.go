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
	googleWebSearchToolID          = "google-web-search"
	googleWebSearchToolDescription = "Search the web for recent information."
)

// Example_googleGenerateContent_webSearchAndThinkingStreaming demonstrates:
//
//   - catalog-based Google Gemini provider setup
//   - preset capability resolver
//   - server-side Google web search grounding
//   - streaming text + thinking
//   - normalized webSearch outputs synthesized from grounding metadata
//
// For Gemini GenerateContent, web search is server-side grounding, not a
// client-side tool-output round trip like function tools.
func Example_googleGenerateContent_webSearchAndThinkingStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	pp, mp, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderGoogleGemini,
		modelpreset.PresetGemini35Flash,
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
		fmt.Fprintln(os.Stderr, "GEMINI_API_KEY/GOOGLE_API_KEY not set; skipping live Google Gemini web-search example")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google Gemini API key:", err)
		return
	}

	modelParam := mp.ModelParam
	modelParam.Stream = true
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 512)
	modelParam.SystemPrompt = "Use web search when helpful. Keep the final answer short. " +
		"If you are unsure, say so plainly."
	modelParam.Reasoning = &spec.ReasoningParam{
		Type:   spec.ReasoningTypeHybridWithTokens,
		Tokens: 1024,
	}

	webSearchTool := spec.ToolChoice{
		Type:               spec.ToolTypeWebSearch,
		ID:                 googleWebSearchToolID,
		Name:               spec.DefaultWebSearchToolName,
		Description:        googleWebSearchToolDescription,
		WebSearchArguments: &spec.WebSearchToolChoiceItem{
			// Intentionally minimal. The Gemini adapter exposes web search
			// primarily as an enable/disable grounding capability.
		},
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}
	opts.StreamHandler = func(ev spec.StreamEvent) error {
		switch ev.Kind {
		case spec.StreamContentKindThinking:
			if ev.Thinking != nil {
				fmt.Fprintf(os.Stderr, "\n[thinking] %s\n", ev.Thinking.Text)
			}
		case spec.StreamContentKindText:
			if ev.Text != nil {
				fmt.Fprintf(os.Stderr, "\n[text] %s\n", ev.Text.Text)
			}
		}
		return nil
	}

	resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{
			newUserTextInput(
				"What is the latest stable Go release? If unknown, say unknown. Then list notable features and why they matter.",
			),
		},
		ToolChoices: []spec.ToolChoice{webSearchTool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeAuto,
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
		case spec.OutputKindWebSearchToolCall:
			if out.WebSearchToolCall != nil {
				fmt.Fprintf(os.Stderr, "Web search call: %+v\n", out.WebSearchToolCall.WebSearchToolCallItems)
			}

		case spec.OutputKindWebSearchToolOutput:
			if out.WebSearchToolOutput != nil {
				for _, item := range out.WebSearchToolOutput.WebSearchToolOutputItems {
					if item.Kind == spec.WebSearchToolOutputKindSearch && item.SearchItem != nil {
						fmt.Fprintf(os.Stderr, "Search result: %s (%s)\n",
							item.SearchItem.Title,
							item.SearchItem.URL,
						)
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
