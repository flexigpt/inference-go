package spec

type CitationKind string

const (
	CitationKindURL CitationKind = "urlCitation"
)

type URLCitation struct {
	URL       string `json:"url"`
	Title     string `json:"title,omitzero"`
	CitedText string `json:"citedText,omitzero"`

	// Start and end indexes may have different from inclusiveness and byte vs char boundaries in different sdks.
	StartIndex     int64  `json:"startIndex,omitzero"`
	EndIndex       int64  `json:"endIndex,omitzero"`
	EncryptedIndex string `json:"encryptedIndex,omitzero"`
}

type Citation struct {
	Kind CitationKind `json:"kind"`

	// Exactly one of the below should be non-nil, depending on Kind.
	URLCitation *URLCitation `json:"urlCitation,omitempty"`
}

type CitationConfig struct {
	Enabled bool `json:"enabled"`
}

type CacheControlKind string

const (
	CacheControlKindEphemeral CacheControlKind = "ephemeral"
)

type CacheControlEphemeral struct {
	TTL string `json:"ttl,omitzero"`
}

type CacheControl struct {
	Kind CacheControlKind `json:"kind"`

	// Exactly one of the below should be non-nil, depending on Kind.
	CacheControlEphemeral *CacheControlEphemeral `json:"cacheControlEphemeral,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
