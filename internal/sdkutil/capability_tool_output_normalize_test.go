package sdkutil

import (
	"strings"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestNormalizeRequestForSDK_CollapsesRichClientToolOutputsWhenStringOnly(t *testing.T) {
	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{Name: "test-model"},
		Inputs: toolOutputNormalizationTestHistory(&spec.ToolOutput{
			Type:   spec.ToolTypeFunction,
			Role:   spec.RoleUser,
			CallID: "call_1",
			Name:   "demo_tool",
			Contents: []spec.ToolOutputItemUnion{
				{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "hello"},
				},
				{
					Kind: spec.ContentItemKindImage,
					ImageItem: &spec.ContentItemImage{
						ImageURL: "https://example.com/a.png",
					},
				},
			},
		}),
	}

	caps := spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn},
		ToolCapabilities: &spec.ToolCapabilities{
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	}

	got, _, warns, err := NormalizeRequestForSDK(
		t.Context(),
		req,
		nil,
		spec.ProviderSDKTypeOpenAIResponses,
		caps,
	)
	if err != nil {
		t.Fatalf("NormalizeRequestForSDK error: %v", err)
	}

	if gotCount, want := len(got.Inputs), 3; gotCount != want {
		t.Fatalf("len(got.Inputs) = %d, want %d", gotCount, want)
	}

	out := got.Inputs[2].FunctionToolOutput
	if out == nil {
		t.Fatalf("FunctionToolOutput = nil")
	}
	if len(out.Contents) != 1 {
		t.Fatalf("len(out.Contents) = %d, want 1", len(out.Contents))
	}
	if out.Contents[0].Kind != spec.ContentItemKindText || out.Contents[0].TextItem == nil {
		t.Fatalf("collapsed output = %#v, want single text item", out.Contents[0])
	}
	if !strings.Contains(out.Contents[0].TextItem.Text, `"kind":"image"`) {
		t.Fatalf("collapsed text = %q, want JSON stringified original contents", out.Contents[0].TextItem.Text)
	}

	found := false
	for _, w := range warns {
		if w.Code == toolOutputCollapsedToString {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected toolOutput_collapsed_to_string warning, got %#v", warns)
	}
	if toolOutputNormalizationTestHasWarning(warns, conversationStartTrimmedWarningCode) {
		t.Fatalf("valid tool history unexpectedly trimmed: %#v", warns)
	}
}

func TestNormalizeRequestForSDK_LeavesSingleTextToolOutputUntouchedWhenStringOnly(t *testing.T) {
	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{Name: "test-model"},
		Inputs: toolOutputNormalizationTestHistory(&spec.ToolOutput{
			Type:   spec.ToolTypeFunction,
			Role:   spec.RoleUser,
			CallID: "call_1",
			Name:   "demo_tool",
			Contents: []spec.ToolOutputItemUnion{{
				Kind:     spec.ContentItemKindText,
				TextItem: &spec.ContentItemText{Text: "plain text"},
			}},
		}),
	}

	caps := spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{spec.ModalityTextIn},
		ToolCapabilities: &spec.ToolCapabilities{
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	}

	got, _, warns, err := NormalizeRequestForSDK(
		t.Context(),
		req,
		nil,
		spec.ProviderSDKTypeOpenAIResponses,
		caps,
	)
	if err != nil {
		t.Fatalf("NormalizeRequestForSDK error: %v", err)
	}
	if gotCount, want := len(got.Inputs), 3; gotCount != want {
		t.Fatalf("len(got.Inputs) = %d, want %d", gotCount, want)
	}

	out := got.Inputs[2].FunctionToolOutput
	if out == nil || len(out.Contents) != 1 || out.Contents[0].TextItem == nil {
		t.Fatalf("FunctionToolOutput = %#v, want one text item", out)
	}
	if out.Contents[0].TextItem.Text != "plain text" {
		t.Fatalf("single text output was unexpectedly rewritten")
	}
	if toolOutputNormalizationTestHasWarning(warns, conversationStartTrimmedWarningCode) {
		t.Fatalf("valid tool history unexpectedly trimmed: %#v", warns)
	}
}

func toolOutputNormalizationTestHistory(
	output *spec.ToolOutput,
) []spec.InputUnion {
	return []spec.InputUnion{
		{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "run the tool"},
				}},
			},
		},
		{
			Kind: spec.InputKindFunctionToolCall,
			FunctionToolCall: &spec.ToolCall{
				Type:      spec.ToolTypeFunction,
				Role:      spec.RoleAssistant,
				ID:        output.CallID,
				CallID:    output.CallID,
				Name:      output.Name,
				Arguments: "{}",
			},
		},
		{
			Kind:               spec.InputKindFunctionToolOutput,
			FunctionToolOutput: output,
		},
	}
}

func toolOutputNormalizationTestHasWarning(
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
