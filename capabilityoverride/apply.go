package capabilityoverride

import (
	"slices"

	"github.com/flexigpt/inference-go/spec"
)

func DeriveModelCapabilities(
	base spec.ModelCapabilities,
	overrides ...*ModelCapabilitiesOverride,
) spec.ModelCapabilities {
	out := CloneModelCapabilities(base)
	ApplyModelCapabilitiesOverrides(&out, overrides...)
	return out
}

func ApplyModelCapabilitiesOverrides(
	dst *spec.ModelCapabilities,
	overrides ...*ModelCapabilitiesOverride,
) {
	if dst == nil {
		return
	}

	for _, ov := range overrides {
		applyModelCapabilitiesOverride(dst, ov)
	}
}

func applyModelCapabilitiesOverride(
	dst *spec.ModelCapabilities,
	ov *ModelCapabilitiesOverride,
) {
	if dst == nil || ov == nil {
		return
	}

	if ov.ModalitiesIn != nil {
		dst.ModalitiesIn = slices.Clone(ov.ModalitiesIn)
	}
	if ov.ModalitiesOut != nil {
		dst.ModalitiesOut = slices.Clone(ov.ModalitiesOut)
	}

	applyParamDialectOverride(dst, ov.ParamDialect)
	applyReasoningCapabilitiesOverride(dst, ov.ReasoningCapabilities)
	applyStopSequenceCapabilitiesOverride(dst, ov.StopSequenceCapabilities)
	applyOutputCapabilitiesOverride(dst, ov.OutputCapabilities)
	applyToolCapabilitiesOverride(dst, ov.ToolCapabilities)
	applyCacheCapabilitiesOverride(dst, ov.CacheCapabilities)
}

func applyParamDialectOverride(
	dst *spec.ModelCapabilities,
	ov *ParamDialectOverride,
) {
	if ov == nil {
		return
	}

	if ov.MaxOutputTokensParamName == nil && ov.ToolChoiceParamStyle == nil {
		return
	}

	if dst.ParamDialect == nil {
		dst.ParamDialect = &spec.ParamDialect{}
	}

	if ov.MaxOutputTokensParamName != nil {
		dst.ParamDialect.MaxOutputTokensParamName = *ov.MaxOutputTokensParamName
	}

	if ov.ToolChoiceParamStyle != nil {
		dst.ParamDialect.ToolChoiceParamStyle = *ov.ToolChoiceParamStyle
	}
}

func applyReasoningCapabilitiesOverride(
	dst *spec.ModelCapabilities,
	ov *ReasoningCapabilitiesOverride,
) {
	if !hasReasoningCapabilitiesOverride(ov) {
		return
	}

	if dst.ReasoningCapabilities == nil {
		dst.ReasoningCapabilities = &spec.ReasoningCapabilities{}
	}

	if ov.SupportedReasoningTypes != nil {
		dst.ReasoningCapabilities.SupportedReasoningTypes = slices.Clone(ov.SupportedReasoningTypes)
	}
	if ov.SupportedReasoningLevels != nil {
		dst.ReasoningCapabilities.SupportedReasoningLevels = slices.Clone(ov.SupportedReasoningLevels)
	}
	if ov.SupportsReasoningConfig != nil {
		dst.ReasoningCapabilities.SupportsReasoningConfig = *ov.SupportsReasoningConfig
	}
	applyReasoningTokenBudgetCapabilitiesOverride(
		&dst.ReasoningCapabilities.HybridTokenBudgetCapabilities,
		ov.HybridTokenBudgetCapabilities,
	)
	if ov.SupportsSummaryStyle != nil {
		dst.ReasoningCapabilities.SupportsSummaryStyle = *ov.SupportsSummaryStyle
	}
	if ov.SupportsEncryptedReasoningInput != nil {
		dst.ReasoningCapabilities.SupportsEncryptedReasoningInput = *ov.SupportsEncryptedReasoningInput
	}
	if ov.TemperatureDisallowedWhenEnabled != nil {
		dst.ReasoningCapabilities.TemperatureDisallowedWhenEnabled = *ov.TemperatureDisallowedWhenEnabled
	}
}

