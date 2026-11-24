package googleinteractionssdk

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"

	"google.golang.org/genai/interactions/models/interactions"
)

func TestResponseFromGoogleInteraction(t *testing.T) {
	interaction := &interactions.Interaction{
		ID:     new("interaction_1"),
		Status: "completed",
		Usage: &interactions.Usage{
			TotalInputTokens:   new(int(10)),
			TotalCachedTokens:  new(int(3)),
			TotalOutputTokens:  new(int(7)),
			TotalThoughtTokens: new(int(4)),
		},
		Steps: []interactions.Step{
			interactions.NewStep(interactions.ThoughtStep{
				Signature: new("thought-signature"),
				Summary: []interactions.ThoughtSummaryContent{
					interactions.NewThoughtSummaryContent(
						interactions.TextContent{Text: "thought summary"},
					),
				},
			}),
			interactions.NewStep(interactions.FunctionCallStep{
				ID:   "call_1",
				Name: "echo_text",
				Arguments: map[string]any{
					"text": "hello",
				},
			}),
			interactions.NewStep(interactions.GoogleSearchCallStep{
				ID: "search_1",
				Arguments: interactions.GoogleSearchCallArguments{
					Queries: []string{"latest Go release"},
				},
				Signature: new("search-call-signature"),
			}),
			interactions.NewStep(interactions.GoogleSearchResultStep{
				CallID: "search_1",
				Result: []interactions.GoogleSearchResult{{
					SearchSuggestions: new("Go release search suggestion"),
				}},
				Signature: new("search-result-signature"),
			}),
			interactions.NewStep(interactions.ModelOutputStep{
				Content: []interactions.Content{
					interactions.NewContent(
						interactions.TextContent{Text: "final answer"},
					),
				},
			}),
		},
	}

	response, err := responseFromGoogleInteraction(
		interaction,
		map[string]spec.ToolChoice{
			"echo_text": {
				ID:   "echo",
				Name: "echo text",
				Type: spec.ToolTypeCustom,
			},
		},
		"search",
	)
	if err != nil {
		t.Fatalf("responseFromGoogleInteraction: %v", err)
	}

	if response.Usage.InputTokensTotal != 10 ||
		response.Usage.InputTokensCached != 3 ||
		response.Usage.InputTokensUncached != 7 ||
		response.Usage.OutputTokens != 7 ||
		response.Usage.ReasoningTokens != 4 {
		t.Fatalf("unexpected usage %#v", response.Usage)
	}

	if len(response.Outputs) != 5 {
		t.Fatalf("output count got %d want 5", len(response.Outputs))
	}

	reasoning := response.Outputs[0].ReasoningMessage
	if reasoning == nil ||
		reasoning.Signature != "thought-signature" ||
		len(reasoning.Summary) != 1 {
		t.Fatalf("unexpected reasoning %#v", reasoning)
	}

	customCall := response.Outputs[1].CustomToolCall
	if customCall == nil ||
		customCall.ChoiceID != "echo" ||
		customCall.CallID != "call_1" ||
		customCall.Arguments != `{"text":"hello"}` {
		t.Fatalf("unexpected custom call %#v", customCall)
	}

	searchCall := response.Outputs[2].WebSearchToolCall
	if searchCall == nil ||
		searchCall.Signature != "search-call-signature" ||
		len(searchCall.WebSearchToolCallItems) != 1 {
		t.Fatalf("unexpected search call %#v", searchCall)
	}

	searchResult := response.Outputs[3].WebSearchToolOutput
	if searchResult == nil ||
		searchResult.Signature != "search-result-signature" ||
		len(searchResult.WebSearchToolOutputItems) != 1 {
		t.Fatalf("unexpected search result %#v", searchResult)
	}

	message := response.Outputs[4].OutputMessage
	if message == nil ||
		len(message.Contents) != 1 ||
		message.Contents[0].TextItem == nil ||
		message.Contents[0].TextItem.Text != "final answer" {
		t.Fatalf("unexpected final message %#v", message)
	}

	if !googleInteractionsHasWarning(
		response.Warnings,
		"googleSearch_citations_unavailable",
	) {
		t.Fatalf("expected citation availability warning, got %#v", response.Warnings)
	}
}

