package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderOpenAIResponses spec.ProviderName = "openairesponses"

	DisplayNameProviderOpenAIResponses = "OpenAI Responses API"
)

const (
	ModelNameGPT6Astra     spec.ModelName = "gpt-6-astra"
	ModelNameGPT56Sol      spec.ModelName = "gpt-5.6-sol"
	ModelNameGPT56Terra    spec.ModelName = "gpt-5.6-terra"
	ModelNameGPT56Luna     spec.ModelName = "gpt-5.6-luna"
	ModelNameGPT55         spec.ModelName = "gpt-5.5"
	ModelNameGPT54         spec.ModelName = "gpt-5.4"
	ModelNameGPT54Mini     spec.ModelName = "gpt-5.4-mini"
	ModelNameGPT54Nano     spec.ModelName = "gpt-5.4-nano"
	ModelNameGPT53Codex    spec.ModelName = "gpt-5.3-codex"
	ModelNameGPT52         spec.ModelName = "gpt-5.2"
	ModelNameGPT52Codex    spec.ModelName = "gpt-5.2-codex"
	ModelNameGPT51         spec.ModelName = "gpt-5.1"
	ModelNameGPT51Codex    spec.ModelName = "gpt-5.1-codex"
	ModelNameGPT51CodexMax spec.ModelName = "gpt-5.1-codex-max"
	ModelNameGPT5Mini      spec.ModelName = "gpt-5-mini"

	ModelNameGPT41     spec.ModelName = "gpt-4.1"
	ModelNameGPT41Mini spec.ModelName = "gpt-4.1-mini"
	ModelNameGPT4o     spec.ModelName = "gpt-4o"
	ModelNameGPT4oMini spec.ModelName = "gpt-4o-mini"

	ModelNameGPTOSS20BRepo  spec.ModelName = "openai/gpt-oss-20b"
	ModelNameGPTOSS20BLocal spec.ModelName = "gpt-oss-20b"

	ModelNameGPTOSS120BFireworksAI spec.ModelName = "openai/gpt-oss-120b:fireworks-ai"
	ModelNameGPTOSS20BFireworksAI  spec.ModelName = "openai/gpt-oss-20b:fireworks-ai"

	ModelNameGPTOSS20BOllama spec.ModelName = "gpt-oss:20b"
)

const (
	DisplayNameGPT6Astra     = "GPT 6 Astra"
	DisplayNameGPT56Sol      = "GPT 5.6 Sol"
	DisplayNameGPT56Terra    = "GPT 5.6 Terra"
	DisplayNameGPT56Luna     = "GPT 5.6 Luna"
	DisplayNameGPT55         = "GPT 5.5"
	DisplayNameGPT54         = "GPT 5.4"
	DisplayNameGPT54Mini     = "GPT 5.4 Mini"
	DisplayNameGPT54Nano     = "GPT 5.4 Nano"
	DisplayNameGPT53Codex    = "GPT 5.3 Codex"
	DisplayNameGPT52         = "GPT 5.2"
	DisplayNameGPT52Codex    = "GPT 5.2 Codex"
	DisplayNameGPT51         = "GPT 5.1"
	DisplayNameGPT51Codex    = "GPT 5.1 Codex"
	DisplayNameGPT51CodexMax = "GPT 5.1 Codex Max"
	DisplayNameGPT5Mini      = "GPT 5 Mini"

	DisplayNameGPT41     = "GPT 4.1"
	DisplayNameGPT41Mini = "GPT 4.1 Mini"
	DisplayNameGPT4o     = "GPT 4o"
	DisplayNameGPT4oMini = "GPT 4o Mini"

	DisplayNameGPTOSS120B = "gpt-oss 120B"
	DisplayNameGPTOSS20B  = "gpt-oss 20B"

	DisplayNameGPTOSS120BFireworksAI = "gpt-oss 120B (Fireworks AI)"
	DisplayNameGPTOSS20BFireworksAI  = "gpt-oss 20B (Fireworks AI)"
)

const (
	PresetGPT6Astra     ModelPresetID = "gpt6Astra"
	PresetGPT56Sol      ModelPresetID = "gpt56sol"
	PresetGPT56Terra    ModelPresetID = "gpt56terra"
	PresetGPT56Luna     ModelPresetID = "gpt56luna"
	PresetGPT55         ModelPresetID = "gpt55"
	PresetGPT54         ModelPresetID = "gpt54"
	PresetGPT54Mini     ModelPresetID = "gpt54mini"
	PresetGPT54Nano     ModelPresetID = "gpt54nano"
	PresetGPT53Codex    ModelPresetID = "gpt53Codex"
	PresetGPT52         ModelPresetID = "gpt52"
	PresetGPT52Codex    ModelPresetID = "gpt52Codex"
	PresetGPT51         ModelPresetID = "gpt51"
	PresetGPT51Codex    ModelPresetID = "gpt51Codex"
	PresetGPT51CodexMax ModelPresetID = "gpt51CodexMax"
	PresetGPT5Mini      ModelPresetID = "gpt5Mini"

	PresetGPT41     ModelPresetID = "gpt41"
	PresetGPT41Mini ModelPresetID = "gpt41Mini"
	PresetGPT4o     ModelPresetID = "gpt4o"
	PresetGPT4oMini ModelPresetID = "gpt4oMini"

	PresetGPTOSS120B ModelPresetID = "gptoss120b"
	PresetGPTOSS20B  ModelPresetID = "gptoss20b"

	PresetGPTOSS120BFireworksAI ModelPresetID = "gptoss120bFireworksAI"
	PresetGPTOSS20BFireworksAI  ModelPresetID = "gptoss20bFireworksAI"
)

var modelOpenAIResponsesGPT6Astra = ModelPreset{
	ID:          PresetGPT6Astra,
	Name:        ModelNameGPT6Astra,
	DisplayName: DisplayNameGPT6Astra,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT6Astra,
		Stream:          true,
		MaxPromptLength: 1000000,
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
		spec.ReasoningLevelMax,
	}),
}

var modelOpenAIResponsesGPT56Sol = ModelPreset{
	ID:          PresetGPT56Sol,
	Name:        ModelNameGPT56Sol,
	DisplayName: DisplayNameGPT56Sol,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT56Sol,
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
		spec.ReasoningLevelMax,
	}),
}

var modelOpenAIResponsesGPT56Terra = ModelPreset{
	ID:          PresetGPT56Terra,
	Name:        ModelNameGPT56Terra,
	DisplayName: DisplayNameGPT56Terra,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT56Terra,
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
		spec.ReasoningLevelMax,
	}),
}

var modelOpenAIResponsesGPT56Luna = ModelPreset{
	ID:          PresetGPT56Luna,
	Name:        ModelNameGPT56Luna,
	DisplayName: DisplayNameGPT56Luna,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT56Luna,
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
		spec.ReasoningLevelMax,
	}),
}

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
			SupportsReasoningContext:         new(true),
			SupportsReasoningMode:            new(true),
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
		PresetGPT6Astra:     modelOpenAIResponsesGPT6Astra,
		PresetGPT56Sol:      modelOpenAIResponsesGPT56Sol,
		PresetGPT56Terra:    modelOpenAIResponsesGPT56Terra,
		PresetGPT56Luna:     modelOpenAIResponsesGPT56Luna,
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