func applyReasoningTokenBudgetCapabilitiesOverride(
	dst **spec.ReasoningTokenBudgetCapabilities,
	ov *ReasoningTokenBudgetCapabilitiesOverride,
) {
	if !hasReasoningTokenBudgetCapabilitiesOverride(ov) {
		return
	}

	if *dst == nil {
		*dst = &spec.ReasoningTokenBudgetCapabilities{}
	}

	if ov.MinAllowed != nil {
		(*dst).MinAllowed = *ov.MinAllowed
	}
	if ov.MaxAllowed != nil {
		(*dst).MaxAllowed = *ov.MaxAllowed
	}
	if ov.ZeroAllowed != nil {
		(*dst).ZeroAllowed = *ov.ZeroAllowed
	}
	if ov.MinusOneAllowed != nil {
		(*dst).MinusOneAllowed = *ov.MinusOneAllowed
	}
}

func applyStopSequenceCapabilitiesOverride(
	dst *spec.ModelCapabilities,
	ov *StopSequenceCapabilitiesOverride,
) {
	if !hasStopSequenceCapabilitiesOverride(ov) {
		return
	}
	if dst.StopSequenceCapabilities == nil {
		dst.StopSequenceCapabilities = &spec.StopSequenceCapabilities{}
	}

	if ov.IsSupported != nil {
		dst.StopSequenceCapabilities.IsSupported = *ov.IsSupported
	}
	if ov.DisallowedWithReasoning != nil {
		dst.StopSequenceCapabilities.DisallowedWithReasoning = *ov.DisallowedWithReasoning
	}
	if ov.MaxSequences != nil {
		dst.StopSequenceCapabilities.MaxSequences = *ov.MaxSequences
	}
}

func applyOutputCapabilitiesOverride(
	dst *spec.ModelCapabilities,
	ov *OutputCapabilitiesOverride,
) {
	if !hasOutputCapabilitiesOverride(ov) {
		return
	}
	if dst.OutputCapabilities == nil {
		dst.OutputCapabilities = &spec.OutputCapabilities{}
	}

	if ov.SupportedOutputFormats != nil {
		dst.OutputCapabilities.SupportedOutputFormats = slices.Clone(ov.SupportedOutputFormats)
	}
	if ov.SupportsVerbosity != nil {
		dst.OutputCapabilities.SupportsVerbosity = *ov.SupportsVerbosity
	}
}

func applyToolCapabilitiesOverride(
	dst *spec.ModelCapabilities,
	ov *ToolCapabilitiesOverride,
) {
	if !hasToolCapabilitiesOverride(ov) {
		return
	}

	if dst.ToolCapabilities == nil {
		dst.ToolCapabilities = &spec.ToolCapabilities{}
	}

	if ov.SupportedToolTypes != nil {
		dst.ToolCapabilities.SupportedToolTypes = slices.Clone(ov.SupportedToolTypes)
	}
	if ov.SupportedToolPolicyModes != nil {
		dst.ToolCapabilities.SupportedToolPolicyModes = slices.Clone(ov.SupportedToolPolicyModes)
	}
	if ov.SupportedClientToolOutputFormats != nil {
		dst.ToolCapabilities.SupportedClientToolOutputFormats = slices.Clone(ov.SupportedClientToolOutputFormats)
	}
	if ov.SupportsParallelToolCalls != nil {
		dst.ToolCapabilities.SupportsParallelToolCalls = *ov.SupportsParallelToolCalls
	}
	if ov.MaxForcedTools != nil {
		dst.ToolCapabilities.MaxForcedTools = *ov.MaxForcedTools
	}
}

func applyCacheCapabilitiesOverride(
	dst *spec.ModelCapabilities,
	ov *CacheCapabilitiesOverride,
) {
	if !hasCacheCapabilitiesOverride(ov) {
		return
	}

	if dst.CacheCapabilities == nil {
		dst.CacheCapabilities = &spec.CacheCapabilities{}
	}

	if ov.SupportsAutomaticCaching != nil {
		dst.CacheCapabilities.SupportsAutomaticCaching = *ov.SupportsAutomaticCaching
	}

	applyCacheControlCapabilitiesOverride(
		&dst.CacheCapabilities.TopLevel,
		ov.TopLevel,
	)
	applyCacheControlCapabilitiesOverride(
		&dst.CacheCapabilities.InputOutputContent,
		ov.InputOutputContent,
	)
	applyCacheControlCapabilitiesOverride(
		&dst.CacheCapabilities.ReasoningContent,
		ov.ReasoningContent,
	)
	applyCacheControlCapabilitiesOverride(
		&dst.CacheCapabilities.ToolChoice,
		ov.ToolChoice,
	)
	applyCacheControlCapabilitiesOverride(
		&dst.CacheCapabilities.ToolCall,
		ov.ToolCall,
	)
	applyCacheControlCapabilitiesOverride(
		&dst.CacheCapabilities.ToolOutput,
		ov.ToolOutput,
	)
}

