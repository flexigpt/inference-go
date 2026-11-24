package googleinteractionssdk

import (
	"context"
	"errors"
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

	state := &googleInteractionsStreamState{}
	var streamErr error

	for stream.Next() {
		event := stream.Value()
		if event == nil {
			continue
		}

		stepDelta := event.GetDataStepDelta()
		if stepDelta != nil {
			textDelta := stepDelta.GetDeltaText()
			if textDelta != nil {
				text := googleInteractionsString(textDelta.GetText())
				state.text.WriteString(text)
				if err := writeText(text); err != nil {
					streamErr = err
					break
				}
			}

			thoughtSummaryDelta := stepDelta.GetDeltaThoughtSummary()
			if thoughtSummaryDelta != nil {
				text := googleInteractionsStreamText(
					thoughtSummaryDelta.GetContent(),
				)
				state.thinking.WriteString(text)
				if err := writeThinking(text); err != nil {
					streamErr = err
					break
				}
			}

			thoughtSignatureDelta := stepDelta.GetDeltaThoughtSignature()
			if thoughtSignatureDelta != nil {
				state.thoughtSignature = googleInteractionsString(
					thoughtSignatureDelta.GetSignature(),
				)
			}
		}

		if completed := event.GetDataInteractionCompleted(); completed != nil {
			in := completed.GetInteraction()
			state.response, state.responseErr = responseFromGoogleInteractionSSE(
				&in,
				toolChoiceNameMap,
				webSearchChoiceID,
			)
			state.terminal = true
			break
		}

		if event.GetDataError() != nil {
			state.eventErr = errors.New(
				"google interactions: stream emitted an error event",
			)
			state.terminal = true
			streamErr = state.eventErr
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
		streamErr = errors.New(
			"google interactions: stream ended before interaction.completed",
		)
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
	text             strings.Builder
	thinking         strings.Builder
	thoughtSignature string
	terminal         bool
	eventErr         error
	response         *spec.FetchCompletionResponse
	responseErr      error
}

func (state *googleInteractionsStreamState) partialResponse() *spec.FetchCompletionResponse {
	response := &spec.FetchCompletionResponse{
		Usage: &spec.Usage{},
	}

	if state.thoughtSignature != "" || state.thinking.Len() > 0 {
		reasoning := &spec.ReasoningContent{
			Role:      spec.RoleAssistant,
			Status:    spec.StatusIncomplete,
			Signature: state.thoughtSignature,
		}
		if thought := state.thinking.String(); thought != "" {
			reasoning.Summary = []string{thought}
		}
		response.Outputs = append(response.Outputs, spec.OutputUnion{
			Kind:             spec.OutputKindReasoningMessage,
			ReasoningMessage: reasoning,
		})
	}

	if text := state.text.String(); text != "" {
		response.Outputs = append(response.Outputs, spec.OutputUnion{
			Kind: spec.OutputKindOutputMessage,
			OutputMessage: &spec.InputOutputContent{
				Role:   spec.RoleAssistant,
				Status: spec.StatusIncomplete,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: text},
				}},
			},
		})
	}

	return response
}

func appendGoogleInteractionsStreamFallback(
	response *spec.FetchCompletionResponse,
	state *googleInteractionsStreamState,
) {
	if response == nil || state == nil {
		return
	}

	if !hasGoogleInteractionsReasoningOutput(response.Outputs) &&
		(state.thoughtSignature != "" || state.thinking.Len() > 0) {
		reasoning := &spec.ReasoningContent{
			Role:      spec.RoleAssistant,
			Status:    spec.StatusIncomplete,
			Signature: state.thoughtSignature,
		}
		if thought := state.thinking.String(); thought != "" {
			reasoning.Summary = []string{thought}
		}
		response.Outputs = append(response.Outputs, spec.OutputUnion{
			Kind:             spec.OutputKindReasoningMessage,
			ReasoningMessage: reasoning,
		})
	}

	if !hasGoogleInteractionsTextOutput(response.Outputs) &&
		state.text.Len() > 0 {
		response.Outputs = append(response.Outputs, spec.OutputUnion{
			Kind: spec.OutputKindOutputMessage,
			OutputMessage: &spec.InputOutputContent{
				Role:   spec.RoleAssistant,
				Status: spec.StatusIncomplete,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: state.text.String()},
				}},
			},
		})
	}
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
