package openairesponsessdk

import "github.com/flexigpt/inference-go/spec"

var openairesponsessdkCapability = spec.ModelCapabilities{
	ModalitiesIn:  []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn, spec.ModalityFileIn},
	ModalitiesOut: []spec.Modality{spec.ModalityTextOut},

	Reasoning: &spec.ReasoningCapabilities{
		SupportedTypes: []spec.ReasoningType{spec.ReasoningTypeSingleWithLevels},
		SupportedLevels: []spec.ReasoningLevel{
			spec.ReasoningLevelNone,
			spec.ReasoningLevelMinimal,
			spec.ReasoningLevelLow,
			spec.ReasoningLevelMedium,
			spec.ReasoningLevelHigh,
			spec.ReasoningLevelXHigh,
		},
		SupportsSummaryStyle:             true,
		TemperatureDisallowedWhenEnabled: false,
		SupportsEncryptedReasoningInput:  true,
	},
	StopSequences: &spec.StopSequenceCapabilities{Supported: false},
	Output: &spec.OutputCapabilities{
		SupportedFormats:  []spec.OutputFormatKind{spec.OutputFormatKindText, spec.OutputFormatKindJSONSchema},
		SupportsVerbosity: true,
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
