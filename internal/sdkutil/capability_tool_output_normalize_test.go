package sdkutil

import (
	"strings"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestNormalizeRequestForSDK_CollapsesRichClientToolOutputsWhenStringOnly(t *testing.T) {
	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{Name: "test-model"},
		Inputs: []spec.InputUnion{
			{
				Kind: spec.InputKindFunctionToolOutput,
				FunctionToolOutput: &spec.ToolOutput{
					Type:   spec.ToolTypeFunction,
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
				},
			},
		},
	}

	caps := spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn},
		ToolCapabilities: &spec.ToolCapabilities{
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	}

	got, warns, err := NormalizeRequestForSDK(
		t.Context(),
		req,
		nil,
		spec.ProviderSDKTypeOpenAIResponses,
		caps,
	)
	if err != nil {
		t.Fatalf("NormalizeRequestForSDK error: %v", err)
	}

	out := got.Inputs[0].FunctionToolOutput
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
		if w.Code == "toolOutput_collapsed_to_string" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected toolOutput_collapsed_to_string warning, got %#v", warns)
	}
}

func TestNormalizeRequestForSDK_LeavesSingleTextToolOutputUntouchedWhenStringOnly(t *testing.T) {
	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{Name: "test-model"},
		Inputs: []spec.InputUnion{
			{
				Kind: spec.InputKindFunctionToolOutput,
				FunctionToolOutput: &spec.ToolOutput{
					Type:   spec.ToolTypeFunction,
					CallID: "call_1",
					Name:   "demo_tool",
					Contents: []spec.ToolOutputItemUnion{{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: "plain text"},
					}},
				},
			},
		},
	}

	caps := spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{spec.ModalityTextIn},
		ToolCapabilities: &spec.ToolCapabilities{
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	}

	got, _, err := NormalizeRequestForSDK(t.Context(), req, nil, spec.ProviderSDKTypeOpenAIResponses, caps)
	if err != nil {
		t.Fatalf("NormalizeRequestForSDK error: %v", err)
	}
	if got.Inputs[0].FunctionToolOutput.Contents[0].TextItem.Text != "plain text" {
		t.Fatalf("single text output was unexpectedly rewritten")
	}
}
