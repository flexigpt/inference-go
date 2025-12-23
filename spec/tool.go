package spec

type ToolType string

const (
	ToolTypeFunction  ToolType = "function"
	ToolTypeCustom    ToolType = "custom"
	ToolTypeWebSearch ToolType = "webSearch"
)

const DefaultWebSearchToolName string = "webSearchToolChoice"

type WebSearchToolChoiceItemUserLocation struct {
	// The city of the user.
	City string `json:"city,omitzero"`
	// The two letter [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the user.
	Country string `json:"country,omitzero"`
	// The region of the user.
	Region string `json:"region,omitzero"`
	// The [IANA timezone](https://nodatime.org/TimeZones) of the user.
	Timezone string `json:"timezone,omitzero"`
}

type WebSearchToolChoiceItem struct {
	// Maximum number of times the tool can be used in the API request.
	MaxUses int64 `json:"max_uses,omitzero"`
	// One of `low`, `medium`, or `high`. `medium` is the default.
	SearchContextSize string `json:"searchContextSize,omitzero"`
	// If provided, only these domains will be included in results. Cannot be used alongside `blocked_domains`.
	AllowedDomains []string `json:"allowed_domains,omitzero"`
	// If provided, these domains will never appear in results. Cannot be used alongside `allowed_domains`.
	BlockedDomains []string `json:"blocked_domains,omitzero"`
	// Parameters for the user's location. Used to provide more relevant search results.
	UserLocation *WebSearchToolChoiceItemUserLocation `json:"user_location,omitempty"`
}

type ToolChoice struct {
	Type ToolType `json:"type"`

	ID           string        `json:"id"`
	CacheControl *CacheControl `json:"cacheControl,omitempty"`

	Name        string `json:"name"`
	Description string `json:"description,omitzero"`

	// Only one below can be non nil.
	// Function/Custom tools can have arguments of any shape. Respective properties are passed to relevant sdk params.
	// Expected to be a map conversion of JSON Schema object in most cases.
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
	// Can be opaque map in some cases.
	Input map[string]any `json:"input,omitempty"`
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

	// Only one can be non nil.
	SearchItem   *WebSearchToolCallSearch   `json:"searchItem,omitempty"`
	OpenPageItem *WebSearchToolCallOpenPage `json:"openPageItem,omitempty"`
	FindItem     *WebSearchToolCallFind     `json:"findItem,omitempty"`
}

type ToolCall struct {
	Type ToolType `json:"type"`

	// The ID of the toolchoice will be mapped using the name output from the call and populated here.
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

	// Only one can be non nil.
	SearchItem *WebSearchToolOutputSearch `json:"searchItem,omitempty"`
	ErrorItem  *WebSearchToolOutputError  `json:"errorItem,omitempty"`
}

type ToolOutputItemUnion struct {
	Kind ContentItemKind `json:"kind"`

	// Only one can be non nil.
	TextItem  *ContentItemText  `json:"textItem,omitempty"`
	ImageItem *ContentItemImage `json:"imageItem,omitempty"`
	FileItem  *ContentItemFile  `json:"fileItem,omitempty"`
}

type ToolOutput struct {
	Type ToolType `json:"type"`

	ChoiceID     string        `json:"choiceID"`
	ID           string        `json:"id"` // This ID and CallID should be linked from ToolCall.
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
