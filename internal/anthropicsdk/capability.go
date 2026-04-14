package anthropicsdk

import "github.com/flexigpt/inference-go/spec"

var anthropicsdkCapability = spec.ModelCapabilities{
	ModalitiesIn:  []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn, spec.ModalityFileIn},
	ModalitiesOut: []spec.Modality{spec.ModalityTextOut},

	ReasoningCapabilities: &spec.ReasoningCapabilities{
		SupportedReasoningTypes: []spec.ReasoningType{
			spec.ReasoningTypeHybridWithTokens,
			spec.ReasoningTypeSingleWithLevels,
		},

		SupportedReasoningLevels: []spec.ReasoningLevel{
			spec.ReasoningLevelNone,
			spec.ReasoningLevelMinimal,
			spec.ReasoningLevelLow,
			spec.ReasoningLevelMedium,
			spec.ReasoningLevelHigh,
			spec.ReasoningLevelXHigh,
		},
		HybridTokenBudgetCapabilities: &spec.ReasoningTokenBudgetCapabilities{
			MinAllowed:      1024,
			ZeroAllowed:     false,
			MinusOneAllowed: false,
		},
		SupportsSummaryStyle: false,

		SupportsEncryptedReasoningInput:  false,
		TemperatureDisallowedWhenEnabled: true,
	},
	StopSequenceCapabilities: &spec.StopSequenceCapabilities{
		IsSupported:             true,
		DisallowedWithReasoning: false,
		MaxSequences:            0,
	},
	OutputCapabilities: &spec.OutputCapabilities{
		SupportedOutputFormats: []spec.OutputFormatKind{spec.OutputFormatKindText, spec.OutputFormatKindJSONSchema},
		SupportsVerbosity:      true, // Maps to effort.
	},

	ToolCapabilities: &spec.ToolCapabilities{
		SupportedToolTypes: []spec.ToolType{spec.ToolTypeFunction, spec.ToolTypeCustom, spec.ToolTypeWebSearch},

		SupportedToolPolicyModes: []spec.ToolPolicyMode{
			spec.ToolPolicyModeAuto,
			spec.ToolPolicyModeAny,
			spec.ToolPolicyModeTool,
			spec.ToolPolicyModeNone,
		},
		SupportsParallelToolCalls: true,
		MaxForcedTools:            1,
	},

	CacheCapabilities: &spec.CacheCapabilities{
		SupportsAutomaticCaching: false,
		TopLevel: &spec.CacheControlCapabilities{
			SupportedKinds: []spec.CacheControlKind{spec.CacheControlKindEphemeral},
			SupportedTTLs:  []spec.CacheControlTTL{spec.CacheControlTTL5m, spec.CacheControlTTL1h},
			SupportsKey:    false,
		},
		InputOutputContent: &spec.CacheControlCapabilities{
			SupportedKinds: []spec.CacheControlKind{spec.CacheControlKindEphemeral},
			SupportedTTLs:  []spec.CacheControlTTL{spec.CacheControlTTL5m, spec.CacheControlTTL1h},
			SupportsKey:    false,
		},
		ToolChoice: &spec.CacheControlCapabilities{
			SupportedKinds: []spec.CacheControlKind{spec.CacheControlKindEphemeral},
			SupportedTTLs:  []spec.CacheControlTTL{spec.CacheControlTTL5m, spec.CacheControlTTL1h},
			SupportsKey:    false,
		},
		ToolCall: &spec.CacheControlCapabilities{
			SupportedKinds: []spec.CacheControlKind{spec.CacheControlKindEphemeral},
			SupportedTTLs:  []spec.CacheControlTTL{spec.CacheControlTTL5m, spec.CacheControlTTL1h},
			SupportsKey:    false,
		},
		ToolOutput: &spec.CacheControlCapabilities{
			SupportedKinds: []spec.CacheControlKind{spec.CacheControlKindEphemeral},
			SupportedTTLs:  []spec.CacheControlTTL{spec.CacheControlTTL5m, spec.CacheControlTTL1h},
			SupportsKey:    false,
		},
	},
}
