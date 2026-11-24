package googleinteractionssdk

import (
	"strings"
	"testing"

	"google.golang.org/genai/interactions/models/interactions"

	"github.com/flexigpt/inference-go/spec"
)

func TestGoogleInteractionsStreamTextAcceptsContentValueAndPointer(t *testing.T) {
	content := interactions.NewContent(
		interactions.TextContent{Text: "streamed thought"},
	)
	thoughtSummary := interactions.NewThoughtSummaryContent(
		interactions.TextContent{Text: "streamed thought summary"},
	)

	for _, input := range []any{content, &content} {
		if got := googleInteractionsStreamText(input); got != "streamed thought" {
			t.Fatalf("stream text got %q want %q", got, "streamed thought")
		}
	}

	for _, input := range []any{thoughtSummary, &thoughtSummary} {
		if got := googleInteractionsStreamText(input); got != "streamed thought summary" {
			t.Fatalf(
				"thought summary stream text got %q want %q",
				got,
				"streamed thought summary",
			)
		}
	}
}

func TestGoogleInteractionsStreamPartialResponsePreservesTextAndThought(t *testing.T) {
	state := newGoogleInteractionsStreamState()
	state.interactionID = "interaction_1"
	state.status = spec.StatusIncomplete
	entry := state.ensureStep(0, false)
	entry.text.WriteString("partial answer")
	entry.thoughtSummary.WriteString("partial thought summary")
	entry.thoughtSignature = "opaque-thought-signature"

	response := state.partialResponse()
	if response == nil || response.Usage == nil {
		t.Fatalf("partial response is incomplete: %#v", response)
	}
	if len(response.Outputs) != 2 {
		t.Fatalf("output count got %d want 2", len(response.Outputs))
	}
}

func TestGoogleInteractionsStreamUnfinishedContinuationUsesStatelessError(
	t *testing.T,
) {
	state := newGoogleInteractionsStreamState()
	state.setRawStatus("requires_continuation")

	err := state.unfinishedStreamError()
	if err == nil {
		t.Fatal("expected continuation error")
	}
	if !strings.Contains(err.Error(), "requires continuation") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGoogleInteractionsStreamFallbackDoesNotDuplicateSplitTerminalText(
	t *testing.T,
) {
	state := newGoogleInteractionsStreamState()
	state.interactionID = "interaction_1"
	state.status = spec.StatusCompleted
	state.ensureStep(0, false).text.WriteString("hello world")

	response := &spec.FetchCompletionResponse{
		Outputs: []spec.OutputUnion{{
			Kind: spec.OutputKindOutputMessage,
			OutputMessage: &spec.InputOutputContent{
				ID:     "interaction_1",
				Role:   spec.RoleAssistant,
				Status: spec.StatusCompleted,
				Contents: []spec.InputOutputContentItemUnion{
					{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: "hello"},
					},
					{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: " world"},
					},
				},
			},
		}},
	}

	appendGoogleInteractionsStreamFallback(response, state)
	if len(response.Outputs) != 1 {
		t.Fatalf("duplicate text fallback: %#v", response.Outputs)
	}
}

func TestGoogleInteractionsStreamOutputsPreserveRawArgumentFragments(
	t *testing.T,
) {
	state := newGoogleInteractionsStreamState()
	state.interactionID = "interaction_1"
	state.status = spec.StatusInProgress

	entry := state.ensureStep(0, false)
	entry.functionCall = &spec.ToolCall{
		Type:      spec.ToolTypeFunction,
		ChoiceID:  "echo-tool",
		ID:        "call_1",
		CallID:    "call_1",
		Name:      "echo",
		Arguments: "{}",
	}
	entry.functionArgs.WriteString(` { "text": "spaced value" } `)

	outputs, _ := state.streamOutputs()
	if len(outputs) != 1 || outputs[0].FunctionToolCall == nil {
		t.Fatalf("unexpected outputs: %#v", outputs)
	}
	if got := outputs[0].FunctionToolCall.Arguments; got != ` { "text": "spaced value" } ` {
		t.Fatalf("arguments got %q", got)
	}
}

func TestGoogleInteractionsStreamFallbackUsesTerminalMetadata(t *testing.T) {
	state := newGoogleInteractionsStreamState()
	state.interactionID = "interaction_1"
	state.status = spec.StatusCompleted
	state.usage = &spec.Usage{
		InputTokensTotal:  12,
		InputTokensCached: 4,
		OutputTokens:      5,
		ReasoningTokens:   3,
	}
	entry := state.ensureStep(0, false)
	entry.text.WriteString("streamed answer")
	entry.thoughtSummary.WriteString("streamed thought")

	response := &spec.FetchCompletionResponse{
		Usage: state.responseUsage(),
	}
	appendGoogleInteractionsStreamFallback(response, state)

	if response.Usage.InputTokensTotal != 12 ||
		response.Usage.OutputTokens != 5 {
		t.Fatalf("unexpected usage %#v", response.Usage)
	}
	if len(response.Outputs) != 2 {
		t.Fatalf("output count got %d want 2", len(response.Outputs))
	}
}

func TestGoogleInteractionsStreamOutputsPreserveToolAndGroundingSteps(t *testing.T) {
	state := newGoogleInteractionsStreamState()
	state.interactionID = "interaction_1"
	state.status = spec.StatusInProgress

	toolChoiceNameMap := map[string]spec.ToolChoice{
		"echo": {
			ID:   "echo-tool",
			Name: "echo",
			Type: spec.ToolTypeFunction,
		},
	}

	state.recordStep(
		0,
		interactions.NewStep(interactions.FunctionCallStep{
			ID:        "call_1",
			Name:      "echo",
			Arguments: map[string]any{},
		}),
		toolChoiceNameMap,
		"search",
		false,
	)
	state.steps[0].functionArgs.WriteString(`{"text":"streamed"}`)

	state.recordStep(
		1,
		interactions.NewStep(interactions.GoogleSearchCallStep{
			ID: "search_1",
			Arguments: interactions.GoogleSearchCallArguments{
				Queries: []string{"latest Go release"},
			},
		}),
		toolChoiceNameMap,
		"search",
		false,
	)

	outputs, _ := state.streamOutputs()
	if len(outputs) != 2 {
		t.Fatalf("output count got %d want 2", len(outputs))
	}
	if outputs[0].FunctionToolCall == nil ||
		outputs[0].FunctionToolCall.CallID != "call_1" ||
		outputs[0].FunctionToolCall.Arguments != `{"text":"streamed"}` {
		t.Fatalf("unexpected function call %#v", outputs[0])
	}
	if outputs[1].WebSearchToolCall == nil ||
		outputs[1].WebSearchToolCall.CallID != "search_1" ||
		len(outputs[1].WebSearchToolCall.WebSearchToolCallItems) != 1 {
		t.Fatalf("unexpected Google Search call %#v", outputs[1])
	}
}
