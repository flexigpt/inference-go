package integration

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

const (
	googleInteractionsCustomToolID   = "interactions-custom-echo"
	googleInteractionsCustomToolName = "custom_echo"
)

var errGoogleInteractionsNilResponse = errors.New("nil response")

// Example_googleInteractions_streamingCustomToolRichResultRoundTrip validates
// stateless replay of:
// - signed reasoning and streamed thought summaries
// - a custom tool represented by an Interactions function declaration
// - a rich custom-tool result with text, image, and document subcontent
// - JSON Schema output on the completion turn.
func Example_googleInteractions_streamingCustomToolRichResultRoundTrip() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	providerSet, provider, model, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Hydrogen,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	tool := newEchoToolChoice(
		googleInteractionsCustomToolID,
		googleInteractionsCustomToolName,
	)
	tool.Type = spec.ToolTypeCustom

	options, err := presetFetchOptions(ctx, providerSet, provider, model)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}

	var streamEvents atomic.Int32
	options.StreamHandler = func(event spec.StreamEvent) error {
		switch event.Kind {
		case spec.StreamContentKindThinking:
			if event.Thinking != nil && event.Thinking.Text != "" {
				streamEvents.Add(1)
				fmt.Fprintf(os.Stderr, "[thinking] %s", event.Thinking.Text)
			}
		case spec.StreamContentKindText:
			if event.Text != nil && event.Text.Text != "" {
				streamEvents.Add(1)
				fmt.Fprintf(os.Stderr, "[text] %s", event.Text.Text)
			}
		}
		return nil
	}

	summaryStyle := spec.ReasoningSummaryStyleAuto
	firstParam := model.ModelParam
	firstParam.Stream = true
	firstParam.MaxOutputLength = min(firstParam.MaxOutputLength, 1024)
	firstParam.SystemPrompt = "Call the requested custom tool before answering."
	firstParam.Reasoning = &spec.ReasoningParam{
		Type:         spec.ReasoningTypeSingleWithLevels,
		Level:        spec.ReasoningLevelLow,
		SummaryStyle: &summaryStyle,
	}

	initialUser := newUserTextInput(
		`Call custom_echo with text "rich stateless custom-tool round trip".`,
	)
	firstResponse, err := providerSet.FetchCompletion(
		ctx,
		provider.Name,
		&spec.FetchCompletionRequest{
			ModelParam:  firstParam,
			Inputs:      []spec.InputUnion{initialUser},
			ToolChoices: []spec.ToolChoice{tool},
			ToolPolicy: &spec.ToolPolicy{
				Mode: spec.ToolPolicyModeTool,
				AllowedTools: []spec.AllowedTool{
					{ToolChoiceID: tool.ID},
				},
				DisableParallel: true,
			},
		},
		options,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "first custom-tool interaction error:", err)
		return
	}
	if !googleInteractionsResponseHasWarning(
		firstResponse,
		"toolPolicy_disableParallel_unsupported",
	) {
		fmt.Fprintln(os.Stderr, "missing expected DisableParallel warning")
		return
	}

	call, err := firstCustomToolCall(firstResponse)
	if err != nil {
		fmt.Fprintln(os.Stderr, "expected custom tool call:", err)
		return
	}

	toolResult, err := runEchoTool(call)
	if err != nil {
		fmt.Fprintln(os.Stderr, "custom tool execution error:", err)
		return
	}
	toolResult.Contents = append(
		toolResult.Contents,
		spec.ToolOutputItemUnion{
			Kind: spec.ContentItemKindImage,
			ImageItem: &spec.ContentItemImage{
				ImageMIME: "image/png",
				ImageData: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQIHWP4z8DwHwAFgAI/ScLq9wAAAABJRU5ErkJggg==",
			},
		},
		spec.ToolOutputItemUnion{
			Kind: spec.ContentItemKindFile,
			FileItem: &spec.ContentItemFile{
				FileMIME: "text/csv",
				FileData: base64.StdEncoding.EncodeToString(
					[]byte("item,count\ncustom-tool,1\n"),
				),
			},
		},
	)

	history := make(
		[]spec.InputUnion,
		0,
		1+len(firstResponse.Outputs)+2,
	)
	history = append(history, initialUser)
	history = append(history, outputUnionsToInputs(firstResponse.Outputs)...)
	//nolint:gocritic // Multi append.
	history = append(history,
		spec.InputUnion{
			Kind:             spec.InputKindCustomToolOutput,
			CustomToolOutput: toolResult,
		})
	history = append(history,
		newUserTextInput(
			"Use the custom tool result and return JSON with an answer and whether the tool was used.",
		),
	)

	secondParam := model.ModelParam
	secondParam.Stream = true
	secondParam.MaxOutputLength = min(secondParam.MaxOutputLength, 1024)
	secondParam.SystemPrompt = "Use the supplied custom-tool result. Do not call another tool."
	secondParam.Reasoning = &spec.ReasoningParam{
		Type:         spec.ReasoningTypeSingleWithLevels,
		Level:        spec.ReasoningLevelLow,
		SummaryStyle: &summaryStyle,
	}
	secondParam.OutputParam = &spec.OutputParam{
		Format: &spec.OutputFormat{
			Kind: spec.OutputFormatKindJSONSchema,
			JSONSchemaParam: &spec.JSONSchemaParam{
				Name: "custom_tool_result",
				Schema: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"answer":    map[string]any{"type": "string"},
						"used_tool": map[string]any{"type": "boolean"},
					},
					"required":             []string{"answer", "used_tool"},
					"additionalProperties": false,
				},
			},
		},
	}

	secondOptions, err := presetFetchOptions(ctx, providerSet, provider, model)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating second capability resolver:", err)
		return
	}
	secondOptions.StreamHandler = options.StreamHandler

	secondResponse, err := providerSet.FetchCompletion(
		ctx,
		provider.Name,
		&spec.FetchCompletionRequest{
			ModelParam:  secondParam,
			Inputs:      history,
			ToolChoices: []spec.ToolChoice{tool},
			ToolPolicy:  &spec.ToolPolicy{Mode: spec.ToolPolicyModeNone},
		},
		secondOptions,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "second custom-tool interaction error:", err)
		return
	}

	var result struct {
		Answer   string `json:"answer"`
		UsedTool bool   `json:"used_tool"`
	}
	if err := json.Unmarshal([]byte(responseText(secondResponse)), &result); err != nil {
		fmt.Fprintln(os.Stderr, "structured final output was not JSON:", err)
		return
	}
	if strings.TrimSpace(result.Answer) == "" || !result.UsedTool {
		fmt.Fprintln(os.Stderr, "unexpected structured result:", result)
		return
	}
	if streamEvents.Load() == 0 {
		fmt.Fprintln(os.Stderr, "expected streamed text or thought events")
		return
	}

	fmt.Fprintln(os.Stderr, "custom tool final answer:", result.Answer)
	fmt.Println("OK")
	// Output: OK
}

