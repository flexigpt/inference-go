package spec

type CacheControlKind string

const (
	CacheControlKindEphemeral CacheControlKind = "ephemeral"
)

type CacheControlTTL string

const (
	CacheControlTTL5m       CacheControlTTL = "5m"
	CacheControlTTL1h       CacheControlTTL = "1h"
	CacheControlTTL24h      CacheControlTTL = "24h"
	CacheControlTTLInMemory CacheControlTTL = "in-memory"
)

type CacheControl struct {
	Kind CacheControlKind `json:"kind"`

	// Optional. If empty and CacheControl is explicitly set, provider default may apply.
	TTL CacheControlTTL `json:"ttl,omitempty"`

	// Optional request-level cache key. Relevant for OpenAI-style root caching.
	// Irrelevant on unsupported providers/scopes and should be ignored, not errored.
	Key string `json:"key,omitempty"`
}
