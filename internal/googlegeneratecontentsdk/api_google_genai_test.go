package googlegeneratecontentsdk

import (
	"testing"

	"google.golang.org/genai"

	"github.com/flexigpt/inference-go/spec"
)

func TestOutputsFromGenAIResponse_PreservesPartBoundariesAndSignatures(t *testing.T) {
	t.Parallel()

	textSig := []byte("text-sig")
	callSig := []byte("call-sig")

	resp := &genai.GenerateContentResponse{
		ResponseID: "resp-1",
		Candidates: []*genai.Candidate{{
			FinishReason: genai.FinishReasonStop,
			Content: &genai.Content{
				Role: genai.RoleModel,
				Parts: []*genai.Part{
					{Text: "first", ThoughtSignature: textSig},
					{Text: "second"},
					{
						FunctionCall: &genai.FunctionCall{
							ID:   "call-1",
							Name: "echo_text",
							Args: map[string]any{"text": "hi"},
						},
						ThoughtSignature: callSig,
					},
				},
			},
		}},
	}

	outs := outputsFromGenAIResponse(resp, map[string]spec.ToolChoice{
		"echo_text": {
			ID:   "echo-tool",
			Type: spec.ToolTypeFunction,
			Name: "echo_text",
		},
	}, "")

	if len(outs) != 3 {
		t.Fatalf("len(outputs) = %d, want 3", len(outs))
	}

	if outs[0].Kind != spec.OutputKindOutputMessage || outs[0].OutputMessage == nil {
		t.Fatalf("output[0] = %#v, want output message", outs[0])
	}
	if got := outs[0].OutputMessage.Contents[0].TextItem.Signature; got != thoughtSignatureToString(textSig) {
		t.Fatalf("text signature = %q, want %q", got, thoughtSignatureToString(textSig))
	}

	if outs[1].Kind != spec.OutputKindOutputMessage || outs[1].OutputMessage == nil {
		t.Fatalf("output[1] = %#v, want output message", outs[1])
	}
	if got := outs[1].OutputMessage.Contents[0].TextItem.Signature; got != "" {
		t.Fatalf("text signature = %q, want empty", got)
	}

	if outs[2].Kind != spec.OutputKindFunctionToolCall || outs[2].FunctionToolCall == nil {
		t.Fatalf("output[2] = %#v, want function tool call", outs[2])
	}
	if got := outs[2].FunctionToolCall.Signature; got != thoughtSignatureToString(callSig) {
		t.Fatalf("tool call signature = %q, want %q", got, thoughtSignatureToString(callSig))
	}
}
