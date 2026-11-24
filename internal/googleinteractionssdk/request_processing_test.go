package googleinteractionssdk

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

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
	if len(union.ArrayOfFunctionResultSubcontent) != 2 {
		t.Fatalf(
			"expected text plus image subcontents, got %#v",
			union,
		)
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
