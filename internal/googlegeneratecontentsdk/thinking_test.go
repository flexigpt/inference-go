package googlegeneratecontentsdk

import (
	"testing"

	"google.golang.org/genai"

	"github.com/flexigpt/inference-go/spec"
)

func TestSanitizeGoogleGenerateContentReasoningInputs(t *testing.T) {
	t.Parallel()

	signed := thoughtSignatureToString([]byte("sig"))

	cases := []struct {
		name string
		in   []spec.InputUnion
		want int
	}{
		{
			name: "keeps non reasoning input",
			in: []spec.InputUnion{{
				Kind: spec.InputKindInputMessage,
				InputMessage: &spec.InputOutputContent{
					Role: spec.RoleUser,
					Contents: []spec.InputOutputContentItemUnion{{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: "hi"},
					}},
				},
			}},
			want: 1,
		},
		{
			name: "drops unsigned reasoning",
			in: []spec.InputUnion{{
				Kind: spec.InputKindReasoningMessage,
				ReasoningMessage: &spec.ReasoningContent{
					Thinking: []string{"plain text only"},
				},
			}},
			want: 0,
		},
		{
			name: "keeps signature only reasoning",
			in: []spec.InputUnion{{
				Kind: spec.InputKindReasoningMessage,
				ReasoningMessage: &spec.ReasoningContent{
					Signature: signed,
				},
			}},
			want: 1,
		},
		{
			name: "keeps signed reasoning with text",
			in: []spec.InputUnion{{
				Kind: spec.InputKindReasoningMessage,
				ReasoningMessage: &spec.ReasoningContent{
					Signature: signed,
					Thinking:  []string{"thought"},
				},
			}},
			want: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := sanitizeGoogleGenerateContentReasoningInputs(tc.in)
			if len(got) != tc.want {
				t.Fatalf("len(...) = %d, want %d", len(got), tc.want)
			}
		})
	}
}

func TestApplyGoogleGenerateContentThinkingPolicy_LevelNoneDisablesThinking(t *testing.T) {
	t.Parallel()

	cfg := &genai.GenerateContentConfig{}
	mp := &spec.ModelParam{
		Name: "gemini-2.5-flash",
		Reasoning: &spec.ReasoningParam{
			Type:  spec.ReasoningTypeSingleWithLevels,
			Level: spec.ReasoningLevelNone,
		},
	}

	if err := applyGoogleGenerateContentThinkingPolicy(
		cfg,
		mp,
		googleGenerateContentSDKCapability.ReasoningCapabilities,
	); err != nil {
		t.Fatal(err)
	}
	if cfg.ThinkingConfig == nil {
		t.Fatal("ThinkingConfig = nil, want disabled config")
	}
	if cfg.ThinkingConfig.IncludeThoughts {
		t.Fatal("IncludeThoughts = true, want false")
	}
	if cfg.ThinkingConfig.ThinkingBudget == nil || *cfg.ThinkingConfig.ThinkingBudget != 0 {
		t.Fatalf("ThinkingBudget = %#v, want 0", cfg.ThinkingConfig.ThinkingBudget)
	}
}
