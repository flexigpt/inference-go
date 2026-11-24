package googleinteractionssdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/flexigpt/inference-go/spec"
	"google.golang.org/genai/interactions/models/interactions"
)

func responseFromGoogleInteraction(
	interaction *interactions.Interaction,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) (*spec.FetchCompletionResponse, error) {
	if interaction == nil {
		err := errors.New("google interactions: response has no interaction")
		return &spec.FetchCompletionResponse{
			Error: &spec.Error{Message: err.Error()},
		}, err
	}

	return responseFromGoogleInteractionParts(
		googleInteractionsString(interaction.ID),
		googleInteractionsString(interaction.Status),
		googleInteractionsString(interaction.OutputText),
		interaction.Steps,
		interaction.Usage,
		toolChoiceNameMap,
		webSearchChoiceID,
	)
}

func responseFromGoogleInteractionSSE(
	interaction *interactions.InteractionSseEventInteraction,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) (*spec.FetchCompletionResponse, error) {
	if interaction == nil {
		err := errors.New(
			"google interactions: completed stream event has no interaction",
		)
		return &spec.FetchCompletionResponse{
			Error: &spec.Error{Message: err.Error()},
		}, err
	}

	return responseFromGoogleInteractionParts(
		googleInteractionsString(interaction.ID),
		googleInteractionsString(interaction.Status),
		"",
		interaction.Steps,
		interaction.Usage,
		toolChoiceNameMap,
		webSearchChoiceID,
	)
}

func responseFromGoogleInteractionParts(
	interactionID string,
	rawStatus string,
	outputText string,
	steps []interactions.Step,
	usage *interactions.Usage,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) (*spec.FetchCompletionResponse, error) {
	statusText := strings.ToLower(strings.TrimSpace(rawStatus))
	status := googleInteractionsStatus(statusText)

	outputs, warnings := outputsFromGoogleInteractionSteps(
		interactionID,
		steps,
		status,
		toolChoiceNameMap,
		webSearchChoiceID,
	)

	if !hasGoogleInteractionsTextOutput(outputs) &&
		strings.TrimSpace(outputText) != "" {
		outputs = append(outputs, spec.OutputUnion{
			Kind: spec.OutputKindOutputMessage,
			OutputMessage: &spec.InputOutputContent{
				ID:     interactionID,
				Role:   spec.RoleAssistant,
				Status: status,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind: spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{
						Text: outputText,
					},
				}},
			},
		})
	}

	response := &spec.FetchCompletionResponse{
		Outputs:  outputs,
		Usage:    usageFromGoogleInteraction(usage),
		Warnings: warnings,
	}

	if err := googleInteractionsTerminalError(statusText); err != nil {
		response.Error = &spec.Error{Message: err.Error()}
		return response, err
	}

	return response, nil
}

func outputsFromGoogleInteractionSteps(
	interactionID string,
	steps []interactions.Step,
	status spec.Status,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) ([]spec.OutputUnion, []spec.Warning) {
	var (
		outputs  []spec.OutputUnion
		warnings []spec.Warning
	)

	for index, step := range steps {
		switch {
		case step.ModelOutputStep != nil:
			message, dropped := googleInteractionsOutputMessage(
				interactionID,
				status,
				step.ModelOutputStep.Content,
			)
			if dropped {
				warnings = appendGoogleInteractionsWarning(
					warnings,
					"output_content_dropped_unsupported",
					"Interactions output outside the normalized text output surface was dropped.",
				)
			}
			if message != nil {
				outputs = append(outputs, spec.OutputUnion{
					Kind:          spec.OutputKindOutputMessage,
					OutputMessage: message,
				})
			}

		case step.ThoughtStep != nil:
			reasoning := &spec.ReasoningContent{
				ID:        interactionID,
				Role:      spec.RoleAssistant,
				Status:    status,
				Signature: googleInteractionsString(step.ThoughtStep.Signature),
			}
			for _, content := range step.ThoughtStep.Summary {
				if text, ok := googleInteractionsThoughtSummaryText(content); ok {
					reasoning.Summary = append(reasoning.Summary, text)
				}
			}
			if reasoning.Signature != "" || len(reasoning.Summary) > 0 {
				outputs = append(outputs, spec.OutputUnion{
					Kind:             spec.OutputKindReasoningMessage,
					ReasoningMessage: reasoning,
				})
			}

		case step.FunctionCallStep != nil:
			call, kind, warning := googleInteractionsFunctionToolCall(
				step.FunctionCallStep,
				index,
				status,
				toolChoiceNameMap,
			)
			if warning != nil {
				warnings = append(warnings, *warning)
			}
			if call == nil {
				continue
			}

			output := spec.OutputUnion{Kind: kind}
			if kind == spec.OutputKindCustomToolCall {
				output.CustomToolCall = call
			} else {
				output.FunctionToolCall = call
			}
			outputs = append(outputs, output)

		case step.GoogleSearchCallStep != nil:
			if webSearchChoiceID == "" {
				continue
			}
			call := googleInteractionsWebSearchToolCall(
				step.GoogleSearchCallStep,
				status,
				webSearchChoiceID,
			)
			if call != nil {
				outputs = append(outputs, spec.OutputUnion{
					Kind:              spec.OutputKindWebSearchToolCall,
					WebSearchToolCall: call,
				})
			}

		case step.GoogleSearchResultStep != nil:
			if webSearchChoiceID == "" {
				continue
			}
			output, hasSuggestions := googleInteractionsWebSearchToolOutput(
				step.GoogleSearchResultStep,
				status,
				webSearchChoiceID,
			)
			if output != nil {
				outputs = append(outputs, spec.OutputUnion{
					Kind:                spec.OutputKindWebSearchToolOutput,
					WebSearchToolOutput: output,
				})
			}
			if hasSuggestions {
				warnings = appendGoogleInteractionsWarning(
					warnings,
					"googleSearch_citations_unavailable",
					"Interactions Google Search result steps expose suggestions but not source URL, title, or span data required for normalized URL citations.",
				)
			}
		}
	}

	return outputs, warnings
}

