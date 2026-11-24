package googleinteractionssdk

import (
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
	state := &googleInteractionsStreamState{}
	state.text.WriteString("partial answer")
	state.thinking.WriteString("partial thought summary")
	state.thoughtSignature = "opaque-thought-signature"

	response := state.partialResponse()
	if response == nil || response.Usage == nil {
		t.Fatalf("partial response is incomplete: %#v", response)
	}
	if len(response.Outputs) != 2 {
		t.Fatalf("output count got %d want 2", len(response.Outputs))
	}

	reasoning := response.Outputs[0].ReasoningMessage
	if response.Outputs[0].Kind != spec.OutputKindReasoningMessage ||
		reasoning == nil ||
		reasoning.Signature != "opaque-thought-signature" ||
		len(reasoning.Summary) != 1 ||
		reasoning.Summary[0] != "partial thought summary" {
		t.Fatalf("unexpected reasoning output %#v", response.Outputs[0])
	}

	message := response.Outputs[1].OutputMessage
	if response.Outputs[1].Kind != spec.OutputKindOutputMessage ||
		message == nil ||
		len(message.Contents) != 1 ||
		message.Contents[0].TextItem == nil ||
		message.Contents[0].TextItem.Text != "partial answer" {
		t.Fatalf("unexpected text output %#v", response.Outputs[1])
	}
}

func TestGoogleInteractionsStreamFallbackUsesTerminalMetadata(t *testing.T) {
	state := &googleInteractionsStreamState{
		interactionID: "interaction_1",
		status:        spec.StatusCompleted,
		usage: &spec.Usage{
			InputTokensTotal:  12,
			InputTokensCached: 4,
			OutputTokens:      5,
			ReasoningTokens:   3,
		},
	}
	state.text.WriteString("streamed answer")
	state.thinking.WriteString("streamed thought")

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

	reasoning := response.Outputs[0].ReasoningMessage
	if reasoning == nil ||
		reasoning.ID != "interaction_1" ||
		reasoning.Status != spec.StatusCompleted {
		t.Fatalf("unexpected reasoning output %#v", reasoning)
	}

	message := response.Outputs[1].OutputMessage
	if message == nil ||
		message.ID != "interaction_1" ||
		message.Status != spec.StatusCompleted ||
		message.Contents[0].TextItem == nil ||
		message.Contents[0].TextItem.Text != "streamed answer" {
		t.Fatalf("unexpected output message %#v", message)
	}
}
