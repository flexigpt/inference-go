package anthropicsdk

import (
	"slices"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/flexigpt/inference-go/spec"
)

func applyAnthropicTopLevelCacheControl(
	params *anthropic.MessageNewParams,
	cc *spec.CacheControl,
) {
	if params == nil {
		return
	}
	cache := anthropicCacheControlParam(cc)
	if cache == nil {
		return
	}
	params.CacheControl = *cache
}

func applyAnthropicContentBlockCacheControl(
	blocks []anthropic.ContentBlockParamUnion,
	cc *spec.CacheControl,
) []anthropic.ContentBlockParamUnion {
	cache := anthropicCacheControlParam(cc)
	if cache == nil || len(blocks) == 0 {
		return blocks
	}

	for _, v := range slices.Backward(blocks) {
		switch {
		case v.OfText != nil:
			v.OfText.CacheControl = *cache
			return blocks
		case v.OfImage != nil:
			v.OfImage.CacheControl = *cache
			return blocks
		case v.OfDocument != nil:
			v.OfDocument.CacheControl = *cache
			return blocks
		}
	}

	return blocks
}

func applyAnthropicToolUseCacheControl(
	block *anthropic.ContentBlockParamUnion,
	cc *spec.CacheControl,
) {
	if block == nil {
		return
	}
	cache := anthropicCacheControlParam(cc)
	if cache == nil {
		return
	}

	switch {
	case block.OfToolUse != nil:
		block.OfToolUse.CacheControl = *cache
	case block.OfServerToolUse != nil:
		block.OfServerToolUse.CacheControl = *cache
	}
}

func applyAnthropicToolResultCacheControl(
	block *anthropic.ContentBlockParamUnion,
	cc *spec.CacheControl,
) {
	if block == nil {
		return
	}
	cache := anthropicCacheControlParam(cc)
	if cache == nil {
		return
	}

	switch {
	case block.OfToolResult != nil:
		block.OfToolResult.CacheControl = *cache
	case block.OfWebSearchToolResult != nil:
		block.OfWebSearchToolResult.CacheControl = *cache
	}
}

func applyAnthropicToolCacheControl(
	tool *anthropic.ToolUnionParam,
	cc *spec.CacheControl,
) {
	if tool == nil {
		return
	}
	cache := anthropicCacheControlParam(cc)
	if cache == nil {
		return
	}

	switch {
	case tool.OfTool != nil:
		tool.OfTool.CacheControl = *cache
	case tool.OfWebSearchTool20250305 != nil:
		tool.OfWebSearchTool20250305.CacheControl = *cache
	}
}

func anthropicCacheControlParam(cc *spec.CacheControl) *anthropic.CacheControlEphemeralParam {
	if cc == nil {
		return nil
	}
	if cc.Kind != "" && cc.Kind != spec.CacheControlKindEphemeral {
		return nil
	}

	out := anthropic.CacheControlEphemeralParam{
		Type: "ephemeral",
	}

	if cc.TTL != "" {
		switch cc.TTL {
		case spec.CacheControlTTL1h:
			out.TTL = anthropic.CacheControlEphemeralTTLTTL1h
		default:
			out.TTL = anthropic.CacheControlEphemeralTTLTTL5m
		}
	}

	return &out
}
