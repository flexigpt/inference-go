package anthropicsdk

import (
	"testing"

	"github.com/anthropics/anthropic-sdk-go"

	"github.com/flexigpt/inference-go/spec"
)

func TestAnthropicAdaptiveThinkingSummaryStyle(t *testing.T) {
	omitted := spec.ReasoningSummaryStyleOmitted
	auto := spec.ReasoningSummaryStyleAuto
	concise := spec.ReasoningSummaryStyleConcise
	detailed := spec.ReasoningSummaryStyleDetailed

	tests := []struct {
		name         string
		summaryStyle *spec.ReasoningSummaryStyle
		wantDisplay  string
	}{
		{
			name:        "default is summarized",
			wantDisplay: "summarized",
		},
		{
			name:         "omitted maps to omitted",
			summaryStyle: &omitted,
			wantDisplay:  "omitted",
		},
		{
			name:         "auto maps to summarized",
			summaryStyle: &auto,
			wantDisplay:  "summarized",
		},
		{
			name:         "concise maps to summarized",
			summaryStyle: &concise,
			wantDisplay:  "summarized",
		},
		{
			name:         "detailed maps to summarized",
			summaryStyle: &detailed,
			wantDisplay:  "summarized",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			params := anthropic.MessageNewParams{
				MaxTokens: 4096,
			}
			modelParam := &spec.ModelParam{
				Name: "claude-test",
				Reasoning: &spec.ReasoningParam{
					Type:         spec.ReasoningTypeSingleWithLevels,
					Level:        spec.ReasoningLevelMedium,
					SummaryStyle: tc.summaryStyle,
				},
			}

			applyAnthropicThinkingPolicy(
				&params,
				modelParam,
				anthropicThinkingAnalysis{},
			)

			if params.Thinking.OfAdaptive == nil {
				t.Fatalf("expected adaptive thinking config, got %#v", params.Thinking)
			}
			if got := string(params.Thinking.OfAdaptive.Display); got != tc.wantDisplay {
				t.Fatalf("adaptive display got %q want %q", got, tc.wantDisplay)
			}
		})
	}
}

func TestAnthropicReasoningCapabilities(t *testing.T) {
	caps := anthropicsdkCapability.ReasoningCapabilities
	if caps == nil {
		t.Fatal("expected Anthropic reasoning capabilities")
	}
	if !caps.SupportsSummaryStyle {
		t.Fatal("expected Anthropic summary-style support")
	}
	if caps.SupportsReasoningContext {
		t.Fatal("Anthropic must not support reasoning context")
	}
	if caps.SupportsReasoningMode {
		t.Fatal("Anthropic must not support reasoning mode")
	}
}