func TestResponseFromGoogleInteractionPreservesPayloadWhitespaceAndSignatures(
	t *testing.T,
) {
	response, err := responseFromGoogleInteraction(
		&interactions.Interaction{
			ID:     new("interaction_1"),
			Status: "completed",
			Steps: []interactions.Step{
				interactions.NewStep(interactions.ThoughtStep{
					Signature: new(" opaque signature "),
					Summary: []interactions.ThoughtSummaryContent{
						interactions.NewThoughtSummaryContent(
							interactions.TextContent{Text: "  thought summary  "},
						),
					},
				}),
				interactions.NewStep(interactions.ModelOutputStep{
					Content: []interactions.Content{
						interactions.NewContent(
							interactions.TextContent{Text: "  answer with spacing  "},
						),
					},
				}),
			},
		},
		nil,
		"",
	)
	if err != nil {
		t.Fatalf("responseFromGoogleInteraction: %v", err)
	}
	if len(response.Outputs) != 2 {
		t.Fatalf("output count got %d want 2", len(response.Outputs))
	}

	reasoning := response.Outputs[0].ReasoningMessage
	if reasoning == nil ||
		reasoning.Signature != " opaque signature " ||
		len(reasoning.Summary) != 1 ||
		reasoning.Summary[0] != "  thought summary  " {
		t.Fatalf("unexpected reasoning: %#v", reasoning)
	}

	message := response.Outputs[1].OutputMessage
	if message == nil ||
		len(message.Contents) != 1 ||
		message.Contents[0].TextItem == nil ||
		message.Contents[0].TextItem.Text != "  answer with spacing  " {
		t.Fatalf("unexpected output message: %#v", message)
	}

	fallback, err := responseFromGoogleInteraction(
		&interactions.Interaction{
			ID:         new("interaction_2"),
			Status:     "completed",
			OutputText: new("  fallback answer  "),
		},
		nil,
		"",
	)
	if err != nil {
		t.Fatalf("responseFromGoogleInteraction fallback: %v", err)
	}
	if got := fallback.Outputs[0].OutputMessage.Contents[0].TextItem.Text; got != "  fallback answer  " {
		t.Fatalf("fallback text got %q", got)
	}
}

func TestResponseFromGoogleInteractionRequiresContinuation(t *testing.T) {
	response, err := responseFromGoogleInteraction(
		&interactions.Interaction{
			ID:     new("interaction_1"),
			Status: "requires_continuation",
			Steps: []interactions.Step{
				interactions.NewStep(interactions.ModelOutputStep{
					Content: []interactions.Content{
						interactions.NewContent(
							interactions.TextContent{Text: "partial"},
						),
					},
				}),
			},
		},
		nil,
		"",
	)
	if err == nil {
		t.Fatal("expected stateless continuation error")
	}
	if response == nil ||
		response.Error == nil ||
		len(response.Outputs) != 1 ||
		response.Outputs[0].OutputMessage == nil {
		t.Fatalf("expected partial response with error, got %#v", response)
	}
}

func TestResponseFromGoogleInteractionRequiresActionIsToolRoundTrip(t *testing.T) {
	response, err := responseFromGoogleInteraction(
		&interactions.Interaction{
			ID:     new("interaction_1"),
			Status: "requires_action",
			Steps: []interactions.Step{
				interactions.NewStep(interactions.FunctionCallStep{
					ID:        "call_1",
					Name:      "echo",
					Arguments: map[string]any{},
				}),
			},
		},
		map[string]spec.ToolChoice{
			"echo": {
				ID:   "echo",
				Name: "echo",
				Type: spec.ToolTypeFunction,
			},
		},
		"",
	)
	if err != nil {
		t.Fatalf("requires_action must support local function round trip: %v", err)
	}
	if response.Error != nil ||
		len(response.Outputs) != 1 ||
		response.Outputs[0].FunctionToolCall == nil ||
		response.Outputs[0].FunctionToolCall.Status != spec.StatusInProgress {
		t.Fatalf("unexpected requires_action response %#v", response)
	}
}

func TestGoogleInteractionsStatusAndTerminalErrorMapping(t *testing.T) {
	tests := []struct {
		name              string
		rawStatus         string
		wantStatus        spec.Status
		wantTerminalError bool
	}{
		{
			name:       "completed",
			rawStatus:  googleInteractionsRawStatusCompleted,
			wantStatus: spec.StatusCompleted,
		},
		{
			name:       "requires action",
			rawStatus:  googleInteractionsRawStatusRequiresAction,
			wantStatus: spec.StatusInProgress,
		},
		{
			name:              "in progress",
			rawStatus:         googleInteractionsRawStatusInProgress,
			wantStatus:        spec.StatusInProgress,
			wantTerminalError: true,
		},
		{
			name:       "incomplete",
			rawStatus:  googleInteractionsRawStatusIncomplete,
			wantStatus: spec.StatusIncomplete,
		},
		{
			name:              "requires continuation",
			rawStatus:         googleInteractionsRawStatusRequiresContinuation,
			wantStatus:        spec.StatusIncomplete,
			wantTerminalError: true,
		},
		{
			name:              "failed",
			rawStatus:         googleInteractionsRawStatusFailed,
			wantStatus:        spec.StatusFailed,
			wantTerminalError: true,
		},
		{
			name:              "cancelled",
			rawStatus:         googleInteractionsRawStatusCancelled,
			wantStatus:        spec.StatusCancelled,
			wantTerminalError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := googleInteractionsStatus(tc.rawStatus); got != tc.wantStatus {
				t.Fatalf("status got %q want %q", got, tc.wantStatus)
			}
			if gotError := googleInteractionsTerminalError(tc.rawStatus) != nil; gotError != tc.wantTerminalError {
				t.Fatalf(
					"terminal error got %v want %v",
					gotError,
					tc.wantTerminalError,
				)
			}
		})
	}
}
