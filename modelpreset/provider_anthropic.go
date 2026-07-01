package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetAnthropicFable5   ModelPresetID = "fable5"
	PresetAnthropicOpus48   ModelPresetID = "opus48"
	PresetAnthropicOpus47   ModelPresetID = "opus47"
	PresetAnthropicOpus46   ModelPresetID = "opus46"
	PresetAnthropicOpus45   ModelPresetID = "opus45"
	PresetAnthropicOpus41   ModelPresetID = "opus41"
	PresetAnthropicSonnet5  ModelPresetID = "sonnet5"
	PresetAnthropicSonnet46 ModelPresetID = "sonnet46"
	PresetAnthropicSonnet45 ModelPresetID = "sonnet45"
	PresetAnthropicSonnet4  ModelPresetID = "sonnet4"
	PresetAnthropicHaiku45  ModelPresetID = "haiku45"
)

const (
	ModelNameAnthropicFable5   spec.ModelName = "claude-fable-5"
	ModelNameAnthropicOpus48   spec.ModelName = "claude-opus-4-8"
	ModelNameAnthropicOpus47   spec.ModelName = "claude-opus-4-7"
	ModelNameAnthropicOpus46   spec.ModelName = "claude-opus-4-6"
	ModelNameAnthropicOpus45   spec.ModelName = "claude-opus-4-5-20251101"
	ModelNameAnthropicOpus41   spec.ModelName = "claude-opus-4-1-20250805"
	ModelNameAnthropicSonnet5  spec.ModelName = "claude-sonnet-5"
	ModelNameAnthropicSonnet46 spec.ModelName = "claude-sonnet-4-6"
	ModelNameAnthropicSonnet45 spec.ModelName = "claude-sonnet-4-5-20250929"
	ModelNameAnthropicSonnet4  spec.ModelName = "claude-sonnet-4-20250514"
	ModelNameAnthropicHaiku45  spec.ModelName = "claude-haiku-4-5-20251001"
)

var modelAnthropicFable5 = ModelPreset{
	ID:          PresetAnthropicFable5,
	Name:        ModelNameAnthropicFable5,
	DisplayName: "Claude Fable 5",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicFable5,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax,
			},
		},
	},
}

var modelAnthropicOpus48 = ModelPreset{
	ID:          PresetAnthropicOpus48,
	Name:        ModelNameAnthropicOpus48,
	DisplayName: "Claude Opus 4.8",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicOpus48,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax,
			},
		},
	},
}

var modelAnthropicOpus47 = ModelPreset{
	ID:          PresetAnthropicOpus47,
	Name:        ModelNameAnthropicOpus47,
	DisplayName: "Claude Opus 4.7",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicOpus47,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax,
			},
		},
	},
}

var modelAnthropicOpus46 = ModelPreset{
	ID:          PresetAnthropicOpus46,
	Name:        ModelNameAnthropicOpus46,
	DisplayName: "Claude Opus 4.6",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicOpus46,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelMax,
			},
		},
	},
}

var modelAnthropicOpus45 = ModelPreset{
	ID:          PresetAnthropicOpus45,
	Name:        ModelNameAnthropicOpus45,
	DisplayName: "Claude Opus 4.5",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicOpus45,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 64000,
		Temperature:     new(1.0),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
		},
	},
}

var modelAnthropicOpus41 = ModelPreset{
	ID:          PresetAnthropicOpus41,
	Name:        ModelNameAnthropicOpus41,
	DisplayName: "Claude Opus 4.1",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicOpus41,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 32000,
		Temperature:     new(0.1),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
		},
	},
}

var modelAnthropicSonnet5 = ModelPreset{
	ID:          PresetAnthropicSonnet5,
	Name:        ModelNameAnthropicSonnet5,
	DisplayName: "Claude Sonnet 5",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicSonnet5,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     new(0.1),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelMax,
			},
		},
	},
}

var modelAnthropicSonnet46 = ModelPreset{
	ID:          PresetAnthropicSonnet46,
	Name:        ModelNameAnthropicSonnet46,
	DisplayName: "Claude Sonnet 4.6",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicSonnet46,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 64000,
		Temperature:     new(0.1),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelMax,
			},
		},
	},
}

var modelAnthropicSonnet45 = ModelPreset{
	ID:          PresetAnthropicSonnet45,
	Name:        ModelNameAnthropicSonnet45,
	DisplayName: "Claude Sonnet 4.5",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicSonnet45,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 64000,
		Temperature:     new(0.1),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
		},
	},
}

var modelAnthropicSonnet4 = ModelPreset{
	ID:          PresetAnthropicSonnet4,
	Name:        ModelNameAnthropicSonnet4,
	DisplayName: "Claude Sonnet 4",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicSonnet4,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 64000,
		Temperature:     new(0.1),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
}

var modelAnthropicHaiku45 = ModelPreset{
	ID:          PresetAnthropicHaiku45,
	Name:        ModelNameAnthropicHaiku45,
	DisplayName: "Claude Haiku 4.5",
	ModelParam: spec.ModelParam{
		Name:            ModelNameAnthropicHaiku45,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 64000,
		Temperature:     new(0.1),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
		CacheControl:    cacheEphemeral5m(),
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
		},
	},
}

var providerAnthropic = ProviderPreset{
	Name:                     ProviderAnthropic,
	DisplayName:              "Anthropic",
	SDKType:                  spec.ProviderSDKTypeAnthropic,
	Origin:                   spec.DefaultAnthropicOrigin,
	ChatCompletionPathPrefix: spec.DefaultAnthropicChatCompletionPrefix,
	APIKeyHeaderKey:          spec.DefaultAnthropicAuthorizationHeaderKey,
	DefaultHeaders: map[string]string{
		spec.DefaultContentTypeHeaderKey: spec.DefaultContentTypeHeader,
		"accept":                         spec.DefaultContentTypeHeader,
		"anthropic-version":              "2023-06-01",
	},
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
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(true),
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
			SupportsVerbosity: new(true),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
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
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetAnthropicFable5:   modelAnthropicFable5,
		PresetAnthropicOpus48:   modelAnthropicOpus48,
		PresetAnthropicOpus47:   modelAnthropicOpus47,
		PresetAnthropicOpus46:   modelAnthropicOpus46,
		PresetAnthropicOpus45:   modelAnthropicOpus45,
		PresetAnthropicOpus41:   modelAnthropicOpus41,
		PresetAnthropicSonnet5:  modelAnthropicSonnet5,
		PresetAnthropicSonnet46: modelAnthropicSonnet46,
		PresetAnthropicSonnet45: modelAnthropicSonnet45,
		PresetAnthropicSonnet4:  modelAnthropicSonnet4,
		PresetAnthropicHaiku45:  modelAnthropicHaiku45,
	},
}
