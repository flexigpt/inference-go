package openairesponsessdk

import (
	"strings"

	"github.com/flexigpt/inference-go/spec"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

func applyOpenAIResponsesCacheControl(
	params *responses.ResponseNewParams,
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
			params.PromptCacheRetention = responses.ResponseNewParamsPromptCacheRetention("24h")
		default:
			params.PromptCacheRetention = responses.ResponseNewParamsPromptCacheRetention("in-memory")
		}
	}
}
