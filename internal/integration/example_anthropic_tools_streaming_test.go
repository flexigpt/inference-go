package integration

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

const (
	anthropicExtractKeyPointsToolID          = "extract-key-points"
	anthropicExtractKeyPointsToolName        = "extract_key_points"
	anthropicExtractKeyPointsToolDescription = "Extract 3 key points from the provided text."

	anthropicWebSearchToolID = "web-search"

	anthropicEchoToolID   = "echo-tool"
	anthropicEchoToolName = "echo_text"
)

// Example_anthropic_toolsAndThinkingStreaming demonstrates:
//
//   - catalog-based Anthropic provider setup
//   - preset model defaults
//   - preset capability resolver
//   - streaming text + thinking
//   - function tools + Anthropic server web search
//   - JSON schema output request
func Example_anthropic_toolsAndThinkingStreaming() {
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
		modelpreset.ProviderAnthropic,
		modelpreset.PresetClaudeSonnet46,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Anthropic preset provider:", err)
		return
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "ANTHROPIC_API_KEY not set; skipping live Anthropic call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Anthropic API key:", err)
		return
	}

	modelParam := mp.ModelParam
	modelParam.Stream = true
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 1024)
	modelParam.SystemPrompt = "Use tools when helpful. Keep the final answer short."
	modelParam.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelMedium,
	}
	modelParam.OutputParam = &spec.OutputParam{
		Format: &spec.OutputFormat{
			Kind: spec.OutputFormatKindJSONSchema,
			JSONSchemaParam: &spec.JSONSchemaParam{
				Name: toolJSONSchemaName,
				Schema: map[string]any{
					toolJSONKeyType: toolJSONValueObject,
					toolJSONKeyProperties: map[string]any{
						"summary": map[string]any{
							toolJSONKeyType: toolJSONValueString,
						},
						"source_used": map[string]any{
							toolJSONKeyType: toolJSONValueBoolean,
						},
					},
					toolJSONKeyRequired:             []any{"summary", "source_used"},
					toolJSONKeyAdditionalProperties: false,
				},
			},
		},
	}

	tools := []spec.ToolChoice{
		{
			Type:        spec.ToolTypeFunction,
			ID:          anthropicExtractKeyPointsToolID,
			Name:        anthropicExtractKeyPointsToolName,
			Description: anthropicExtractKeyPointsToolDescription,
			Arguments: map[string]any{
				toolJSONKeyType: toolJSONValueObject,
				toolJSONKeyProperties: map[string]any{
					toolJSONKeyText: map[string]any{toolJSONKeyType: toolJSONValueString},
				},
				toolJSONKeyRequired:             []any{toolJSONKeyText},
				toolJSONKeyAdditionalProperties: false,
			},
		},
		{
			Type: spec.ToolTypeWebSearch,
			ID:   anthropicWebSearchToolID,
			Name: spec.DefaultWebSearchToolName,
			WebSearchArguments: &spec.WebSearchToolChoiceItem{
				MaxUses:           1,
				SearchContextSize: spec.WebSearchContextSizeMedium,
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
		case spec.StreamContentKindThinking:
			if ev.Thinking != nil {
				fmt.Fprintf(os.Stderr, "[thinking] %s\n", ev.Thinking.Text)
			}
		case spec.StreamContentKindText:
			if ev.Text != nil {
				fmt.Fprint(os.Stderr, ev.Text.Text)
			}
		}
		return nil
	}

	_, err = ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{
			newUserTextInput("What is the latest stable Go version? If unknown, say so."),
		},
		ToolChoices: tools,
		ToolPolicy: &spec.ToolPolicy{
			Mode:            spec.ToolPolicyModeAuto,
			DisableParallel: true,
		},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nFetchCompletion error:", err)
		return
	}

	fmt.Println("OK")
	// Output: OK
}

// Example_anthropic_functionToolRoundTrip demonstrates the full Anthropic
// client-tool flow that requires strict ordering:
//
//  1. user message + tool choice
//  2. assistant emits a function tool call
//  3. caller executes the tool locally
//  4. next request sends:
//     - original user turn
//     - assistant tool call
//     - user tool_result
//     - extra user text
//  5. assistant returns the final answer
//
// For Anthropic, the tool_result must immediately follow the assistant tool-use
// turn as the next user turn. The Anthropic adapter normalizes this ordering.
func Example_anthropic_functionToolRoundTrip() {
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
		modelpreset.ProviderAnthropic,
		modelpreset.PresetClaudeSonnet46,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Anthropic preset provider:", err)
		return
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "ANTHROPIC_API_KEY not set; skipping live Anthropic call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Anthropic API key:", err)
		return
	}

	tool := newEchoToolChoice(anthropicEchoToolID, anthropicEchoToolName)

	initialUser := newUserTextInput(
		fmt.Sprintf(
			`Use the %s tool with text %q. Do not answer yet; just call the tool.`,
			anthropicEchoToolName,
			"anthropic tool round trip",
		),
	)

	firstParam := mp.ModelParam
	firstParam.MaxOutputLength = min(firstParam.MaxOutputLength, 512)
	firstParam.Reasoning = nil
	firstParam.SystemPrompt = strings.Join([]string{
		"You are validating a client tool round trip.",
		"When the tool is forced, emit only the tool call in the first response.",
		"Do not provide the final answer until after the tool result is returned.",
	}, " ")

	firstOpts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
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

	call, err := firstFunctionToolCall(firstResp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "expected a function tool call, got error:", err)
		return
	}
	fmt.Fprintf(os.Stderr, "tool call: name=%s id=%s args=%s\n",
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

	secondParam := mp.ModelParam
	secondParam.MaxOutputLength = min(secondParam.MaxOutputLength, 2048)
	secondParam.SystemPrompt = strings.Join([]string{
		"You have now received the tool result.",
		"Answer briefly in plain text.",
		"Do not call any tool again.",
	}, " ")

	secondOpts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}

	secondResp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: secondParam,
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
