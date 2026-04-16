package openairesponsessdk

import "github.com/flexigpt/inference-go/spec"

var openairesponsessdkCapability = spec.ModelCapabilities{
	ModalitiesIn:  []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn, spec.ModalityFileIn},
	ModalitiesOut: []spec.Modality{spec.ModalityTextOut},

	ReasoningCapabilities: &spec.ReasoningCapabilities{
		SupportsReasoningConfig: true,
		SupportedReasoningTypes: []spec.ReasoningType{spec.ReasoningTypeSingleWithLevels},
		SupportedReasoningLevels: []spec.ReasoningLevel{
			spec.ReasoningLevelNone,
			spec.ReasoningLevelMinimal,
			spec.ReasoningLevelLow,
			spec.ReasoningLevelMedium,
			spec.ReasoningLevelHigh,
			spec.ReasoningLevelXHigh,
		},
		SupportsSummaryStyle: true,

		SupportsEncryptedReasoningInput:  true,
		TemperatureDisallowedWhenEnabled: false,
	},

	StopSequenceCapabilities: &spec.StopSequenceCapabilities{
		IsSupported:             false,
		DisallowedWithReasoning: false,
		MaxSequences:            0,
	},
	OutputCapabilities: &spec.OutputCapabilities{
		SupportedOutputFormats: []spec.OutputFormatKind{spec.OutputFormatKindText, spec.OutputFormatKindJSONSchema},
		SupportsVerbosity:      true,
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
		SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
			spec.ToolOutputFormatKindString,
			spec.ToolOutputFormatKindContentItemList,
		},
	},

	CacheCapabilities: &spec.CacheCapabilities{
		SupportsAutomaticCaching: true,
		TopLevel: &spec.CacheControlCapabilities{
			SupportsTTL:    true,
			SupportedKinds: []spec.CacheControlKind{spec.CacheControlKindEphemeral},
			SupportedTTLs:  []spec.CacheControlTTL{spec.CacheControlTTLInMemory, spec.CacheControlTTL24h},
			SupportsKey:    true,
		},
	},
}
