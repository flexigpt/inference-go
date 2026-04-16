package openaichatsdk

import (
	"github.com/flexigpt/inference-go/spec"
	"github.com/openai/openai-go/v3"
)

func resolveOpenAIChatParamDialect(
	caps *spec.ModelCapabilities,
) spec.ParamDialect {
	out := spec.ParamDialect{
		MaxOutputTokensParamName: spec.MaxOutputTokensParamNameMaxCompletionTokens,
		ToolChoiceParamStyle:     spec.ToolChoiceParamStyleAllowedTools,
	}

	if caps == nil || caps.ParamDialect == nil {
		return out
	}

	if caps.ParamDialect.MaxOutputTokensParamName == spec.MaxOutputTokensParamNameMaxTokens {
		out.MaxOutputTokensParamName = spec.MaxOutputTokensParamNameMaxTokens
	}
	if caps.ParamDialect.ToolChoiceParamStyle == spec.ToolChoiceParamStyleRequiredNamed {
		out.ToolChoiceParamStyle = spec.ToolChoiceParamStyleRequiredNamed
	}

	return out
}

func applyOpenAIChatMaxOutputLength(
	params *openai.ChatCompletionNewParams,
	maxOutputLength int,
	dialect spec.ParamDialect,
) {
	if params == nil || maxOutputLength <= 0 {
		return
	}

	switch dialect.MaxOutputTokensParamName {
	case spec.MaxOutputTokensParamNameMaxTokens:
		params.MaxTokens = openai.Int(int64(maxOutputLength))
	default:
		params.MaxCompletionTokens = openai.Int(int64(maxOutputLength))
	}
}
