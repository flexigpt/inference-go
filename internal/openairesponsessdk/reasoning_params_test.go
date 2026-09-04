package openairesponsessdk

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestOpenAIResponsesReasoningRequestParameters(t *testing.T) {
	omitted := spec.ReasoningSummaryStyleOmitted
	auto := spec.ReasoningSummaryStyleAuto
	concise := spec.ReasoningSummaryStyleConcise
	detailed := spec.ReasoningSummaryStyleDetailed

	tests := []struct {
		name         string
		summaryStyle *spec.ReasoningSummaryStyle
		context      *spec.ReasoningContext
		mode         *spec.ReasoningMode
		wantSummary  string
		wantContext  string
		wantMode     string
	}{
		{
			name:        "default summary is auto with no context or mode",
			wantSummary: "auto",
		},
		{
			name:         "omitted maps to concise with auto and standard",
			summaryStyle: &omitted,
			context:      new(spec.ReasoningContextAuto),
			mode:         new(spec.ReasoningModeStandard),
			wantSummary:  "concise",
			wantContext:  "auto",
			wantMode:     "standard",
		},
		{
			name:         "auto retains current turn and pro",
			summaryStyle: &auto,
			context:      new(spec.ReasoningContextCurrentTurn),
			mode:         new(spec.ReasoningModePro),
			wantSummary:  "auto",
			wantContext:  "current_turn",
			wantMode:     "pro",
		},
		{
			name:         "concise retains all turns and standard",
			summaryStyle: &concise,
			context:      new(spec.ReasoningContextAllTurns),
			mode:         new(spec.ReasoningModeStandard),
			wantSummary:  "concise",
			wantContext:  "all_turns",
			wantMode:     "standard",
		},
		{
			name:         "detailed retains auto and pro",
			summaryStyle: &detailed,
			context:      new(spec.ReasoningContextAuto),
			mode:         new(spec.ReasoningModePro),
			wantSummary:  "detailed",
			wantContext:  "auto",
			wantMode:     "pro",
		},
	}

	requestBodies := make(chan map[string]any, len(tests))
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		raw, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var body map[string]any
		if err := json.Unmarshal(raw, &body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		requestBodies <- body

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":{"message":"intentional test response"}}`))
	}))
	defer server.Close()

	api, err := NewOpenAIResponsesAPI(spec.ProviderParam{
		Name:                     "openai-responses-test",
		SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
		APIKey:                   "test-api-key",
		Origin:                   server.URL,
		ChatCompletionPathPrefix: "/v1/responses",
	}, nil)
	if err != nil {
		t.Fatalf("NewOpenAIResponsesAPI: %v", err)
	}
	if err := api.InitLLM(t.Context()); err != nil {
		t.Fatalf("InitLLM: %v", err)
	}
	t.Cleanup(func() {
		_ = api.DeInitLLM(t.Context())
	})

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := api.FetchCompletion(
				t.Context(),
				openAIResponsesReasoningTestRequest(
					tc.summaryStyle,
					tc.context,
					tc.mode,
				),
				nil,
			)
			if err == nil {
				t.Fatal("expected intentional HTTP error")
			}

			body := <-requestBodies
			reasoning, ok := body["reasoning"].(map[string]any)
			if !ok {
				t.Fatalf("request reasoning missing or invalid: %#v", body["reasoning"])
			}
			assertOpenAIReasoningString(t, reasoning, "summary", tc.wantSummary)
			assertOpenAIReasoningOptionalString(t, reasoning, "context", tc.wantContext)
			assertOpenAIReasoningOptionalString(t, reasoning, "mode", tc.wantMode)
		})
	}
}

func TestOpenAIResponsesReasoningCapabilities(t *testing.T) {
	caps := openairesponsessdkCapability.ReasoningCapabilities
	if caps == nil {
		t.Fatal("expected OpenAI Responses reasoning capabilities")
	}
	if !caps.SupportsSummaryStyle {
		t.Fatal("expected OpenAI Responses summary-style support")
	}
	if !caps.SupportsReasoningContext {
		t.Fatal("expected OpenAI Responses reasoning-context support")
	}
	if !caps.SupportsReasoningMode {
		t.Fatal("expected OpenAI Responses reasoning-mode support")
	}
}

func openAIResponsesReasoningTestRequest(
	summaryStyle *spec.ReasoningSummaryStyle,
	context *spec.ReasoningContext,
	mode *spec.ReasoningMode,
) *spec.FetchCompletionRequest {
	return &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name: "test-model",
			Reasoning: &spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelLow,
				SummaryStyle: summaryStyle,
				Context:      context,
				Mode:         mode,
			},
		},
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "test"},
				}},
			},
		}},
	}
}

func assertOpenAIReasoningOptionalString(
	t *testing.T,
	reasoning map[string]any,
	key string,
	want string,
) {
	t.Helper()

	got, exists := reasoning[key]
	if want == "" {
		if exists {
			t.Fatalf("reasoning.%s got %#v, expected omission", key, got)
		}
		return
	}

	assertOpenAIReasoningString(t, reasoning, key, want)
}

func assertOpenAIReasoningString(
	t *testing.T,
	reasoning map[string]any,
	key string,
	want string,
) {
	t.Helper()

	got, ok := reasoning[key].(string)
	if !ok || got != want {
		t.Fatalf("reasoning.%s got %#v want %q", key, reasoning[key], want)
	}
}
