package googlegeneratecontentsdk

import (
	"bytes"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestToGoogleGenerateContentContents_PreservesSignedReasoningBeforeToolCall(t *testing.T) {
	t.Parallel()

	sig := thoughtSignatureToString([]byte("sig-bytes"))

	inputs := []spec.InputUnion{
		{
			Kind: spec.InputKindReasoningMessage,
			ReasoningMessage: &spec.ReasoningContent{
				Role:      spec.RoleAssistant,
				Signature: sig,
			},
		},
		{
			Kind: spec.InputKindFunctionToolCall,
			FunctionToolCall: &spec.ToolCall{
				Type:      spec.ToolTypeFunction,
				Role:      spec.RoleAssistant,
				ID:        testCallID,
				CallID:    testCallID,
				Name:      testCallNameValue,
				Arguments: `{"text":"hi"}`,
			},
		},
	}

	contents, _, err := toGoogleGenerateContentContents(t.Context(), "", inputs)
	if err != nil {
		t.Fatal(err)
	}
	if len(contents) != 1 {
		t.Fatalf("len(contents) = %d, want 1", len(contents))
	}
	if len(contents[0].Parts) != 2 {
		t.Fatalf("len(parts) = %d, want 2", len(contents[0].Parts))
	}
	if !contents[0].Parts[0].Thought {
		t.Fatal("first part is not a thought")
	}

	wantSig, ok := decodeThoughtSignature(sig)
	if !ok {
		t.Fatal("decodeThoughtSignature failed")
	}
	if !bytes.Equal(contents[0].Parts[0].ThoughtSignature, wantSig) {
		t.Fatalf("ThoughtSignature = %q, want %q", contents[0].Parts[0].ThoughtSignature, wantSig)
	}
	if contents[0].Parts[1].FunctionCall == nil {
		t.Fatal("second part FunctionCall = nil")
	}
	if contents[0].Parts[1].FunctionCall.Name != testCallNameValue {
		t.Fatalf("FunctionCall.Name = %q, want echo_text", contents[0].Parts[1].FunctionCall.Name)
	}
}
