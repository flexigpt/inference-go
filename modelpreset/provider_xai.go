package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetXAIGrokBuild01        ModelPresetID = "grokBuild01"
	PresetXAIGrok43             ModelPresetID = "grok43"
	PresetXAIGrok42Reasoning    ModelPresetID = "grok42Reasoning"
	PresetXAIGrok42NonReasoning ModelPresetID = "grok42NonReasoning"
)

const (
	ModelNameXAIGrokBuild01        spec.ModelName = "grok-build-0.1"
	ModelNameXAIGrok43             spec.ModelName = "grok-4.3"
	ModelNameXAIGrok42Reasoning    spec.ModelName = "grok-4.20-0309-reasoning"
	ModelNameXAIGrok42NonReasoning spec.ModelName = "grok-4.20-0309-non-reasoning"
)

var modelXAIGrok43 = ModelPreset{
	ID:          PresetXAIGrok43,
	Name:        ModelNameXAIGrok43,
	DisplayName: "xAI Grok 4.3",
	ModelParam: spec.ModelParam{
		Name:            ModelNameXAIGrok43,
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
	ID:          PresetXAIGrokBuild01,
	Name:        ModelNameXAIGrokBuild01,
	DisplayName: "xAI Build 0.1",
	ModelParam: spec.ModelParam{
		Name:            ModelNameXAIGrokBuild01,
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
	ID:          PresetXAIGrok42Reasoning,
	Name:        ModelNameXAIGrok42Reasoning,
	DisplayName: "xAI Grok 4.2 Reasoning",
	ModelParam: spec.ModelParam{
		Name:            ModelNameXAIGrok42Reasoning,
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
	ID:          PresetXAIGrok42NonReasoning,
	Name:        ModelNameXAIGrok42NonReasoning,
	DisplayName: "xAI Grok 4.2 Non-Reasoning",
	ModelParam: spec.ModelParam{
		Name:            ModelNameXAIGrok42NonReasoning,
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
	DisplayName:              "xAI",
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
		PresetXAIGrokBuild01:        modelXAIGrokBuild01,
		PresetXAIGrok43:             modelXAIGrok43,
		PresetXAIGrok42Reasoning:    modelXAIGrok42Reasoning,
		PresetXAIGrok42NonReasoning: modelXAIGrok42NonReasoning,
	},
}
