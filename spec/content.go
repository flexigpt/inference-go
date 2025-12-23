package spec

type RoleEnum string

const (
	RoleSystem    RoleEnum = "system"
	RoleDeveloper RoleEnum = "developer"
	RoleUser      RoleEnum = "user"
	RoleAssistant RoleEnum = "assistant"
	RoleFunction  RoleEnum = "function"
	RoleTool      RoleEnum = "tool"
)

type Status string

const (
	StatusInProgress Status = "inProgress"
	StatusCompleted  Status = "completed"
	StatusIncomplete Status = "incomplete"
	StatusFailed     Status = "failed"
	StatusCancelled  Status = "cancelled"
	StatusQueued     Status = "queued"
	StatusSearching  Status = "searching"
)

type ContentItemKind string

const (
	ContentItemKindText    ContentItemKind = "text"
	ContentItemKindImage   ContentItemKind = "image"
	ContentItemKindFile    ContentItemKind = "file"
	ContentItemKindRefusal ContentItemKind = "refusal"
)

type ContentItemText struct {
	Text string `json:"text"`

	// Any additional references to this text item.
	Citations []Citation `json:"citations,omitempty"`
}

type ContentItemRefusal struct {
	Refusal string `json:"refusal"`
}

type ImageDetail string

const (
	ImageDetailHigh ImageDetail = "high"
	ImageDetailLow  ImageDetail = "low"
	ImageDetailAuto ImageDetail = "auto"
)

type ContentItemImage struct {
	ID        string      `json:"id,omitzero"`
	Detail    ImageDetail `json:"detail,omitzero"`
	ImageName string      `json:"imageName,omitzero"`
	ImageMIME string      `json:"imageMIME,omitzero"`
	ImageURL  string      `json:"imageURL,omitzero"`
	// Base64 encoded data.
	ImageData string `json:"imageData,omitzero"`
}

type ContentItemFile struct {
	ID       string `json:"id,omitzero"`
	FileName string `json:"fileName,omitzero"`
	FileMIME string `json:"fileMIME,omitzero"`
	FileURL  string `json:"fileURL,omitzero"`
	// Base64 encoded data.
	FileData          string          `json:"fileData,omitzero"`
	AdditionalContext string          `json:"additionalContext,omitzero"`
	CitationConfig    *CitationConfig `json:"citationConfig"`
}

type InputOutputContentItemUnion struct {
	Kind ContentItemKind `json:"kind"`

	// Only one can be non nil.
	TextItem    *ContentItemText    `json:"textItem,omitempty"`
	RefusalItem *ContentItemRefusal `json:"refusalItem,omitempty"`
	ImageItem   *ContentItemImage   `json:"imageItem,omitempty"`
	FileItem    *ContentItemFile    `json:"fileItem,omitempty"`
}

type InputOutputContent struct {
	ID           string        `json:"id"`
	Role         RoleEnum      `json:"role"`
	Status       Status        `json:"status,omitzero"`
	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	Contents []InputOutputContentItemUnion `json:"contents,omitempty"`
}

type ReasoningContent struct {
	// An unique identifier for this content.
	ID           string        `json:"id"`
	Role         RoleEnum      `json:"role"`
	Status       Status        `json:"status,omitzero"`
	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	Signature        string   `json:"signature,omitzero"`
	Summary          []string `json:"summary,omitempty"`
	Thinking         []string `json:"thinking,omitempty"`
	RedactedThinking []string `json:"redactedThinking,omitempty"`

	// The encrypted content of the reasoning item.
	// In case of openai responses api, it is generated if `reasoning.encrypted_content` is specified in the `include`.
	EncryptedContent []string `json:"encryptedContent,omitempty"`
}

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
