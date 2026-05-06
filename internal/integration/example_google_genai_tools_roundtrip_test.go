package integration

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/spec"
)

const (
	googleToolsRoundTripProviderName                = "google-tools"
	googleToolsRoundTripModelName                   = "gemini-2.5-flash"
	googleToolsRoundTripToolID                      = "echo-tool"
	googleToolsRoundTripToolName                    = "echo_text"
	googleToolsRoundTripToolDescription             = "Echo the provided text back in a deterministic tool result."
	googleToolsRoundTripJSONKeyType                 = "type"
	googleToolsRoundTripJSONValueObject             = "object"
	googleToolsRoundTripJSONKeyProperties           = "properties"
	googleToolsRoundTripJSONKeyText                 = "text"
	googleToolsRoundTripJSONValueString             = "string"
	googleToolsRoundTripJSONKeyRequired             = "required"
	googleToolsRoundTripJSONKeyAdditionalProperties = "additionalProperties"
)

// Example_googleGenerateContent_functionToolRoundTrip demonstrates a full
// Gemini function-tool round trip:
//
//  1. user message + tool definition
//  2. model emits a function tool call
//  3. caller executes the tool locally
//  4. next request sends:
//     - original user turn
//     - assistant function tool call
//     - user function tool output
//     - extra user text
//  5. model returns the final answer
func Example_googleGenerateContent_functionToolRoundTrip() {
	// Deadline greater than default timeout to simulate the sdk bug of context cancellation.
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	_, err = ps.AddProvider(ctx, googleToolsRoundTripProviderName, &inference.AddProviderConfig{
		SDKType: spec.ProviderSDKTypeGoogleGenerateContent,
		Origin:  spec.DefaultGoogleGenerateContentOrigin,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding Google GenAI provider:", err)
		return
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "GEMINI_API_KEY/GOOGLE_API_KEY not set; skipping live Google GenAI tool example")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, googleToolsRoundTripProviderName, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting Google GenAI API key:", err)
		return
	}

	tool := spec.ToolChoice{
		Type:        spec.ToolTypeFunction,
		ID:          googleToolsRoundTripToolID,
		Name:        googleToolsRoundTripToolName,
		Description: googleToolsRoundTripToolDescription,
		Arguments: map[string]any{
			googleToolsRoundTripJSONKeyType: googleToolsRoundTripJSONValueObject,
			googleToolsRoundTripJSONKeyProperties: map[string]any{
				googleToolsRoundTripJSONKeyText: map[string]any{
					googleToolsRoundTripJSONKeyType: googleToolsRoundTripJSONValueString,
				},
			},
			googleToolsRoundTripJSONKeyRequired:             []any{googleToolsRoundTripJSONKeyText},
			googleToolsRoundTripJSONKeyAdditionalProperties: false,
		},
	}

	initialUser := newUserTextInput(
		fmt.Sprintf(
			`this is a test. think about what 20 new words you can say and say that in text. then call the %s tool with that text.`,
			googleToolsRoundTripToolName,
		),
	)

	firstReq := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            googleToolsRoundTripModelName,
			MaxOutputLength: 512,
			Stream:          true,
			Reasoning: &spec.ReasoningParam{
				Type:   spec.ReasoningTypeHybridWithTokens,
				Tokens: 256,
			},
			SystemPrompt: "You are validating a Gemini function tool round trip. " +
				"When the tool is forced, emit only the tool call in the first response.",
		},
		Inputs:      []spec.InputUnion{initialUser},
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeTool,
			AllowedTools: []spec.AllowedTool{
				{ToolChoiceID: tool.ID},
			},
			DisableParallel: true,
		},
	}

	firstResp, err := ps.FetchCompletion(ctx, googleToolsRoundTripProviderName, firstReq, &spec.FetchCompletionOptions{
		CompletionKey: googleToolsRoundTripModelName,
		StreamHandler: func(ev spec.StreamEvent) error {
			switch ev.Kind {
			case spec.StreamContentKindThinking:
				if ev.Thinking != nil {
					fmt.Fprintf(os.Stderr, "\n\n#######[thinking 1] %s\n", ev.Thinking.Text)
				}
			case spec.StreamContentKindText:
				if ev.Text != nil {
					fmt.Fprintf(os.Stderr, "\n\n#######[text 1] %s\n", ev.Text.Text)
				}
			}
			return nil
		},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "first FetchCompletion error:", err)
		return
	}

	firstText := responseText(firstResp)
	if firstText != "" {
		fmt.Fprintln(os.Stderr, "FIRST assistant text:", firstText)
	}

	call, err := firstFunctionToolCall(firstResp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "expected a function tool call, got error:", err)
		return
	}
	fmt.Fprintf(
		os.Stderr,
		"\ntool call 1: name=%s id=%s args=%s\n",
		call.Name,
		nonEmpty(call.CallID, call.ID),
		call.Arguments,
	)

	toolOutput, err := runEchoTool(call)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error executing local tool:", err)
		return
	}
	fmt.Fprintf(os.Stderr, "\ntool result 1 for %s: %s\n", toolOutput.CallID, firstToolOutputText(toolOutput))

	history := make([]spec.InputUnion, 0, len(firstResp.Outputs)+2)
	history = append(history, initialUser)
	history = append(history, outputUnionsToInputs(firstResp.Outputs)...)
	history = append(history, spec.InputUnion{
		Kind:               spec.InputKindFunctionToolOutput,
		FunctionToolOutput: toolOutput,
	})

	secondReq := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name:            googleToolsRoundTripModelName,
			MaxOutputLength: 256,
			Stream:          true,
			Reasoning: &spec.ReasoningParam{
				Type:   spec.ReasoningTypeHybridWithTokens,
				Tokens: 256,
			},
			SystemPrompt: "You have now received the tool result. " +
				"Answer with a sonnet of it. Do not call any tool again.",
		},
		Inputs:      append(history, newUserTextInput("Now finish in one short sentence.")),
		ToolChoices: []spec.ToolChoice{tool},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeNone,
		},
	}

	secondResp, err := ps.FetchCompletion(
		ctx,
		googleToolsRoundTripProviderName,
		secondReq,
		&spec.FetchCompletionOptions{
			CompletionKey: googleToolsRoundTripModelName,
			StreamHandler: func(ev spec.StreamEvent) error {
				switch ev.Kind {
				case spec.StreamContentKindThinking:
					if ev.Thinking != nil {
						fmt.Fprintf(os.Stderr, "\n\n#######[thinking 2] %s\n", ev.Thinking.Text)
					}
				case spec.StreamContentKindText:
					if ev.Text != nil {
						fmt.Fprintf(os.Stderr, "\n\n#######[text 2] %s\n", ev.Text.Text)
					}
				}
				return nil
			},
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "second FetchCompletion error:", err)
		return
	}

	finalText := responseText(secondResp)
	if finalText != "" {
		fmt.Fprintln(os.Stderr, "\nFINAL assistant text:", finalText)
	}

	fmt.Println("OK")
	// Output: OK
}
