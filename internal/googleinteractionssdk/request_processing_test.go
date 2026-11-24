package googleinteractionssdk

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

func TestBuildGoogleInteractionsCreateInteraction(t *testing.T) {
	omitted := spec.ReasoningSummaryStyleOmitted
	temperature := 0.4
	rawParams := `{"store":true,"background":true}`
	imageData := base64.StdEncoding.EncodeToString([]byte("image"))
	fileData := base64.StdEncoding.EncodeToString([]byte("name,count\ncat,1\n"))

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            "gemini-4-hydrogen",
			Stream:          true,
			MaxOutputLength: 512,
			Temperature:     &temperature,
			Reasoning: &spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelXHigh,
				SummaryStyle: &omitted,
			},
			SystemPrompt:  "Return a concise answer.",
			StopSequences: []string{"STOP"},
			OutputParam: &spec.OutputParam{
				Format: &spec.OutputFormat{
					Kind: spec.OutputFormatKindJSONSchema,
					JSONSchemaParam: &spec.JSONSchemaParam{
						Name:        "answer",
						Description: "Structured answer",
						Strict:      true,
						Schema: map[string]any{
							"type": "object",
						},
					},
				},
			},
			AdditionalParametersRawJSON: &rawParams,
		},
		Inputs: []spec.InputUnion{
			{
				Kind: spec.InputKindInputMessage,
				InputMessage: &spec.InputOutputContent{
					Role: spec.RoleUser,
					Contents: []spec.InputOutputContentItemUnion{
						{
							Kind:     spec.ContentItemKindText,
							TextItem: &spec.ContentItemText{Text: "Describe these inputs."},
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
								FileData: fileData,
								FileMIME: "text/csv",
							},
						},
					},
				},
			},
		},
		ToolChoices: []spec.ToolChoice{
			{
				Type:        spec.ToolTypeFunction,
				ID:          "echo",
				Name:        "echo text",
				Description: "Echo a text value.",
				Arguments: map[string]any{
					"type": "object",
				},
			},
			{
				Type: spec.ToolTypeWebSearch,
				ID:   "search",
				Name: "google_search",
			},
		},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeTool,
			AllowedTools: []spec.AllowedTool{
				{ToolChoiceID: "echo"},
			},
			DisableParallel: true,
		},
	}

	build, err := buildGoogleInteractionsCreateInteraction(req, true)
	if err != nil {
		t.Fatalf("buildGoogleInteractionsCreateInteraction: %v", err)
	}

	raw, err := json.Marshal(build.Interaction)
	if err != nil {
		t.Fatal(err)
	}
	payload := string(raw)

	for _, want := range []string{
		`"store":false`,
		`"background":false`,
		`"stream":true`,
		`"thinking_level":"xhigh"`,
		`"thinking_summaries":"none"`,
		`"max_output_tokens":512`,
		`"stop_sequences":["STOP"]`,
		`"application/json"`,
		`"echo_text"`,
		`"image"`,
		`"document"`,
	} {
		if !strings.Contains(payload, want) {
			t.Errorf("payload does not contain %q: %s", want, payload)
		}
	}

	for _, unwanted := range []string{
		`"temperature"`,
		`"previous_interaction_id"`,
		`"agent"`,
		rawParams,
		`"google_search"`,
		`"continuation_token"`,
		`"service_tier"`,
		`"safety_settings"`,
		`"video_config"`,
	} {
		if strings.Contains(payload, unwanted) {
			t.Errorf("payload unexpectedly contains %q: %s", unwanted, payload)
		}
	}

	if build.WebSearchChoiceID != "" {
		t.Fatalf(
			"forced callable policy must omit web search, got %q",
			build.WebSearchChoiceID,
		)
	}
	if got := build.ToolChoiceNameMap["echo_text"].ID; got != "echo" {
		t.Fatalf("tool map ID got %q want echo", got)
	}

	for _, code := range []string{
		"temperature_dropped_unsupported",
		"additionalParameters_dropped_unsupported",
		"jsonSchema_metadata_not_forwarded",
		"toolPolicy_disableParallel_unsupported",
	} {
		if !googleInteractionsHasWarning(build.Warnings, code) {
			t.Errorf("missing warning %q in %#v", code, build.Warnings)
		}
	}
}

func TestGoogleInteractionsInputRejectsUnsupportedMedia(t *testing.T) {
	_, err := googleInteractionsContents(
		[]spec.InputOutputContentItemUnion{{
			Kind: spec.ContentItemKindImage,
			ImageItem: &spec.ContentItemImage{
				ImageData: "abc",
				ImageMIME: "image/svg+xml",
			},
		}},
		new([]spec.Warning),
	)
	if err == nil {
		t.Fatal("expected unsupported image MIME error")
	}

	_, err = googleInteractionsContents(
		[]spec.InputOutputContentItemUnion{{
			Kind: spec.ContentItemKindFile,
			FileItem: &spec.ContentItemFile{
				FileData: "abc",
				FileMIME: "application/octet-stream",
			},
		}},
		new([]spec.Warning),
	)
	if err == nil {
		t.Fatal("expected unsupported document MIME error")
	}
}

