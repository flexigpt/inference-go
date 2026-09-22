package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderXiaomi spec.ProviderName = "xiaomi"

	DisplayNameProviderXiaomi = "Xiaomi"
)

const (
	ModelNameMiMoV26Pro   spec.ModelName = "mimo-v2.6-pro"
	ModelNameMiMoV26Flash spec.ModelName = "mimo-v2.6-flash"
	ModelNameMiMoV25Pro   spec.ModelName = "mimo-v2.5-pro"
	ModelNameMiMoV25      spec.ModelName = "mimo-v2.5"

	ModelNameOpenRouterXiaomiMiMoV26Pro   spec.ModelName = "xiaomi/mimo-v2.6-pro"
	ModelNameOpenRouterXiaomiMiMoV26Flash spec.ModelName = "xiaomi/mimo-v2.6-flash"
	ModelNameOpenRouterXiaomiMiMoV25Pro   spec.ModelName = "xiaomi/mimo-v2.5-pro"
	ModelNameOpenRouterXiaomiMiMoV25      spec.ModelName = "xiaomi/mimo-v2.5"

	ModelNameMiMoV25ProDeepInfra      spec.ModelName = "XiaomiMiMo/MiMo-V2.5-Pro:deepinfra"
	ModelNameMiMoV2FlashFeatherlessAI spec.ModelName = "XiaomiMiMo/MiMo-V2-Flash:featherless-ai"
)

const (
	DisplayNameMiMoV26Pro   = "MiMo V2.6 Pro"
	DisplayNameMiMoV26Flash = "MiMo V2.6 Flash"
	DisplayNameMiMoV25Pro   = "MiMo V2.5 Pro"
	DisplayNameMiMoV25      = "MiMo V2.5"

	DisplayNameXiaomiMiMoV26Pro   = "Xiaomi MiMo V2.6 Pro"
	DisplayNameXiaomiMiMoV26Flash = "Xiaomi MiMo V2.5 Flash"
	DisplayNameXiaomiMiMoV25      = "Xiaomi MiMo V2.5"
	DisplayNameXiaomiMiMoV25Pro   = "Xiaomi MiMo V2.5 Pro"

	DisplayNameMiMoV2Flash = "MiMo V2 Flash"

	DisplayNameMiMoV2FlashFeatherlessAI = "MiMo V2 Flash (Featherless AI)"
	DisplayNameMiMoV25ProDeepInfra      = "MiMo V2.5 Pro (DeepInfra)"
)

const (
	PresetMiMoV26Pro   ModelPresetID = "mimov26pro"
	PresetMiMoV26Flash ModelPresetID = "mimov26flash"
	PresetMiMoV25Pro   ModelPresetID = "mimov25pro"
	PresetMiMoV25      ModelPresetID = "mimov25"

	PresetXiaomiMiMoV26Pro   ModelPresetID = "xiaomiMiMoV26Pro"
	PresetXiaomiMiMoV26Flash ModelPresetID = "xiaomiMiMoV26Flash"
	PresetXiaomiMiMoV25      ModelPresetID = "xiaomiMiMoV25"
	PresetXiaomiMiMoV25Pro   ModelPresetID = "xiaomiMiMoV25Pro"

	PresetMiMoV2Flash ModelPresetID = "mimov2flash"

	PresetMiMoV25ProDeepInfra      ModelPresetID = "mimov25proDeepInfra"
	PresetMiMoV2FlashFeatherlessAI ModelPresetID = "mimov2flashFeatherlessAI"
)

var modelMiMoV26Pro = ModelPreset{
	ID:          PresetMiMoV26Pro,
	Name:        ModelNameMiMoV26Pro,
	DisplayName: DisplayNameMiMoV26Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV26Pro,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     nil,
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
	},
}

var modelMiMoV26Flash = ModelPreset{
	ID:          PresetMiMoV26Flash,
	Name:        ModelNameMiMoV26Flash,
	DisplayName: DisplayNameMiMoV26Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV26Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     nil,
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
	},
}

var modelMiMoV25Pro = ModelPreset{
	ID:          PresetMiMoV25Pro,
	Name:        ModelNameMiMoV25Pro,
	DisplayName: DisplayNameMiMoV25Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV25Pro,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelMiMoV25 = ModelPreset{
	ID:          PresetMiMoV25,
	Name:        ModelNameMiMoV25,
	DisplayName: DisplayNameMiMoV25,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV25,
		Stream:          true,
		MaxPromptLength: 32000,
		MaxOutputLength: 32000,
		Temperature:     nil,
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
	},
}

var providerXiaomi = ProviderPreset{
	Name:                     ProviderXiaomi,
	DisplayName:              DisplayNameProviderXiaomi,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://api.xiaomimimo.com",
	ChatCompletionPathPrefix: spec.DefaultOpenAIResponsesPrefix,
	APIKeyHeaderKey:          "api-key",
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
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle:             new(false),
			SupportsReasoningContext:         new(false),
			SupportsReasoningMode:            new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(true),
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
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
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
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetMiMoV26Pro:   modelMiMoV26Pro,
		PresetMiMoV26Flash: modelMiMoV26Flash,
		PresetMiMoV25Pro:   modelMiMoV25Pro,
		PresetMiMoV25:      modelMiMoV25,
	},
}
