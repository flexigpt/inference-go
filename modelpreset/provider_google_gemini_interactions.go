package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderGoogleGeminiInteractions spec.ProviderName = "googlegeminiinteractions"

	DisplayNameProviderGoogleGeminiInteractions = "Google Gemini Interactions API"
)

const (
	ModelNameGemini4Neon     spec.ModelName = "gemini-4-neon"
	ModelNameGemini4Carbon   spec.ModelName = "gemini-4-carbon"
	ModelNameGemini4Hydrogen spec.ModelName = "gemini-4-hydrogen"
	ModelNameGemini4Argon    spec.ModelName = "gemini-4-argon"
)

const (
	DisplayNameGemini4Neon     = "Gemini 4 Neon"
	DisplayNameGemini4Carbon   = "Gemini 4 Carbon"
	DisplayNameGemini4Hydrogen = "Gemini 4 Hydrogen"
	DisplayNameGemini4Argon    = "Gemini 4 Argon"
)

const (
	PresetGemini4Neon     ModelPresetID = "gemini4Neon"
	PresetGemini4Carbon   ModelPresetID = "gemini4Carbon"
	PresetGemini4Hydrogen ModelPresetID = "gemini4Hydrogen"
	PresetGemini4Argon    ModelPresetID = "gemini4Argon"
)

var modelGoogleGemini4Neon = ModelPreset{
	ID:          PresetGemini4Neon,
	Name:        ModelNameGemini4Neon,
	DisplayName: DisplayNameGemini4Neon,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini4Neon,
		Stream:          true,
		MaxPromptLength: 983040,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
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
				spec.ReasoningLevelXHigh,
			},
		},
	},
}

var modelGoogleGemini4Carbon = ModelPreset{
	ID:          PresetGemini4Carbon,
	Name:        ModelNameGemini4Carbon,
	DisplayName: DisplayNameGemini4Carbon,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini4Carbon,
		Stream:          true,
		MaxPromptLength: 983040,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
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
		},
	},
}

var modelGoogleGemini4Hydrogen = ModelPreset{
	ID:          PresetGemini4Hydrogen,
	Name:        ModelNameGemini4Hydrogen,
	DisplayName: DisplayNameGemini4Hydrogen,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini4Hydrogen,
		Stream:          true,
		MaxPromptLength: 983040,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
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
				spec.ReasoningLevelXHigh,
			},
		},
	},
}

var modelGoogleGemini4Argon = ModelPreset{
	ID:          PresetGemini4Argon,
	Name:        ModelNameGemini4Argon,
	DisplayName: DisplayNameGemini4Argon,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini4Argon,
		Stream:          true,
		MaxPromptLength: 983040,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
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
		},
	},
}

// Gemini 4 uses the Interactions API. The adapter owns the base capability
// profile. Model-specific reasoning limits remain on the existing Gemini 4
// model presets.
var providerGoogleGeminiInteractions = ProviderPreset{
	Name:                     ProviderGoogleGeminiInteractions,
	DisplayName:              DisplayNameProviderGoogleGeminiInteractions,
	SDKType:                  spec.ProviderSDKTypeGoogleInteractions,
	Origin:                   spec.DefaultGoogleInteractionsOrigin,
	ChatCompletionPathPrefix: spec.DefaultGoogleInteractionsPrefix,
	APIKeyHeaderKey:          spec.DefaultGoogleInteractionsAPIKeyHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetGemini4Neon:     modelGoogleGemini4Neon,
		PresetGemini4Carbon:   modelGoogleGemini4Carbon,
		PresetGemini4Hydrogen: modelGoogleGemini4Hydrogen,
		PresetGemini4Argon:    modelGoogleGemini4Argon,
	},
}
