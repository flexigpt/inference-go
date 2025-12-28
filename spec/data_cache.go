package spec

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
