package sdkutil

import (
	"errors"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

var conversationStartTestSDKTypes = []spec.ProviderSDKType{
	spec.ProviderSDKTypeAnthropic,
	spec.ProviderSDKTypeGoogleGenerateContent,
	spec.ProviderSDKTypeOpenAIChatCompletions,
	spec.ProviderSDKTypeOpenAIResponses,
}

func TestNormalizeRequestForSDKSanitizesConversationStart(t *testing.T) {
	for _, sdkType := range conversationStartTestSDKTypes {
		t.Run(string(sdkType), func(t *testing.T) {
			normalized, _, warnings, err := NormalizeRequestForSDK(
				t.Context(),
				&spec.FetchCompletionRequest{
					ModelParam: spec.ModelParam{Name: "test-model"},
					Inputs: []spec.InputUnion{
						conversationStartTestFunctionToolCall("prefix-call"),
						conversationStartTestFunctionToolOutput("prefix-call"),
						conversationStartTestUserMessage("first retained user message"),
						conversationStartTestFunctionToolOutput("drop-from-first-user-turn"),
						conversationStartTestFunctionToolCall("later-call"),
						conversationStartTestFunctionToolOutput("later-call"),
					},
				},
				nil,
				sdkType,
				conversationStartTestCapabilities(),
			)
			if err != nil {
				t.Fatalf("NormalizeRequestForSDK() error = %v", err)
			}

			if got, want := len(normalized.Inputs), 3; got != want {
				t.Fatalf("normalized input count = %d, want %d", got, want)
			}
			if normalized.Inputs[0].Kind != spec.InputKindInputMessage ||
				normalized.Inputs[1].Kind != spec.InputKindFunctionToolCall ||
				normalized.Inputs[2].Kind != spec.InputKindFunctionToolOutput {
				t.Fatalf(
					"normalized input kinds = %#v, want user message, assistant tool call, tool output",
					normalized.Inputs,
				)
			}

			if !conversationStartTestHasWarning(
				warnings,
				conversationStartTrimmedWarningCode,
			) {
				t.Errorf(
					"warnings = %#v, want warning code %q",
					warnings,
					conversationStartTrimmedWarningCode,
				)
			}
		})
	}
}

func TestNormalizeRequestForSDKRejectsOutputOnlyHistory(t *testing.T) {
	for _, sdkType := range conversationStartTestSDKTypes {
		t.Run(string(sdkType), func(t *testing.T) {
			_, _, _, err := NormalizeRequestForSDK(
				t.Context(),
				&spec.FetchCompletionRequest{
					ModelParam: spec.ModelParam{Name: "test-model"},
					Inputs: []spec.InputUnion{
						conversationStartTestFunctionToolOutput("orphaned-call"),
					},
				},
				nil,
				sdkType,
				conversationStartTestCapabilities(),
			)

			if !errors.Is(err, errNoUserMessageInInputHistory) {
				t.Fatalf(
					"NormalizeRequestForSDK() error = %v, want no-user-message error",
					err,
				)
			}
		})
	}
}

func conversationStartTestCapabilities() spec.ModelCapabilities {
	return spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{spec.ModalityTextIn},
	}
}

func conversationStartTestUserMessage(text string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindInputMessage,
		InputMessage: &spec.InputOutputContent{
			Role: spec.RoleUser,
			Contents: []spec.InputOutputContentItemUnion{
				{
					Kind: spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{
						Text: text,
					},
				},
			},
		},
	}
}

func conversationStartTestFunctionToolCall(callID string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindFunctionToolCall,
		FunctionToolCall: &spec.ToolCall{
			Type:      spec.ToolTypeFunction,
			Role:      spec.RoleAssistant,
			ID:        callID,
			CallID:    callID,
			Name:      "lookup",
			Arguments: "{}",
		},
	}
}

func conversationStartTestFunctionToolOutput(callID string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindFunctionToolOutput,
		FunctionToolOutput: &spec.ToolOutput{
			Type:   spec.ToolTypeFunction,
			Role:   spec.RoleUser,
			CallID: callID,
			Contents: []spec.ToolOutputItemUnion{
				{
					Kind: spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{
						Text: "tool result",
					},
				},
			},
		},
	}
}

func conversationStartTestHasWarning(warnings []spec.Warning, code string) bool {
	for _, warning := range warnings {
		if warning.Code == code {
			return true
		}
	}
	return false
}
