package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetOpenAIResponsesGPT55         ModelPresetID = "gpt55"
	PresetOpenAIResponsesGPT54         ModelPresetID = "gpt54"
	PresetOpenAIResponsesGPT54Mini     ModelPresetID = "gpt54mini"
	PresetOpenAIResponsesGPT54Nano     ModelPresetID = "gpt54nano"
	PresetOpenAIResponsesGPT53Codex    ModelPresetID = "gpt53Codex"
	PresetOpenAIResponsesGPT52         ModelPresetID = "gpt52"
	PresetOpenAIResponsesGPT52Codex    ModelPresetID = "gpt52Codex"
	PresetOpenAIResponsesGPT51         ModelPresetID = "gpt51"
	PresetOpenAIResponsesGPT51Codex    ModelPresetID = "gpt51Codex"
	PresetOpenAIResponsesGPT51CodexMax ModelPresetID = "gpt51CodexMax"
	PresetOpenAIResponsesGPT5Mini      ModelPresetID = "gpt5Mini"
)

const (
	ModelNameOpenAIResponsesGPT55         spec.ModelName = "gpt-5.5"
	ModelNameOpenAIResponsesGPT54         spec.ModelName = "gpt-5.4"
	ModelNameOpenAIResponsesGPT54Mini     spec.ModelName = "gpt-5.4-mini"
	ModelNameOpenAIResponsesGPT54Nano     spec.ModelName = "gpt-5.4-nano"
	ModelNameOpenAIResponsesGPT53Codex    spec.ModelName = "gpt-5.3-codex"
	ModelNameOpenAIResponsesGPT52         spec.ModelName = "gpt-5.2"
	ModelNameOpenAIResponsesGPT52Codex    spec.ModelName = "gpt-5.2-codex"
	ModelNameOpenAIResponsesGPT51         spec.ModelName = "gpt-5.1"
	ModelNameOpenAIResponsesGPT51Codex    spec.ModelName = "gpt-5.1-codex"
	ModelNameOpenAIResponsesGPT51CodexMax spec.ModelName = "gpt-5.1-codex-max"
	ModelNameOpenAIResponsesGPT5Mini      spec.ModelName = "gpt-5-mini"
)

var modelOpenAIResponsesGPT55 = ModelPreset{
	ID:          PresetOpenAIResponsesGPT55,
	Name:        ModelNameOpenAIResponsesGPT55,
	DisplayName: "OpenAI GPT 5.5",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT55,
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
	ID:          PresetOpenAIResponsesGPT54,
	Name:        ModelNameOpenAIResponsesGPT54,
	DisplayName: "OpenAI GPT 5.4",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT54,
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
	ID:          PresetOpenAIResponsesGPT54Mini,
	Name:        ModelNameOpenAIResponsesGPT54Mini,
	DisplayName: "OpenAI GPT 5.4 Mini",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT54Mini,
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
	ID:          PresetOpenAIResponsesGPT54Nano,
	Name:        ModelNameOpenAIResponsesGPT54Nano,
	DisplayName: "OpenAI GPT 5.4 Nano",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT54Nano,
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
	ID:          PresetOpenAIResponsesGPT53Codex,
	Name:        ModelNameOpenAIResponsesGPT53Codex,
	DisplayName: "OpenAI GPT 5.3 Codex",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT53Codex,
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
	ID:          PresetOpenAIResponsesGPT52,
	Name:        ModelNameOpenAIResponsesGPT52,
	DisplayName: "OpenAI GPT 5.2",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT52,
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
	ID:          PresetOpenAIResponsesGPT52Codex,
	Name:        ModelNameOpenAIResponsesGPT52Codex,
	DisplayName: "OpenAI GPT 5.2 Codex",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT52Codex,
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
	ID:          PresetOpenAIResponsesGPT51,
	Name:        ModelNameOpenAIResponsesGPT51,
	DisplayName: "OpenAI GPT 5.1",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT51,
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
	ID:          PresetOpenAIResponsesGPT51Codex,
	Name:        ModelNameOpenAIResponsesGPT51Codex,
	DisplayName: "OpenAI GPT 5.1 Codex",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT51Codex,
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
	ID:          PresetOpenAIResponsesGPT51CodexMax,
	Name:        ModelNameOpenAIResponsesGPT51CodexMax,
	DisplayName: "OpenAI GPT 5.1 Codex Max",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT51CodexMax,
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
	ID:          PresetOpenAIResponsesGPT5Mini,
	Name:        ModelNameOpenAIResponsesGPT5Mini,
	DisplayName: "OpenAI GPT 5 Mini",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIResponsesGPT5Mini,
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
	DisplayName:              "OpenAI Responses API",
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
		PresetOpenAIResponsesGPT55:         modelOpenAIResponsesGPT55,
		PresetOpenAIResponsesGPT54:         modelOpenAIResponsesGPT54,
		PresetOpenAIResponsesGPT54Mini:     modelOpenAIResponsesGPT54Mini,
		PresetOpenAIResponsesGPT54Nano:     modelOpenAIResponsesGPT54Nano,
		PresetOpenAIResponsesGPT53Codex:    modelOpenAIResponsesGPT53Codex,
		PresetOpenAIResponsesGPT52:         modelOpenAIResponsesGPT52,
		PresetOpenAIResponsesGPT52Codex:    modelOpenAIResponsesGPT52Codex,
		PresetOpenAIResponsesGPT51:         modelOpenAIResponsesGPT51,
		PresetOpenAIResponsesGPT51Codex:    modelOpenAIResponsesGPT51Codex,
		PresetOpenAIResponsesGPT51CodexMax: modelOpenAIResponsesGPT51CodexMax,
		PresetOpenAIResponsesGPT5Mini:      modelOpenAIResponsesGPT5Mini,
	},
}
