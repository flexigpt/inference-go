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
