package spec

type InputKind string

const (
	InputKindInputMessage        InputKind = "inputMessage"
	InputKindOutputMessage       InputKind = "outputMessage"
	InputKindReasoningMessage    InputKind = "reasoningMessage"
	InputKindFunctionToolCall    InputKind = "functionToolCall"
	InputKindFunctionToolOutput  InputKind = "functionToolOutput"
	InputKindCustomToolCall      InputKind = "customToolCall"
	InputKindCustomToolOutput    InputKind = "customToolOutput"
	InputKindWebSearchToolCall   InputKind = "webSearchToolCall"
	InputKindWebSearchToolOutput InputKind = "webSearchToolOutput"
)

type InputUnion struct {
	Kind InputKind `json:"kind"`

	InputMessage        *InputOutputContent `json:"inputMessage,omitempty"`
	OutputMessage       *InputOutputContent `json:"outputMessage,omitempty"`
	ReasoningMessage    *ReasoningContent   `json:"reasoningMessage,omitempty"`
	FunctionToolCall    *ToolCall           `json:"functionToolCall,omitempty"`
	FunctionToolOutput  *ToolOutput         `json:"functionToolOutput,omitempty"`
	CustomToolCall      *ToolCall           `json:"customToolCall,omitempty"`
	CustomToolOutput    *ToolOutput         `json:"customToolOutput,omitempty"`
	WebSearchToolCall   *ToolCall           `json:"webSearchToolCall,omitempty"`
	WebSearchToolOutput *ToolOutput         `json:"webSearchToolOutput,omitempty"`
}

type OutputKind string

const (
	OutputKindOutputMessage       OutputKind = "outputMessage"
	OutputKindReasoningMessage    OutputKind = "reasoningMessage"
	OutputKindFunctionToolCall    OutputKind = "functionToolCall"
	OutputKindCustomToolCall      OutputKind = "customToolCall"
	OutputKindWebSearchToolCall   OutputKind = "webSearchToolCall"
	OutputKindWebSearchToolOutput OutputKind = "webSearchToolOutput"
)

type OutputUnion struct {
	Kind OutputKind `json:"kind"`

	OutputMessage       *InputOutputContent `json:"outputMessage,omitempty"`
	ReasoningMessage    *ReasoningContent   `json:"reasoningMessage,omitempty"`
	FunctionToolCall    *ToolCall           `json:"functionToolCall,omitempty"`
	CustomToolCall      *ToolCall           `json:"customToolCall,omitempty"`
	WebSearchToolCall   *ToolCall           `json:"webSearchToolCall,omitempty"`
	WebSearchToolOutput *ToolOutput         `json:"webSearchToolOutput,omitempty"`
}
