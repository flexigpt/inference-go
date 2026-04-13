package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
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
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, "google-tools", &inference.AddProviderConfig{
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
		fmt.Fprintln(os.Stderr, "GEMINI_API_KEY/GOOGLE_API_KEY not set; skipping live Google GenAI tool example")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, "google-tools", apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google GenAI API key:", err)
		return
	}

	tool := spec.ToolChoice{
		Type:        spec.ToolTypeFunction,
		ID:          "echo-tool",
		Name:        "echo_text",
		Description: "Echo the provided text back in a deterministic tool result.",
		Arguments: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"text": map[string]any{
					"type": "string",
				},
			},
			"required":             []any{"text"},
			"additionalProperties": false,
		},
	}

	initialUser := newUserTextInput(
		`Use the echo_text tool with text "google tool round trip". Do not answer yet; only call the tool.`,
	)

	firstReq := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "gemini-3-flash-preview",
			MaxOutputLength: 512,
			SystemPrompt: "You are validating a Gemini function tool round trip. " +
				"When the tool is forced, emit only the tool call in the first response.",
		},
		Inputs:      []spec.InputUnion{initialUser},
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeTool,
			AllowedTools: []spec.AllowedTool{
				{ToolChoiceID: tool.ID},
			},
			DisableParallel: true,
		},
	}

	firstResp, err := ps.FetchCompletion(ctx, "google-tools", firstReq, &spec.FetchCompletionOptions{
		CompletionKey: "gemini-3-flash-preview",
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "first FetchCompletion error:", err)
		return
	}

	call, err := firstFunctionToolCall(firstResp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "expected a function tool call, got error:", err)
		return
	}
	fmt.Fprintf(
		os.Stderr,
		"tool call: name=%s id=%s args=%s\n",
		call.Name,
		nonEmpty(call.CallID, call.ID),
		call.Arguments,
	)

	toolOutput, err := runEchoTool(call)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error executing local tool:", err)
		return
	}
	fmt.Fprintf(os.Stderr, "tool result for %s: %s\n", toolOutput.CallID, firstToolOutputText(toolOutput))

	secondReq := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "gemini-2.5-flash",
			MaxOutputLength: 256,
			SystemPrompt: "You have now received the tool result. " +
				"Answer briefly in plain text. Do not call any tool again.",
		},
		Inputs: []spec.InputUnion{
			initialUser,
			{
				Kind:             spec.InputKindFunctionToolCall,
				FunctionToolCall: call,
			},
			{
				Kind:               spec.InputKindFunctionToolOutput,
				FunctionToolOutput: toolOutput,
			},
			newUserTextInput("Now finish in one short sentence."),
		},
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeNone,
		},
	}

	secondResp, err := ps.FetchCompletion(ctx, "google-tools", secondReq, &spec.FetchCompletionOptions{
		CompletionKey: "gemini-2.5-flash",
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "second FetchCompletion error:", err)
		return
	}

	finalText := responseText(secondResp)
	if finalText != "" {
		fmt.Fprintln(os.Stderr, "final assistant text:", finalText)
	}

	fmt.Println("OK")
	// Output: OK
}