func googleInteractionsOutputMessage(
	id string,
	status spec.Status,
	content []interactions.Content,
) (*spec.InputOutputContent, bool) {
	message := &spec.InputOutputContent{
		ID:     id,
		Role:   spec.RoleAssistant,
		Status: status,
	}
	dropped := false

	for _, item := range content {
		text, ok := googleInteractionsTextContent(item)
		if !ok {
			dropped = true
			continue
		}
		message.Contents = append(message.Contents, spec.InputOutputContentItemUnion{
			Kind: spec.ContentItemKindText,
			TextItem: &spec.ContentItemText{
				Text: text,
			},
		})
	}

	if len(message.Contents) == 0 {
		return nil, dropped
	}
	return message, dropped
}

func googleInteractionsTextContent(
	content interactions.Content,
) (string, bool) {
	if content.TextContent == nil {
		return "", false
	}

	text := googleInteractionsString(content.TextContent.Text)
	if text == "" {
		return "", false
	}
	return text, true
}

func googleInteractionsThoughtSummaryText(
	content interactions.ThoughtSummaryContent,
) (string, bool) {
	if content.TextContent == nil {
		return "", false
	}

	text := googleInteractionsString(content.TextContent.Text)
	if text == "" {
		return "", false
	}
	return text, true
}

func googleInteractionsFunctionToolCall(
	functionCall *interactions.FunctionCallStep,
	index int,
	status spec.Status,
	toolChoiceNameMap map[string]spec.ToolChoice,
) (*spec.ToolCall, spec.OutputKind, *spec.Warning) {
	if functionCall == nil {
		return nil, "", nil
	}

	name := strings.TrimSpace(googleInteractionsString(functionCall.Name))
	choice, ok := toolChoiceNameMap[name]
	if !ok {
		return nil, "", &spec.Warning{
			Code:    "toolCall_dropped_unknown",
			Message: "Interactions returned a function call without a matching ToolChoice.",
		}
	}

	callID := strings.TrimSpace(googleInteractionsString(functionCall.ID))
	if callID == "" {
		callID = fmt.Sprintf("%s_%d", name, index)
	}

	arguments := "{}"
	if functionCall.Arguments != nil {
		raw, err := json.Marshal(functionCall.Arguments)
		if err == nil {
			arguments = string(raw)
		}
	}

	call := &spec.ToolCall{
		Type:      choice.Type,
		ChoiceID:  choice.ID,
		ID:        callID,
		CallID:    callID,
		Name:      name,
		Arguments: arguments,
		Role:      spec.RoleAssistant,
		Status:    status,
	}

	if choice.Type == spec.ToolTypeCustom {
		return call, spec.OutputKindCustomToolCall, nil
	}
	return call, spec.OutputKindFunctionToolCall, nil
}

func googleInteractionsWebSearchToolCall(
	searchCall *interactions.GoogleSearchCallStep,
	status spec.Status,
	choiceID string,
) *spec.ToolCall {
	if searchCall == nil {
		return nil
	}

	callID := strings.TrimSpace(googleInteractionsString(searchCall.ID))
	if callID == "" {
		return nil
	}

	call := &spec.ToolCall{
		Type:      spec.ToolTypeWebSearch,
		ChoiceID:  choiceID,
		ID:        callID,
		CallID:    callID,
		Name:      spec.DefaultWebSearchToolName,
		Role:      spec.RoleAssistant,
		Status:    status,
		Signature: googleInteractionsString(searchCall.Signature),
	}

	for _, query := range searchCall.Arguments.Queries {
		if strings.TrimSpace(query) == "" {
			continue
		}
		call.WebSearchToolCallItems = append(
			call.WebSearchToolCallItems,
			spec.WebSearchToolCallItemUnion{
				Kind: spec.WebSearchToolCallKindSearch,
				SearchItem: &spec.WebSearchToolCallSearch{
					Query: query,
				},
			},
		)
	}

	return call
}

