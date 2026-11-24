package googleinteractionssdk

import (
	"context"
	"errors"
	"slices"
	"strings"
	"sync"

	"google.golang.org/genai"
	"google.golang.org/genai/interactions/models/interactions"
	"google.golang.org/genai/interactions/models/operations"

	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

func (api *GoogleInteractionsAPI) doStreaming(
	ctx context.Context,
	client *genai.Client,
	provider spec.ProviderName,
	model spec.ModelName,
	request operations.CreateInteractionRequest,
	opts *spec.FetchCompletionOptions,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) (*spec.FetchCompletionResponse, any, error) {
	streamCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Interactions streaming is selected by CreateModelInteraction.Stream.
	// The Create response then exposes InteractionSSEStreamEvent.
	result, err := client.Interactions.Create(
		streamCtx,
		request,
		operations.WithAutomaticContinuation(false),
	)
	if err != nil {
		response := &spec.FetchCompletionResponse{
			Error: &spec.Error{Message: err.Error()},
		}
		return response, nil, err
	}
	if result == nil || result.InteractionSSEStreamEvent == nil {
		err := errors.New(
			"google interactions: stream request returned no InteractionSSEStreamEvent",
		)
		return &spec.FetchCompletionResponse{
			Error: &spec.Error{Message: err.Error()},
		}, result, err
	}

	stream := result.InteractionSSEStreamEvent
	defer func() {
		_ = stream.Close()
	}()

	streamConfig := sdkutil.ResolveStreamConfig(opts)
	var emitMu sync.Mutex
	emit := func(kind spec.StreamContentKind, text string) error {
		if text == "" {
			return nil
		}
		emitMu.Lock()
		defer emitMu.Unlock()

		event := spec.StreamEvent{
			Kind:          kind,
			Provider:      provider,
			Model:         model,
			CompletionKey: opts.CompletionKey,
		}
		switch kind {
		case spec.StreamContentKindText:
			event.Text = &spec.StreamTextChunk{Text: text}
		case spec.StreamContentKindThinking:
			event.Thinking = &spec.StreamThinkingChunk{Text: text}
		}

		err := sdkutil.SafeCallStreamHandler(opts.StreamHandler, event)
		if err != nil {
			cancel()
		}
		return err
	}
	writeText, flushText := sdkutil.NewBufferedStreamerWithError(
		func(text string) error {
			return emit(spec.StreamContentKindText, text)
		},
		streamConfig.FlushInterval,
		streamConfig.FlushChunkSize,
	)
	writeThinking, flushThinking := sdkutil.NewBufferedStreamerWithError(
		func(text string) error {
			return emit(spec.StreamContentKindThinking, text)
		},
		streamConfig.FlushInterval,
		streamConfig.FlushChunkSize,
	)

	state := newGoogleInteractionsStreamState()
	var streamErr error

	for stream.Next() {
		event := stream.Value()
		if event == nil {
			continue
		}

		if created := event.GetDataInteractionCreated(); created != nil {
			interaction := created.GetInteraction()
			state.setInteraction(
				&interaction,
				toolChoiceNameMap,
				webSearchChoiceID,
				false,
			)
		}

		if statusUpdate := event.GetDataInteractionStatusUpdate(); statusUpdate != nil {
			state.setRawStatus(googleInteractionsString(statusUpdate.GetStatus()))
			state.recordUsage(statusUpdate.GetUsage())
		}

		if start := event.GetDataStepStart(); start != nil {
			state.recordStep(
				start.GetIndex(),
				start.GetStep(),
				toolChoiceNameMap,
				webSearchChoiceID,
				false,
			)
		}

		if delta := event.GetDataStepDelta(); delta != nil {
			if err := state.applyDelta(delta, writeText, writeThinking); err != nil {
				streamErr = err
				break
			}
		}

		if stop := event.GetDataStepStop(); stop != nil {
			state.recordUsage(stop.GetUsage())
		}

		if completed := event.GetDataInteractionCompleted(); completed != nil {
			interaction := completed.GetInteraction()
			state.setInteraction(
				&interaction,
				toolChoiceNameMap,
				webSearchChoiceID,
				true,
			)

			state.response, state.responseErr = responseFromGoogleInteractionSSE(
				&interaction,
				toolChoiceNameMap,
				webSearchChoiceID,
			)
			if interaction.Usage == nil && state.response != nil && state.usage != nil {
				state.response.Usage = state.responseUsage()
			}
			state.terminal = true
			break
		}

		if event.GetDataError() != nil {
			state.eventErr = errors.New(
				"google interactions: stream emitted an error event",
			)
			state.terminal = true
			break
		}
	}

	streamErr = errors.Join(
		streamErr,
		stream.Err(),
		flushText(),
		flushThinking(),
	)
	if !state.terminal && streamErr == nil {
		streamErr = state.unfinishedStreamError()
	}

	response := state.response
	if response == nil {
		response = state.partialResponse()
	}
	appendGoogleInteractionsStreamFallback(response, state)

	callErr := errors.Join(
		streamErr,
		state.eventErr,
		state.responseErr,
	)
	if callErr != nil && response.Error == nil {
		response.Error = &spec.Error{Message: callErr.Error()}
	}

	return response, result, callErr
}

type googleInteractionsStreamState struct {
	interactionID string
	status        spec.Status
	usage         *spec.Usage
	rawStatus     string
	terminal      bool
	eventErr      error
	response      *spec.FetchCompletionResponse
	responseErr   error
	steps         map[int]*googleInteractionsStreamStep
	order         []int
	warnings      []spec.Warning
}

type googleInteractionsStreamStep struct {
	text             strings.Builder
	thoughtSummary   strings.Builder
	thoughtSignature string
	functionCall     *spec.ToolCall
	functionArgs     strings.Builder
	webSearchCall    *spec.ToolCall
	webSearchOutput  *spec.ToolOutput
}

func newGoogleInteractionsStreamState() *googleInteractionsStreamState {
	return &googleInteractionsStreamState{
		steps: make(map[int]*googleInteractionsStreamStep),
	}
}

func (state *googleInteractionsStreamState) partialResponse() *spec.FetchCompletionResponse {
	outputs, warnings := state.streamOutputs()
	response := &spec.FetchCompletionResponse{
		Outputs:  outputs,
		Usage:    state.responseUsage(),
		Warnings: warnings,
	}

	return response
}

func (state *googleInteractionsStreamState) setInteraction(
	interaction *interactions.InteractionSseEventInteraction,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
	replaceSteps bool,
) {
	if interaction == nil {
		return
	}

	if id := googleInteractionsString(interaction.ID); id != "" {
		state.interactionID = id
	}
	state.setRawStatus(googleInteractionsString(interaction.Status))
	state.recordUsage(interaction.Usage)

	if len(interaction.Steps) == 0 {
		return
	}
	if replaceSteps {
		state.steps = make(map[int]*googleInteractionsStreamStep, len(interaction.Steps))
		state.order = nil
	}

	for index, step := range interaction.Steps {
		state.recordStep(
			index,
			step,
			toolChoiceNameMap,
			webSearchChoiceID,
			replaceSteps,
		)
	}
}

func (state *googleInteractionsStreamState) setStatus(status spec.Status) {
	if status != spec.StatusNone {
		state.status = status
	}
}

func (state *googleInteractionsStreamState) setRawStatus(rawStatus string) {
	rawStatus = strings.ToLower(strings.TrimSpace(rawStatus))
	if rawStatus == "" {
		return
	}

	state.rawStatus = rawStatus
	state.setStatus(googleInteractionsStatus(rawStatus))
}

func (state *googleInteractionsStreamState) recordUsage(
	usage *interactions.Usage,
) {
	if usage == nil {
		return
	}
	state.usage = usageFromGoogleInteraction(usage)
}

func (state *googleInteractionsStreamState) recordStep(
	index int,
	step interactions.Step,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
	replace bool,
) {
	entry := state.ensureStep(index, replace)

	switch {
	case step.ModelOutputStep != nil:
		if replace || entry.text.Len() == 0 {
			for _, content := range step.ModelOutputStep.Content {
				text, ok := googleInteractionsTextContent(content)
				if !ok {
					state.warnings = appendGoogleInteractionsWarning(
						state.warnings,
						"output_content_dropped_unsupported",
						"Interactions output outside the normalized text output surface was dropped.",
					)
					continue
				}
				entry.text.WriteString(text)
			}
		}

	case step.ThoughtStep != nil:
		if replace || entry.thoughtSummary.Len() == 0 {
			entry.thoughtSignature = googleInteractionsRawString(
				step.ThoughtStep.Signature,
			)
			for _, content := range step.ThoughtStep.Summary {
				if text, ok := googleInteractionsThoughtSummaryText(content); ok {
					entry.thoughtSummary.WriteString(text)
				}
			}
		}

	case step.FunctionCallStep != nil:
		if replace || entry.functionCall == nil {
			call, _, warning := googleInteractionsFunctionToolCall(
				step.FunctionCallStep,
				index,
				state.responseStatus(),
				toolChoiceNameMap,
			)
			if warning != nil {
				state.warnings = append(state.warnings, *warning)
			}
			entry.functionCall = call
			if replace {
				entry.functionArgs.Reset()
			}
		}

	case step.GoogleSearchCallStep != nil:
		if replace || entry.webSearchCall == nil {
			entry.webSearchCall = googleInteractionsWebSearchToolCall(
				step.GoogleSearchCallStep,
				state.responseStatus(),
				webSearchChoiceID,
			)
		}

	case step.GoogleSearchResultStep != nil:
		if replace || entry.webSearchOutput == nil {
			output, hasSuggestions := googleInteractionsWebSearchToolOutput(
				step.GoogleSearchResultStep,
				state.responseStatus(),
				webSearchChoiceID,
			)
			entry.webSearchOutput = output
			if hasSuggestions {
				state.warnings = appendGoogleInteractionsWarning(
					state.warnings,
					"googleSearch_citations_unavailable",
					"Interactions Google Search result steps expose suggestions but not source URL, title, or span data required for normalized URL citations.",
				)
			}
		}
	}
}

func (state *googleInteractionsStreamState) applyDelta(
	delta *interactions.StepDelta,
	writeText func(string) error,
	writeThinking func(string) error,
) error {
	if delta == nil {
		return nil
	}

	if metadata := delta.GetMetadata(); metadata != nil {
		state.recordUsage(metadata.GetTotalUsage())
	}

	entry := state.ensureStep(delta.GetIndex(), false)

	if textDelta := delta.GetDeltaText(); textDelta != nil {
		text := googleInteractionsRawString(textDelta.GetText())
		entry.text.WriteString(text)
		if err := writeText(text); err != nil {
			return err
		}
	}

	if thoughtSummaryDelta := delta.GetDeltaThoughtSummary(); thoughtSummaryDelta != nil {
		text := googleInteractionsStreamText(thoughtSummaryDelta.GetContent())
		entry.thoughtSummary.WriteString(text)
		if err := writeThinking(text); err != nil {
			return err
		}
	}

	if thoughtSignatureDelta := delta.GetDeltaThoughtSignature(); thoughtSignatureDelta != nil {
		entry.thoughtSignature = googleInteractionsRawString(
			thoughtSignatureDelta.GetSignature(),
		)
	}

	if argumentsDelta := delta.GetDeltaArgumentsDelta(); argumentsDelta != nil {
		if entry.functionCall == nil {
			state.warnings = appendGoogleInteractionsWarning(
				state.warnings,
				"stream_arguments_delta_without_function_call",
				"Interactions emitted an arguments delta before a function call step.",
			)
		} else {
			entry.functionArgs.WriteString(
				googleInteractionsRawString(argumentsDelta.GetArguments()),
			)
		}
	}

	if searchCallDelta := delta.GetDeltaGoogleSearchCall(); searchCallDelta != nil {
		state.applyGoogleSearchCallDelta(entry, searchCallDelta)
	}
	if searchResultDelta := delta.GetDeltaGoogleSearchResult(); searchResultDelta != nil {
		state.applyGoogleSearchResultDelta(entry, searchResultDelta)
	}

	if delta.GetDeltaTextAnnotationDelta() != nil {
		state.warnings = appendGoogleInteractionsWarning(
			state.warnings,
			"textAnnotations_dropped_unsupported",
			"Interactions emitted text annotations, but no normalized citation mapping is configured for this SDK annotation type.",
		)
	}

	return nil
}

func (state *googleInteractionsStreamState) applyGoogleSearchCallDelta(
	entry *googleInteractionsStreamStep,
	delta *interactions.GoogleSearchCallDelta,
) {
	if entry == nil || delta == nil {
		return
	}
	if entry.webSearchCall == nil {
		state.warnings = appendGoogleInteractionsWarning(
			state.warnings,
			"stream_googleSearchCall_delta_without_start",
			"Interactions emitted a Google Search call delta before a Google Search call step.",
		)
		return
	}

	if signature := googleInteractionsRawString(delta.GetSignature()); signature != "" {
		entry.webSearchCall.Signature = signature
	}
	for _, query := range googleInteractionsGoogleSearchQueries(delta.GetArguments()) {
		appendGoogleInteractionsSearchQuery(entry.webSearchCall, query)
	}
}

func (state *googleInteractionsStreamState) applyGoogleSearchResultDelta(
	entry *googleInteractionsStreamStep,
	delta *interactions.GoogleSearchResultDelta,
) {
	if entry == nil || delta == nil {
		return
	}
	if entry.webSearchOutput == nil {
		state.warnings = appendGoogleInteractionsWarning(
			state.warnings,
			"stream_googleSearchResult_delta_without_start",
			"Interactions emitted a Google Search result delta before a Google Search result step.",
		)
		return
	}

	if signature := googleInteractionsRawString(delta.GetSignature()); signature != "" {
		entry.webSearchOutput.Signature = signature
	}
	if isError := delta.GetIsError(); isError != nil {
		entry.webSearchOutput.IsError = *isError
	}
	for _, result := range delta.GetResult() {
		suggestion := strings.TrimSpace(
			googleInteractionsString(result.GetSearchSuggestions()),
		)
		appendGoogleInteractionsSearchSuggestion(
			entry.webSearchOutput,
			suggestion,
		)
	}
}

func googleInteractionsGoogleSearchQueries(value any) []string {
	switch value := value.(type) {
	case interactions.GoogleSearchCallArguments:
		return slices.Clone(value.Queries)
	case *interactions.GoogleSearchCallArguments:
		if value == nil {
			return nil
		}
		return slices.Clone(value.Queries)
	default:
		return nil
	}
}

func appendGoogleInteractionsSearchQuery(
	call *spec.ToolCall,
	query string,
) {
	if call == nil || strings.TrimSpace(query) == "" {
		return
	}

	for _, item := range call.WebSearchToolCallItems {
		if item.Kind == spec.WebSearchToolCallKindSearch &&
			item.SearchItem != nil &&
			item.SearchItem.Query == query {
			return
		}
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

func appendGoogleInteractionsSearchSuggestion(
	output *spec.ToolOutput,
	suggestion string,
) {
	if output == nil || strings.TrimSpace(suggestion) == "" {
		return
	}

	for _, item := range output.WebSearchToolOutputItems {
		if item.Kind == spec.WebSearchToolOutputKindSearch &&
			item.SearchItem != nil &&
			item.SearchItem.RenderedContent == suggestion {
			return
		}
	}

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

func (state *googleInteractionsStreamState) ensureStep(
	index int,
	replace bool,
) *googleInteractionsStreamStep {
	entry, exists := state.steps[index]
	if !exists {
		state.order = append(state.order, index)
	}
	if !exists || replace {
		entry = &googleInteractionsStreamStep{}
		state.steps[index] = entry
	}
	return entry
}

func (state *googleInteractionsStreamState) streamOutputs() (
	[]spec.OutputUnion,
	[]spec.Warning,
) {
	status := state.responseStatus()
	indexes := slices.Clone(state.order)
	slices.Sort(indexes)

	outputs := make([]spec.OutputUnion, 0, len(indexes))
	warnings := slices.Clone(state.warnings)

	for _, index := range indexes {
		entry := state.steps[index]
		if entry == nil {
			continue
		}

		if text := entry.text.String(); text != "" {
			outputs = append(outputs, spec.OutputUnion{
				Kind: spec.OutputKindOutputMessage,
				OutputMessage: &spec.InputOutputContent{
					ID:     state.interactionID,
					Role:   spec.RoleAssistant,
					Status: status,
					Contents: []spec.InputOutputContentItemUnion{{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: text},
					}},
				},
			})
		}

		if entry.thoughtSignature != "" || entry.thoughtSummary.Len() > 0 {
			reasoning := &spec.ReasoningContent{
				ID:        state.interactionID,
				Role:      spec.RoleAssistant,
				Status:    status,
				Signature: entry.thoughtSignature,
			}
			if summary := entry.thoughtSummary.String(); summary != "" {
				reasoning.Summary = []string{summary}
			}
			outputs = append(outputs, spec.OutputUnion{
				Kind:             spec.OutputKindReasoningMessage,
				ReasoningMessage: reasoning,
			})
		}

		if entry.functionCall != nil {
			call := *entry.functionCall
			if arguments := entry.functionArgs.String(); arguments != "" {
				call.Arguments = arguments
			}
			call.Status = status

			if call.Type == spec.ToolTypeCustom {
				outputs = append(outputs, spec.OutputUnion{
					Kind:           spec.OutputKindCustomToolCall,
					CustomToolCall: &call,
				})
			} else {
				outputs = append(outputs, spec.OutputUnion{
					Kind:             spec.OutputKindFunctionToolCall,
					FunctionToolCall: &call,
				})
			}
		}

		if entry.webSearchCall != nil {
			call := *entry.webSearchCall
			call.Status = status
			outputs = append(outputs, spec.OutputUnion{
				Kind:              spec.OutputKindWebSearchToolCall,
				WebSearchToolCall: &call,
			})
		}

		if entry.webSearchOutput != nil {
			output := *entry.webSearchOutput
			output.Status = status
			outputs = append(outputs, spec.OutputUnion{
				Kind:                spec.OutputKindWebSearchToolOutput,
				WebSearchToolOutput: &output,
			})
			warnings = appendGoogleInteractionsWarning(
				warnings,
				"googleSearch_citations_unavailable",
				"Interactions Google Search result steps expose suggestions but not source URL, title, or span data required for normalized URL citations.",
			)
		}
	}

	return outputs, warnings
}

func (state *googleInteractionsStreamState) unfinishedStreamError() error {
	if state != nil {
		switch state.rawStatus {
		case googleInteractionsRawStatusRequiresContinuation,
			googleInteractionsRawStatusFailed,
			googleInteractionsRawStatusCancelled:
			if err := googleInteractionsTerminalError(state.rawStatus); err != nil {
				return err
			}
		}
	}

	return errors.New(
		"google interactions: stream ended before interaction.completed",
	)
}

func appendGoogleInteractionsStreamFallback(
	response *spec.FetchCompletionResponse,
	state *googleInteractionsStreamState,
) {
	if response == nil || state == nil {
		return
	}

	streamOutputs, streamWarnings := state.streamOutputs()
	for _, warning := range streamWarnings {
		response.Warnings = appendGoogleInteractionsWarning(
			response.Warnings,
			warning.Code,
			warning.Message,
		)
	}

	for _, output := range streamOutputs {
		switch output.Kind {
		case spec.OutputKindOutputMessage:
			text := googleInteractionsOutputUnionText(output)
			if text != "" &&
				!googleInteractionsResponseContainsText(response.Outputs, text) {
				response.Outputs = append(response.Outputs, output)
			}

		case spec.OutputKindReasoningMessage:
			appendGoogleInteractionsReasoningFallback(
				&response.Outputs,
				output.ReasoningMessage,
			)

		case spec.OutputKindFunctionToolCall,
			spec.OutputKindCustomToolCall,
			spec.OutputKindWebSearchToolCall,
			spec.OutputKindWebSearchToolOutput:
			if !googleInteractionsContainsMatchingToolOutput(
				response.Outputs,
				output,
			) {
				response.Outputs = append(response.Outputs, output)
			}
		}
	}
}

func appendGoogleInteractionsReasoningFallback(
	outputs *[]spec.OutputUnion,
	reasoning *spec.ReasoningContent,
) {
	if outputs == nil || reasoning == nil {
		return
	}

	for index := range *outputs {
		output := &(*outputs)[index]
		if output.Kind != spec.OutputKindReasoningMessage ||
			output.ReasoningMessage == nil ||
			reasoning.Signature == "" ||
			output.ReasoningMessage.Signature != reasoning.Signature {
			continue
		}

		for _, summary := range reasoning.Summary {
			if !slices.Contains(output.ReasoningMessage.Summary, summary) {
				output.ReasoningMessage.Summary = append(
					output.ReasoningMessage.Summary,
					summary,
				)
			}
		}
		return
	}

	*outputs = append(*outputs, spec.OutputUnion{
		Kind:             spec.OutputKindReasoningMessage,
		ReasoningMessage: reasoning,
	})
}

func googleInteractionsResponseContainsText(
	outputs []spec.OutputUnion,
	text string,
) bool {
	for _, output := range outputs {
		if output.Kind != spec.OutputKindOutputMessage ||
			output.OutputMessage == nil {
			continue
		}
		if strings.Contains(googleInteractionsOutputUnionText(output), text) {
			return true
		}
	}
	return false
}

func googleInteractionsOutputUnionText(output spec.OutputUnion) string {
	if output.Kind != spec.OutputKindOutputMessage ||
		output.OutputMessage == nil {
		return ""
	}

	var text strings.Builder
	for _, content := range output.OutputMessage.Contents {
		if content.TextItem != nil {
			text.WriteString(content.TextItem.Text)
		}
	}
	return text.String()
}

func googleInteractionsContainsMatchingToolOutput(
	outputs []spec.OutputUnion,
	target spec.OutputUnion,
) bool {
	targetKind, targetCallID := googleInteractionsToolOutputIdentity(target)
	if targetKind == "" || targetCallID == "" {
		return false
	}

	for _, output := range outputs {
		kind, callID := googleInteractionsToolOutputIdentity(output)
		if kind == targetKind && callID == targetCallID {
			return true
		}
	}
	return false
}

func googleInteractionsToolOutputIdentity(
	output spec.OutputUnion,
) (targetKind spec.OutputKind, targetCallID string) {
	switch output.Kind {
	case spec.OutputKindFunctionToolCall:
		if output.FunctionToolCall != nil {
			return output.Kind, output.FunctionToolCall.CallID
		}
	case spec.OutputKindCustomToolCall:
		if output.CustomToolCall != nil {
			return output.Kind, output.CustomToolCall.CallID
		}
	case spec.OutputKindWebSearchToolCall:
		if output.WebSearchToolCall != nil {
			return output.Kind, output.WebSearchToolCall.CallID
		}
	case spec.OutputKindWebSearchToolOutput:
		if output.WebSearchToolOutput != nil {
			return output.Kind, output.WebSearchToolOutput.CallID
		}
	default:
	}
	return "", ""
}

func googleInteractionsStreamText(content any) string {
	switch content := content.(type) {
	case interactions.Content:
		text, _ := googleInteractionsTextContent(content)
		return text
	case *interactions.Content:
		if content == nil {
			return ""
		}
		text, _ := googleInteractionsTextContent(*content)
		return text
	case interactions.ThoughtSummaryContent:
		text, _ := googleInteractionsThoughtSummaryText(content)
		return text
	case *interactions.ThoughtSummaryContent:
		if content == nil {
			return ""
		}
		text, _ := googleInteractionsThoughtSummaryText(*content)
		return text
	default:
		return ""
	}
}

func (state *googleInteractionsStreamState) responseStatus() spec.Status {
	if state != nil && state.status != spec.StatusNone {
		return state.status
	}
	return spec.StatusIncomplete
}

func (state *googleInteractionsStreamState) responseUsage() *spec.Usage {
	if state != nil && state.usage != nil {
		return state.usage
	}
	return &spec.Usage{}
}
