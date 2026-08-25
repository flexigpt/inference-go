package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderMeta spec.ProviderName = "meta"

	DisplayNameProviderMeta = "Meta"
)

const (
	ModelNameMuseSpark11            spec.ModelName = "muse-spark-1.1"
	ModelNameMuseSpark12            spec.ModelName = "muse-spark-1.2"
	ModelNameMuseSpark12Contributor spec.ModelName = "muse-spark-1.2-contributor"

	ModelNameLlama4BehemothLocal spec.ModelName = "llama4-behemoth"
	ModelNameLlama4MaverickLocal spec.ModelName = "llama4-maverick"
	ModelNameLlama4ScoutLocal    spec.ModelName = "llama4-scout"
)

const (
	DisplayNameMuseSpark11            = "Muse Spark 1.1"
	DisplayNameMuseSpark12            = "Muse Spark 1.2"
	DisplayNameMuseSpark12Contributor = "Muse Spark 1.2 Contributor"

	DisplayNameLlama4Behemoth = "Llama 4 Behemoth"
	DisplayNameLlama4Maverick = "Llama 4 Maverick"
	DisplayNameLlama4Scout    = "Llama 4 Scout"
)

const (
	PresetMuseSpark11            ModelPresetID = "museSpark11"
	PresetMuseSpark12            ModelPresetID = "museSpark12"
	PresetMuseSpark12Contributor ModelPresetID = "museSpark12Contributor"

	PresetLlama4Behemoth ModelPresetID = "llama4Behemoth"
	PresetLlama4Maverick ModelPresetID = "llama4Maverick"
	PresetLlama4Scout    ModelPresetID = "llama4Scout"
)

var modelMetaMuseSpark11 = ModelPreset{
	ID:          PresetMuseSpark11,
	Name:        ModelNameMuseSpark11,
	DisplayName: DisplayNameMuseSpark11,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMuseSpark11,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelMetaMuseSpark12 = ModelPreset{
	ID:          PresetMuseSpark12,
	Name:        ModelNameMuseSpark12,
	DisplayName: DisplayNameMuseSpark12,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMuseSpark12,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelMetaMuseSpark12Contributor = ModelPreset{
	ID:          PresetMuseSpark12Contributor,
	Name:        ModelNameMuseSpark12Contributor,
	DisplayName: DisplayNameMuseSpark12Contributor,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMuseSpark12Contributor,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var providerMeta = ProviderPreset{
	Name:                     ProviderMeta,
	DisplayName:              DisplayNameProviderMeta,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://api.meta.ai",
	ChatCompletionPathPrefix: spec.DefaultOpenAIResponsesPrefix,
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
			spec.ModalityFileIn,
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
				spec.ReasoningLevelXHigh,
			},
			SupportsSummaryStyle:             new(true),
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
			},
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
		CacheCapabilities: &capabilityoverride.CacheCapabilitiesOverride{
			SupportsAutomaticCaching: new(true),
			TopLevel: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL: new(true),
				SupportedKinds: []spec.CacheControlKind{
					spec.CacheControlKindEphemeral,
				},
				SupportedTTLs: []spec.CacheControlTTL{
					spec.CacheControlTTLInMemory,
					spec.CacheControlTTL24h,
				},
				SupportsKey: new(true),
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetMuseSpark11:            modelMetaMuseSpark11,
		PresetMuseSpark12:            modelMetaMuseSpark12,
		PresetMuseSpark12Contributor: modelMetaMuseSpark12Contributor,
	},
}
