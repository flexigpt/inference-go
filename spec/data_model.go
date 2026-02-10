package spec

type (
	ModelName       string
	ReasoningLevel  string
	ReasoningType   string
	ProviderName    string
	ProviderSDKType string
)

const (
	ReasoningTypeHybridWithTokens ReasoningType = "hybridWithTokens"
	ReasoningTypeSingleWithLevels ReasoningType = "singleWithLevels"
)

const (
	ReasoningLevelNone    ReasoningLevel = "none"
	ReasoningLevelMinimal ReasoningLevel = "minimal"
	ReasoningLevelLow     ReasoningLevel = "low"
	ReasoningLevelMedium  ReasoningLevel = "medium"
	ReasoningLevelHigh    ReasoningLevel = "high"
	ReasoningLevelXHigh   ReasoningLevel = "xhigh"
)

type ReasoningSummaryStyle string

const (
	ReasoningSummaryStyleAuto     ReasoningSummaryStyle = "auto"
	ReasoningSummaryStyleConcise  ReasoningSummaryStyle = "concise"
	ReasoningSummaryStyleDetailed ReasoningSummaryStyle = "detailed"
)

type ReasoningParam struct {
	Type   ReasoningType  `json:"type"`
	Level  ReasoningLevel `json:"level"`
	Tokens int            `json:"tokens"`
	// SummaryStyle - what kind of summary should be emitted for the reasoning performed by the model.
	// SummaryStyle is supported by OpenAI responses only.
	SummaryStyle *ReasoningSummaryStyle `json:"summaryStyle,omitempty"`
}

// OutputVerbosity constrains the verbosity of the model's response.
// Lower values will result in more concise responses, while higher values will result in more verbose responses.
type OutputVerbosity string

const (
	OutputVerbosityLow    OutputVerbosity = "low"
	OutputVerbosityMedium OutputVerbosity = "medium"
	OutputVerbosityHigh   OutputVerbosity = "high"
)

type OutputFormatKind string

const (
	OutputFormatKindText       OutputFormatKind = "text"
	OutputFormatKindJSONSchema OutputFormatKind = "jsonSchema"
)

type JSONSchemaParam struct {
	// Name - must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64..
	Name string `json:"name"`

	// Description: Optional description of what the response format is for, used by the model to determine how to
	// respond in the format.
	Description string `json:"description,omitempty"`

	// JSON Schema payload when Type == jsonSchema.
	Schema map[string]any `json:"schema,omitempty"`

	// Strict requests stricter adherence where supported.
	Strict bool `json:"strict,omitempty"`
}

type OutputFormat struct {
	// Can be text or JSONSchema. We don't support JSONObject as it is recommended to move to JSONSchema type.
	Kind OutputFormatKind `json:"kind"`

	JSONSchemaParam *JSONSchemaParam `json:"jsonSchemaParam,omitempty"`
}

type OutputParam struct {
	Format *OutputFormat `json:"format,omitempty"`
	// Verbosity is supported by OpenAI, not supported by Anthropic.
	Verbosity *OutputVerbosity `json:"verbosity,omitempty"`
}

type ModelParam struct {
	Name            ModelName       `json:"name"`
	Stream          bool            `json:"stream"`
	MaxPromptLength int             `json:"maxPromptLength"`
	MaxOutputLength int             `json:"maxOutputLength"`
	Temperature     *float64        `json:"temperature,omitempty"`
	Reasoning       *ReasoningParam `json:"reasoning,omitempty"`
	SystemPrompt    string          `json:"systemPrompt"`
	Timeout         int             `json:"timeout"`

	// OutputParam controls the model's output format.
	//
	// Cross-provider notes:
	//   - OpenAI Chat Completions: maps to response_format + verbosity.
	//   - OpenAI Responses: maps to text.
	//   - Anthropic Messages: supports jsonSchema only via output_config.format. verbosity is not supported.
	OutputParam *OutputParam `json:"outputParam,omitempty"`

	// StopSequences requests provider-side stop sequences when supported.
	// Cross-provider notes:
	//   - OpenAI Chat Completions: maps to stop. Up to 4 sequences supported. Not supported by reasoning models
	//   - OpenAI Responses: Not supported.
	//   - Anthropic Messages: maps to stop_sequences.
	StopSequences []string `json:"stopSequences,omitempty"`

	AdditionalParametersRawJSON *string `json:"additionalParametersRawJSON"`
}

type Usage struct {
	InputTokensTotal    int64 `json:"inputTokensTotal"`
	InputTokensCached   int64 `json:"inputTokensCached"`
	InputTokensUncached int64 `json:"inputTokensUncached"`
	OutputTokens        int64 `json:"outputTokens"`
	ReasoningTokens     int64 `json:"reasoningTokens"`
}
