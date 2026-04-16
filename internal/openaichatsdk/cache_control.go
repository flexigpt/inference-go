package openaichatsdk

import (
	"strings"

	"github.com/flexigpt/inference-go/spec"
	"github.com/openai/openai-go/v3"
)

func applyOpenAIChatCacheControl(
	params *openai.ChatCompletionNewParams,
	cc *spec.CacheControl,
) {
	if params == nil || cc == nil {
		return
	}
	if cc.Kind != "" && cc.Kind != spec.CacheControlKindEphemeral {
		return
	}

	if key := strings.TrimSpace(cc.Key); key != "" {
		params.PromptCacheKey = openai.String(key)
	}

	if cc.TTL != "" {
		switch cc.TTL {
		case spec.CacheControlTTL24h:
			params.PromptCacheRetention = openai.ChatCompletionNewParamsPromptCacheRetention("24h")
		default:
			params.PromptCacheRetention = openai.ChatCompletionNewParamsPromptCacheRetention("in-memory")
		}
	}
}
