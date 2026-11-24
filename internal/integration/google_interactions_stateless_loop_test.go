package integration

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

const (
	googleInteractionsLoopToolID   = "interactions-loop-echo"
	googleInteractionsLoopToolName = "echo_text"
)

func TestGoogleInteractions_FunctionToolRoundTripLoop(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY/GOOGLE_API_KEY not set")
	}

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Minute)
	defer cancel()

	providerSet, err := newProviderSetWithDebug(slog.LevelInfo)
	if err != nil {
		t.Fatal(err)
	}

	provider, model, err := addCatalogModelProvider(
		ctx,
		providerSet,
		modelpreset.ProviderGoogleGeminiInteractions,
		modelpreset.PresetGemini4Hydrogen,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := providerSet.SetProviderAPIKey(ctx, provider.Name, apiKey); err != nil {
		t.Fatal(err)
	}

	sawSignedReasoning := false
	for iteration := range 3 {
		saw, err := runGoogleInteractionsEchoRoundTrip(
			ctx,
			providerSet,
			provider,
			model,
			fmt.Sprintf("interactions loop %d", iteration),
		)
		if err != nil {
			t.Fatalf("iteration %d failed: %v", iteration, err)
		}
		sawSignedReasoning = sawSignedReasoning || saw
	}

	if !sawSignedReasoning {
		t.Fatal(
			"did not observe a signed Interactions thought; reasoning replay was not exercised",
		)
	}
}

func runGoogleInteractionsEchoRoundTrip(
	ctx context.Context,
	providerSet *inference.ProviderSetAPI,
	provider modelpreset.ProviderPreset,
	model modelpreset.ModelPreset,
	payload string,
) (bool, error) {
	tool := newEchoToolChoice(
		googleInteractionsLoopToolID,
		googleInteractionsLoopToolName,
	)

	initialUser := newUserTextInput(
		fmt.Sprintf(
			`Think briefly, then call %s with text %q. Do not answer until after the first tool result.`,
			googleInteractionsLoopToolName,
			payload,
		),
	)
	history := make([]spec.InputUnion, 0, 8)
	history = append(history, initialUser)

	var sawStreamedText atomic.Bool
	sawSignedReasoning := false
	toolCalls := 0
	seenCallIDs := map[string]struct{}{}

	for turn := range 5 {
		policy := &spec.ToolPolicy{Mode: spec.ToolPolicyModeNone}
		if toolCalls < 2 {
			policy = &spec.ToolPolicy{
				Mode: spec.ToolPolicyModeTool,
				AllowedTools: []spec.AllowedTool{
					{ToolChoiceID: tool.ID},
				},
				DisableParallel: true,
			}
		}

		summaryStyle := spec.ReasoningSummaryStyleAuto
		modelParam := model.ModelParam
		modelParam.Stream = true
		modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 2048)
		modelParam.SystemPrompt = "Preserve prior assistant thought and tool state. " +
			"Call the forced function on each requested tool turn. " +
			"After the second tool result, answer plainly and do not call another tool."
		modelParam.Reasoning = &spec.ReasoningParam{
			Type:         spec.ReasoningTypeSingleWithLevels,
			Level:        spec.ReasoningLevelLow,
			SummaryStyle: &summaryStyle,
		}

		options, err := presetFetchOptions(ctx, providerSet, provider, model)
		if err != nil {
			return sawSignedReasoning, fmt.Errorf(
				"turn %d preset capability resolver: %w",
				turn,
				err,
			)
		}
		options.StreamHandler = func(event spec.StreamEvent) error {
			if event.Kind == spec.StreamContentKindText &&
				event.Text != nil &&
				event.Text.Text != "" {
				sawStreamedText.Store(true)
			}
			return nil
		}

		response, err := providerSet.FetchCompletion(
			ctx,
			provider.Name,
			&spec.FetchCompletionRequest{
				ModelParam:  modelParam,
				Inputs:      history,
				ToolChoices: []spec.ToolChoice{tool},
				ToolPolicy:  policy,
			},
			options,
		)
		if err != nil {
			return sawSignedReasoning, fmt.Errorf("turn %d fetch: %w", turn, err)
		}

		sawSignedReasoning = sawSignedReasoning ||
			responseHasGoogleSignature(response)
		history = append(history, outputUnionsToInputs(response.Outputs)...)

		call, _ := firstFunctionToolCall(response)
		if call != nil {
			if toolCalls >= 2 {
				return sawSignedReasoning, fmt.Errorf(
					"turn %d: received an unexpected function call after tool policy none",
					turn,
				)
			}

			callID := nonEmpty(call.CallID, call.ID)
			if callID == "" {
				return sawSignedReasoning, fmt.Errorf(
					"turn %d: function call has no call id",
					turn,
				)
			}
			if _, exists := seenCallIDs[callID]; exists {
				return sawSignedReasoning, fmt.Errorf(
					"turn %d: duplicate function call id %q",
					turn,
					callID,
				)
			}
			seenCallIDs[callID] = struct{}{}

			toolOutput, err := runEchoTool(call)
			if err != nil {
				return sawSignedReasoning, fmt.Errorf(
					"turn %d tool execution: %w",
					turn,
					err,
				)
			}
			if toolOutput.CallID != callID {
				return sawSignedReasoning, fmt.Errorf(
					"turn %d: tool output call id got %q want %q",
					turn,
					toolOutput.CallID,
					callID,
				)
			}

			toolCalls++
			nextUser := newUserTextInput(
				"Use both tool results and answer in one short sentence without another tool call.",
			)
			if toolCalls == 1 {
				nextUser = newUserTextInput(
					fmt.Sprintf(
						`Now call %s with text %q. Do not answer yet.`,
						googleInteractionsLoopToolName,
						payload+" second",
					),
				)
			}

			history = append(history,
				spec.InputUnion{
					Kind:               spec.InputKindFunctionToolOutput,
					FunctionToolOutput: toolOutput,
				},
				nextUser,
			)
			continue
		}

		if toolCalls < 2 {
			return sawSignedReasoning, fmt.Errorf(
				"turn %d: expected forced function call %d",
				turn,
				toolCalls+1,
			)
		}
		if responseText(response) == "" {
			return sawSignedReasoning, fmt.Errorf(
				"turn %d: expected final assistant text",
				turn,
			)
		}
		if !sawStreamedText.Load() {
			return sawSignedReasoning, errors.New(
				"expected at least one streamed text event",
			)
		}

		return sawSignedReasoning, nil
	}

	return sawSignedReasoning, errors.New(
		"interaction conversation did not complete within five turns",
	)
}