func applyCacheControlCapabilitiesOverride(
	dst **spec.CacheControlCapabilities,
	ov *CacheControlCapabilitiesOverride,
) {
	if !hasCacheControlCapabilitiesOverride(ov) {
		return
	}

	if *dst == nil {
		*dst = &spec.CacheControlCapabilities{}
	}

	if ov.SupportedKinds != nil {
		(*dst).SupportedKinds = slices.Clone(ov.SupportedKinds)
	}
	if ov.SupportedTTLs != nil {
		(*dst).SupportedTTLs = slices.Clone(ov.SupportedTTLs)
	}
	if ov.SupportsKey != nil {
		(*dst).SupportsKey = *ov.SupportsKey
	}
	if ov.SupportsTTL != nil {
		(*dst).SupportsTTL = *ov.SupportsTTL
	}
}

func hasReasoningCapabilitiesOverride(ov *ReasoningCapabilitiesOverride) bool {
	return ov != nil &&
		(ov.SupportsReasoningConfig != nil ||
			ov.SupportedReasoningTypes != nil ||
			ov.SupportedReasoningLevels != nil ||
			hasReasoningTokenBudgetCapabilitiesOverride(ov.HybridTokenBudgetCapabilities) ||
			ov.SupportsSummaryStyle != nil ||
			ov.SupportsEncryptedReasoningInput != nil ||
			ov.TemperatureDisallowedWhenEnabled != nil)
}

func hasReasoningTokenBudgetCapabilitiesOverride(ov *ReasoningTokenBudgetCapabilitiesOverride) bool {
	return ov != nil &&
		(ov.MinAllowed != nil ||
			ov.MaxAllowed != nil ||
			ov.ZeroAllowed != nil ||
			ov.MinusOneAllowed != nil)
}

func hasStopSequenceCapabilitiesOverride(ov *StopSequenceCapabilitiesOverride) bool {
	return ov != nil &&
		(ov.IsSupported != nil ||
			ov.DisallowedWithReasoning != nil ||
			ov.MaxSequences != nil)
}

func hasOutputCapabilitiesOverride(ov *OutputCapabilitiesOverride) bool {
	return ov != nil &&
		(ov.SupportedOutputFormats != nil ||
			ov.SupportsVerbosity != nil)
}

func hasToolCapabilitiesOverride(ov *ToolCapabilitiesOverride) bool {
	return ov != nil &&
		(ov.SupportedToolTypes != nil ||
			ov.SupportedToolPolicyModes != nil ||
			ov.SupportsParallelToolCalls != nil ||
			ov.MaxForcedTools != nil ||
			ov.SupportedClientToolOutputFormats != nil)
}

func hasCacheCapabilitiesOverride(ov *CacheCapabilitiesOverride) bool {
	return ov != nil &&
		(ov.SupportsAutomaticCaching != nil ||
			hasCacheControlCapabilitiesOverride(ov.TopLevel) ||
			hasCacheControlCapabilitiesOverride(ov.InputOutputContent) ||
			hasCacheControlCapabilitiesOverride(ov.ReasoningContent) ||
			hasCacheControlCapabilitiesOverride(ov.ToolChoice) ||
			hasCacheControlCapabilitiesOverride(ov.ToolCall) ||
			hasCacheControlCapabilitiesOverride(ov.ToolOutput))
}

func hasCacheControlCapabilitiesOverride(ov *CacheControlCapabilitiesOverride) bool {
	return ov != nil &&
		(ov.SupportsTTL != nil ||
			ov.SupportedKinds != nil ||
			ov.SupportedTTLs != nil ||
			ov.SupportsKey != nil)
}
