package googlegeneratecontentsdk

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/genai"

	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

func resolveGoogleGenerateContentReasoningCapabilities(
	ctx context.Context,
	opts *spec.FetchCompletionOptions,
	modelName spec.ModelName,
) *spec.ReasoningCapabilities {
	if opts == nil || opts.CapabilityResolver == nil || strings.TrimSpace(string(modelName)) == "" {
		return nil
	}

	caps, err := opts.CapabilityResolver.ResolveModelCapabilities(ctx, spec.ResolveModelCapabilitiesRequest{
		ProviderSDKType: spec.ProviderSDKTypeGoogleGenerateContent,
		ModelName:       modelName,
		CompletionKey:   opts.CompletionKey,
	})
	if err != nil {
		logutil.Debug(
			"googleGenerateContent: failed to resolve model capabilities for reasoning policy",
			"model", modelName,
			"err", err,
		)
		return nil
	}
	if caps == nil {
		return nil
	}
	return caps.ReasoningCapabilities
}

// applyGoogleGenerateContentThinkingPolicy sets ThinkingConfig in GenerateContentConfig
// based on ModelParam.Reasoning and optional model-specific capabilities.
func applyGoogleGenerateContentThinkingPolicy(
	config *genai.GenerateContentConfig,
	mp *spec.ModelParam,
	caps *spec.ReasoningCapabilities,
) error {
	if config == nil || mp == nil || mp.Reasoning == nil {
		return nil
	}
	rp := mp.Reasoning

	switch rp.Type {
	case spec.ReasoningTypeHybridWithTokens:
		budget, err := validateGoogleThinkingBudget(rp.Tokens, caps, mp.Name)
		if err != nil {
			return err
		}
		if budget == 0 {
			config.ThinkingConfig = disabledGoogleGenerateContentThinkingConfig()
			return nil
		}
		config.ThinkingConfig = &genai.ThinkingConfig{
			IncludeThoughts: true,
			ThinkingBudget:  &budget,
		}

	case spec.ReasoningTypeSingleWithLevels:
		if rp.Level == spec.ReasoningLevelNone {
			config.ThinkingConfig = disabledGoogleGenerateContentThinkingConfig()
			return nil
		}
		level, ok := googleThinkingLevelFromSpec(rp.Level)
		if !ok {
			return nil
		}
		config.ThinkingConfig = &genai.ThinkingConfig{
			IncludeThoughts: true,
			ThinkingLevel:   level,
		}
	default:
		// Unknown reasoning type → no ThinkingConfig.
	}
	return nil
}

func disabledGoogleGenerateContentThinkingConfig() *genai.ThinkingConfig {
	zero := int32(0)
	return &genai.ThinkingConfig{
		IncludeThoughts: false,
		ThinkingBudget:  &zero,
	}
}

func validateGoogleThinkingBudget(
	tokens int,
	caps *spec.ReasoningCapabilities,
	modelName spec.ModelName,
) (int32, error) {
	budget := sdkutil.ClampIntToInt32(tokens)

	var budgetCaps *spec.ReasoningTokenBudgetCapabilities
	if caps != nil {
		budgetCaps = caps.HybridTokenBudgetCapabilities
	}

	switch {
	case budget < -1:
		return 0, fmt.Errorf("googleGenerateContent: reasoning tokens must be >= -1 for model %q", modelName)
	case budget == -1:
		if budgetCaps != nil && !budgetCaps.MinusOneAllowed {
			return 0, fmt.Errorf("googleGenerateContent: reasoning tokens=-1 is not allowed for model %q", modelName)
		}
	case budget == 0:
		if budgetCaps != nil && !budgetCaps.ZeroAllowed {
			return 0, fmt.Errorf("googleGenerateContent: reasoning tokens=0 is not allowed for model %q", modelName)
		}
	default:
		if budgetCaps != nil && budgetCaps.MinAllowed > 0 && int(budget) < budgetCaps.MinAllowed {
			return 0, fmt.Errorf(
				"googleGenerateContent: reasoning tokens=%d is below minAllowed=%d for model %q",
				budget,
				budgetCaps.MinAllowed,
				modelName,
			)
		}
		if budgetCaps != nil && budgetCaps.MaxAllowed > 0 && int(budget) > budgetCaps.MaxAllowed {
			return 0, fmt.Errorf(
				"googleGenerateContent: reasoning tokens=%d exceeds maxAllowed=%d for model %q",
				budget,
				budgetCaps.MaxAllowed,
				modelName,
			)
		}
	}

	return budget, nil
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
	case spec.ReasoningLevelHigh, spec.ReasoningLevelXHigh, spec.ReasoningLevelMax:
		return genai.ThinkingLevelHigh, true
	default:
		return "", false
	}
}

// sanitizeGoogleGenerateContentReasoningInputs enforces the Google GenAI policy for
// reasoning history pass-back:
//
//   - Keep only Google-native signed thoughts: entries with a valid non-empty
//     Signature. Visible thought text may be empty; signature-only reasoning
//     parts still need to be passed back.
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

		if in.ReasoningMessage == nil {
			dropped++
			continue
		}

		if _, ok := decodeThoughtSignature(in.ReasoningMessage.Signature); !ok {
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

// thoughtSignatureToString base64-encodes the raw ThoughtSignature bytes
// returned by the genai SDK into a plain string suitable for
// spec.ReasoningContent.Signature.
func thoughtSignatureToString(sig []byte) string {
	if len(sig) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(sig)
}

func decodeThoughtSignature(sig string) ([]byte, bool) {
	sig = strings.TrimSpace(sig)
	if sig == "" {
		return nil, false
	}
	if b, err := base64.StdEncoding.DecodeString(sig); err == nil && len(b) > 0 {
		return b, true
	}

	if b, err := base64.RawStdEncoding.DecodeString(sig); err == nil && len(b) > 0 {
		return b, true
	}
	return nil, false
}
