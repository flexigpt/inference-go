package capabilityoverride

import (
	"slices"

	"github.com/flexigpt/inference-go/spec"
)

func CloneModelCapabilities(in spec.ModelCapabilities) spec.ModelCapabilities {
	out := spec.ModelCapabilities{
		ModalitiesIn:  slices.Clone(in.ModalitiesIn),
		ModalitiesOut: slices.Clone(in.ModalitiesOut),
	}

	if in.ReasoningCapabilities != nil {
		c := *in.ReasoningCapabilities
		c.SupportedReasoningTypes = slices.Clone(c.SupportedReasoningTypes)
		c.SupportedReasoningLevels = slices.Clone(c.SupportedReasoningLevels)
		if c.HybridTokenBudgetCapabilities != nil {
			h := *c.HybridTokenBudgetCapabilities
			c.HybridTokenBudgetCapabilities = &h
		}
		out.ReasoningCapabilities = &c
	}

	if in.StopSequenceCapabilities != nil {
		c := *in.StopSequenceCapabilities
		out.StopSequenceCapabilities = &c
	}

	if in.OutputCapabilities != nil {
		c := *in.OutputCapabilities
		c.SupportedOutputFormats = slices.Clone(c.SupportedOutputFormats)
		out.OutputCapabilities = &c
	}

	if in.ToolCapabilities != nil {
		c := *in.ToolCapabilities
		c.SupportedToolTypes = slices.Clone(c.SupportedToolTypes)
		c.SupportedToolPolicyModes = slices.Clone(c.SupportedToolPolicyModes)
		c.SupportedClientToolOutputFormats = slices.Clone(c.SupportedClientToolOutputFormats)
		out.ToolCapabilities = &c
	}

	if in.CacheCapabilities != nil {
		c := *in.CacheCapabilities
		c.TopLevel = cloneCacheControlCapabilities(in.CacheCapabilities.TopLevel)
		c.InputOutputContent = cloneCacheControlCapabilities(in.CacheCapabilities.InputOutputContent)
		c.ReasoningContent = cloneCacheControlCapabilities(in.CacheCapabilities.ReasoningContent)
		c.ToolChoice = cloneCacheControlCapabilities(in.CacheCapabilities.ToolChoice)
		c.ToolCall = cloneCacheControlCapabilities(in.CacheCapabilities.ToolCall)
		c.ToolOutput = cloneCacheControlCapabilities(in.CacheCapabilities.ToolOutput)
		out.CacheCapabilities = &c
	}

	if in.ParamDialect != nil {
		c := *in.ParamDialect
		out.ParamDialect = &c
	}

	return out
}

func CloneModelCapabilitiesOverride(in *ModelCapabilitiesOverride) *ModelCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &ModelCapabilitiesOverride{
		ModalitiesIn:             slices.Clone(in.ModalitiesIn),
		ModalitiesOut:            slices.Clone(in.ModalitiesOut),
		ReasoningCapabilities:    cloneReasoningCapabilitiesOverride(in.ReasoningCapabilities),
		StopSequenceCapabilities: cloneStopSequenceCapabilitiesOverride(in.StopSequenceCapabilities),
		OutputCapabilities:       cloneOutputCapabilitiesOverride(in.OutputCapabilities),
		ToolCapabilities:         cloneToolCapabilitiesOverride(in.ToolCapabilities),
		CacheCapabilities:        cloneCacheCapabilitiesOverride(in.CacheCapabilities),
		ParamDialect:             cloneParamDialectOverride(in.ParamDialect),
	}
}

func cloneCacheControlCapabilities(
	in *spec.CacheControlCapabilities,
) *spec.CacheControlCapabilities {
	if in == nil {
		return nil
	}
	out := *in
	out.SupportedKinds = slices.Clone(in.SupportedKinds)
	out.SupportedTTLs = slices.Clone(in.SupportedTTLs)
	return &out
}

func cloneParamDialectOverride(in *ParamDialectOverride) *ParamDialectOverride {
	if in == nil {
		return nil
	}

	out := &ParamDialectOverride{}

	if in.MaxOutputTokensParamName != nil {
		v := *in.MaxOutputTokensParamName
		out.MaxOutputTokensParamName = &v
	}

	if in.ToolChoiceParamStyle != nil {
		v := *in.ToolChoiceParamStyle
		out.ToolChoiceParamStyle = &v
	}

	return out
}

