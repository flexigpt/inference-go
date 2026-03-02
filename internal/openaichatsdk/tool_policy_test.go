package openaichatsdk

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"
	"github.com/openai/openai-go/v3"
)

func TestApplyOpenAIChatToolPolicy_ModeAny_AllToolsWhenAllowedToolsEmpty(t *testing.T) {
	params := openai.ChatCompletionNewParams{
		Tools: []openai.ChatCompletionToolUnionParam{
			// Contents don't matter for this test; only non-empty Tools is required.
			{OfFunction: &openai.ChatCompletionFunctionToolParam{}},
		},
	}
	nameMap := map[string]spec.ToolChoice{
		"t1": {Type: spec.ToolTypeFunction, ID: "t1"},
	}
	err := applyOpenAIChatToolPolicy(&params, &spec.ToolPolicy{Mode: spec.ToolPolicyModeAny}, nameMap)
	if err != nil {
		t.Fatalf("expected no error; got %v", err)
	}
}

func TestApplyOpenAIChatToolPolicy_ModeTool_WebSearchOnlyDoesNotPanic(t *testing.T) {
	params := openai.ChatCompletionNewParams{
		Tools: []openai.ChatCompletionToolUnionParam{
			{OfFunction: &openai.ChatCompletionFunctionToolParam{}},
		},
	}
	nameMap := map[string]spec.ToolChoice{
		"web_search": {Type: spec.ToolTypeWebSearch, ID: "ws"},
	}
	policy := &spec.ToolPolicy{
		Mode: spec.ToolPolicyModeTool,
		AllowedTools: []spec.AllowedTool{{
			ToolChoiceName: "web_search",
		}},
	}
	if err := applyOpenAIChatToolPolicy(&params, policy, nameMap); err == nil {
		t.Fatalf("expected error when forcing web_search in chat.completions; got nil")
	}
}

func TestApplyOpenAIChatToolPolicy_ModeNone_ClearsWebSearchOptions(t *testing.T) {
	params := openai.ChatCompletionNewParams{}
	params.WebSearchOptions.SearchContextSize = string(spec.WebSearchContextSizeMedium)

	if err := applyOpenAIChatToolPolicy(
		&params,
		&spec.ToolPolicy{Mode: spec.ToolPolicyModeNone},
		map[string]spec.ToolChoice{"t1": {Type: spec.ToolTypeFunction, ID: "t1"}},
	); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if params.WebSearchOptions.SearchContextSize != "" {
		t.Fatalf("expected web search options cleared; got %#v", params.WebSearchOptions)
	}
}
