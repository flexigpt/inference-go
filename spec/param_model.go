package spec

type ModelName string

type ModelParam struct {
	Name            ModelName       `json:"name"`
	Stream          bool            `json:"stream"`
	MaxPromptLength int             `json:"maxPromptLength"`
	MaxOutputLength int             `json:"maxOutputLength"`
	Temperature     *float64        `json:"temperature,omitempty"`
	Reasoning       *ReasoningParam `json:"reasoning,omitempty"`
	SystemPrompt    string          `json:"systemPrompt"`
	Timeout         int             `json:"timeout"`

	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	// OutputParam controls the model's output format.
	//
	// Cross-provider notes:
	//   - OpenAI Chat Completions: maps to response_format + verbosity.
	//   - OpenAI Responses: maps to text.
	//   - Anthropic Messages: supports jsonSchema via output_config.format; verbosity maps to output_config.effort.
	OutputParam *OutputParam `json:"outputParam,omitempty"`

	// StopSequences requests provider-side stop sequences when supported.
	// Cross-provider notes:
	//   - OpenAI Chat Completions: maps to stop. Up to 4 sequences supported. Not supported by reasoning models
	//   - OpenAI Responses: Not supported.
	//   - Anthropic Messages: maps to stop_sequences.
	StopSequences []string `json:"stopSequences,omitempty"`

	AdditionalParametersRawJSON *string `json:"additionalParametersRawJSON"`
}
