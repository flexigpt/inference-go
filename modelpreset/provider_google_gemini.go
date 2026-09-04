package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderGoogleGemini spec.ProviderName = "googlegemini"

	DisplayNameProviderGoogleGemini = "Google Gemini API"
)

const (
	ModelNameGemini38Flash     spec.ModelName = "gemini-3.8-flash"
	ModelNameGemini37Flash     spec.ModelName = "gemini-3.7-flash"
	ModelNameGemini36Flash     spec.ModelName = "gemini-3.6-flash"
	ModelNameGemini35Flash     spec.ModelName = "gemini-3.5-flash"
	ModelNameGemini35FlashLite spec.ModelName = "gemini-3.5-flash-lite"
	ModelNameGemini31Pro       spec.ModelName = "gemini-3.1-pro-preview"
	ModelNameGemini31FlashLite spec.ModelName = "gemini-3.1-flash-lite"
	ModelNameGemini3Flash      spec.ModelName = "gemini-3-flash-preview"
	ModelNameGemini25Flash     spec.ModelName = "gemini-2.5-flash"
	ModelNameGemini25FlashLite spec.ModelName = "gemini-2.5-flash-lite-preview-06-17"

	ModelNameGemma426BA4BRepo  spec.ModelName = "google/gemma-4-26b-a4b"
	ModelNameGemma426BA4BLocal spec.ModelName = "gemma4-26b-a4b"

	ModelNameGemma426BOllama spec.ModelName = "gemma4:26b"
	ModelNameGemma4E4BOllama spec.ModelName = "gemma4:e4b"
)

const (
	DisplayNameGemini38Flash     = "Gemini 3.8 Flash"
	DisplayNameGemini37Flash     = "Gemini 3.7 Flash"
	DisplayNameGemini36Flash     = "Gemini 3.6 Flash"
	DisplayNameGemini35Flash     = "Gemini 3.5 Flash"
	DisplayNameGemini35FlashLite = "Gemini 3.5 Flash Lite"
	DisplayNameGemini31Pro       = "Gemini 3.1 Pro"
	DisplayNameGemini31FlashLite = "Gemini 3.1 Flash Lite"
	DisplayNameGemini3Flash      = "Gemini 3 Flash"
	DisplayNameGemini25Flash     = "Gemini 2.5 Flash"
	DisplayNameGemini25FlashLite = "Gemini 2.5 Flash Lite"

	DisplayNameGemma426BA4B = "Gemma 4 26B A4B"
	DisplayNameGemma426B    = "Gemma 4 26B"
	DisplayNameGemma4E4B    = "Gemma 4 E4B"
)

const (
	PresetGemini38Flash     ModelPresetID = "gemini38Flash"
	PresetGemini37Flash     ModelPresetID = "gemini37Flash"
	PresetGemini36Flash     ModelPresetID = "gemini36Flash"
	PresetGemini35Flash     ModelPresetID = "gemini35Flash"
	PresetGemini35FlashLite ModelPresetID = "gemini35FlashLite"
	PresetGemini31Pro       ModelPresetID = "gemini31Pro"
	PresetGemini31FlashLite ModelPresetID = "gemini31FlashLite"
	PresetGemini3Flash      ModelPresetID = "gemini3Flash"
	PresetGemini25Flash     ModelPresetID = "gemini25Flash"
	PresetGemini25FlashLite ModelPresetID = "gemini25FlashLite"

	PresetGemma426BA4B ModelPresetID = "gemma426ba4b"
	PresetGemma426B    ModelPresetID = "gemma426b"
	PresetGemma4E4B    ModelPresetID = "gemma4e4b"
)

