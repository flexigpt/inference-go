package googleinteractionssdk

import "github.com/flexigpt/inference-go/spec"

func googleInteractionsSDKCapability() spec.ModelCapabilities {
	return spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
			spec.ModalityFileIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &spec.ReasoningCapabilities{
			SupportsReasoningConfig: true,
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
			},
			SupportsSummaryStyle:             true,
			SupportsReasoningContext:         false,
			SupportsReasoningMode:            false,
			SupportsEncryptedReasoningInput:  false,
			TemperatureDisallowedWhenEnabled: true,
		},
		StopSequenceCapabilities: &spec.StopSequenceCapabilities{
			IsSupported:             true,
			DisallowedWithReasoning: false,
			MaxSequences:            4,
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
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	}
}

func appendGoogleInteractionsWarning(
	warnings []spec.Warning,
	code string,
	message string,
) []spec.Warning {
	for _, warning := range warnings {
		if warning.Code == code && warning.Message == message {
			return warnings
		}
	}

	return append(warnings, spec.Warning{
		Code:    code,
		Message: message,
	})
}
