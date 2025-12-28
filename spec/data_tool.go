package spec

type ToolType string

const (
	ToolTypeFunction  ToolType = "function"
	ToolTypeCustom    ToolType = "custom"
	ToolTypeWebSearch ToolType = "webSearch"
)

const DefaultWebSearchToolName string = "webSearchToolChoice"

type WebSearchToolChoiceItemUserLocation struct {
	City     string `json:"city,omitzero"`
	Country  string `json:"country,omitzero"`
	Region   string `json:"region,omitzero"`
	Timezone string `json:"timezone,omitzero"`
}

type WebSearchToolChoiceItem struct {
	MaxUses           int64                                `json:"max_uses,omitzero"`
	SearchContextSize string                               `json:"searchContextSize,omitzero"`
	AllowedDomains    []string                             `json:"allowed_domains,omitzero"`
	BlockedDomains    []string                             `json:"blocked_domains,omitzero"`
	UserLocation      *WebSearchToolChoiceItemUserLocation `json:"user_location,omitempty"`
}

type ToolChoice struct {
	Type ToolType `json:"type"`

	ID           string        `json:"id"`
	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	Name        string `json:"name"`
	Description string `json:"description,omitzero"`

	Arguments          map[string]any           `json:"arguments,omitempty"`
	WebSearchArguments *WebSearchToolChoiceItem `json:"webSearchArguments,omitempty"`
}

type WebSearchToolCallKind string

const (
	WebSearchToolCallKindSearch   WebSearchToolCallKind = "search"
	WebSearchToolCallKindOpenPage WebSearchToolCallKind = "openPage"
	WebSearchToolCallKindFind     WebSearchToolCallKind = "find"
)

type WebSearchToolCallSearchSource struct {
	URL string `json:"url"`
}

type WebSearchToolCallSearch struct {
	Query   string                          `json:"query"`
	Sources []WebSearchToolCallSearchSource `json:"sources,omitempty"`
	Input   map[string]any                  `json:"input,omitempty"`
}

type WebSearchToolCallOpenPage struct {
	URL string `json:"url"`
}

type WebSearchToolCallFind struct {
	URL     string `json:"url"`
	Pattern string `json:"pattern"`
}

type WebSearchToolCallItemUnion struct {
	Kind WebSearchToolCallKind `json:"kind"`

	SearchItem   *WebSearchToolCallSearch   `json:"searchItem,omitempty"`
	OpenPageItem *WebSearchToolCallOpenPage `json:"openPageItem,omitempty"`
	FindItem     *WebSearchToolCallFind     `json:"findItem,omitempty"`
}

type ToolCall struct {
	Type ToolType `json:"type"`

	ChoiceID     string        `json:"choiceID"`
	ID           string        `json:"id"`
	Role         RoleEnum      `json:"role"`
	Status       Status        `json:"status,omitzero"`
	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	CallID                 string                       `json:"callID"`
	Name                   string                       `json:"name"`
	Arguments              string                       `json:"arguments,omitempty"`
	WebSearchToolCallItems []WebSearchToolCallItemUnion `json:"webSearchToolCallItems,omitempty"`
}

type WebSearchToolOutputKind string

const (
	WebSearchToolOutputKindSearch WebSearchToolOutputKind = "search"
	WebSearchToolOutputKindError  WebSearchToolOutputKind = "error"
)

type WebSearchToolOutputSearch struct {
	URL              string `json:"url"`
	Title            string `json:"title,omitzero"`
	EncryptedContent string `json:"encryptedContent,omitzero"`
	RenderedContent  string `json:"renderedContent,omitzero"`
	PageAge          string `json:"page_age,omitzero"`
}

type WebSearchToolOutputError struct {
	Code string `json:"code"`
}

type WebSearchToolOutputItemUnion struct {
	Kind WebSearchToolOutputKind `json:"kind"`

	SearchItem *WebSearchToolOutputSearch `json:"searchItem,omitempty"`
	ErrorItem  *WebSearchToolOutputError  `json:"errorItem,omitempty"`
}

type ToolOutputItemUnion struct {
	Kind ContentItemKind `json:"kind"`

	TextItem  *ContentItemText  `json:"textItem,omitempty"`
	ImageItem *ContentItemImage `json:"imageItem,omitempty"`
	FileItem  *ContentItemFile  `json:"fileItem,omitempty"`
}

type ToolOutput struct {
	Type ToolType `json:"type"`

	ChoiceID     string        `json:"choiceID"`
	ID           string        `json:"id"`
	Role         RoleEnum      `json:"role"`
	Status       Status        `json:"status,omitzero"`
	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	CallID    string `json:"callID"`
	Name      string `json:"name"`
	IsError   bool   `json:"isError"`
	Signature string `json:"signature,omitzero"`

	Contents                 []ToolOutputItemUnion          `json:"contents,omitempty"`
	WebSearchToolOutputItems []WebSearchToolOutputItemUnion `json:"webSearchToolOutputItems,omitempty"`
}
