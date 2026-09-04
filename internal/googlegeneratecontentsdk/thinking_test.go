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

func TestApplyGoogleGenerateContentThinkingPolicySummaryStyle(t *testing.T) {
	omitted := spec.ReasoningSummaryStyleOmitted
	auto := spec.ReasoningSummaryStyleAuto
	concise := spec.ReasoningSummaryStyleConcise
	detailed := spec.ReasoningSummaryStyleDetailed

	tests := []struct {
		name         string
		reasoning    spec.ReasoningParam
		wantIncluded bool
	}{
		{
			name: "level reasoning default includes thoughts",
			reasoning: spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelMedium,
			},
			wantIncluded: true,
		},
		{
			name: "level reasoning omitted hides thoughts",
			reasoning: spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelMedium,
				SummaryStyle: &omitted,
			},
			wantIncluded: false,
		},
		{
			name: "level reasoning auto includes thoughts",
			reasoning: spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelMedium,
				SummaryStyle: &auto,
			},
			wantIncluded: true,
		},
		{
			name: "level reasoning concise includes thoughts",
			reasoning: spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelMedium,
				SummaryStyle: &concise,
			},
			wantIncluded: true,
		},
		{
			name: "level reasoning detailed includes thoughts",
			reasoning: spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelMedium,
				SummaryStyle: &detailed,
			},
			wantIncluded: true,
		},
		{
			name: "token budget default includes thoughts",
			reasoning: spec.ReasoningParam{
				Type:   spec.ReasoningTypeHybridWithTokens,
				Tokens: 1024,
			},
			wantIncluded: true,
		},
		{
			name: "token budget omitted hides thoughts",
			reasoning: spec.ReasoningParam{
				Type:         spec.ReasoningTypeHybridWithTokens,
				Tokens:       1024,
				SummaryStyle: &omitted,
			},
			wantIncluded: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &genai.GenerateContentConfig{}
			modelParam := &spec.ModelParam{
				Name:      "gemini-test",
				Reasoning: &tc.reasoning,
			}

			err := applyGoogleGenerateContentThinkingPolicy(
				config,
				modelParam,
				googleGenerateContentSDKCapability.ReasoningCapabilities,
			)
			if err != nil {
				t.Fatalf("applyGoogleGenerateContentThinkingPolicy: %v", err)
			}
			if config.ThinkingConfig == nil {
				t.Fatal("expected ThinkingConfig")
			}
			if config.ThinkingConfig.IncludeThoughts != tc.wantIncluded {
				t.Fatalf(
					"IncludeThoughts got %v want %v",
					config.ThinkingConfig.IncludeThoughts,
					tc.wantIncluded,
				)
			}
		})
	}
}

func TestGoogleGenerateContentReasoningCapabilities(t *testing.T) {
	caps := googleGenerateContentSDKCapability.ReasoningCapabilities
	if caps == nil {
		t.Fatal("expected Google reasoning capabilities")
	}
	if !caps.SupportsSummaryStyle {
		t.Fatal("expected Google summary-style support")
	}
	if caps.SupportsReasoningContext {
		t.Fatal("Google must not support reasoning context")
	}
	if caps.SupportsReasoningMode {
		t.Fatal("Google must not support reasoning mode")
	}
}
