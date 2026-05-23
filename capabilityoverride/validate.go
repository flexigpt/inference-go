package capabilityoverride

import (
	"errors"
	"fmt"
	"strings"

	"github.com/flexigpt/inference-go/spec"
)

func ValidateModelCapabilitiesOverride(o *ModelCapabilitiesOverride) error {
	if o == nil {
		return nil
	}

	if err := validateInputModalities(o.ModalitiesIn); err != nil {
		return fmt.Errorf("modalitiesIn: %w", err)
	}
	if err := validateOutputModalities(o.ModalitiesOut); err != nil {
		return fmt.Errorf("modalitiesOut: %w", err)
	}

	if o.ReasoningCapabilities != nil {
		if err := validateReasoningCapabilitiesOverride(o.ReasoningCapabilities); err != nil {
			return fmt.Errorf("reasoningCapabilities: %w", err)
		}
	}

	if o.StopSequenceCapabilities != nil {
		if err := validateStopSequenceCapabilitiesOverride(o.StopSequenceCapabilities); err != nil {
			return fmt.Errorf("stopSequenceCapabilities: %w", err)
		}
	}

	if o.OutputCapabilities != nil {
		if err := validateOutputCapabilitiesOverride(o.OutputCapabilities); err != nil {
			return fmt.Errorf("outputCapabilities: %w", err)
		}
	}

	if o.ToolCapabilities != nil {
		if err := validateToolCapabilitiesOverride(o.ToolCapabilities); err != nil {
			return fmt.Errorf("toolCapabilities: %w", err)
		}
	}

	if o.CacheCapabilities != nil {
		if err := validateCacheCapabilitiesOverride(o.CacheCapabilities); err != nil {
			return fmt.Errorf("cacheCapabilities: %w", err)
		}
	}

	if o.ParamDialect != nil {
		if err := validateParamDialectOverride(o.ParamDialect); err != nil {
			return fmt.Errorf("paramDialect: %w", err)
		}
	}

	return nil
}

func validateInputModalities(mm []spec.Modality) error {
	if mm == nil {
		return nil
	}

	seen := map[spec.Modality]struct{}{}

	for i, m := range mm {
		if strings.TrimSpace(string(m)) == "" {
			return fmt.Errorf("[%d] empty modality", i)
		}
		switch m {
		case spec.ModalityTextIn,
			spec.ModalityImageIn,
			spec.ModalityFileIn,
			spec.ModalityAudioIn,
			spec.ModalityVideoIn:
		default:
			return fmt.Errorf("[%d] invalid input modality %q", i, m)
		}

		if _, ok := seen[m]; ok {
			return fmt.Errorf("[%d] duplicate modality %q", i, m)
		}
		seen[m] = struct{}{}
	}

	return nil
}

func validateOutputModalities(mm []spec.Modality) error {
	if mm == nil {
		return nil
	}

	seen := map[spec.Modality]struct{}{}

	for i, m := range mm {
		if strings.TrimSpace(string(m)) == "" {
			return fmt.Errorf("[%d] empty modality", i)
		}

		switch m {
		case spec.ModalityTextOut,
			spec.ModalityImageOut,
			spec.ModalityFileOut,
			spec.ModalityAudioOut,
			spec.ModalityVideoOut:
		default:
			return fmt.Errorf("[%d] invalid output modality %q", i, m)
		}

		if _, ok := seen[m]; ok {
			return fmt.Errorf("[%d] duplicate modality %q", i, m)
		}
		seen[m] = struct{}{}
	}

	return nil
}

func validateReasoningCapabilitiesOverride(o *ReasoningCapabilitiesOverride) error {
	if o == nil {
		return nil
	}

	if o.SupportedReasoningTypes != nil {
		seen := map[spec.ReasoningType]struct{}{}

		for i, t := range o.SupportedReasoningTypes {
			switch t {
			case spec.ReasoningTypeHybridWithTokens,
				spec.ReasoningTypeSingleWithLevels:
			default:
				return fmt.Errorf("supportedReasoningTypes[%d] unknown type %q", i, t)
			}

			if _, ok := seen[t]; ok {
				return fmt.Errorf("supportedReasoningTypes[%d] duplicate %q", i, t)
			}
			seen[t] = struct{}{}
		}
	}

	if o.SupportedReasoningLevels != nil {
		seen := map[spec.ReasoningLevel]struct{}{}

		for i, l := range o.SupportedReasoningLevels {
			switch l {
			case spec.ReasoningLevelNone,
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax:
			default:
				return fmt.Errorf("supportedReasoningLevels[%d] unknown level %q", i, l)
			}

			if _, ok := seen[l]; ok {
				return fmt.Errorf("supportedReasoningLevels[%d] duplicate %q", i, l)
			}
			seen[l] = struct{}{}
		}
	}

	if err := validateReasoningTokenBudgetCapabilitiesOverride(o.HybridTokenBudgetCapabilities); err != nil {
		return fmt.Errorf("hybridTokenBudgetCapabilities: %w", err)
	}

	return nil
}

func validateReasoningTokenBudgetCapabilitiesOverride(
	o *ReasoningTokenBudgetCapabilitiesOverride,
) error {
	if o == nil {
		return nil
	}

	if o.MinAllowed != nil && *o.MinAllowed < 0 {
		return errors.New("minAllowed must be >= 0")
	}

	if o.MaxAllowed != nil && *o.MaxAllowed < 0 {
		return errors.New("maxAllowed must be >= 0")
	}

	if o.MinAllowed != nil && o.MaxAllowed != nil && *o.MaxAllowed < *o.MinAllowed {
		return errors.New("maxAllowed must be >= minAllowed")
	}
	return nil
}

