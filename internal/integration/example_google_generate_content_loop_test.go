package integration

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

const (
	googleGenerateContentLoopToolID   = "echo-tool"
	googleGenerateContentLoopToolName = "echo_text"
)

func TestGoogleGenerateContent_FunctionToolRoundTripLoop(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY/GOOGLE_API_KEY not set")
	}

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Minute)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		t.Fatal(err)
	}

	pp, mp, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderGoogleGemini,
		modelpreset.PresetGoogleGemini25Flash,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		t.Fatal(err)
	}

	sawSignedReasoning := false
	for i := range 3 {
		saw, err := runGoogleGenerateContentEchoRoundTrip(
			ctx,
			ps,
			pp,
			mp,
			fmt.Sprintf("google loop %d", i),
		)
		if err != nil {
			t.Fatalf("iteration %d failed: %v", i, err)
		}
		if saw {
			sawSignedReasoning = true
		}
	}

	if !sawSignedReasoning {
		t.Fatalf("did not observe any signed reasoning across loop; thought-signature replay path was not exercised")
	}
}

func runGoogleGenerateContentEchoRoundTrip(
	ctx context.Context,
	ps *inference.ProviderSetAPI,
	pp modelpreset.ProviderPreset,
	mp modelpreset.ModelPreset,
	payload string,
) (bool, error) {
	tool := newEchoToolChoice(googleGenerateContentLoopToolID, googleGenerateContentLoopToolName)

	history := []spec.InputUnion{
		newUserTextInput(
			fmt.Sprintf(
				`Think briefly, then call %s with text %q. After the tool result arrives, answer in one short sentence.`,
				googleGenerateContentLoopToolName,
				payload,
			),
		),
	}

	sawToolCall := false
	sawSignedReasoning := false

	for turn := range 4 {
		policy := &spec.ToolPolicy{Mode: spec.ToolPolicyModeNone}
		if !sawToolCall {
			policy = &spec.ToolPolicy{
				Mode: spec.ToolPolicyModeTool,
				AllowedTools: []spec.AllowedTool{
					{ToolChoiceID: tool.ID},
				},
				DisableParallel: true,
			}
		}

		modelParam := mp.ModelParam
		modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 4096)
		modelParam.SystemPrompt = "Preserve prior assistant reasoning/tool state across turns. " +
			"After the tool result is available, answer plainly and do not call the tool again."
		modelParam.Reasoning = &spec.ReasoningParam{
			Type:   spec.ReasoningTypeHybridWithTokens,
			Tokens: 2048,
		}

		opts, err := presetFetchOptions(ctx, ps, pp, mp)
		if err != nil {
			return sawSignedReasoning, fmt.Errorf("turn %d preset capability resolver: %w", turn, err)
		}

		resp, err := ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
			ModelParam:  modelParam,
			Inputs:      history,
			ToolChoices: []spec.ToolChoice{tool},
			ToolPolicy:  policy,
		}, opts)
		if err != nil {
			return sawSignedReasoning, fmt.Errorf("turn %d fetch: %w", turn, err)
		}

		if responseHasGoogleSignature(resp) {
			sawSignedReasoning = true
		}

		history = append(history, outputUnionsToInputs(resp.Outputs)...)

		call, _ := firstFunctionToolCall(resp)
		if call != nil && !sawToolCall {
			sawToolCall = true

			toolOutput, err := runEchoTool(call)
			if err != nil {
				return sawSignedReasoning, fmt.Errorf("turn %d tool exec: %w", turn, err)
			}

			history = append(history, spec.InputUnion{
				Kind:               spec.InputKindFunctionToolOutput,
				FunctionToolOutput: toolOutput,
			})
			continue
		}

		if !sawToolCall {
			return sawSignedReasoning, fmt.Errorf("turn %d: expected tool call before final answer", turn)
		}
		if responseText(resp) == "" {
			return sawSignedReasoning, fmt.Errorf("turn %d: expected final assistant text", turn)
		}

		return sawSignedReasoning, nil
	}

	return sawSignedReasoning, errors.New("conversation did not complete within 4 turns")
}

func responseHasGoogleSignature(resp *spec.FetchCompletionResponse) bool {
	if resp == nil {
		return false
	}

	for _, out := range resp.Outputs {
		switch out.Kind {
		case spec.OutputKindReasoningMessage:
			if out.ReasoningMessage != nil && out.ReasoningMessage.Signature != "" {
				return true
			}

		case spec.OutputKindFunctionToolCall:
			if out.FunctionToolCall != nil && out.FunctionToolCall.Signature != "" {
				return true
			}

		case spec.OutputKindCustomToolCall:
			if out.CustomToolCall != nil && out.CustomToolCall.Signature != "" {
				return true
			}

		case spec.OutputKindOutputMessage:
			if out.OutputMessage == nil {
				continue
			}
			for _, c := range out.OutputMessage.Contents {
				if c.Kind == spec.ContentItemKindText &&
					c.TextItem != nil &&
					c.TextItem.Signature != "" {
					return true
				}
			}
		default:
		}
	}

	return false
}
