package sdkutil

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestSanitizeConversationStartTrimsLeadingPartialHistory(t *testing.T) {
	sanitized, leadingDropped, firstUserToolOutputsDropped, err := sanitizeConversationStart(
		[]spec.InputUnion{
			conversationStartFunctionalAssistantMessage("old answer"),
			conversationStartTestFunctionToolCall("prefix-call"),
			conversationStartTestFunctionToolOutput("prefix-call"),
			conversationStartTestUserMessage("new question"),
			conversationStartFunctionalAssistantMessage("new answer"),
		},
	)
	if err != nil {
		t.Fatalf("sanitizeConversationStart() error = %v", err)
	}
	if got, want := leadingDropped, 3; got != want {
		t.Fatalf("leadingDropped = %d, want %d", got, want)
	}
	if got, want := firstUserToolOutputsDropped, 0; got != want {
		t.Fatalf("firstUserToolOutputsDropped = %d, want %d", got, want)
	}
	if len(sanitized) != 2 ||
		sanitized[0].Kind != spec.InputKindInputMessage ||
		sanitized[1].Kind != spec.InputKindOutputMessage {
		t.Fatalf("sanitized inputs = %#v, want user message followed by assistant message", sanitized)
	}
}

func TestSanitizeConversationStartDropsAllToolOutputKindsFromFirstUserTurn(t *testing.T) {
	tests := []struct {
		name       string
		toolOutput spec.InputUnion
	}{
		{
			name:       "function tool output",
			toolOutput: conversationStartTestFunctionToolOutput("function-call"),
		},
		{
			name:       "custom tool output",
			toolOutput: conversationStartFunctionalCustomToolOutput("custom-call"),
		},
		{
			name:       "web search tool output",
			toolOutput: conversationStartFunctionalWebSearchToolOutput("web-search-call"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sanitized, leadingDropped, firstUserToolOutputsDropped, err := sanitizeConversationStart(
				[]spec.InputUnion{
					conversationStartTestUserMessage("question"),
					tc.toolOutput,
					conversationStartFunctionalAssistantMessage("answer"),
				},
			)
			if err != nil {
				t.Fatalf("sanitizeConversationStart() error = %v", err)
			}
			if got, want := leadingDropped, 0; got != want {
				t.Fatalf("leadingDropped = %d, want %d", got, want)
			}
			if got, want := firstUserToolOutputsDropped, 1; got != want {
				t.Fatalf("firstUserToolOutputsDropped = %d, want %d", got, want)
			}
			if len(sanitized) != 2 ||
				sanitized[0].Kind != spec.InputKindInputMessage ||
				sanitized[1].Kind != spec.InputKindOutputMessage {
				t.Fatalf(
					"sanitized inputs = %#v, want user message followed by assistant message",
					sanitized,
				)
			}
		})
	}
}

func TestSanitizeConversationStartStopsRemovingOutputsAfterAssistantTurn(t *testing.T) {
	tests := []struct {
		name           string
		assistantInput spec.InputUnion
		toolOutput     spec.InputUnion
	}{
		{
			name:           "assistant output message",
			assistantInput: conversationStartFunctionalAssistantMessage("answer"),
			toolOutput:     conversationStartTestFunctionToolOutput("function-call"),
		},
		{
			name:           "reasoning message",
			assistantInput: conversationStartFunctionalReasoningMessage(),
			toolOutput:     conversationStartTestFunctionToolOutput("function-call"),
		},
		{
			name:           "function tool call",
			assistantInput: conversationStartTestFunctionToolCall("function-call"),
			toolOutput:     conversationStartTestFunctionToolOutput("function-call"),
		},
		{
			name:           "custom tool call",
			assistantInput: conversationStartFunctionalCustomToolCall("custom-call"),
			toolOutput:     conversationStartFunctionalCustomToolOutput("custom-call"),
		},
		{
			name:           "web search tool call",
			assistantInput: conversationStartFunctionalWebSearchToolCall("web-search-call"),
			toolOutput:     conversationStartFunctionalWebSearchToolOutput("web-search-call"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sanitized, leadingDropped, firstUserToolOutputsDropped, err := sanitizeConversationStart(
				[]spec.InputUnion{
					conversationStartTestUserMessage("question"),
					tc.assistantInput,
					tc.toolOutput,
				},
			)
			if err != nil {
				t.Fatalf("sanitizeConversationStart() error = %v", err)
			}
			if leadingDropped != 0 || firstUserToolOutputsDropped != 0 {
				t.Fatalf(
					"unexpected drops: leading=%d firstUserToolOutputs=%d",
					leadingDropped,
					firstUserToolOutputsDropped,
				)
			}
			if len(sanitized) != 3 ||
				sanitized[0].Kind != spec.InputKindInputMessage ||
				sanitized[1].Kind != tc.assistantInput.Kind ||
				sanitized[2].Kind != tc.toolOutput.Kind {
				t.Fatalf("sanitized inputs = %#v, want original complete turn", sanitized)
			}
		})
	}
}

