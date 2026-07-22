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
	googleToolsRoundTripToolID   = "echo-tool"
	googleToolsRoundTripToolName = "echo_text"
)

// Example_googleGenerateContent_functionToolRoundTrip demonstrates a full
// Gemini function-tool round trip:
//
//  1. user message + tool definition
//  2. model emits a function tool call
//  3. caller executes the tool locally
//  4. next request sends:
//     - original user turn
//     - assistant function tool call
//     - user function tool output
//     - extra user text
//  5. model returns the final answer
func Example_googleGenerateContent_functionToolRoundTrip() {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
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
		modelpreset.PresetGemini35FlashLite,
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
		fmt.Fprintln(os.Stderr, "GEMINI_API_KEY/GOOGLE_API_KEY not set; skipping live Google Gemini tool example")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google Gemini API key:", err)
		return
	}

	tool := newEchoToolChoice(googleToolsRoundTripToolID, googleToolsRoundTripToolName)

	initialUser := newUserTextInput(
		fmt.Sprintf(
			`Think briefly, then call the %s tool with text %q. Do not answer until after the tool result.`,
			googleToolsRoundTripToolName,
			"google function tool round trip",
		),
	)

	firstParam := mp.ModelParam
	firstParam.Stream = true
	firstParam.MaxOutputLength = min(firstParam.MaxOutputLength, 4096)
	firstParam.Reasoning = &spec.ReasoningParam{
		Type:   spec.ReasoningTypeHybridWithTokens,
		Tokens: 2048,
	}
	firstParam.SystemPrompt = "You are validating a Gemini function tool round trip. " +
		"When the tool is forced, emit only the tool call in the first response."

	firstOpts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}
	firstOpts.StreamHandler = func(ev spec.StreamEvent) error {
		switch ev.Kind {
		case spec.StreamContentKindThinking:
			if ev.Thinking != nil {
				fmt.Fprintf(os.Stderr, "\n[thinking 1] %s\n", ev.Thinking.Text)
			}
		case spec.StreamContentKindText:
			if ev.Text != nil {
				fmt.Fprintf(os.Stderr, "\n[text 1] %s\n", ev.Text.Text)
			}
		}
		return nil
	}

	firstResp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam:  firstParam,
		Inputs:      []spec.InputUnion{initialUser},
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeTool,
			AllowedTools: []spec.AllowedTool{
				{ToolChoiceID: tool.ID},
			},
			DisableParallel: true,
		},
	}, firstOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "first FetchCompletion error:", err)
		return
	}

	if firstText := responseText(firstResp); firstText != "" {
		fmt.Fprintln(os.Stderr, "first assistant text:", firstText)
	}

	call, err := firstFunctionToolCall(firstResp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "expected a function tool call, got error:", err)
		return
	}
	fmt.Fprintf(os.Stderr, "tool call 1: name=%s id=%s args=%s\n",
		call.Name,
		nonEmpty(call.CallID, call.ID),
		call.Arguments,
	)

	toolOutput, err := runEchoTool(call)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error executing local tool:", err)
		return
	}
	fmt.Fprintf(os.Stderr, "tool result 1 for %s: %s\n", toolOutput.CallID, firstToolOutputText(toolOutput))

	history := make([]spec.InputUnion, 0, len(firstResp.Outputs)+2)
	history = append(history, initialUser)
	history = append(history, outputUnionsToInputs(firstResp.Outputs)...)
	history = append(history, spec.InputUnion{
		Kind:               spec.InputKindFunctionToolOutput,
		FunctionToolOutput: toolOutput,
	})

	secondParam := mp.ModelParam
	secondParam.Stream = true
	secondParam.MaxOutputLength = min(secondParam.MaxOutputLength, 4096)
	secondParam.Reasoning = &spec.ReasoningParam{
		Type:   spec.ReasoningTypeHybridWithTokens,
		Tokens: 2048,
	}
	secondParam.SystemPrompt = "You have now received the tool result. " +
		"Answer briefly and do not call any tool again."

	secondOpts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}
	secondOpts.StreamHandler = func(ev spec.StreamEvent) error {
		switch ev.Kind {
		case spec.StreamContentKindThinking:
			if ev.Thinking != nil {
				fmt.Fprintf(os.Stderr, "\n[thinking 2] %s\n", ev.Thinking.Text)
			}
		case spec.StreamContentKindText:
			if ev.Text != nil {
				fmt.Fprintf(os.Stderr, "\n[text 2] %s\n", ev.Text.Text)
			}
		}
		return nil
	}

	secondResp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam:  secondParam,
		Inputs:      append(history, newUserTextInput("Now finish in one short sentence.")),
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeNone,
		},
	}, secondOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "second FetchCompletion error:", err)
		return
	}

	if finalText := responseText(secondResp); finalText != "" {
		fmt.Fprintln(os.Stderr, "final assistant text:", finalText)
	}

	fmt.Println("OK")
	// Output: OK
}
