package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderMoonshot spec.ProviderName = "moonshot"

	DisplayNameProviderMoonshot = "Moonshot AI"
)

const (
	ModelNameMoonshotKimiK3               spec.ModelName = "kimi-k3"
	ModelNameMoonshotKimiK27Code          spec.ModelName = "kimi-k2.7-code"
	ModelNameMoonshotKimiK27CodeHighspeed spec.ModelName = "kimi-k2.7-code-highspeed"
	ModelNameMoonshotKimiK26              spec.ModelName = "kimi-k2.6"

	ModelNameOpenRouterMoonshotKimiK26     spec.ModelName = "moonshotai/kimi-k2.6"
	ModelNameOpenRouterMoonshotKimiK27Code spec.ModelName = "moonshotai/kimi-k2.7-code"

	ModelNameKimiK2InstructNovita        spec.ModelName = "moonshotai/Kimi-K2-Instruct:novita"
	ModelNameKimiK2Instruct0905Novita    spec.ModelName = "moonshotai/Kimi-K2-Instruct-0905:novita"
	ModelNameKimiK2ThinkingFeatherlessAI spec.ModelName = "moonshotai/Kimi-K2-Thinking:featherless-ai"
)

const (
	DisplayNameMoonshotKimiK3               = "Kimi K3"
	DisplayNameMoonshotKimiK27Code          = "Kimi K2.7 Code"
	DisplayNameMoonshotKimiK27CodeHighspeed = "Kimi K2.7 Code Highspeed"
	DisplayNameMoonshotKimiK26              = "Kimi K2.6"

	DisplayNameKimiK2Thinking     = "Kimi K2 Thinking"
	DisplayNameKimiK2Instruct     = "Kimi K2 Instruct"
	DisplayNameKimiK2Instruct0905 = "Kimi K2 Instruct 0905"

	DisplayNameKimiK2InstructNovita        = "Kimi K2 Instruct (Novita)"
	DisplayNameKimiK2Instruct0905Novita    = "Kimi K2 Instruct 0905 (Novita)"
	DisplayNameKimiK2ThinkingFeatherlessAI = "Kimi K2 Thinking (Featherless AI)"
)

const (
	PresetMoonshotKimiK3               ModelPresetID = "kimiK3"
	PresetMoonshotKimiK27Code          ModelPresetID = "kimiK27Code"
	PresetMoonshotKimiK27CodeHighspeed ModelPresetID = "kimiK27CodeHighspeed"
	PresetMoonshotKimiK26              ModelPresetID = "kimiK26"

	PresetKimiK2Thinking     ModelPresetID = "kimik2thinking"
	PresetKimiK2Instruct     ModelPresetID = "kimik2instruct"
	PresetKimiK2Instruct0905 ModelPresetID = "kimik2instruct0905"

	PresetKimiK2InstructNovita        ModelPresetID = "kimik2instructNovita"
	PresetKimiK2Instruct0905Novita    ModelPresetID = "kimik2instruct0905Novita"
	PresetKimiK2ThinkingFeatherlessAI ModelPresetID = "kimik2thinkingFeatherlessAI"
)

var modelMoonshotKimiK3 = ModelPreset{
	ID:          PresetMoonshotKimiK3,
	Name:        ModelNameMoonshotKimiK3,
	DisplayName: DisplayNameMoonshotKimiK3,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMoonshotKimiK3,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMax),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelMax,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(true),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
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
}

var modelMoonshotKimiK27Code = ModelPreset{
	ID:          PresetMoonshotKimiK27Code,
	Name:        ModelNameMoonshotKimiK27Code,
	DisplayName: DisplayNameMoonshotKimiK27Code,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMoonshotKimiK27Code,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 32768,
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1024),
				ZeroAllowed:     new(false),
				MinusOneAllowed: new(false),
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(true),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelMoonshotKimiK27CodeHighspeed = ModelPreset{
	ID:          PresetMoonshotKimiK27CodeHighspeed,
	Name:        ModelNameMoonshotKimiK27CodeHighspeed,
	DisplayName: DisplayNameMoonshotKimiK27CodeHighspeed,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMoonshotKimiK27CodeHighspeed,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 32768,
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1024),
				ZeroAllowed:     new(false),
				MinusOneAllowed: new(false),
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(true),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelMoonshotKimiK26 = ModelPreset{
	ID:          PresetMoonshotKimiK26,
	Name:        ModelNameMoonshotKimiK26,
	DisplayName: DisplayNameMoonshotKimiK26,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMoonshotKimiK26,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 32768,
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1024),
				ZeroAllowed:     new(false),
				MinusOneAllowed: new(false),
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(true),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var providerMoonshot = ProviderPreset{
	Name:                     ProviderMoonshot,
	DisplayName:              DisplayNameProviderMoonshot,
	SDKType:                  spec.ProviderSDKTypeAnthropic,
	Origin:                   "https://api.moonshot.ai",
	ChatCompletionPathPrefix: "/anthropic/v1/messages",
	APIKeyHeaderKey:          spec.DefaultAnthropicAuthorizationHeaderKey,
	DefaultHeaders: map[string]string{
		spec.DefaultContentTypeHeaderKey:      spec.DefaultContentTypeHeader,
		spec.DefaultAcceptHeaderKey:           spec.DefaultContentTypeHeader,
		spec.DefaultAnthropicVersionHeaderKey: spec.DefaultAnthropicVersionHeader,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:          new(false),
			SupportedReasoningTypes:          []spec.ReasoningType{},
			SupportedReasoningLevels:         []spec.ReasoningLevel{},
			SupportsSummaryStyle:             new(false),
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
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindContentItemList,
			},
		},
		CacheCapabilities: &capabilityoverride.CacheCapabilitiesOverride{
			SupportsAutomaticCaching: new(true),
			TopLevel: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
			InputOutputContent: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
			ReasoningContent: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
			ToolChoice: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
			ToolCall: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
			ToolOutput: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetMoonshotKimiK3:               modelMoonshotKimiK3,
		PresetMoonshotKimiK27Code:          modelMoonshotKimiK27Code,
		PresetMoonshotKimiK27CodeHighspeed: modelMoonshotKimiK27CodeHighspeed,
		PresetMoonshotKimiK26:              modelMoonshotKimiK26,
	},
}
