package googlegeneratecontentsdk

import "github.com/flexigpt/inference-go/spec"

var googleGenerateContentSDKCapability = spec.ModelCapabilities{
	ModalitiesIn:  []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn, spec.ModalityFileIn},
	ModalitiesOut: []spec.Modality{spec.ModalityTextOut},

	ReasoningCapabilities: &spec.ReasoningCapabilities{
		SupportsReasoningConfig: true,
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
			spec.ReasoningLevelXHigh, // mapped to ThinkingLevelHigh
		},
		HybridTokenBudgetCapabilities: &spec.ReasoningTokenBudgetCapabilities{
			MinAllowed:      1,
			MaxAllowed:      32768,
			ZeroAllowed:     true,
			MinusOneAllowed: true,
		},
		SupportsSummaryStyle:             false,
		SupportsEncryptedReasoningInput:  false,
		TemperatureDisallowedWhenEnabled: false,
	},

	StopSequenceCapabilities: &spec.StopSequenceCapabilities{
		IsSupported:             true,
		DisallowedWithReasoning: false,
		MaxSequences:            5,
	},

	OutputCapabilities: &spec.OutputCapabilities{
		SupportedOutputFormats: []spec.OutputFormatKind{
			spec.OutputFormatKindText,
			spec.OutputFormatKindJSONSchema,
		},
		SupportsVerbosity: false,
	},

	ToolCapabilities: &spec.ToolCapabilities{
		SupportedToolTypes: []spec.ToolType{
			spec.ToolTypeFunction,
			spec.ToolTypeCustom,
			spec.ToolTypeWebSearch,
		},
		SupportedToolPolicyModes: []spec.ToolPolicyMode{
			spec.ToolPolicyModeAuto,
			spec.ToolPolicyModeAny,
			spec.ToolPolicyModeTool,
			spec.ToolPolicyModeNone,
		},
		SupportsParallelToolCalls: true,
		MaxForcedTools:            1,
		SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
			spec.ToolOutputFormatKindString,
		},
	},

	// Google GenAI does not support per-message cache_control parameters.
	// Context caching is a separate resource-based API (CachedContent) that is
	// out of scope for the spec.CacheControl mechanism.
	// The normalizer will drop any caller-provided cache controls with a warning.
	CacheCapabilities: nil,
}
