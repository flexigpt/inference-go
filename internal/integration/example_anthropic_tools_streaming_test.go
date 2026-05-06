package integration

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

const (
	anthropicToolsProviderName                    = "anthropic"
	anthropicToolsExtractKeyPointsToolID          = "extract-key-points"
	anthropicToolsExtractKeyPointsToolName        = "extract_key_points"
	anthropicToolsExtractKeyPointsToolDescription = "Extract 3 key points from the provided text."
	anthropicToolsJSONKeyType                     = "type"
	anthropicToolsJSONValueObject                 = "object"
	anthropicToolsJSONKeyProperties               = "properties"
	anthropicToolsJSONKeyText                     = "text"
	anthropicToolsJSONValueString                 = "string"
	anthropicToolsJSONKeyRequired                 = "required"
	anthropicToolsJSONKeyAdditionalProperties     = "additionalProperties"
	anthropicToolsModelName                       = "claude-sonnet-4-6"
	anthropicToolsJSONSchemaName                  = "answer"
	anthropicToolsJSONKeySummary                  = "summary"
	anthropicToolsJSONKeySourceUsed               = "source_used"
	anthropicToolsJSONValueBoolean                = "boolean"
	anthropicToolsCompletionKey                   = "sonnet46"
	anthropicToolsWebSearchToolID                 = "web-search"
	anthropicToolsEchoToolID                      = "echo-tool"
	anthropicToolsEchoToolName                    = "echo_text"
	anthropicToolsEchoToolDescription             = "Echo the provided text back in a deterministic tool result."
)

