package sdkutil

import (
	"errors"

	"github.com/flexigpt/inference-go/spec"
)

const conversationStartTrimmedWarningCode = "conversation_start_trimmed"

var errNoUserMessageInInputHistory = errors.New(
	"input history must contain a non-empty user message",
)

// sanitizeConversationStart makes a sliced history look like a new
// conversation:
//
// - Everything before the first ordinary user InputMessage is removed.
// - Tool outputs in that first user turn are removed.
//
// Tool output handling after the first assistant-side item is intentionally
// unchanged. This function only protects the first user turn.
func sanitizeConversationStart(
	inputs []spec.InputUnion,
) (
	sanitized []spec.InputUnion,
	leadingInputsDropped int,
	firstUserToolOutputsDropped int,
	err error,
) {
	if len(inputs) == 0 {
		return inputs, 0, 0, nil
	}

	firstUserIndex := -1
	for i, input := range inputs {
		if isOrdinaryUserInputMessage(input) {
			firstUserIndex = i
			break
		}
	}

	if firstUserIndex < 0 {
		return nil, 0, 0, errNoUserMessageInInputHistory
	}

	sanitized = make([]spec.InputUnion, 0, len(inputs)-firstUserIndex)
	firstUserTurn := true

	for _, input := range inputs[firstUserIndex:] {
		if firstUserTurn && isToolOutputInput(input) {
			firstUserToolOutputsDropped++
			continue
		}

		sanitized = append(sanitized, input)

		if firstUserTurn && startsAssistantTurn(input) {
			firstUserTurn = false
		}
	}

	return sanitized, firstUserIndex, firstUserToolOutputsDropped, nil
}

func isOrdinaryUserInputMessage(input spec.InputUnion) bool {
	return input.Kind == spec.InputKindInputMessage &&
		input.InputMessage != nil &&
		input.InputMessage.Role == spec.RoleUser &&
		!IsInputUnionEmpty(input)
}

func isToolOutputInput(input spec.InputUnion) bool {
	if IsInputUnionEmpty(input) {
		return false
	}

	switch input.Kind {
	case spec.InputKindFunctionToolOutput,
		spec.InputKindCustomToolOutput,
		spec.InputKindWebSearchToolOutput:
		return true
	default:
		return false
	}
}

func startsAssistantTurn(input spec.InputUnion) bool {
	if IsInputUnionEmpty(input) {
		return false
	}

	switch input.Kind {
	case spec.InputKindOutputMessage:
		return input.OutputMessage != nil &&
			input.OutputMessage.Role == spec.RoleAssistant
	case spec.InputKindReasoningMessage:
		return input.ReasoningMessage != nil
	case spec.InputKindFunctionToolCall:
		return input.FunctionToolCall != nil
	case spec.InputKindCustomToolCall:
		return input.CustomToolCall != nil
	case spec.InputKindWebSearchToolCall:
		return input.WebSearchToolCall != nil
	default:
		return false
	}
}
