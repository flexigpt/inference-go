package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderAnthropic spec.ProviderName = "anthropic"

	DisplayNameProviderAnthropic = "Anthropic"
)

const (
	ModelNameClaudeFable5   spec.ModelName = "claude-fable-5"
	ModelNameClaudeOpus5    spec.ModelName = "claude-opus-5"
	ModelNameClaudeOpus48   spec.ModelName = "claude-opus-4-8"
	ModelNameClaudeOpus47   spec.ModelName = "claude-opus-4-7"
	ModelNameClaudeOpus46   spec.ModelName = "claude-opus-4-6"
	ModelNameClaudeOpus45   spec.ModelName = "claude-opus-4-5-20251101"
	ModelNameClaudeOpus41   spec.ModelName = "claude-opus-4-1-20250805"
	ModelNameClaudeSonnet5  spec.ModelName = "claude-sonnet-5"
	ModelNameClaudeSonnet46 spec.ModelName = "claude-sonnet-4-6"
	ModelNameClaudeSonnet45 spec.ModelName = "claude-sonnet-4-5-20250929"
	ModelNameClaudeSonnet4  spec.ModelName = "claude-sonnet-4-20250514"
	ModelNameClaudeHaiku45  spec.ModelName = "claude-haiku-4-5-20251001"
)

const (
	DisplayNameClaudeFable5   = "Claude Fable 5"
	DisplayNameClaudeOpus5    = "Claude Opus 5"
	DisplayNameClaudeOpus48   = "Claude Opus 4.8"
	DisplayNameClaudeOpus47   = "Claude Opus 4.7"
	DisplayNameClaudeOpus46   = "Claude Opus 4.6"
	DisplayNameClaudeOpus45   = "Claude Opus 4.5"
	DisplayNameClaudeOpus41   = "Claude Opus 4.1"
	DisplayNameClaudeSonnet5  = "Claude Sonnet 5"
	DisplayNameClaudeSonnet46 = "Claude Sonnet 4.6"
	DisplayNameClaudeSonnet45 = "Claude Sonnet 4.5"
	DisplayNameClaudeSonnet4  = "Claude Sonnet 4"
	DisplayNameClaudeHaiku45  = "Claude Haiku 4.5"
)

const (
	PresetClaudeFable5   ModelPresetID = "fable5"
	PresetClaudeOpus5    ModelPresetID = "opus5"
	PresetClaudeOpus48   ModelPresetID = "opus48"
	PresetClaudeOpus47   ModelPresetID = "opus47"
	PresetClaudeOpus46   ModelPresetID = "opus46"
	PresetClaudeOpus45   ModelPresetID = "opus45"
	PresetClaudeOpus41   ModelPresetID = "opus41"
	PresetClaudeSonnet5  ModelPresetID = "sonnet5"
	PresetClaudeSonnet46 ModelPresetID = "sonnet46"
	PresetClaudeSonnet45 ModelPresetID = "sonnet45"
	PresetClaudeSonnet4  ModelPresetID = "sonnet4"
	PresetClaudeHaiku45  ModelPresetID = "haiku45"
)

var modelAnthropicFable5 = ModelPreset{
	ID:          PresetClaudeFable5,
	Name:        ModelNameClaudeFable5,
	DisplayName: DisplayNameClaudeFable5,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeFable5,
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

var modelAnthropicOpus5 = ModelPreset{
	ID:          PresetClaudeOpus5,
	Name:        ModelNameClaudeOpus5,
	DisplayName: DisplayNameClaudeOpus5,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeOpus5,
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
	ID:          PresetClaudeOpus48,
	Name:        ModelNameClaudeOpus48,
	DisplayName: DisplayNameClaudeOpus48,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeOpus48,
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
	ID:          PresetClaudeOpus47,
	Name:        ModelNameClaudeOpus47,
	DisplayName: DisplayNameClaudeOpus47,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeOpus47,
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
	ID:          PresetClaudeOpus46,
	Name:        ModelNameClaudeOpus46,
	DisplayName: DisplayNameClaudeOpus46,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeOpus46,
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
	ID:          PresetClaudeOpus45,
	Name:        ModelNameClaudeOpus45,
	DisplayName: DisplayNameClaudeOpus45,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeOpus45,
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
	ID:          PresetClaudeOpus41,
	Name:        ModelNameClaudeOpus41,
	DisplayName: DisplayNameClaudeOpus41,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeOpus41,
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
	ID:          PresetClaudeSonnet5,
	Name:        ModelNameClaudeSonnet5,
	DisplayName: DisplayNameClaudeSonnet5,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeSonnet5,
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
	ID:          PresetClaudeSonnet46,
	Name:        ModelNameClaudeSonnet46,
	DisplayName: DisplayNameClaudeSonnet46,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeSonnet46,
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
	ID:          PresetClaudeSonnet45,
	Name:        ModelNameClaudeSonnet45,
	DisplayName: DisplayNameClaudeSonnet45,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeSonnet45,
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
	ID:          PresetClaudeSonnet4,
	Name:        ModelNameClaudeSonnet4,
	DisplayName: DisplayNameClaudeSonnet4,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeSonnet4,
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
	ID:          PresetClaudeHaiku45,
	Name:        ModelNameClaudeHaiku45,
	DisplayName: DisplayNameClaudeHaiku45,
	ModelParam: spec.ModelParam{
		Name:            ModelNameClaudeHaiku45,
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
	DisplayName:              DisplayNameProviderAnthropic,
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
		PresetClaudeFable5:   modelAnthropicFable5,
		PresetClaudeOpus5:    modelAnthropicOpus5,
		PresetClaudeOpus48:   modelAnthropicOpus48,
		PresetClaudeOpus47:   modelAnthropicOpus47,
		PresetClaudeOpus46:   modelAnthropicOpus46,
		PresetClaudeOpus45:   modelAnthropicOpus45,
		PresetClaudeOpus41:   modelAnthropicOpus41,
		PresetClaudeSonnet5:  modelAnthropicSonnet5,
		PresetClaudeSonnet46: modelAnthropicSonnet46,
		PresetClaudeSonnet45: modelAnthropicSonnet45,
		PresetClaudeSonnet4:  modelAnthropicSonnet4,
		PresetClaudeHaiku45:  modelAnthropicHaiku45,
	},
}