// Example_googleInteractions_webSearchGroundingRoundTrip validates that
// server-side Google Search outputs can be retained as stateless step history
// before a later no-tool completion.
func Example_googleInteractions_webSearchGroundingRoundTrip() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	providerSet, provider, model, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Hydrogen,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	options, err := presetFetchOptions(ctx, providerSet, provider, model)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}
	options.StreamHandler = func(event spec.StreamEvent) error {
		if event.Kind == spec.StreamContentKindText && event.Text != nil {
			fmt.Fprint(os.Stderr, event.Text.Text)
		}
		return nil
	}

	webSearch := spec.ToolChoice{
		Type: spec.ToolTypeWebSearch,
		ID:   "interactions-google-search",
		Name: spec.DefaultWebSearchToolName,
		WebSearchArguments: &spec.WebSearchToolChoiceItem{
			SearchContextSize: spec.WebSearchContextSizeMedium,
		},
	}

	firstParam := model.ModelParam
	firstParam.Stream = true
	firstParam.MaxOutputLength = min(firstParam.MaxOutputLength, 1024)
	firstParam.SystemPrompt = "Use Google Search for current facts before answering."
	firstParam.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}

	initialUser := newUserTextInput(
		"What is the latest stable Go release? Search before you answer.",
	)
	firstResponse, err := providerSet.FetchCompletion(
		ctx,
		provider.Name,
		&spec.FetchCompletionRequest{
			ModelParam:  firstParam,
			Inputs:      []spec.InputUnion{initialUser},
			ToolChoices: []spec.ToolChoice{webSearch},
			ToolPolicy:  &spec.ToolPolicy{Mode: spec.ToolPolicyModeAuto},
		},
		options,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Google Search interaction error:", err)
		return
	}
	if !googleInteractionsHasWebSearchCall(firstResponse) {
		fmt.Fprintln(os.Stderr, "expected a normalized Google Search call")
		return
	}
	if !googleInteractionsResponseHasWarning(
		firstResponse,
		"webSearch_options_not_forwarded",
	) {
		fmt.Fprintln(os.Stderr, "missing expected Google Search option warning")
		return
	}

	history := make([]spec.InputUnion, 0, 1+len(firstResponse.Outputs)+1)
	history = append(history, initialUser)
	history = append(history, outputUnionsToInputs(firstResponse.Outputs)...)
	history = append(
		history,
		newUserTextInput(
			"Using the preceding grounded conversation, state the release in one sentence without another search.",
		),
	)

	secondParam := model.ModelParam
	secondParam.Stream = false
	secondParam.MaxOutputLength = min(secondParam.MaxOutputLength, 512)
	secondParam.SystemPrompt = "Use the supplied history only. Do not call tools."

	secondResponse, err := providerSet.FetchCompletion(
		ctx,
		provider.Name,
		&spec.FetchCompletionRequest{
			ModelParam:  secondParam,
			Inputs:      history,
			ToolChoices: []spec.ToolChoice{webSearch},
			ToolPolicy:  &spec.ToolPolicy{Mode: spec.ToolPolicyModeNone},
		},
		options,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "grounding-history follow-up error:", err)
		return
	}
	if strings.TrimSpace(responseText(secondResponse)) == "" {
		fmt.Fprintln(os.Stderr, "expected follow-up answer")
		return
	}

	fmt.Fprintln(os.Stderr, "grounded follow-up:", responseText(secondResponse))
	fmt.Println("OK")
	// Output: OK
}

