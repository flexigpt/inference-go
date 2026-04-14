package googlegeneratecontentsdk

import (
	"bytes"
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

func TestConsolidateGoogleGenerateContentStreamParts(t *testing.T) {
	t.Parallel()

	sig1 := []byte("sig-1")
	sig2 := []byte("sig-2")
	fc := &genai.FunctionCall{Name: "echo_text", Args: map[string]any{"text": "hi"}}

	tests := []struct {
		name     string
		in       []*genai.Part
		wantLen  int
		validate func(t *testing.T, out []*genai.Part)
	}{
		{
			name:    "nil input",
			in:      nil,
			wantLen: 0,
		},
		{
			name:    "empty input",
			in:      []*genai.Part{},
			wantLen: 0,
		},
		{
			name:    "single text part unchanged",
			in:      []*genai.Part{{Text: "hello"}},
			wantLen: 1,
			validate: func(t *testing.T, out []*genai.Part) {
				t.Helper()
				if out[0].Text != "hello" {
					t.Errorf("text = %q, want hello", out[0].Text)
				}
			},
		},
		{
			name: "adjacent text parts merged with last signature",
			in: []*genai.Part{
				{Text: "hello"},
				{Text: " world"},
				{Text: "!", ThoughtSignature: sig1},
			},
			wantLen: 1,
			validate: func(t *testing.T, out []*genai.Part) {
				t.Helper()
				if out[0].Text != "hello world!" {
					t.Errorf("text = %q, want 'hello world!'", out[0].Text)
				}
				if !bytes.Equal(out[0].ThoughtSignature, sig1) {
					t.Errorf("ThoughtSignature mismatch: got %q, want %q", out[0].ThoughtSignature, sig1)
				}
			},
		},
		{
			name: "adjacent thought parts merged; last sig wins",
			in: []*genai.Part{
				{Thought: true, Text: "thinking..."},
				{Thought: true, Text: "more thinking", ThoughtSignature: sig1},
				{Thought: true, ThoughtSignature: sig2}, // sig-only final part
			},
			wantLen: 1,
			validate: func(t *testing.T, out []*genai.Part) {
				t.Helper()
				if !out[0].Thought {
					t.Error("want Thought=true")
				}
				if out[0].Text != "thinking...more thinking" {
					t.Errorf("text = %q, want 'thinking...more thinking'", out[0].Text)
				}
				if !bytes.Equal(out[0].ThoughtSignature, sig2) {
					t.Errorf("ThoughtSignature = %q, want sig2", out[0].ThoughtSignature)
				}
			},
		},
		{
			name: "thought then text become separate groups",
			in: []*genai.Part{
				{Thought: true, Text: "thinking"},
				{Text: "answer"},
			},
			wantLen: 2,
			validate: func(t *testing.T, out []*genai.Part) {
				t.Helper()
				if !out[0].Thought || out[0].Text != "thinking" {
					t.Errorf("[0]: Thought=%v Text=%q", out[0].Thought, out[0].Text)
				}
				if out[1].Thought || out[1].Text != "answer" {
					t.Errorf("[1]: Thought=%v Text=%q", out[1].Thought, out[1].Text)
				}
			},
		},
		{
			name: "function call is a boundary; adjacent text groups stay separate",
			in: []*genai.Part{
				{Text: "before1"},
				{Text: "before2"},
				{FunctionCall: fc},
				{Text: "after1"},
				{Text: "after2"},
			},
			wantLen: 3,
			validate: func(t *testing.T, out []*genai.Part) {
				t.Helper()
				if out[0].Text != "before1before2" {
					t.Errorf("[0].Text = %q", out[0].Text)
				}
				if out[1].FunctionCall == nil || out[1].FunctionCall.Name != "echo_text" {
					t.Errorf("[1] FunctionCall = %v", out[1].FunctionCall)
				}
				if out[2].Text != "after1after2" {
					t.Errorf("[2].Text = %q", out[2].Text)
				}
			},
		},
		{
			name: "full streaming scenario: thought chunks + text chunks + trailing sigs",
			in: []*genai.Part{
				{Thought: true, Text: "chunk 1"},
				{Thought: true, Text: " chunk 2"},
				{Thought: true, ThoughtSignature: sig1}, // final thought sig on sig-only part
				{Text: "answer part 1"},
				{Text: " answer part 2", ThoughtSignature: sig2}, // text sig on last text chunk
			},
			wantLen: 2,
			validate: func(t *testing.T, out []*genai.Part) {
				t.Helper()
				if !out[0].Thought {
					t.Error("[0]: want Thought=true")
				}
				if out[0].Text != "chunk 1 chunk 2" {
					t.Errorf("[0].Text = %q, want 'chunk 1 chunk 2'", out[0].Text)
				}
				if !bytes.Equal(out[0].ThoughtSignature, sig1) {
					t.Errorf("[0].ThoughtSignature mismatch")
				}

				if out[1].Thought {
					t.Error("[1]: want Thought=false")
				}
				if out[1].Text != "answer part 1 answer part 2" {
					t.Errorf("[1].Text = %q, want 'answer part 1 answer part 2'", out[1].Text)
				}
				if !bytes.Equal(out[1].ThoughtSignature, sig2) {
					t.Errorf("[1].ThoughtSignature mismatch")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := consolidateGoogleGenerateContentStreamParts(tc.in)
			if len(out) != tc.wantLen {
				t.Fatalf("len(out) = %d, want %d", len(out), tc.wantLen)
			}
			if tc.validate != nil {
				tc.validate(t, out)
			}
		})
	}
}
