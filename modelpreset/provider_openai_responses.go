package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelOpenAIResponsesGPT55 = ModelPreset{
	ID:          PresetGPT55,
	Name:        ModelNameGPT55,
	DisplayName: DisplayNameGPT55,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT55,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelNone,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh,
	}),
}

var modelOpenAIResponsesGPT54 = ModelPreset{
	ID:          PresetGPT54,
	Name:        ModelNameGPT54,
	DisplayName: DisplayNameGPT54,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT54,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelNone,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh,
	}),
}

var modelOpenAIResponsesGPT54Mini = ModelPreset{
	ID:          PresetGPT54Mini,
	Name:        ModelNameGPT54Mini,
	DisplayName: DisplayNameGPT54Mini,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT54Mini,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelNone,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh,
	}),
}

var modelOpenAIResponsesGPT54Nano = ModelPreset{
	ID:          PresetGPT54Nano,
	Name:        ModelNameGPT54Nano,
	DisplayName: DisplayNameGPT54Nano,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT54Nano,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelNone,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh,
	}),
}

var modelOpenAIResponsesGPT53Codex = ModelPreset{
	ID:          PresetGPT53Codex,
	Name:        ModelNameGPT53Codex,
	DisplayName: DisplayNameGPT53Codex,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT53Codex,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh,
	}),
}

var modelOpenAIResponsesGPT52 = ModelPreset{
	ID:          PresetGPT52,
	Name:        ModelNameGPT52,
	DisplayName: DisplayNameGPT52,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT52,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelNone,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh,
	}),
}

var modelOpenAIResponsesGPT52Codex = ModelPreset{
	ID:          PresetGPT52Codex,
	Name:        ModelNameGPT52Codex,
	DisplayName: DisplayNameGPT52Codex,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT52Codex,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
		spec.ReasoningLevelXHigh,
	}),
}

var modelOpenAIResponsesGPT51 = ModelPreset{
	ID:          PresetGPT51,
	Name:        ModelNameGPT51,
	DisplayName: DisplayNameGPT51,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT51,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelNone,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
	}),
}

var modelOpenAIResponsesGPT51Codex = ModelPreset{
	ID:          PresetGPT51Codex,
	Name:        ModelNameGPT51Codex,
	DisplayName: DisplayNameGPT51Codex,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT51Codex,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
	}),
}

var modelOpenAIResponsesGPT51CodexMax = ModelPreset{
	ID:          PresetGPT51CodexMax,
	Name:        ModelNameGPT51CodexMax,
	DisplayName: DisplayNameGPT51CodexMax,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT51CodexMax,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
	}),
}

var modelOpenAIResponsesGPT5Mini = ModelPreset{
	ID:          PresetGPT5Mini,
	Name:        ModelNameGPT5Mini,
	DisplayName: DisplayNameGPT5Mini,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT5Mini,
		Stream:          true,
		MaxPromptLength: 400000,
		MaxOutputLength: 128000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIResponsesReasoningOverride([]spec.ReasoningLevel{
		spec.ReasoningLevelMinimal,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
	}),
}

func openAIResponsesReasoningOverride(levels []spec.ReasoningLevel) *capabilityoverride.ModelCapabilitiesOverride {
	return &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: levels,
			SupportsSummaryStyle:     new(true),
		},
	}
}

var providerOpenAIResponses = ProviderPreset{
	Name:                     ProviderOpenAIResponses,
	DisplayName:              DisplayNameProviderOpenAIResponses,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   spec.DefaultOpenAIOrigin,
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
			SupportedReasoningTypes: []spec.ReasoningType{
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
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetGPT55:         modelOpenAIResponsesGPT55,
		PresetGPT54:         modelOpenAIResponsesGPT54,
		PresetGPT54Mini:     modelOpenAIResponsesGPT54Mini,
		PresetGPT54Nano:     modelOpenAIResponsesGPT54Nano,
		PresetGPT53Codex:    modelOpenAIResponsesGPT53Codex,
		PresetGPT52:         modelOpenAIResponsesGPT52,
		PresetGPT52Codex:    modelOpenAIResponsesGPT52Codex,
		PresetGPT51:         modelOpenAIResponsesGPT51,
		PresetGPT51Codex:    modelOpenAIResponsesGPT51Codex,
		PresetGPT51CodexMax: modelOpenAIResponsesGPT51CodexMax,
		PresetGPT5Mini:      modelOpenAIResponsesGPT5Mini,
	},
}