// Example_googleInteractions_multimodalConversationRoundTrip validates
// stateless replay of a user turn that contains text, image, and document
// content before a follow-up question.
func Example_googleInteractions_multimodalConversationRoundTrip() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	providerSet, provider, model, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Hydrogen,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	options, err := presetFetchOptions(ctx, providerSet, provider, model)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}

	firstInputs := []spec.InputUnion{{
		Kind: spec.InputKindInputMessage,
		InputMessage: &spec.InputOutputContent{
			Role: spec.RoleUser,
			Contents: []spec.InputOutputContentItemUnion{
				{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "Inspect the image and CSV attachment."},
				},
				{
					Kind: spec.ContentItemKindImage,
					ImageItem: &spec.ContentItemImage{
						ImageMIME: "image/png",
						ImageData: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQIHWP4z8DwHwAFgAI/ScLq9wAAAABJRU5ErkJggg==",
					},
				},
				{
					Kind: spec.ContentItemKindFile,
					FileItem: &spec.ContentItemFile{
						FileMIME: "text/csv",
						FileData: base64.StdEncoding.EncodeToString(
							[]byte("animal,count\ncat,2\ndog,3\n"),
						),
					},
				},
			},
		},
	}}

	param := model.ModelParam
	param.Stream = false
	param.MaxOutputLength = min(param.MaxOutputLength, 1024)
	param.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}

	firstResponse, err := providerSet.FetchCompletion(
		ctx,
		provider.Name,
		&spec.FetchCompletionRequest{
			ModelParam: param,
			Inputs:     firstInputs,
		},
		options,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "first multimodal interaction error:", err)
		return
	}

	secondInputs := replayBasicConversationInputs(
		firstInputs,
		firstResponse,
		newUserTextInput("What total animal count did the CSV contain?"),
	)
	secondResponse, err := providerSet.FetchCompletion(
		ctx,
		provider.Name,
		&spec.FetchCompletionRequest{
			ModelParam: param,
			Inputs:     secondInputs,
		},
		options,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "multimodal follow-up error:", err)
		return
	}
	if strings.TrimSpace(responseText(secondResponse)) == "" {
		fmt.Fprintln(os.Stderr, "expected multimodal follow-up answer")
		return
	}

	fmt.Fprintln(os.Stderr, "multimodal follow-up:", responseText(secondResponse))
	fmt.Println("OK")
	// Output: OK
}

func firstCustomToolCall(
	response *spec.FetchCompletionResponse,
) (*spec.ToolCall, error) {
	if response == nil {
		return nil, errGoogleInteractionsNilResponse
	}
	for _, output := range response.Outputs {
		if output.Kind == spec.OutputKindCustomToolCall &&
			output.CustomToolCall != nil {
			return output.CustomToolCall, nil
		}
	}
	return nil, fmt.Errorf(
		"no custom tool call found in %d outputs",
		len(response.Outputs),
	)
}

func googleInteractionsResponseHasWarning(
	response *spec.FetchCompletionResponse,
	code string,
) bool {
	if response == nil {
		return false
	}
	for _, warning := range response.Warnings {
		if warning.Code == code {
			return true
		}
	}
	return false
}

func googleInteractionsHasWebSearchCall(
	response *spec.FetchCompletionResponse,
) bool {
	if response == nil {
		return false
	}
	for _, output := range response.Outputs {
		if output.Kind == spec.OutputKindWebSearchToolCall &&
			output.WebSearchToolCall != nil {
			return true
		}
	}
	return false
}