func TestGoogleInteractionsTextFormatUsesSDKDefault(t *testing.T) {
	format, warnings, err := googleInteractionsResponseFormat(
		&spec.OutputParam{
			Format: &spec.OutputFormat{
				Kind: spec.OutputFormatKindText,
			},
		},
	)
	if err != nil {
		t.Fatalf("googleInteractionsResponseFormat: %v", err)
	}
	if format != nil {
		t.Fatalf("text format must use SDK default, got %#v", format)
	}
	if len(warnings) != 0 {
		t.Fatalf("unexpected text format warnings %#v", warnings)
	}
}

func TestGoogleInteractionsFunctionResultUsesStringUnion(t *testing.T) {
	output := &spec.ToolOutput{
		Type:   spec.ToolTypeFunction,
		CallID: "call_1",
		Name:   "echo",
		Contents: []spec.ToolOutputItemUnion{{
			Kind:     spec.ContentItemKindText,
			TextItem: &spec.ContentItemText{Text: "tool result"},
		}},
	}

	step, err := googleInteractionsFunctionResultStep(output, nil)
	if err != nil {
		t.Fatalf("googleInteractionsFunctionResultStep: %v", err)
	}
	if step == nil || step.FunctionResultStep == nil {
		t.Fatalf("missing function result step %#v", step)
	}
	if step.FunctionResultStep.Result.Str == nil ||
		*step.FunctionResultStep.Result.Str != "tool result" {
		t.Fatalf(
			"unexpected FunctionResultStepResultUnion %#v",
			step.FunctionResultStep.Result,
		)
	}
}

func TestGoogleInteractionsFunctionResultSupportsRichContentList(t *testing.T) {
	output := &spec.ToolOutput{
		Type:   spec.ToolTypeFunction,
		CallID: "call_1",
		Contents: []spec.ToolOutputItemUnion{
			{
				Kind:     spec.ContentItemKindText,
				TextItem: &spec.ContentItemText{Text: "text"},
			},
			{
				Kind: spec.ContentItemKindImage,
				ImageItem: &spec.ContentItemImage{
					ImageData: "aGVsbG8=",
					ImageMIME: "image/png",
				},
			},
			{
				Kind: spec.ContentItemKindFile,
				FileItem: &spec.ContentItemFile{
					FileData: "Y3N2",
					FileMIME: "text/csv",
				},
			},
		},
	}

	step, err := googleInteractionsFunctionResultStep(output, nil)
	if err != nil {
		t.Fatalf("googleInteractionsFunctionResultStep: %v", err)
	}
	if step == nil || step.FunctionResultStep == nil {
		t.Fatalf("missing function result step %#v", step)
	}
	union := step.FunctionResultStep.Result
	if len(union.ArrayOfFunctionResultSubcontent) != 3 {
		t.Fatalf(
			"expected text, image, and document subcontents, got %#v",
			union,
		)
	}
	if union.ArrayOfFunctionResultSubcontent[0].TextContent == nil {
		t.Fatalf("missing text subcontent %#v", union)
	}
	if union.ArrayOfFunctionResultSubcontent[1].ImageContent == nil {
		t.Fatalf("missing image subcontent %#v", union)
	}
	if union.ArrayOfFunctionResultSubcontent[2].DocumentContent == nil {
		t.Fatalf("missing document subcontent %#v", union)
	}
}

func TestSanitizeGoogleInteractionsReasoningInputs(t *testing.T) {
	inputs := []spec.InputUnion{
		{
			Kind: spec.InputKindReasoningMessage,
			ReasoningMessage: &spec.ReasoningContent{
				Thinking: []string{"unsigned thought"},
			},
		},
		{
			Kind: spec.InputKindReasoningMessage,
			ReasoningMessage: &spec.ReasoningContent{
				Signature:        "opaque-google-signature",
				Summary:          []string{"summary"},
				Thinking:         []string{"must be removed"},
				RedactedThinking: []string{"must be removed"},
				EncryptedContent: []string{"must be removed"},
			},
		},
	}

	got := sanitizeGoogleInteractionsReasoningInputs(inputs)
	if len(got) != 1 {
		t.Fatalf("reasoning input count got %d want 1", len(got))
	}

	reasoning := got[0].ReasoningMessage
	if reasoning == nil ||
		reasoning.Signature != "opaque-google-signature" ||
		len(reasoning.Summary) != 1 ||
		len(reasoning.Thinking) != 0 ||
		len(reasoning.RedactedThinking) != 0 ||
		len(reasoning.EncryptedContent) != 0 {
		t.Fatalf("unexpected sanitized reasoning %#v", reasoning)
	}
}

