package integration

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

const (
	googleInteractionsEchoToolID   = "interactions-echo"
	googleInteractionsEchoToolName = "echo_text"
)

func Example_googleInteractions_basicConversation() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, pp, mp, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Hydrogen,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}

	param := mp.ModelParam
	param.Stream = false
	param.MaxPromptLength = min(param.MaxPromptLength, 8192)
	param.MaxOutputLength = min(param.MaxOutputLength, 1024)
	param.SystemPrompt = "You are concise and helpful."
	param.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}

	firstInputs := []spec.InputUnion{
		newUserTextInput("Say hello from Gemini Interactions in one sentence."),
	}
	first, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: param,
		Inputs:     firstInputs,
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "first interaction error:", err)
		return
	}
	fmt.Fprintln(os.Stderr, "first answer:", responseText(first))

	followUpInputs := replayBasicConversationInputs(
		firstInputs,
		first,
		newUserTextInput("Repeat the greeting exactly."),
	)
	second, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: param,
		Inputs:     followUpInputs,
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "follow-up interaction error:", err)
		return
	}
	fmt.Fprintln(os.Stderr, "second answer:", responseText(second))

	fmt.Println("OK")
	// Output: OK
}

func Example_googleInteractions_structuredOutput() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, pp, mp, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Carbon,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}

	param := mp.ModelParam
	param.Stream = false
	param.MaxOutputLength = min(param.MaxOutputLength, 1024)
	param.OutputParam = &spec.OutputParam{
		Format: &spec.OutputFormat{
			Kind: spec.OutputFormatKindJSONSchema,
			JSONSchemaParam: &spec.JSONSchemaParam{
				Name:        "release_summary",
				Description: "A concise release summary.",
				Strict:      true,
				Schema: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"summary": map[string]any{
							"type": "string",
						},
						"important": map[string]any{
							"type": "boolean",
						},
					},
					"required": []string{"summary", "important"},
				},
			},
		},
	}

	initialUser := newUserTextInput(
		"Summarize why semantic versioning matters to library users.",
	)
	resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: param,
		Inputs:     []spec.InputUnion{initialUser},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "structured-output interaction error:", err)
		return
	}

	text := responseText(resp)
	var value struct {
		Summary   string `json:"summary"`
		Important bool   `json:"important"`
	}
	if err := json.Unmarshal([]byte(text), &value); err != nil {
		fmt.Fprintln(os.Stderr, "structured output was not JSON:", err)
		return
	}
	if strings.TrimSpace(value.Summary) == "" {
		fmt.Fprintln(os.Stderr, "structured output summary was empty")
		return
	}

	followUpInputs := replayBasicConversationInputs(
		[]spec.InputUnion{initialUser},
		resp,
		newUserTextInput(
			"Using the previous JSON response, return a new JSON summary in simpler language.",
		),
	)
	followUp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: param,
		Inputs:     followUpInputs,
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "structured-output follow-up error:", err)
		return
	}

	followUpText := responseText(followUp)
	var followUpValue struct {
		Summary   string `json:"summary"`
		Important bool   `json:"important"`
	}
	if err := json.Unmarshal([]byte(followUpText), &followUpValue); err != nil {
		fmt.Fprintln(os.Stderr, "structured follow-up was not JSON:", err)
		return
	}
	if strings.TrimSpace(followUpValue.Summary) == "" {
		fmt.Fprintln(os.Stderr, "structured follow-up summary was empty")
		return
	}

	fmt.Fprintln(os.Stderr, "summary:", value.Summary)
	fmt.Fprintln(os.Stderr, "important:", value.Important)
	fmt.Fprintln(os.Stderr, "follow-up summary:", followUpValue.Summary)
	fmt.Fprintln(os.Stderr, "follow-up important:", followUpValue.Important)

	fmt.Println("OK")
	// Output: OK
}

func Example_googleInteractions_multimodalImageAndDocumentInput() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, pp, mp, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Hydrogen,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}

	// A valid 1x1 PNG and a small CSV document avoid external fixture paths.
	imageData := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQIHWP4z8DwHwAFgAI/ScLq9wAAAABJRU5ErkJggg=="
	csvData := base64.StdEncoding.EncodeToString(
		[]byte("animal,count\ncat,2\ndog,3\n"),
	)

	param := mp.ModelParam
	param.Stream = false
	param.MaxOutputLength = min(param.MaxOutputLength, 1024)
	param.SystemPrompt = "Describe the image briefly, then summarize the CSV values."

	resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: param,
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{
					{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: "Inspect both attachments."},
					},
					{
						Kind: spec.ContentItemKindImage,
						ImageItem: &spec.ContentItemImage{
							ImageData: imageData,
							ImageMIME: "image/png",
						},
					},
					{
						Kind: spec.ContentItemKindFile,
						FileItem: &spec.ContentItemFile{
							FileName: "animals.csv",
							FileData: csvData,
							FileMIME: "text/csv",
						},
					},
				},
			},
		}},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "multimodal interaction error:", err)
		return
	}

	fmt.Fprintln(os.Stderr, "multimodal answer:", responseText(resp))
	fmt.Println("OK")
	// Output: OK
}