// Example_anthropic_toolsAndThinkingStreaming demonstrates:
//   - streaming text + thinking
//   - function tools + anthropic server web search
//   - JSON schema output request
func Example_anthropic_toolsAndThinkingStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, anthropicToolsProviderName, &inference.AddProviderConfig{
		SDKType:                  spec.ProviderSDKTypeAnthropic,
		Origin:                   spec.DefaultAnthropicOrigin,
		ChatCompletionPathPrefix: spec.DefaultAnthropicChatCompletionPrefix,
		APIKeyHeaderKey:          spec.DefaultAnthropicAuthorizationHeaderKey,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Anthropic provider:", err)
		return
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "ANTHROPIC_API_KEY not set; skipping live Anthropic call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, anthropicToolsProviderName, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Anthropic API key:", err)
		return
	}

	tools := []spec.ToolChoice{
		{
			Type:        spec.ToolTypeFunction,
			ID:          anthropicToolsExtractKeyPointsToolID,
			Name:        anthropicToolsExtractKeyPointsToolName,
			Description: anthropicToolsExtractKeyPointsToolDescription,
			Arguments: map[string]any{
				anthropicToolsJSONKeyType: anthropicToolsJSONValueObject,
				anthropicToolsJSONKeyProperties: map[string]any{
					anthropicToolsJSONKeyText: map[string]any{anthropicToolsJSONKeyType: anthropicToolsJSONValueString},
				},
				anthropicToolsJSONKeyRequired:             []any{anthropicToolsJSONKeyText},
				anthropicToolsJSONKeyAdditionalProperties: false,
			},
		},
		{
			Type: spec.ToolTypeWebSearch,
			ID:   anthropicToolsWebSearchToolID,
			Name: spec.DefaultWebSearchToolName,
			WebSearchArguments: &spec.WebSearchToolChoiceItem{
				MaxUses:           1,
				SearchContextSize: spec.WebSearchContextSizeMedium,
			},
		},
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:         anthropicToolsModelName,
			Stream:       true,
			SystemPrompt: "Use tools when helpful. Keep the final answer short.",
			Reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelMedium,
			},
			OutputParam: &spec.OutputParam{
				Format: &spec.OutputFormat{
					Kind: spec.OutputFormatKindJSONSchema,
					JSONSchemaParam: &spec.JSONSchemaParam{
						Name: anthropicToolsJSONSchemaName,
						Schema: map[string]any{
							anthropicToolsJSONKeyType: anthropicToolsJSONValueObject,
							anthropicToolsJSONKeyProperties: map[string]any{
								anthropicToolsJSONKeySummary: map[string]any{
									anthropicToolsJSONKeyType: anthropicToolsJSONValueString,
								},
								anthropicToolsJSONKeySourceUsed: map[string]any{
									anthropicToolsJSONKeyType: anthropicToolsJSONValueBoolean,
								},
							},
							anthropicToolsJSONKeyRequired: []any{
								anthropicToolsJSONKeySummary,
								anthropicToolsJSONKeySourceUsed,
							},
							anthropicToolsJSONKeyAdditionalProperties: false,
						},
					},
				},
			},
		},
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "What is the latest stable Go version? If unknown, say so."},
				}},
			},
		}},
		ToolChoices: tools,
		ToolPolicy: &spec.ToolPolicy{
			Mode:            spec.ToolPolicyModeAuto,
			DisableParallel: true,
		},
	}

	_, err = ps.FetchCompletion(ctx, anthropicToolsProviderName, req, &spec.FetchCompletionOptions{
		CompletionKey: anthropicToolsCompletionKey,
		StreamHandler: func(ev spec.StreamEvent) error {
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
		},
	})
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
// This is the flow where, for Anthropic, tool_result must immediately follow
// the assistant tool-use turn as the next user turn.
func Example_anthropic_functionToolRoundTrip() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, anthropicToolsProviderName, &inference.AddProviderConfig{
		SDKType:                  spec.ProviderSDKTypeAnthropic,
		Origin:                   spec.DefaultAnthropicOrigin,
		ChatCompletionPathPrefix: spec.DefaultAnthropicChatCompletionPrefix,
		APIKeyHeaderKey:          spec.DefaultAnthropicAuthorizationHeaderKey,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Anthropic provider:", err)
		return
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "ANTHROPIC_API_KEY not set; skipping live Anthropic call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, anthropicToolsProviderName, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Anthropic API key:", err)
		return
	}

	tool := spec.ToolChoice{
		Type:        spec.ToolTypeFunction,
		ID:          anthropicToolsEchoToolID,
		Name:        anthropicToolsEchoToolName,
		Description: anthropicToolsEchoToolDescription,
		Arguments: map[string]any{
			anthropicToolsJSONKeyType: anthropicToolsJSONValueObject,
			anthropicToolsJSONKeyProperties: map[string]any{
				anthropicToolsJSONKeyText: map[string]any{
					anthropicToolsJSONKeyType: anthropicToolsJSONValueString,
				},
			},
			anthropicToolsJSONKeyRequired:             []any{anthropicToolsJSONKeyText},
			anthropicToolsJSONKeyAdditionalProperties: false,
		},
	}

	initialUser := newUserTextInput(
		fmt.Sprintf(
			`Use the %s tool with text %q. Do not answer yet; just call the tool.`,
			anthropicToolsEchoToolName,
			"anthropic tool round trip",
		),
	)

	firstReq := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            anthropicToolsModelName,
			MaxOutputLength: 512,
			SystemPrompt: strings.Join([]string{
				"You are validating a client tool round trip.",
				"When the tool is forced, emit only the tool call in the first response.",
				"Do not provide the final answer until after the tool result is returned.",
			}, " "),
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

	firstResp, err := ps.FetchCompletion(ctx, anthropicToolsProviderName, firstReq, &spec.FetchCompletionOptions{
		CompletionKey: anthropicToolsCompletionKey,
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

	// This second request intentionally sends:
	//   - prior original user turn
	//   - assistant tool call
	//   - tool output
	//   - extra user text
	//
	// The Anthropic adapter should normalize the last two into the immediate next
	// user turn, with tool_result first and user text after it.
	secondReq := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            anthropicToolsModelName,
			MaxOutputLength: 256,
			SystemPrompt: strings.Join([]string{
				"You have now received the tool result.",
				"Answer briefly in plain text.",
				"Do not call any tool again.",
			}, " "),
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

	secondResp, err := ps.FetchCompletion(ctx, anthropicToolsProviderName, secondReq, &spec.FetchCompletionOptions{
		CompletionKey: anthropicToolsCompletionKey,
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
