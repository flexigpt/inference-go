package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelXAIGrok43 = ModelPreset{
	ID:          PresetGrok43,
	Name:        ModelNameGrok43,
	DisplayName: DisplayNameGrok43,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGrok43,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         3600,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelXAIGrokBuild01 = ModelPreset{
	ID:          PresetBuild01,
	Name:        ModelNameGrokBuild01,
	DisplayName: DisplayNameBuild01,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGrokBuild01,
		Stream:          true,
		MaxPromptLength: 2000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         3600,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:         new(false),
			SupportsEncryptedReasoningInput: new(true),
		},
	},
}

var modelXAIGrok42Reasoning = ModelPreset{
	ID:          PresetGrok42Reasoning,
	Name:        ModelNameGrok42Reasoning,
	DisplayName: DisplayNameGrok42Reasoning,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGrok42Reasoning,
		Stream:          true,
		MaxPromptLength: 2000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         3600,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:         new(false),
			SupportsEncryptedReasoningInput: new(true),
		},
	},
}

var modelXAIGrok42NonReasoning = ModelPreset{
	ID:          PresetGrok42NonReasoning,
	Name:        ModelNameGrok42NonReasoning,
	DisplayName: DisplayNameGrok42NonReasoning,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGrok42NonReasoning,
		Stream:          true,
		MaxPromptLength: 2000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:         new(false),
			SupportsEncryptedReasoningInput: new(false),
		},
	},
}

var providerXAI = ProviderPreset{
	Name:                     ProviderXAI,
	DisplayName:              DisplayNameProviderXAI,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://api.x.ai",
	ChatCompletionPathPrefix: spec.DefaultOpenAIResponsesPrefix,
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(true),
			TemperatureDisallowedWhenEnabled: new(false),
		},
		StopSequenceCapabilities: &capabilityoverride.StopSequenceCapabilitiesOverride{
			IsSupported:             new(false),
			DisallowedWithReasoning: new(false),
			MaxSequences:            new(0),
		},
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
				spec.ToolTypeWebSearch,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
		CacheCapabilities: &capabilityoverride.CacheCapabilitiesOverride{
			SupportsAutomaticCaching: new(true),
			TopLevel: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportedKinds: []spec.CacheControlKind{
					spec.CacheControlKindEphemeral,
				},
				SupportsTTL: new(false),
				SupportsKey: new(true),
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetBuild01:            modelXAIGrokBuild01,
		PresetGrok43:             modelXAIGrok43,
		PresetGrok42Reasoning:    modelXAIGrok42Reasoning,
		PresetGrok42NonReasoning: modelXAIGrok42NonReasoning,
	},
}
