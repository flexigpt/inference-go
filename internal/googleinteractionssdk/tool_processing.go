package googleinteractionssdk

import (
	"errors"
	"fmt"
	"maps"

	"google.golang.org/genai/interactions/models/interactions"

	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

//nolint:gocritic // Too many results are ok for this.
func buildGoogleInteractionsTools(
	toolChoices []spec.ToolChoice,
	policy *spec.ToolPolicy,
	config *interactions.GenerationConfig,
) (
	[]interactions.Tool,
	map[string]spec.ToolChoice,
	string,
	bool,
	[]spec.Warning,
	error,
) {
	if policy != nil && policy.Mode == spec.ToolPolicyModeNone {
		return nil, nil, "", false, nil, nil
	}
	if len(toolChoices) == 0 {
		return nil, nil, "", false, nil, nil
	}

	ordered, allNames := sdkutil.BuildToolChoiceNameMapping(toolChoices)
	var warnings []spec.Warning

	if policy != nil && policy.DisableParallel {
		warnings = appendGoogleInteractionsWarning(
			warnings,
			"toolPolicy_disableParallel_unsupported",
			"Interactions does not expose a single-tool-call control; DisableParallel was not enforced.",
		)
	}

	forcedCallable := policy != nil &&
		(policy.Mode == spec.ToolPolicyModeAny ||
			policy.Mode == spec.ToolPolicyModeTool)

	tools := make([]interactions.Tool, 0, len(ordered))
	nameMap := make(map[string]spec.ToolChoice, len(ordered))
	var webSearchChoiceID string
	var callableNames []string

	for _, named := range ordered {
		choice := named.Choice
		switch choice.Type {
		case spec.ToolTypeFunction, spec.ToolTypeCustom:
			schema := choice.Arguments
			if schema == nil {
				schema = sdkutil.EmptyJSONArgs
			}

			function := interactions.Function{
				Name:        new(named.Name),
				Description: new(sdkutil.ToolDescription(choice)),
				Parameters:  maps.Clone(schema),
			}
			tools = append(tools, interactions.NewTool(function))
			nameMap[named.Name] = choice
			callableNames = append(callableNames, named.Name)

		case spec.ToolTypeWebSearch:
			if forcedCallable || webSearchChoiceID != "" {
				continue
			}

			tools = append(tools, interactions.NewTool(
				interactions.GoogleSearch{},
			))
			webSearchChoiceID = choice.ID

			if choice.WebSearchArguments != nil {
				warnings = appendGoogleInteractionsWarning(
					warnings,
					"webSearch_options_not_forwarded",
					"Interactions enables Google Search but does not map normalized search filters, location, context size, or limits.",
				)
			}
		}
	}

	if policy == nil {
		return tools, nameMap, webSearchChoiceID, false, warnings, nil
	}
	if config == nil {
		return nil, nil, "", false, warnings, errors.New(
			"google interactions: nil generation config for tool policy",
		)
	}

	switch policy.Mode {
	case spec.ToolPolicyModeAuto:
		return tools, nameMap, webSearchChoiceID, false, warnings, nil

	case spec.ToolPolicyModeAny:
		if len(callableNames) == 0 {
			return nil, nil, "", false, warnings, errors.New(
				"google interactions: toolPolicy=any requires a function or custom tool; Google Search cannot be forced",
			)
		}

		allowedNames, err := googleInteractionsAllowedFunctionNames(
			policy.AllowedTools,
			allNames,
			callableNames,
		)
		if err != nil {
			return nil, nil, "", false, warnings, err
		}
		config.ToolChoice = googleInteractionsToolChoice(
			interactions.ToolChoiceTypeAny,
			allowedNames,
		)
		return tools, nameMap, "", true, warnings, nil

	case spec.ToolPolicyModeTool:
		allowedNames, err := googleInteractionsAllowedFunctionNames(
			policy.AllowedTools,
			allNames,
			nil,
		)
		if err != nil {
			return nil, nil, "", false, warnings, err
		}
		if len(allowedNames) != 1 {
			return nil, nil, "", false, warnings, errors.New(
				"google interactions: toolPolicy=tool requires exactly one function or custom tool",
			)
		}

		config.ToolChoice = googleInteractionsToolChoice(
			interactions.ToolChoiceTypeAny,
			allowedNames,
		)
		return tools, nameMap, "", true, warnings, nil

	default:
		return nil, nil, "", false, warnings, fmt.Errorf(
			"google interactions: unsupported tool policy mode %q",
			policy.Mode,
		)
	}
}

func googleInteractionsAllowedFunctionNames(
	allowed []spec.AllowedTool,
	allNames map[string]spec.ToolChoice,
	defaultNames []string,
) ([]string, error) {
	if len(allowed) == 0 {
		if len(defaultNames) == 0 {
			return nil, errors.New(
				"google interactions: forced tool policy requires allowed callable tools",
			)
		}
		return append([]string(nil), defaultNames...), nil
	}

	resolved, err := sdkutil.ResolveAllowedTools(allowed, allNames)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(resolved))
	for _, tool := range resolved {
		switch tool.Type {
		case spec.ToolTypeFunction, spec.ToolTypeCustom:
			out = append(out, tool.Name)
		default:
		}
	}

	if len(out) == 0 {
		return nil, errors.New(
			"google interactions: forced policy resolved no callable function/custom tools",
		)
	}
	return out, nil
}

func googleInteractionsToolChoice(
	mode interactions.ToolChoiceType,
	allowedNames []string,
) *interactions.ToolChoice {
	if len(allowedNames) == 0 {
		choice := interactions.NewToolChoice(mode)
		return &choice
	}

	choice := interactions.NewToolChoice(
		interactions.ToolChoiceConfig{
			AllowedTools: &interactions.AllowedTools{
				Mode:  interactions.ToolChoiceTypeAny.ToPointer(),
				Tools: allowedNames,
			},
		},
	)
	return &choice
}
