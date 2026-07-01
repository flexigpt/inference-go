package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetMistralMedium35  ModelPresetID = "mistralMedium35"
	PresetMistralSmall4    ModelPresetID = "mistralSmall4"
	PresetMistralLarge3    ModelPresetID = "mistralLarge3"
	PresetMistralDevstral2 ModelPresetID = "devstral2"
)

const (
	ModelNameMistralMedium35  spec.ModelName = "mistral-medium-3-5"
	ModelNameMistralSmall4    spec.ModelName = "mistral-small-2603"
	ModelNameMistralLarge3    spec.ModelName = "mistral-large-2512"
	ModelNameMistralDevstral2 spec.ModelName = "devstral-2512"
)

var modelMistralMedium35 = ModelPreset{
	ID:          PresetMistralMedium35,
	Name:        ModelNameMistralMedium35,
	DisplayName: "Mistral Medium 3.5",
	ModelParam: spec.ModelParam{
		Name:            ModelNameMistralMedium35,
		Stream:          true,
		MaxPromptLength: 256000,
		MaxOutputLength: 32768,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMistralSmall4 = ModelPreset{
	ID:          PresetMistralSmall4,
	Name:        ModelNameMistralSmall4,
	DisplayName: "Mistral Small 4",
	ModelParam: spec.ModelParam{
		Name:            ModelNameMistralSmall4,
		Stream:          true,
		MaxPromptLength: 256000,
		MaxOutputLength: 32768,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMistralLarge3 = ModelPreset{
	ID:          PresetMistralLarge3,
	Name:        ModelNameMistralLarge3,
	DisplayName: "Mistral Large 3",
	ModelParam: spec.ModelParam{
		Name:            ModelNameMistralLarge3,
		Stream:          true,
		MaxPromptLength: 256000,
		MaxOutputLength: 32768,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelMistralDevstral2 = ModelPreset{
	ID:          PresetMistralDevstral2,
	Name:        ModelNameMistralDevstral2,
	DisplayName: "Mistral Devstral 2",
	ModelParam: spec.ModelParam{
		Name:            ModelNameMistralDevstral2,
		Stream:          true,
		MaxPromptLength: 256000,
		MaxOutputLength: 32768,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var providerMistral = ProviderPreset{
	Name:                     ProviderMistral,
	DisplayName:              "Mistral AI",
	SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
	Origin:                   "https://api.mistral.ai",
	ChatCompletionPathPrefix: spec.DefaultOpenAIChatCompletionsPrefix,
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
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
		StopSequenceCapabilities: &capabilityoverride.StopSequenceCapabilitiesOverride{
			IsSupported:             new(true),
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
			SupportsAutomaticCaching: new(false),
			TopLevel: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
		},
		ParamDialect: &capabilityoverride.ParamDialectOverride{
			MaxOutputTokensParamName: new(spec.MaxOutputTokensParamNameMaxTokens),
			ToolChoiceParamStyle:     new(spec.ToolChoiceParamStyleRequiredNamed),
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetMistralMedium35:  modelMistralMedium35,
		PresetMistralSmall4:    modelMistralSmall4,
		PresetMistralLarge3:    modelMistralLarge3,
		PresetMistralDevstral2: modelMistralDevstral2,
	},
}
