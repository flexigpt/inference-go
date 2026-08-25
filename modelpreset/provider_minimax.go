package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderMiniMax spec.ProviderName = "minimax"

	DisplayNameProviderMiniMax = "MiniMax"
)

const (
	ModelNameMiniMaxM3                spec.ModelName = "MiniMax-M3"
	ModelNameMiniMaxM27               spec.ModelName = "MiniMax-M2.7"
	ModelNameMiniMaxM27Highspeed      spec.ModelName = "MiniMax-M2.7-highspeed"
	ModelNameMiniMaxM25               spec.ModelName = "MiniMax-M2.5"
	ModelNameMiniMaxM25Highspeed      spec.ModelName = "MiniMax-M2.5-highspeed"
	ModelNameMiniMaxM21               spec.ModelName = "MiniMax-M2.1"
	ModelNameMiniMaxM21Highspeed      spec.ModelName = "MiniMax-M2.1-highspeed"
	ModelNameMiniMaxM2                spec.ModelName = "MiniMax-M2"
	ModelNameOpenRouterMiniMaxM3      spec.ModelName = "minimax/minimax-m3"
	ModelNameOpenRouterMiniMaxM27     spec.ModelName = "minimax/minimax-m2.7"
	ModelNameOpenRouterMiniMaxM25Free spec.ModelName = "minimax/minimax-m2.5:free"

	ModelNameMiniMaxM27FireworksAI spec.ModelName = "MiniMaxAI/MiniMax-M2.7:fireworks-ai"
	ModelNameMiniMaxM25Novita      spec.ModelName = "MiniMaxAI/MiniMax-M2.5:novita"
)

const (
	DisplayNameMiniMaxM2             = "MiniMax M2"
	DisplayNameMiniMaxM21            = "MiniMax M2.1"
	DisplayNameMiniMaxM21Highspeed   = "MiniMax M2.1 Highspeed"
	DisplayNameMiniMaxM25            = "MiniMax M2.5"
	DisplayNameMiniMaxM25Highspeed   = "MiniMax M2.5 Highspeed"
	DisplayNameMiniMaxM25Free        = "MiniMax M2.5 Free"
	DisplayNameMiniMaxM27            = "MiniMax M2.7"
	DisplayNameMiniMaxM27Highspeed   = "MiniMax M2.7 Highspeed"
	DisplayNameMiniMaxM3             = "MiniMax M3"
	DisplayNameMiniMaxM25Novita      = "MiniMax M2.5 (Novita)"
	DisplayNameMiniMaxM27FireworksAI = "MiniMax M2.7 (Fireworks AI)"
)

const (
	PresetMiniMaxM2           ModelPresetID = "minimaxm2"
	PresetMiniMaxM21          ModelPresetID = "minimaxm21"
	PresetMiniMaxM21Highspeed ModelPresetID = "minimaxm21Highspeed"
	PresetMiniMaxM25          ModelPresetID = "minimaxm25"
	PresetMiniMaxM25Highspeed ModelPresetID = "minimaxm25Highspeed"
	PresetMiniMaxM25Free      ModelPresetID = "minimaxm25free"
	PresetMiniMaxM27          ModelPresetID = "minimaxm27"
	PresetMiniMaxM27Highspeed ModelPresetID = "minimaxm27Highspeed"
	PresetMiniMaxM3           ModelPresetID = "minimaxM3"

	PresetMiniMaxM27FireworksAI ModelPresetID = "minimaxm27FireworksAI"
	PresetMiniMaxM25Novita      ModelPresetID = "minimaxm25Novita"
)

var modelMiniMaxM3 = ModelPreset{
	ID:          PresetMiniMaxM3,
	Name:        ModelNameMiniMaxM3,
	DisplayName: DisplayNameMiniMaxM3,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM3,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 524288,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMiniMaxM27 = ModelPreset{
	ID:          PresetMiniMaxM27,
	Name:        ModelNameMiniMaxM27,
	DisplayName: DisplayNameMiniMaxM27,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM27,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 204800,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMiniMaxM27Highspeed = ModelPreset{
	ID:          PresetMiniMaxM27Highspeed,
	Name:        ModelNameMiniMaxM27Highspeed,
	DisplayName: DisplayNameMiniMaxM27Highspeed,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM27Highspeed,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 204800,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMiniMaxM25 = ModelPreset{
	ID:          PresetMiniMaxM25,
	Name:        ModelNameMiniMaxM25,
	DisplayName: DisplayNameMiniMaxM25,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM25,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 204800,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMiniMaxM25Highspeed = ModelPreset{
	ID:          PresetMiniMaxM25Highspeed,
	Name:        ModelNameMiniMaxM25Highspeed,
	DisplayName: DisplayNameMiniMaxM25Highspeed,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM25Highspeed,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 204800,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMiniMaxM21 = ModelPreset{
	ID:          PresetMiniMaxM21,
	Name:        ModelNameMiniMaxM21,
	DisplayName: DisplayNameMiniMaxM21,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM21,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 204800,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMiniMaxM21Highspeed = ModelPreset{
	ID:          PresetMiniMaxM21Highspeed,
	Name:        ModelNameMiniMaxM21Highspeed,
	DisplayName: DisplayNameMiniMaxM21Highspeed,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM21Highspeed,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 204800,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelMiniMaxM2 = ModelPreset{
	ID:          PresetMiniMaxM2,
	Name:        ModelNameMiniMaxM2,
	DisplayName: DisplayNameMiniMaxM2,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM2,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 204800,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var providerMiniMax = ProviderPreset{
	Name:                     ProviderMiniMax,
	DisplayName:              DisplayNameProviderMiniMax,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://api.minimax.io",
	ChatCompletionPathPrefix: spec.DefaultOpenAIResponsesPrefix,
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
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
			},
			SupportsVerbosity: new(false),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
		CacheCapabilities: &capabilityoverride.CacheCapabilitiesOverride{
			SupportsAutomaticCaching: new(false),
			TopLevel: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(true),
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetMiniMaxM3:           modelMiniMaxM3,
		PresetMiniMaxM27:          modelMiniMaxM27,
		PresetMiniMaxM27Highspeed: modelMiniMaxM27Highspeed,
		PresetMiniMaxM25:          modelMiniMaxM25,
		PresetMiniMaxM25Highspeed: modelMiniMaxM25Highspeed,
		PresetMiniMaxM21:          modelMiniMaxM21,
		PresetMiniMaxM21Highspeed: modelMiniMaxM21Highspeed,
		PresetMiniMaxM2:           modelMiniMaxM2,
	},
}
