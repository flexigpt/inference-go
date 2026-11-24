package googleinteractionssdk

import (
	"slices"
	"testing"

	"google.golang.org/genai/interactions/models/interactions"

	"github.com/flexigpt/inference-go/spec"
)

func TestApplyGoogleInteractionsThinkingConfig(t *testing.T) {
	omitted := spec.ReasoningSummaryStyleOmitted
	auto := spec.ReasoningSummaryStyleAuto
	concise := spec.ReasoningSummaryStyleConcise
	detailed := spec.ReasoningSummaryStyleDetailed

	tests := []struct {
		name            string
		reasoning       *spec.ReasoningParam
		wantConfigured  bool
		wantLevel       string
		wantSummary     string
		wantWarningCode []string
	}{
		{
			name: "nil reasoning",
		},
		{
			name: "default summary style",
			reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelLow,
			},
			wantConfigured: true,
			wantLevel:      "low",
			wantSummary:    "auto",
		},
		{
			name: "omitted summary style",
			reasoning: &spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelMedium,
				SummaryStyle: &omitted,
			},
			wantConfigured: true,
			wantLevel:      "medium",
			wantSummary:    "none",
		},
		{
			name: "auto summary style",
			reasoning: &spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelHigh,
				SummaryStyle: &auto,
			},
			wantConfigured: true,
			wantLevel:      "high",
			wantSummary:    "auto",
		},
		{
			name: "concise summary style narrows to auto",
			reasoning: &spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelXHigh,
				SummaryStyle: &concise,
			},
			wantConfigured: true,
			wantLevel:      "xhigh",
			wantSummary:    "auto",
			wantWarningCode: []string{
				"reasoning_summaryStyle_narrowed",
			},
		},
		{
			name: "detailed summary style narrows to auto",
			reasoning: &spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelMinimal,
				SummaryStyle: &detailed,
			},
			wantConfigured: true,
			wantLevel:      "minimal",
			wantSummary:    "auto",
			wantWarningCode: []string{
				"reasoning_summaryStyle_narrowed",
			},
		},
		{
			name: "hybrid token reasoning is unsupported",
			reasoning: &spec.ReasoningParam{
				Type:   spec.ReasoningTypeHybridWithTokens,
				Tokens: 1024,
			},
			wantWarningCode: []string{
				"reasoning_dropped_unsupported",
			},
		},
		{
			name: "unsupported level is rejected",
			reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelMax,
			},
			wantWarningCode: []string{
				"reasoning_dropped_invalid_level",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &interactions.GenerationConfig{}

			configured, warnings := applyGoogleInteractionsThinkingConfig(
				config,
				tc.reasoning,
			)
			if configured != tc.wantConfigured {
				t.Fatalf(
					"configured got %v want %v",
					configured,
					tc.wantConfigured,
				)
			}

			gotWarningCodes := make([]string, 0, len(warnings))
			for _, warning := range warnings {
				gotWarningCodes = append(gotWarningCodes, warning.Code)
			}
			if !slices.Equal(gotWarningCodes, tc.wantWarningCode) {
				t.Fatalf(
					"warning codes got %#v want %#v",
					gotWarningCodes,
					tc.wantWarningCode,
				)
			}

			if !tc.wantConfigured {
				if config.ThinkingLevel != nil ||
					config.ThinkingSummaries != nil {
					t.Fatalf("unexpected thinking config: %#v", config)
				}
				return
			}

			if config.ThinkingLevel == nil ||
				string(*config.ThinkingLevel) != tc.wantLevel {
				t.Fatalf(
					"thinking level got %#v want %q",
					config.ThinkingLevel,
					tc.wantLevel,
				)
			}
			if config.ThinkingSummaries == nil ||
				string(*config.ThinkingSummaries) != tc.wantSummary {
				t.Fatalf(
					"thinking summaries got %#v want %q",
					config.ThinkingSummaries,
					tc.wantSummary,
				)
			}
		})
	}
}
