package spec

// OutputVerbosity constrains the verbosity of the model's response.
// Lower values will result in more concise responses, while higher values will result in more verbose responses.
type OutputVerbosity string

const (
	OutputVerbosityLow    OutputVerbosity = "low"
	OutputVerbosityMedium OutputVerbosity = "medium"
	OutputVerbosityHigh   OutputVerbosity = "high"
	OutputVerbosityXHigh  OutputVerbosity = "xhigh"
	OutputVerbosityMax    OutputVerbosity = "max"
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
	// Maps to "Verbosity" in OpenAI, "Effort" in Anthropic.
	Verbosity *OutputVerbosity `json:"verbosity,omitempty"`
}
