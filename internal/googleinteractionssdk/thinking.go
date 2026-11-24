package googleinteractionssdk

import (
	"strings"

	"google.golang.org/genai/interactions/models/interactions"

	"github.com/flexigpt/inference-go/spec"
)

func sanitizeGoogleInteractionsReasoningInputs(
	inputs []spec.InputUnion,
) []spec.InputUnion {
	if len(inputs) == 0 {
		return nil
	}

	out := make([]spec.InputUnion, 0, len(inputs))
	for _, input := range inputs {
		if input.Kind != spec.InputKindReasoningMessage {
			out = append(out, input)
			continue
		}

		if input.ReasoningMessage == nil ||
			strings.TrimSpace(input.ReasoningMessage.Signature) == "" {
			continue
		}

		// Interactions replays provider-issued thought signatures and optional
		// summaries. It does not accept plaintext, redacted, or encrypted
		// reasoning as a portable thought input.
		reasoning := *input.ReasoningMessage
		reasoning.Thinking = nil
		reasoning.RedactedThinking = nil
		reasoning.EncryptedContent = nil
		input.ReasoningMessage = &reasoning
		out = append(out, input)
	}

	return out
}

func applyGoogleInteractionsThinkingConfig(
	config *interactions.GenerationConfig,
	reasoning *spec.ReasoningParam,
) (bool, []spec.Warning) {
	if config == nil || reasoning == nil {
		return false, nil
	}

	if reasoning.Type != spec.ReasoningTypeSingleWithLevels {
		return false, []spec.Warning{{
			Code:    "reasoning_dropped_unsupported",
			Message: "Interactions supports level-based reasoning, not token-budget reasoning.",
		}}
	}

	switch reasoning.Level {
	case spec.ReasoningLevelMinimal,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh:
	default:
		return false, []spec.Warning{{
			Code:    "reasoning_dropped_invalid_level",
			Message: "Interactions supports minimal, low, medium, high, and xhigh reasoning levels.",
		}}
	}

	level := interactions.ThinkingLevel(reasoning.Level)
	config.ThinkingLevel = level.ToPointer()

	summary := interactions.ThinkingSummariesAuto
	var warnings []spec.Warning

	if reasoning.SummaryStyle != nil {
		switch *reasoning.SummaryStyle {
		case spec.ReasoningSummaryStyleOmitted:
			summary = interactions.ThinkingSummariesNone
		case spec.ReasoningSummaryStyleAuto:
			summary = interactions.ThinkingSummariesAuto
		case spec.ReasoningSummaryStyleConcise,
			spec.ReasoningSummaryStyleDetailed:
			warnings = appendGoogleInteractionsWarning(
				warnings,
				"reasoning_summaryStyle_narrowed",
				"Interactions maps concise and detailed reasoning summaries to auto.",
			)
		default:
			warnings = appendGoogleInteractionsWarning(
				warnings,
				"reasoning_summaryStyle_narrowed",
				"Interactions maps unknown reasoning summary styles to auto.",
			)
		}
	}

	config.ThinkingSummaries = summary.ToPointer()
	return true, warnings
}

func googleInteractionsThoughtSummaryContents(
	summary []string,
) []interactions.ThoughtSummaryContent {
	out := make([]interactions.ThoughtSummaryContent, 0, len(summary))
	for _, text := range summary {
		if strings.TrimSpace(text) == "" {
			continue
		}
		out = append(out, interactions.NewThoughtSummaryContent(
			interactions.TextContent{Text: text},
		))
	}
	return out
}

func googleInteractionsOptionalString(value string) *string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	return new(value)
}