var modelGoogleGemini38Flash = ModelPreset{
	ID:          PresetGemini38Flash,
	Name:        ModelNameGemini38Flash,
	DisplayName: DisplayNameGemini38Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini38Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini37Flash = ModelPreset{
	ID:          PresetGemini37Flash,
	Name:        ModelNameGemini37Flash,
	DisplayName: DisplayNameGemini37Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini37Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini36Flash = ModelPreset{
	ID:          PresetGemini36Flash,
	Name:        ModelNameGemini36Flash,
	DisplayName: DisplayNameGemini36Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini36Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini35Flash = ModelPreset{
	ID:          PresetGemini35Flash,
	Name:        ModelNameGemini35Flash,
	DisplayName: DisplayNameGemini35Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini35Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini35FlashLite = ModelPreset{
	ID:          PresetGemini35FlashLite,
	Name:        ModelNameGemini35FlashLite,
	DisplayName: DisplayNameGemini35FlashLite,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini35FlashLite,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini31Pro = ModelPreset{
	ID:          PresetGemini31Pro,
	Name:        ModelNameGemini31Pro,
	DisplayName: DisplayNameGemini31Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini31Pro,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini31FlashLite = ModelPreset{
	ID:          PresetGemini31FlashLite,
	Name:        ModelNameGemini31FlashLite,
	DisplayName: DisplayNameGemini31FlashLite,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini31FlashLite,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini3Flash = ModelPreset{
	ID:          PresetGemini3Flash,
	Name:        ModelNameGemini3Flash,
	DisplayName: DisplayNameGemini3Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini3Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini25Flash = ModelPreset{
	ID:          PresetGemini25Flash,
	Name:        ModelNameGemini25Flash,
	DisplayName: DisplayNameGemini25Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini25Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1),
				MaxAllowed:      new(24576),
				ZeroAllowed:     new(true),
				MinusOneAllowed: new(true),
			},
		},
	},
}

var modelGoogleGemini25FlashLite = ModelPreset{
	ID:          PresetGemini25FlashLite,
	Name:        ModelNameGemini25FlashLite,
	DisplayName: DisplayNameGemini25FlashLite,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini25FlashLite,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(512),
				MaxAllowed:      new(24576),
				ZeroAllowed:     new(true),
				MinusOneAllowed: new(true),
			},
		},
	},
}

var providerGoogleGemini = ProviderPreset{
	Name:                     ProviderGoogleGemini,
	DisplayName:              DisplayNameProviderGoogleGemini,
	SDKType:                  spec.ProviderSDKTypeGoogleGenerateContent,
	Origin:                   spec.DefaultGoogleGenerateContentOrigin,
	ChatCompletionPathPrefix: spec.DefaultGoogleGenerateContentPrefix,
	APIKeyHeaderKey:          spec.DefaultGoogleGenerateContentAPIKeyHeaderKey,
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
				spec.ReasoningTypeHybridWithTokens,
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
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1),
				MaxAllowed:      new(32768),
				ZeroAllowed:     new(true),
				MinusOneAllowed: new(true),
			},
			SupportsSummaryStyle:             new(true),
			SupportsReasoningContext:         new(false),
			SupportsReasoningMode:            new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
		StopSequenceCapabilities: &capabilityoverride.StopSequenceCapabilitiesOverride{
			IsSupported:             new(true),
			DisallowedWithReasoning: new(false),
			MaxSequences:            new(4),
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
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetGemini38Flash:     modelGoogleGemini38Flash,
		PresetGemini37Flash:     modelGoogleGemini37Flash,
		PresetGemini36Flash:     modelGoogleGemini36Flash,
		PresetGemini35Flash:     modelGoogleGemini35Flash,
		PresetGemini35FlashLite: modelGoogleGemini35FlashLite,
		PresetGemini3Flash:      modelGoogleGemini3Flash,
		PresetGemini31Pro:       modelGoogleGemini31Pro,
		PresetGemini31FlashLite: modelGoogleGemini31FlashLite,
		PresetGemini25Flash:     modelGoogleGemini25Flash,
		PresetGemini25FlashLite: modelGoogleGemini25FlashLite,
	},
}