func googleInteractionsWebSearchToolOutput(
	searchResult *interactions.GoogleSearchResultStep,
	status spec.Status,
	choiceID string,
) (*spec.ToolOutput, bool) {
	if searchResult == nil {
		return nil, false
	}

	callID := strings.TrimSpace(googleInteractionsString(searchResult.CallID))
	if callID == "" {
		return nil, false
	}

	output := &spec.ToolOutput{
		Type:      spec.ToolTypeWebSearch,
		ChoiceID:  choiceID,
		ID:        callID,
		CallID:    callID,
		Name:      spec.DefaultWebSearchToolName,
		Role:      spec.RoleAssistant,
		Status:    status,
		Signature: googleInteractionsString(searchResult.Signature),
		IsError:   googleInteractionsBool(searchResult.IsError),
	}

	hasSuggestions := false
	for _, item := range searchResult.Result {
		suggestion := strings.TrimSpace(
			googleInteractionsString(item.SearchSuggestions),
		)
		if suggestion == "" {
			continue
		}
		hasSuggestions = true
		output.WebSearchToolOutputItems = append(
			output.WebSearchToolOutputItems,
			spec.WebSearchToolOutputItemUnion{
				Kind: spec.WebSearchToolOutputKindSearch,
				SearchItem: &spec.WebSearchToolOutputSearch{
					RenderedContent: suggestion,
				},
			},
		)
	}

	if output.IsError {
		output.WebSearchToolOutputItems = append(
			output.WebSearchToolOutputItems,
			spec.WebSearchToolOutputItemUnion{
				Kind: spec.WebSearchToolOutputKindError,
				ErrorItem: &spec.WebSearchToolOutputError{
					Code: "google_search_error",
				},
			},
		)
	}

	return output, hasSuggestions
}

func usageFromGoogleInteraction(
	usage *interactions.Usage,
) *spec.Usage {
	if usage == nil {
		return &spec.Usage{}
	}

	input := googleInteractionsInt64(usage.TotalInputTokens)
	cached := googleInteractionsInt64(usage.TotalCachedTokens)

	return &spec.Usage{
		InputTokensTotal:    input,
		InputTokensCached:   cached,
		InputTokensUncached: max(input-cached, 0),
		OutputTokens: googleInteractionsInt64(
			usage.TotalOutputTokens,
		),
		ReasoningTokens: googleInteractionsInt64(
			usage.TotalThoughtTokens,
		),
	}
}

func googleInteractionsStatus(status string) spec.Status {
	switch status {
	case "completed":
		return spec.StatusCompleted
	case "requires_action", "in_progress":
		return spec.StatusInProgress
	case "incomplete", "requires_continuation":
		return spec.StatusIncomplete
	case "failed":
		return spec.StatusFailed
	case "cancelled":
		return spec.StatusCancelled
	default:
		return spec.StatusNone
	}
}

func googleInteractionsTerminalError(status string) error {
	switch status {
	case "completed", "incomplete", "requires_action":
		return nil
	case "requires_continuation":
		return errors.New(
			"google interactions: interaction requires continuation, which is outside the stateless adapter surface",
		)
	case "failed", "cancelled":
		return fmt.Errorf("google interactions: interaction %s", status)
	default:
		return fmt.Errorf(
			"google interactions: interaction ended with unexpected status %q",
			status,
		)
	}
}

func hasGoogleInteractionsTextOutput(outputs []spec.OutputUnion) bool {
	for _, output := range outputs {
		if output.Kind != spec.OutputKindOutputMessage ||
			output.OutputMessage == nil {
			continue
		}
		for _, content := range output.OutputMessage.Contents {
			if content.Kind == spec.ContentItemKindText &&
				content.TextItem != nil &&
				content.TextItem.Text != "" {
				return true
			}
		}
	}
	return false
}

func googleInteractionsString(value any) string {
	switch value := value.(type) {
	case nil:
		return ""
	case string:
		return strings.TrimSpace(value)
	case *string:
		if value == nil {
			return ""
		}
		return strings.TrimSpace(*value)
	default:
		return strings.TrimSpace(fmt.Sprint(value))
	}
}

func googleInteractionsBool(value any) bool {
	switch value := value.(type) {
	case bool:
		return value
	case *bool:
		return value != nil && *value
	default:
		return false
	}
}

func googleInteractionsInt64(value any) int64 {
	switch value := value.(type) {
	case int:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case *int:
		if value != nil {
			return int64(*value)
		}
	case *int32:
		if value != nil {
			return int64(*value)
		}
	case *int64:
		if value != nil {
			return *value
		}
	}
	return 0
}