func TestGoogleInteractionsTemperatureIsAlwaysDroppedAtBuildTime(t *testing.T) {
	temperature := 0.4
	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:        "gemini-4-hydrogen",
			Temperature: &temperature,
			Reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelLow,
			},
		},
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "hello"},
				}},
			},
		}},
	}

	normalized, _, normalizationWarnings, err := sdkutil.NormalizeRequestForSDK(
		t.Context(),
		req,
		nil,
		spec.ProviderSDKTypeGoogleInteractions,
		googleInteractionsSDKCapability(),
	)
	if err != nil {
		t.Fatalf("NormalizeRequestForSDK: %v", err)
	}
	if normalized.ModelParam.Temperature == nil {
		t.Fatal("temperature must reach the adapter for unconditional dropping")
	}
	if googleInteractionsHasWarning(
		normalizationWarnings,
		"temperature_dropped_reasoning_enabled",
	) {
		t.Fatalf(
			"temperature must not be described as reasoning-only incompatible: %#v",
			normalizationWarnings,
		)
	}

	build, err := buildGoogleInteractionsCreateInteraction(normalized, false)
	if err != nil {
		t.Fatalf("buildGoogleInteractionsCreateInteraction: %v", err)
	}
	if !googleInteractionsHasWarning(
		build.Warnings,
		"temperature_dropped_unsupported",
	) {
		t.Fatalf("missing unconditional temperature warning: %#v", build.Warnings)
	}

	raw, err := json.Marshal(build.Interaction)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), `"temperature"`) {
		t.Fatalf("temperature leaked into request payload: %s", raw)
	}
}

func TestGoogleInteractionsFunctionResultUsesToolOutputIDAsCallIDFallback(
	t *testing.T,
) {
	step, err := googleInteractionsFunctionResultStep(&spec.ToolOutput{
		Type: spec.ToolTypeFunction,
		ID:   "call_from_id",
		Name: "echo",
		Contents: []spec.ToolOutputItemUnion{{
			Kind:     spec.ContentItemKindText,
			TextItem: &spec.ContentItemText{Text: "result"},
		}},
	}, nil)
	if err != nil {
		t.Fatalf("googleInteractionsFunctionResultStep: %v", err)
	}
	if step == nil || step.FunctionResultStep == nil {
		t.Fatalf("missing function result step: %#v", step)
	}
	if step.FunctionResultStep.CallID != "call_from_id" {
		t.Fatalf(
			"CallID got %q want call_from_id",
			step.FunctionResultStep.CallID,
		)
	}
}

func TestGoogleInteractionsContextOnlyFileIsSafelyConvertedToText(t *testing.T) {
	warnings := []spec.Warning{}
	contents, err := googleInteractionsContents(
		[]spec.InputOutputContentItemUnion{{
			Kind: spec.ContentItemKindFile,
			FileItem: &spec.ContentItemFile{
				AdditionalContext: "The attachment is unavailable. Use this context.",
			},
		}},
		&warnings,
	)
	if err != nil {
		t.Fatalf("googleInteractionsContents: %v", err)
	}
	if len(contents) != 1 || contents[0].TextContent == nil {
		t.Fatalf("context-only file contents: %#v", contents)
	}
	if contents[0].TextContent.Text !=
		"The attachment is unavailable. Use this context." {
		t.Fatalf("unexpected context text: %#v", contents[0].TextContent)
	}
	if !googleInteractionsHasWarning(warnings, "file_dropped_empty") {
		t.Fatalf("expected empty-file warning: %#v", warnings)
	}
}

func TestGoogleInteractionsCapabilitiesAreIndependent(t *testing.T) {
	first := googleInteractionsSDKCapability()
	first.ModalitiesIn[0] = spec.ModalityAudioIn
	first.ReasoningCapabilities.SupportedReasoningLevels[0] = spec.ReasoningLevelMax

	second := googleInteractionsSDKCapability()
	if second.ModalitiesIn[0] != spec.ModalityTextIn {
		t.Fatalf("modalities share mutable state: %#v", second.ModalitiesIn)
	}
	if second.ReasoningCapabilities.SupportedReasoningLevels[0] !=
		spec.ReasoningLevelMinimal {
		t.Fatalf(
			"reasoning levels share mutable state: %#v",
			second.ReasoningCapabilities.SupportedReasoningLevels,
		)
	}
}

func googleInteractionsHasWarning(
	warnings []spec.Warning,
	code string,
) bool {
	for _, warning := range warnings {
		if warning.Code == code {
			return true
		}
	}
	return false
}
