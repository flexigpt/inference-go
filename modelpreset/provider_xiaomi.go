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
	ModelNameMiMoV25Pro spec.ModelName = "mimo-v2.5-pro"
	ModelNameMiMoV25    spec.ModelName = "mimo-v2.5"

	ModelNameOpenRouterXiaomiMiMoV25Pro spec.ModelName = "xiaomi/mimo-v2.5-pro"
	ModelNameOpenRouterXiaomiMiMoV25    spec.ModelName = "xiaomi/mimo-v2.5"

	ModelNameMiMoV25ProDeepInfra      spec.ModelName = "XiaomiMiMo/MiMo-V2.5-Pro:deepinfra"
	ModelNameMiMoV2FlashFeatherlessAI spec.ModelName = "XiaomiMiMo/MiMo-V2-Flash:featherless-ai"
)

const (
	DisplayNameMiMoV2Flash      = "MiMo V2 Flash"
	DisplayNameMiMoV25          = "MiMo V2.5"
	DisplayNameMiMoV25Pro       = "MiMo V2.5 Pro"
	DisplayNameXiaomiMiMoV25    = "Xiaomi MiMo V2.5"
	DisplayNameXiaomiMiMoV25Pro = "Xiaomi MiMo V2.5 Pro"

	DisplayNameMiMoV2FlashFeatherlessAI = "MiMo V2 Flash (Featherless AI)"
	DisplayNameMiMoV25ProDeepInfra      = "MiMo V2.5 Pro (DeepInfra)"
)

const (
	PresetMiMoV2Flash      ModelPresetID = "mimov2flash"
	PresetMiMoV25          ModelPresetID = "mimov25"
	PresetMiMoV25Pro       ModelPresetID = "mimov25pro"
	PresetXiaomiMiMoV25    ModelPresetID = "xiaomiMiMoV25"
	PresetXiaomiMiMoV25Pro ModelPresetID = "xiaomiMiMoV25Pro"

	PresetMiMoV25ProDeepInfra      ModelPresetID = "mimov25proDeepInfra"
	PresetMiMoV2FlashFeatherlessAI ModelPresetID = "mimov2flashFeatherlessAI"
)

var modelXiaomiMiMoV25Pro = ModelPreset{
	ID:          PresetMiMoV25Pro,
	Name:        ModelNameMiMoV25Pro,
	DisplayName: DisplayNameMiMoV25Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV25Pro,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelXiaomiMiMoV25 = ModelPreset{
	ID:          PresetMiMoV25,
	Name:        ModelNameMiMoV25,
	DisplayName: DisplayNameMiMoV25,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV25,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
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
		PresetMiMoV25Pro: modelXiaomiMiMoV25Pro,
		PresetMiMoV25:    modelXiaomiMiMoV25,
	},
}
