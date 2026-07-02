package integration

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

const (
	openAIChatMathToolID          = "math"
	openAIChatMathToolName        = "multiply"
	openAIChatMathToolDescription = "Multiply two integers."
)

// Example_openAIChat_toolsAndJSONSchema demonstrates:
//
//   - catalog-based OpenAI Chat provider setup
//   - preset model defaults
//   - preset capability resolver
//   - streaming text
//   - JSON schema output
//   - function tools
func Example_openAIChat_toolsAndJSONSchema() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	ps, err := newProviderSetWithDebug(slog.LevelDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating ProviderSetAPI:", err)
		return
	}

	pp, mp, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderOpenAIChat,
		modelpreset.PresetGPT41,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error adding OpenAI Chat preset provider:", err)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENAI_API_KEY not set; skipping live OpenAI Chat call")
		fmt.Println("OK")
		return
	}
	if err := ps.SetProviderAPIKey(ctx, pp.Name, apiKey); err != nil {
		fmt.Fprintln(os.Stderr, "error setting OpenAI API key:", err)
		return
	}

	modelParam := mp.ModelParam
	modelParam.Stream = true
	modelParam.MaxOutputLength = min(modelParam.MaxOutputLength, 1024)
	modelParam.SystemPrompt = "Answer directly."
	modelParam.OutputParam = &spec.OutputParam{
		Format: &spec.OutputFormat{
			Kind: spec.OutputFormatKindJSONSchema,
			JSONSchemaParam: &spec.JSONSchemaParam{
				Name: "result",
				Schema: map[string]any{
					toolJSONKeyType: toolJSONValueObject,
					toolJSONKeyProperties: map[string]any{
						toolJSONSchemaName: map[string]any{toolJSONKeyType: toolJSONValueString},
					},
					toolJSONKeyRequired:             []any{toolJSONSchemaName},
					toolJSONKeyAdditionalProperties: false,
				},
				Strict: true,
			},
		},
	}

	tools := []spec.ToolChoice{
		{
			Type:        spec.ToolTypeFunction,
			ID:          openAIChatMathToolID,
			Name:        openAIChatMathToolName,
			Description: openAIChatMathToolDescription,
			Arguments: map[string]any{
				toolJSONKeyType: toolJSONValueObject,
				toolJSONKeyProperties: map[string]any{
					"a": map[string]any{toolJSONKeyType: "integer"},
					"b": map[string]any{toolJSONKeyType: "integer"},
				},
				toolJSONKeyRequired:             []any{"a", "b"},
				toolJSONKeyAdditionalProperties: false,
			},
		},
	}

	opts, err := presetFetchOptions(ctx, ps, pp, mp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating preset capability resolver:", err)
		return
	}
	opts.StreamHandler = func(ev spec.StreamEvent) error {
		if ev.Kind == spec.StreamContentKindText && ev.Text != nil {
			fmt.Fprint(os.Stderr, ev.Text.Text)
		}
		return nil
	}

	_, err = ps.FetchCompletion(ctx, pp.Name, &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{
			newUserTextInput("What is 6*7? Use the multiply tool if useful."),
		},
		ToolChoices: tools,
		ToolPolicy: &spec.ToolPolicy{
			Mode:            spec.ToolPolicyModeAuto,
			DisableParallel: true,
		},
	}, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nFetchCompletion error:", err)
		return
	}

	fmt.Println("OK")
	// Output: OK
}