func TestNormalizeRequestForSDKConversationStartDoesNotMutateCallerInput(t *testing.T) {
	request := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{Name: "test-model"},
		Inputs: []spec.InputUnion{
			conversationStartTestFunctionToolOutput("prefix-call"),
			conversationStartTestUserMessage("question"),
			conversationStartFunctionalCustomToolOutput("first-user-tool-output"),
		},
	}

	normalized, _, _, err := NormalizeRequestForSDK(
		t.Context(),
		request,
		nil,
		spec.ProviderSDKTypeAnthropic,
		conversationStartTestCapabilities(),
	)
	if err != nil {
		t.Fatalf("NormalizeRequestForSDK() error = %v", err)
	}
	if got, want := len(normalized.Inputs), 1; got != want {
		t.Fatalf("len(normalized.Inputs) = %d, want %d", got, want)
	}
	if normalized.Inputs[0].Kind != spec.InputKindInputMessage {
		t.Fatalf("normalized first input kind = %q, want inputMessage", normalized.Inputs[0].Kind)
	}

	if got, want := len(request.Inputs), 3; got != want {
		t.Fatalf("len(request.Inputs) = %d, want %d", got, want)
	}
	if request.Inputs[0].Kind != spec.InputKindFunctionToolOutput ||
		request.Inputs[1].Kind != spec.InputKindInputMessage ||
		request.Inputs[2].Kind != spec.InputKindCustomToolOutput {
		t.Fatalf("original request inputs were mutated: %#v", request.Inputs)
	}
}

func conversationStartFunctionalAssistantMessage(text string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindOutputMessage,
		OutputMessage: &spec.InputOutputContent{
			Role: spec.RoleAssistant,
			Contents: []spec.InputOutputContentItemUnion{{
				Kind:     spec.ContentItemKindText,
				TextItem: &spec.ContentItemText{Text: text},
			}},
		},
	}
}

func conversationStartFunctionalReasoningMessage() spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindReasoningMessage,
		ReasoningMessage: &spec.ReasoningContent{
			Role:     spec.RoleAssistant,
			Thinking: []string{"thought"},
		},
	}
}

func conversationStartFunctionalCustomToolCall(callID string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindCustomToolCall,
		CustomToolCall: &spec.ToolCall{
			Type:      spec.ToolTypeCustom,
			Role:      spec.RoleAssistant,
			ID:        callID,
			CallID:    callID,
			Name:      "custom_tool",
			Arguments: "{}",
		},
	}
}

func conversationStartFunctionalWebSearchToolCall(callID string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindWebSearchToolCall,
		WebSearchToolCall: &spec.ToolCall{
			Type:   spec.ToolTypeWebSearch,
			Role:   spec.RoleAssistant,
			ID:     callID,
			CallID: callID,
			Name:   spec.DefaultWebSearchToolName,
		},
	}
}

func conversationStartFunctionalCustomToolOutput(callID string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindCustomToolOutput,
		CustomToolOutput: &spec.ToolOutput{
			Type:   spec.ToolTypeCustom,
			Role:   spec.RoleUser,
			CallID: callID,
			Name:   "custom_tool",
		},
	}
}

func conversationStartFunctionalWebSearchToolOutput(callID string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindWebSearchToolOutput,
		WebSearchToolOutput: &spec.ToolOutput{
			Type:   spec.ToolTypeWebSearch,
			Role:   spec.RoleUser,
			CallID: callID,
			Name:   spec.DefaultWebSearchToolName,
		},
	}
}