func Example_googleInteractions_functionToolRoundTrip() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	ps, pp, mp, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Hydrogen,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	tool := newEchoToolChoice(
		googleInteractionsEchoToolID,
		googleInteractionsEchoToolName,
	)

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}

	param := mp.ModelParam
	param.Stream = false
	param.MaxOutputLength = min(param.MaxOutputLength, 2048)
	param.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}
	param.SystemPrompt = "Call the requested function before writing an answer."

	initialUser := newUserTextInput(
		`Call echo_text with text "stateless interactions tool round trip".`,
	)
	first, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam:  param,
		Inputs:      []spec.InputUnion{initialUser},
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeTool,
			AllowedTools: []spec.AllowedTool{
				{ToolChoiceID: tool.ID},
			},
		},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "first tool interaction error:", err)
		return
	}

	call, err := firstFunctionToolCall(first)
	if err != nil {
		fmt.Fprintln(os.Stderr, "expected function call:", err)
		return
	}
	result, err := runEchoTool(call)
	if err != nil {
		fmt.Fprintln(os.Stderr, "tool execution error:", err)
		return
	}

	outputs := outputUnionsToInputs(first.Outputs)

	history := make([]spec.InputUnion, 0, 1+len(outputs)+2)

	history = append(history, initialUser)
	history = append(history, outputs...)
	//nolint:gocritic // No append combine.
	history = append(history, spec.InputUnion{
		Kind:               spec.InputKindFunctionToolOutput,
		FunctionToolOutput: result,
	})
	history = append(history, newUserTextInput(
		"Use the tool result and answer in one sentence.",
	))

	param.SystemPrompt = "Use the supplied function result. Do not call another tool."
	second, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam:  param,
		Inputs:      history,
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy:  &spec.ToolPolicy{Mode: spec.ToolPolicyModeNone},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "second tool interaction error:", err)
		return
	}

	fmt.Fprintln(os.Stderr, "tool result:", firstToolOutputText(result))
	fmt.Fprintln(os.Stderr, "final answer:", responseText(second))
	fmt.Println("OK")
	// Output: OK
}

func Example_googleInteractions_webSearchGroundingAndThinkingStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	ps, pp, mp, ok := googleInteractionsExampleSetup(
		ctx,
		modelpreset.PresetGemini4Hydrogen,
	)
	if !ok {
		fmt.Println("OK")
		return
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating capability resolver:", err)
		return
	}
	opts.StreamHandler = func(event spec.StreamEvent) error {
		switch event.Kind {
		case spec.StreamContentKindThinking:
			if event.Thinking != nil {
				fmt.Fprintf(os.Stderr, "[thinking] %s", event.Thinking.Text)
			}
		case spec.StreamContentKindText:
			if event.Text != nil {
				fmt.Fprintf(os.Stderr, "[text] %s", event.Text.Text)
			}
		}
		return nil
	}

	param := mp.ModelParam
	param.Stream = true
	param.MaxOutputLength = min(param.MaxOutputLength, 1024)
	param.Reasoning = &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: spec.ReasoningLevelLow,
	}
	param.SystemPrompt = "Use Google Search for current facts. Keep the final answer short."

	webSearch := spec.ToolChoice{
		Type:        spec.ToolTypeWebSearch,
		ID:          "google-search",
		Name:        spec.DefaultWebSearchToolName,
		Description: "Search current public web information.",
	}

	resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: param,
		Inputs: []spec.InputUnion{
			newUserTextInput(
				"What is the latest stable Go release? State uncertainty if search does not provide an answer.",
			),
		},
		ToolChoices: []spec.ToolChoice{webSearch},
		ToolPolicy:  &spec.ToolPolicy{Mode: spec.ToolPolicyModeAuto},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nweb-search interaction error:", err)
		return
	}

	var (
		searchCalls   int
		searchResults int
		citations     int
	)
	for _, output := range resp.Outputs {
		switch output.Kind {
		case spec.OutputKindWebSearchToolCall:
			if output.WebSearchToolCall != nil {
				searchCalls += len(
					output.WebSearchToolCall.WebSearchToolCallItems,
				)
			}
		case spec.OutputKindWebSearchToolOutput:
			if output.WebSearchToolOutput != nil {
				searchResults += len(
					output.WebSearchToolOutput.WebSearchToolOutputItems,
				)
			}
		case spec.OutputKindOutputMessage:
			if output.OutputMessage == nil {
				continue
			}
			for _, content := range output.OutputMessage.Contents {
				if content.TextItem != nil {
					citations += len(content.TextItem.Citations)
				}
			}
		default:
		}
	}

	fmt.Fprintln(os.Stderr, "\nanswer:", responseText(resp))
	fmt.Fprintln(os.Stderr, "grounding calls:", searchCalls)
	fmt.Fprintln(os.Stderr, "grounding result items:", searchResults)
	fmt.Fprintln(os.Stderr, "normalized URL citations:", citations)
	fmt.Println("OK")
	// Output: OK
}

func googleInteractionsExampleSetup(
	ctx context.Context,
	modelID modelpreset.ModelPresetID,
) (
	*inference.ProviderSetAPI,
	modelpreset.ProviderPreset,
	modelpreset.ModelPreset,
	bool,
) {
	ps, err := newProviderSetWithDebug(slog.LevelInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return nil, modelpreset.ProviderPreset{}, modelpreset.ModelPreset{}, false
	}

	pp, mp, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderGoogleGeminiInteractions,
		modelID,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Google Interactions provider:", err)
		return nil, modelpreset.ProviderPreset{}, modelpreset.ModelPreset{}, false
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		fmt.Fprintln(
			os.Stderr,
			"GEMINI_API_KEY/GOOGLE_API_KEY not set; skipping live Google Interactions example",
		)
		return nil, modelpreset.ProviderPreset{}, modelpreset.ModelPreset{}, false
	}

	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google API key:", err)
		return nil, modelpreset.ProviderPreset{}, modelpreset.ModelPreset{}, false
	}

	return ps, pp, mp, true
}
