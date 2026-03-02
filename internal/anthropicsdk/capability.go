package anthropicsdk

import "github.com/flexigpt/inference-go/spec"

var anthropicsdkCapability = spec.ModelCapabilities{
	ModalitiesIn:  []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn, spec.ModalityFileIn},
	ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	Reasoning: &spec.ReasoningCapabilities{
		SupportedTypes: []spec.ReasoningType{
			spec.ReasoningTypeHybridWithTokens,
			spec.ReasoningTypeSingleWithLevels,
		},
		SupportedLevels: []spec.ReasoningLevel{
			spec.ReasoningLevelNone,
			spec.ReasoningLevelMinimal,
			spec.ReasoningLevelLow,
			spec.ReasoningLevelMedium,
			spec.ReasoningLevelHigh,
			spec.ReasoningLevelXHigh,
		},
		SupportsSummaryStyle:             false,
		TemperatureDisallowedWhenEnabled: true,
		SupportsEncryptedReasoningInput:  false,
	},
	StopSequences: &spec.StopSequenceCapabilities{Supported: true, Max: 0},
	Output: &spec.OutputCapabilities{
		SupportedFormats:  []spec.OutputFormatKind{spec.OutputFormatKindText, spec.OutputFormatKindJSONSchema},
		SupportsVerbosity: true, // maps to effort
	},
	Tools: &spec.ToolCapabilities{
		SupportedToolTypes: []spec.ToolType{spec.ToolTypeFunction, spec.ToolTypeCustom, spec.ToolTypeWebSearch},
		SupportedPolicyModes: []spec.ToolPolicyMode{
			spec.ToolPolicyModeAuto,
			spec.ToolPolicyModeAny,
			spec.ToolPolicyModeTool,
			spec.ToolPolicyModeNone,
		},
		SupportsParallelToolCalls: true,
		MaxForcedTools:            1,
	},
}
