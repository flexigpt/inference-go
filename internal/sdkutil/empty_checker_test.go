package sdkutil

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestIsInputUnionEmpty_ReasoningContent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   spec.InputUnion
		want bool
	}{
		{
			name: "nil reasoning is empty",
			in: spec.InputUnion{
				Kind: spec.InputKindReasoningMessage,
			},
			want: true,
		},
		{
			name: "signature only reasoning is not empty",
			in: spec.InputUnion{
				Kind: spec.InputKindReasoningMessage,
				ReasoningMessage: &spec.ReasoningContent{
					Signature: "c2ln",
				},
			},
			want: false,
		},
		{
			name: "thinking reasoning is not empty",
			in: spec.InputUnion{
				Kind: spec.InputKindReasoningMessage,
				ReasoningMessage: &spec.ReasoningContent{
					Thinking: []string{"hi"},
				},
			},
			want: false,
		},
		{
			name: "signature only text item is not empty",
			in: spec.InputUnion{
				Kind: spec.InputKindOutputMessage,
				OutputMessage: &spec.InputOutputContent{
					Role: spec.RoleAssistant,
					Contents: []spec.InputOutputContentItemUnion{{
						Kind: spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{
							Signature: "c2ln",
						},
					}},
				},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := IsInputUnionEmpty(tc.in); got != tc.want {
				t.Fatalf("IsInputUnionEmpty() = %v, want %v", got, tc.want)
			}
		})
	}
}
