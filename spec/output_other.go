package spec

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Usage struct {
	InputTokensTotal    int64 `json:"inputTokensTotal"`
	InputTokensCached   int64 `json:"inputTokensCached"`
	InputTokensUncached int64 `json:"inputTokensUncached"`
	OutputTokens        int64 `json:"outputTokens"`
	ReasoningTokens     int64 `json:"reasoningTokens"`
}
