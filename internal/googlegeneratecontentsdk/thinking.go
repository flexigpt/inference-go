package googlegeneratecontentsdk

import (
	"encoding/base64"
	"strings"

	"google.golang.org/genai"

	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

// applyGoogleGenerateContentThinkingPolicy sets ThinkingConfig in GenerateContentConfig
// based on ModelParam.Reasoning.
//
//   - ReasoningTypeHybridWithTokens  → fixed ThinkingBudget (token count).
//   - ReasoningTypeSingleWithLevels  → qualitative ThinkingLevel; None means
//     no ThinkingConfig (thinking disabled).
func applyGoogleGenerateContentThinkingPolicy(
	config *genai.GenerateContentConfig,
	mp *spec.ModelParam,
) {
	if config == nil || mp == nil || mp.Reasoning == nil {
		return
	}
	rp := mp.Reasoning

	switch rp.Type {
	case spec.ReasoningTypeHybridWithTokens:
		tokens := sdkutil.ClampIntToInt32(rp.Tokens)
		budget := max(tokens, 1)
		config.ThinkingConfig = &genai.ThinkingConfig{
			IncludeThoughts: true,
			ThinkingBudget:  &budget,
		}

	case spec.ReasoningTypeSingleWithLevels:
		level, ok := googleThinkingLevelFromSpec(rp.Level)
		if !ok {
			// Spec.ReasoningLevelNone or unknown → no ThinkingConfig.
			return
		}
		config.ThinkingConfig = &genai.ThinkingConfig{
			IncludeThoughts: true,
			ThinkingLevel:   level,
		}

	default:
		// Unknown reasoning type → no ThinkingConfig.
	}
}

// googleThinkingLevelFromSpec maps a spec.ReasoningLevel to the
// genai.ThinkingLevel constant.  Returns (level, false) for None/unknown.
func googleThinkingLevelFromSpec(level spec.ReasoningLevel) (genai.ThinkingLevel, bool) {
	switch level {
	case spec.ReasoningLevelNone:
		return "", false
	case spec.ReasoningLevelMinimal:
		return genai.ThinkingLevelMinimal, true
	case spec.ReasoningLevelLow:
		return genai.ThinkingLevelLow, true
	case spec.ReasoningLevelMedium:
		return genai.ThinkingLevelMedium, true
	case spec.ReasoningLevelHigh:
		return genai.ThinkingLevelHigh, true
	case spec.ReasoningLevelXHigh:
		// The genai SDK has no XHigh level; map to the highest available.
		return genai.ThinkingLevelHigh, true
	default:
		return "", false
	}
}

// sanitizeGoogleGenerateContentReasoningInputs enforces the Google GenAI policy for
// reasoning history pass-back:
//
//   - Keep only Google-native signed thoughts: entries where both a non-empty
//     Thinking text and a non-empty Signature are present (the model's own
//     thought + ThoughtSignature bytes, base64-encoded into Signature).
//   - Drop everything else: Anthropic's RedactedThinking, OpenAI's
//     EncryptedContent, or any unsigned/plain-text reasoning content from a
//     different provider.
func sanitizeGoogleGenerateContentReasoningInputs(inputs []spec.InputUnion) []spec.InputUnion {
	if len(inputs) == 0 {
		return nil
	}

	out := make([]spec.InputUnion, 0, len(inputs))
	dropped := 0

	for _, in := range inputs {
		if in.Kind != spec.InputKindReasoningMessage {
			out = append(out, in)
			continue
		}

		if sdkutil.IsInputUnionEmpty(in) || in.ReasoningMessage == nil {
			dropped++
			continue
		}

		if !isGoogleNativeReasoning(in.ReasoningMessage) {
			dropped++
			continue
		}

		out = append(out, in)
	}

	if dropped > 0 {
		logutil.Debug(
			"googleGenerateContent: sanitized non-native reasoning messages from input history",
			"dropped", dropped,
		)
	}

	return out
}

// isGoogleNativeReasoning returns true when the ReasoningContent carries a
// Google-native thought: at least one non-empty Thinking string AND a
// non-empty Signature (base64-encoded ThoughtSignature bytes).
func isGoogleNativeReasoning(r *spec.ReasoningContent) bool {
	if r == nil || strings.TrimSpace(r.Signature) == "" {
		return false
	}
	for _, t := range r.Thinking {
		if strings.TrimSpace(t) != "" {
			return true
		}
	}
	return false
}

// thoughtSignatureToString base64-encodes the raw ThoughtSignature bytes
// returned by the genai SDK into a plain string suitable for
// spec.ReasoningContent.Signature.
func thoughtSignatureToString(sig []byte) string {
	if len(sig) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(sig)
}

// thoughtSignatureFromString decodes a base64-encoded Signature string back
// to the raw bytes required by genai.Part.ThoughtSignature.
// If the string is not valid base64, it falls back to a direct []byte cast.
func thoughtSignatureFromString(sig string) []byte {
	sig = strings.TrimSpace(sig)
	if sig == "" {
		return nil
	}
	b, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return []byte(sig)
	}
	return b
}
