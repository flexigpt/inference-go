package spec

type CitationKind string

const (
	CitationKindURL CitationKind = "urlCitation"
)

type URLCitation struct {
	URL       string `json:"url"`
	Title     string `json:"title,omitzero"`
	CitedText string `json:"citedText,omitzero"`

	StartIndex     int64  `json:"startIndex,omitzero"`
	EndIndex       int64  `json:"endIndex,omitzero"`
	EncryptedIndex string `json:"encryptedIndex,omitzero"`
}

type Citation struct {
	Kind CitationKind `json:"kind"`

	URLCitation *URLCitation `json:"urlCitation,omitempty"`
}

type CitationConfig struct {
	Enabled bool `json:"enabled"`
}
