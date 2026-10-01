package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderModelRunner spec.ProviderName = "modelrunner"

	DisplayNameProviderModelRunner = "ModelRunner"
)

// Model names are the platform endpoint id, `owner/alias`, the same value the
// discovery list returns and the only value the chat route accepts.
const (
	ModelNameModelRunnerGemini35Flash     spec.ModelName = "google/gemini-3.5-flash"
	ModelNameModelRunnerGemini37Flash     spec.ModelName = "google/gemini-3.7-flash"
	ModelNameModelRunnerGemini35FlashLite spec.ModelName = "google/gemini-3.5-flash-lite"
	ModelNameModelRunnerDeepSeekV4Pro     spec.ModelName = "deepseek/v4"
	ModelNameModelRunnerZAIGLM52          spec.ModelName = "z-ai/glm-5.2"
	ModelNameModelRunnerQwen38Max         spec.ModelName = "alibaba/qwen3.8-max"
)

var modelModelRunnerGemini35Flash = ModelPreset{
	ID:          PresetGemini35Flash,
	Name:        ModelNameModelRunnerGemini35Flash,
	DisplayName: DisplayNameGemini35Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameModelRunnerGemini35Flash,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelModelRunnerGemini37Flash = ModelPreset{
	ID:          PresetGemini37Flash,
	Name:        ModelNameModelRunnerGemini37Flash,
	DisplayName: DisplayNameGemini37Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameModelRunnerGemini37Flash,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelModelRunnerGemini35FlashLite = ModelPreset{
	ID:          PresetGemini35FlashLite,
	Name:        ModelNameModelRunnerGemini35FlashLite,
	DisplayName: DisplayNameGemini35FlashLite,
	ModelParam: spec.ModelParam{
		Name:            ModelNameModelRunnerGemini35FlashLite,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelModelRunnerDeepSeekV4Pro = ModelPreset{
	ID:          PresetDeepSeekV4Pro,
	Name:        ModelNameModelRunnerDeepSeekV4Pro,
	DisplayName: DisplayNameDeepSeekV4Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameModelRunnerDeepSeekV4Pro,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 393216,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelModelRunnerZAIGLM52 = ModelPreset{
	ID:          PresetZAIGLM52,
	Name:        ModelNameModelRunnerZAIGLM52,
	DisplayName: DisplayNameZAIGLM52,
	ModelParam: spec.ModelParam{
		Name:            ModelNameModelRunnerZAIGLM52,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

// The one entry that takes images as well as text, so it carries the only
// per-model modality override on this provider.
var modelModelRunnerQwen38Max = ModelPreset{
	ID:          PresetQwen38Max,
	Name:        ModelNameModelRunnerQwen38Max,
	DisplayName: DisplayNameQwen38Max,
	ModelParam: spec.ModelParam{
		Name:            ModelNameModelRunnerQwen38Max,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
	},
}

// ModelRunner is a gateway: one key and one OpenAI chat-completions surface in
// front of models from several labs. The provider-level capabilities are
// therefore the intersection that holds for every listed model, with the one
// exception overridden on its own preset above.
var providerModelRunner = ProviderPreset{
	Name:                     ProviderModelRunner,
	DisplayName:              DisplayNameProviderModelRunner,
	SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
	Origin:                   "https://queue.modelrunner.run",
	ChatCompletionPathPrefix: spec.DefaultOpenAIChatCompletionsPrefix,
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		// Several of these models think, and some expose an effort control
		// upstream, but the gateway documents no reasoning parameter of its own
		// and the models do not agree on a spelling. Declared unsupported rather
		// than offering a control that would only work for part of the list.
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:          new(false),
			SupportedReasoningTypes:          []spec.ReasoningType{},
			SupportedReasoningLevels:         []spec.ReasoningLevel{},
			SupportsSummaryStyle:             new(false),
			SupportsReasoningContext:         new(false),
			SupportsReasoningMode:            new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
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
			SupportsVerbosity: new(false),
		},
		// Function tools with an automatic policy are what the gateway
		// documents and what its own conformance run covers. Forced and
		// parallel calls are left off rather than assumed from the wire format.
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
				spec.ToolOutputFormatKindString,
			},
		},
		// Prompt caching happens upstream and is billed at its own rate on every
		// model here; there is no cache-control parameter to set.
		CacheCapabilities: &capabilityoverride.CacheCapabilitiesOverride{
			SupportsAutomaticCaching: new(true),
			TopLevel: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetGemini35Flash:     modelModelRunnerGemini35Flash,
		PresetGemini37Flash:     modelModelRunnerGemini37Flash,
		PresetGemini35FlashLite: modelModelRunnerGemini35FlashLite,
		PresetDeepSeekV4Pro:     modelModelRunnerDeepSeekV4Pro,
		PresetZAIGLM52:          modelModelRunnerZAIGLM52,
		PresetQwen38Max:         modelModelRunnerQwen38Max,
	},
}