func cloneReasoningCapabilitiesOverride(
	in *ReasoningCapabilitiesOverride,
) *ReasoningCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &ReasoningCapabilitiesOverride{
		SupportsReasoningConfig:  cloneBoolPtr(in.SupportsReasoningConfig),
		SupportedReasoningTypes:  slices.Clone(in.SupportedReasoningTypes),
		SupportedReasoningLevels: slices.Clone(in.SupportedReasoningLevels),
		HybridTokenBudgetCapabilities: cloneReasoningTokenBudgetCapabilitiesOverride(
			in.HybridTokenBudgetCapabilities,
		),
		SupportsSummaryStyle:             cloneBoolPtr(in.SupportsSummaryStyle),
		SupportsEncryptedReasoningInput:  cloneBoolPtr(in.SupportsEncryptedReasoningInput),
		TemperatureDisallowedWhenEnabled: cloneBoolPtr(in.TemperatureDisallowedWhenEnabled),
	}
}

func cloneReasoningTokenBudgetCapabilitiesOverride(
	in *ReasoningTokenBudgetCapabilitiesOverride,
) *ReasoningTokenBudgetCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &ReasoningTokenBudgetCapabilitiesOverride{
		MinAllowed:      cloneIntPtr(in.MinAllowed),
		MaxAllowed:      cloneIntPtr(in.MaxAllowed),
		ZeroAllowed:     cloneBoolPtr(in.ZeroAllowed),
		MinusOneAllowed: cloneBoolPtr(in.MinusOneAllowed),
	}
}

func cloneStopSequenceCapabilitiesOverride(
	in *StopSequenceCapabilitiesOverride,
) *StopSequenceCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &StopSequenceCapabilitiesOverride{
		IsSupported:             cloneBoolPtr(in.IsSupported),
		DisallowedWithReasoning: cloneBoolPtr(in.DisallowedWithReasoning),
		MaxSequences:            cloneIntPtr(in.MaxSequences),
	}
}

func cloneOutputCapabilitiesOverride(in *OutputCapabilitiesOverride) *OutputCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &OutputCapabilitiesOverride{
		SupportedOutputFormats: slices.Clone(in.SupportedOutputFormats),
		SupportsVerbosity:      cloneBoolPtr(in.SupportsVerbosity),
	}
}

func cloneToolCapabilitiesOverride(in *ToolCapabilitiesOverride) *ToolCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &ToolCapabilitiesOverride{
		SupportedToolTypes:               slices.Clone(in.SupportedToolTypes),
		SupportedToolPolicyModes:         slices.Clone(in.SupportedToolPolicyModes),
		SupportsParallelToolCalls:        cloneBoolPtr(in.SupportsParallelToolCalls),
		MaxForcedTools:                   cloneIntPtr(in.MaxForcedTools),
		SupportedClientToolOutputFormats: slices.Clone(in.SupportedClientToolOutputFormats),
	}
}

func cloneCacheCapabilitiesOverride(in *CacheCapabilitiesOverride) *CacheCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &CacheCapabilitiesOverride{
		SupportsAutomaticCaching: cloneBoolPtr(in.SupportsAutomaticCaching),
		TopLevel:                 cloneCacheControlCapabilitiesOverride(in.TopLevel),
		InputOutputContent:       cloneCacheControlCapabilitiesOverride(in.InputOutputContent),
		ReasoningContent:         cloneCacheControlCapabilitiesOverride(in.ReasoningContent),
		ToolChoice:               cloneCacheControlCapabilitiesOverride(in.ToolChoice),
		ToolCall:                 cloneCacheControlCapabilitiesOverride(in.ToolCall),
		ToolOutput:               cloneCacheControlCapabilitiesOverride(in.ToolOutput),
	}
}

func cloneCacheControlCapabilitiesOverride(
	in *CacheControlCapabilitiesOverride,
) *CacheControlCapabilitiesOverride {
	if in == nil {
		return nil
	}

	return &CacheControlCapabilitiesOverride{
		SupportsTTL:    cloneBoolPtr(in.SupportsTTL),
		SupportedKinds: slices.Clone(in.SupportedKinds),
		SupportedTTLs:  slices.Clone(in.SupportedTTLs),
		SupportsKey:    cloneBoolPtr(in.SupportsKey),
	}
}

func cloneBoolPtr(p *bool) *bool {
	if p == nil {
		return nil
	}
	v := *p
	return &v
}

func cloneIntPtr(p *int) *int {
	if p == nil {
		return nil
	}
	v := *p
	return &v
}