func validateStopSequenceCapabilitiesOverride(o *StopSequenceCapabilitiesOverride) error {
	if o == nil {
		return nil
	}

	if o.MaxSequences != nil && *o.MaxSequences < 0 {
		return errors.New("maxSequences must be >= 0")
	}

	return nil
}

func validateOutputCapabilitiesOverride(o *OutputCapabilitiesOverride) error {
	if o == nil {
		return nil
	}

	if o.SupportedOutputFormats != nil {
		seen := map[spec.OutputFormatKind]struct{}{}

		for i, k := range o.SupportedOutputFormats {
			switch k {
			case spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema:
			default:
				return fmt.Errorf("supportedOutputFormats[%d] unknown kind %q", i, k)
			}

			if _, ok := seen[k]; ok {
				return fmt.Errorf("supportedOutputFormats[%d] duplicate %q", i, k)
			}
			seen[k] = struct{}{}
		}
	}

	return nil
}

func validateToolCapabilitiesOverride(o *ToolCapabilitiesOverride) error {
	if o == nil {
		return nil
	}

	if o.MaxForcedTools != nil && *o.MaxForcedTools < 0 {
		return errors.New("maxForcedTools must be >= 0")
	}

	if o.SupportedToolTypes != nil {
		seen := map[spec.ToolType]struct{}{}

		for i, t := range o.SupportedToolTypes {
			switch t {
			case spec.ToolTypeFunction,
				spec.ToolTypeCustom,
				spec.ToolTypeWebSearch:
			default:
				return fmt.Errorf("supportedToolTypes[%d] unknown type %q", i, t)
			}

			if _, ok := seen[t]; ok {
				return fmt.Errorf("supportedToolTypes[%d] duplicate %q", i, t)
			}
			seen[t] = struct{}{}
		}
	}

	if o.SupportedToolPolicyModes != nil {
		seen := map[spec.ToolPolicyMode]struct{}{}

		for i, m := range o.SupportedToolPolicyModes {
			switch m {
			case spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone:
			default:
				return fmt.Errorf("supportedToolPolicyModes[%d] unknown mode %q", i, m)
			}

			if _, ok := seen[m]; ok {
				return fmt.Errorf("supportedToolPolicyModes[%d] duplicate %q", i, m)
			}
			seen[m] = struct{}{}
		}
	}

	if o.SupportedClientToolOutputFormats != nil {
		seen := map[spec.ToolOutputFormatKind]struct{}{}

		for i, k := range o.SupportedClientToolOutputFormats {
			switch k {
			case spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList:
			default:
				return fmt.Errorf("supportedClientToolOutputFormats[%d] unknown kind %q", i, k)
			}

			if _, ok := seen[k]; ok {
				return fmt.Errorf("supportedClientToolOutputFormats[%d] duplicate %q", i, k)
			}
			seen[k] = struct{}{}
		}
	}

	return nil
}

func validateCacheCapabilitiesOverride(o *CacheCapabilitiesOverride) error {
	if o == nil {
		return nil
	}

	scopes := []struct {
		name string
		val  *CacheControlCapabilitiesOverride
	}{
		{"topLevel", o.TopLevel},
		{"inputOutputContent", o.InputOutputContent},
		{"reasoningContent", o.ReasoningContent},
		{"toolChoice", o.ToolChoice},
		{"toolCall", o.ToolCall},
		{"toolOutput", o.ToolOutput},
	}

	for _, s := range scopes {
		if s.val != nil {
			if err := validateCacheControlCapabilitiesOverride(s.val); err != nil {
				return fmt.Errorf("%s: %w", s.name, err)
			}
		}
	}

	return nil
}

func validateCacheControlCapabilitiesOverride(o *CacheControlCapabilitiesOverride) error {
	if o == nil {
		return nil
	}

	if o.SupportedKinds != nil {
		seen := map[spec.CacheControlKind]struct{}{}

		for i, k := range o.SupportedKinds {
			switch k {
			case spec.CacheControlKindEphemeral:
			default:
				return fmt.Errorf("supportedKinds[%d] unknown kind %q", i, k)
			}

			if _, ok := seen[k]; ok {
				return fmt.Errorf("supportedKinds[%d] duplicate %q", i, k)
			}
			seen[k] = struct{}{}
		}
	}

	if o.SupportedTTLs != nil {
		seen := map[spec.CacheControlTTL]struct{}{}

		for i, t := range o.SupportedTTLs {
			switch t {
			case spec.CacheControlTTL5m,
				spec.CacheControlTTL1h,
				spec.CacheControlTTL24h,
				spec.CacheControlTTLInMemory:
			default:
				return fmt.Errorf("supportedTTLs[%d] unknown TTL %q", i, t)
			}

			if _, ok := seen[t]; ok {
				return fmt.Errorf("supportedTTLs[%d] duplicate %q", i, t)
			}
			seen[t] = struct{}{}
		}
	}

	return nil
}

func validateParamDialectOverride(o *ParamDialectOverride) error {
	if o == nil {
		return nil
	}

	if o.MaxOutputTokensParamName != nil {
		switch *o.MaxOutputTokensParamName {
		case spec.MaxOutputTokensParamNameMaxCompletionTokens,
			spec.MaxOutputTokensParamNameMaxTokens:
		default:
			return fmt.Errorf("maxOutputTokensParamName unknown value %q", *o.MaxOutputTokensParamName)
		}
	}

	if o.ToolChoiceParamStyle != nil {
		switch *o.ToolChoiceParamStyle {
		case spec.ToolChoiceParamStyleAllowedTools,
			spec.ToolChoiceParamStyleRequiredNamed:
		default:
			return fmt.Errorf("toolChoiceParamStyle unknown value %q", *o.ToolChoiceParamStyle)
		}
	}

	return nil
}
