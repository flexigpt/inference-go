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
	Text      string     `json:"text"`
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
	ImageData string      `json:"imageData,omitzero"`
}

type ContentItemFile struct {
	ID       string `json:"id,omitzero"`
	FileName string `json:"fileName,omitzero"`
	FileMIME string `json:"fileMIME,omitzero"`
	FileURL  string `json:"fileURL,omitzero"`

	FileData          string          `json:"fileData,omitzero"`
	AdditionalContext string          `json:"additionalContext,omitzero"`
	CitationConfig    *CitationConfig `json:"citationConfig"`
}

type InputOutputContentItemUnion struct {
	Kind ContentItemKind `json:"kind"`

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
	ID           string        `json:"id"`
	Role         RoleEnum      `json:"role"`
	Status       Status        `json:"status,omitzero"`
	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	Signature        string   `json:"signature,omitzero"`
	Summary          []string `json:"summary,omitempty"`
	Thinking         []string `json:"thinking,omitempty"`
	RedactedThinking []string `json:"redactedThinking,omitempty"`
	EncryptedContent []string `json:"encryptedContent,omitempty"`
}
