package anthropicsdk

import (
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const anthropicDefaultThinkingBudget int64 = 1024

type thinkingOverride int

const (
	thinkingOverrideNone thinkingOverride = iota
	thinkingOverrideForceEnabled
	thinkingOverrideForceDisabled
)

func (o thinkingOverride) String() string {
	switch o {
	case thinkingOverrideForceEnabled:
		return "forceEnabled"
	case thinkingOverrideForceDisabled:
		return "forceDisabled"
	default:
		return "none"
	}
}

type anthropicThinkingAnalysis struct {
	Override                    thinkingOverride
	TotalReasoningMessages      int
	SignedOrRedactedReasoning   int
	UnsignedReasoning           int
	LastUserIsToolResult        bool
	PrevAssistantStartsThinking bool
}

// analyzeAnthropicThinkingBehavior enforces the policy:
//   - No reasoning messages: if last user msg is tool_result => force thinking disabled, else honor requested thinking.
//   - Mixed signed+unsigned reasoning => keep signed only (handled by conversion); no override here.
//   - All signed/redacted: force thinking enabled, if last user msg is tool_result and previous "turn" is thinking.
//
// Additionally, we treat "signed/redacted thinking present in input" as a fail-safe requirement:
// if we will send a ThinkingBlock/RedactedThinkingBlock, we ensure thinking is enabled unless explicitly forced off.
func analyzeAnthropicThinkingBehavior(inputs []spec.InputUnion) anthropicThinkingAnalysis {
	var a anthropicThinkingAnalysis
	if len(inputs) == 0 {
		return a
	}

	// Count reasoning messages and classify which ones are usable for Anthropic.
	for _, in := range inputs {
		if in.Kind != spec.InputKindReasoningMessage {
			continue
		}
		if sdkutil.IsInputUnionEmpty(in) || in.ReasoningMessage == nil {
			continue
		}
		a.TotalReasoningMessages++
		if isAnthropicSignedOrRedactedReasoning(in.ReasoningMessage) {
			a.SignedOrRedactedReasoning++
		} else {
			a.UnsignedReasoning++
		}
	}

	lastUserIdx, lastUserIsToolResult := findLastUserMessageIndex(inputs)
	a.LastUserIsToolResult = lastUserIsToolResult
	if lastUserIsToolResult && lastUserIdx >= 0 {
		a.PrevAssistantStartsThinking = prevAssistantTurnStartsWithThinking(inputs, lastUserIdx)
	}

	// Policy overrides.
	switch {
	case a.TotalReasoningMessages == 0:
		// No reasoning messages anywhere.
		if a.LastUserIsToolResult {
			a.Override = thinkingOverrideForceDisabled
		}

	case a.SignedOrRedactedReasoning > 0 && a.UnsignedReasoning == 0:
		// All reasoning is signed/redacted.
		if a.LastUserIsToolResult && a.PrevAssistantStartsThinking {
			a.Override = thinkingOverrideForceEnabled
		}

	default:
		// Mixed signed/unsigned, or all-unsigned: message conversion already drops
		// the unsupported ones. We don't force an override here.
	}

	if a.Override != thinkingOverrideNone {
		logutil.Debug(
			"anthropic: thinking override applied",
			"override", a.Override.String(),
			"reasoningTotal", a.TotalReasoningMessages,
			"reasoningSigned", a.SignedOrRedactedReasoning,
			"reasoningUnsigned", a.UnsignedReasoning,
			"lastUserIsToolResult", a.LastUserIsToolResult,
			"prevAssistantStartsThinking", a.PrevAssistantStartsThinking,
		)
	}

	return a
}

func isAnthropicSignedOrRedactedReasoning(r *spec.ReasoningContent) bool {
	if r == nil {
		return false
	}
	// Redacted thinking is always acceptable.
	for _, s := range r.RedactedThinking {
		if strings.TrimSpace(s) != "" {
			return true
		}
	}
	// Signed thinking requires both signature and non-empty thinking.
	if strings.TrimSpace(r.Signature) == "" {
		return false
	}
	for _, t := range r.Thinking {
		if strings.TrimSpace(t) != "" {
			return true
		}
	}
	return false
}

// findLastUserMessageIndex finds the index of the last user-authored item in the
// interleaved input list (user InputMessage or function/custom tool output).
// It returns (idx, isToolResult).
func findLastUserMessageIndex(inputs []spec.InputUnion) (int, bool) {
	for i := len(inputs) - 1; i >= 0; i-- {
		in := inputs[i]
		if sdkutil.IsInputUnionEmpty(in) {
			continue
		}
		switch in.Kind {
		case spec.InputKindInputMessage:
			if in.InputMessage != nil && in.InputMessage.Role == spec.RoleUser {
				return i, false
			}
		case spec.InputKindFunctionToolOutput:
			if in.FunctionToolOutput != nil {
				return i, true
			}
		case spec.InputKindCustomToolOutput:
			if in.CustomToolOutput != nil {
				return i, true
			}
		default:
			// Not user-authored; keep scanning.
		}
	}
	return -1, false
}

func isUserAuthoredItem(in spec.InputUnion) bool {
	if sdkutil.IsInputUnionEmpty(in) {
		return false
	}
	switch in.Kind {
	case spec.InputKindInputMessage:
		return in.InputMessage != nil && in.InputMessage.Role == spec.RoleUser
	case spec.InputKindFunctionToolOutput:
		return in.FunctionToolOutput != nil
	case spec.InputKindCustomToolOutput:
		return in.CustomToolOutput != nil
	default:
		return false
	}
}

func isAssistantAuthoredItem(in spec.InputUnion) bool {
	if sdkutil.IsInputUnionEmpty(in) {
		return false
	}
	switch in.Kind {
	case spec.InputKindOutputMessage:
		return in.OutputMessage != nil && in.OutputMessage.Role == spec.RoleAssistant
	case spec.InputKindReasoningMessage:
		return in.ReasoningMessage != nil
	case spec.InputKindFunctionToolCall, spec.InputKindCustomToolCall, spec.InputKindWebSearchToolCall:
		return true
	case spec.InputKindWebSearchToolOutput:
		// In our Anthropic adapter, web_search_tool_result is an assistant block.
		return in.WebSearchToolOutput != nil
	default:
		return false
	}
}

// prevAssistantTurnStartsWithThinking checks, for the assistant "turn" immediately
// preceding the given tool_result index, whether the first assistant-authored
// item after the previous user message is a signed/redacted reasoning message.
func prevAssistantTurnStartsWithThinking(inputs []spec.InputUnion, toolResultIdx int) bool {
	if toolResultIdx <= 0 || toolResultIdx > len(inputs)-1 {
		return false
	}

	// Find the user item *before* this tool_result.
	prevUserIdx := -1
	for j := toolResultIdx - 1; j >= 0; j-- {
		if isUserAuthoredItem(inputs[j]) {
			prevUserIdx = j
			break
		}
	}

	// Now find the first assistant-authored item between prevUserIdx and toolResultIdx.
	for k := prevUserIdx + 1; k < toolResultIdx; k++ {
		in := inputs[k]
		if !isAssistantAuthoredItem(in) {
			continue
		}
		// "Starts with thinking" means the first assistant item is signed/redacted reasoning.
		if in.Kind == spec.InputKindReasoningMessage && isAnthropicSignedOrRedactedReasoning(in.ReasoningMessage) {
			return true
		}
		return false
	}
	return false
}

func applyAnthropicThinkingPolicy(
	params *anthropic.MessageNewParams,
	mp *spec.ModelParam,
	a anthropicThinkingAnalysis,
) {
	if params == nil || mp == nil {
		return
	}

	// Derive the requested thinking config from ModelParam.Reasoning.
	requestedEnabled, requestedAdaptive, requestedBudget := requestedAnthropicThinking(mp)
	effectiveEnabled := requestedEnabled
	effectiveAdaptive := requestedAdaptive
	effectiveBudget := requestedBudget

	switch a.Override {
	case thinkingOverrideForceDisabled:
		effectiveEnabled = false
		effectiveAdaptive = false
		effectiveBudget = 0

	case thinkingOverrideForceEnabled:
		effectiveEnabled = true
		// Keep adaptive iff it was explicitly requested by reasoning.type=singleWithLevels.
		// If we are forcing thinking on (i.e. not user-requested), we do NOT invent adaptive mode.
		if effectiveBudget <= 0 {
			effectiveBudget = anthropicDefaultThinkingBudget
		}

	default:
		// Ok.
	}

	// Fail-safe: if we're going to send signed/redacted thinking blocks as part of the prompt,
	// ensure thinking is enabled (unless explicitly forced off).
	if a.Override != thinkingOverrideForceDisabled && !effectiveEnabled && a.SignedOrRedactedReasoning > 0 {
		logutil.Warn(
			"anthropic: signed/redacted reasoning present in input but thinking is disabled; enabling thinking as a fail-safe",
			"provider",
			"anthropic",
			"model",
			string(mp.Name),
		)
		effectiveEnabled = true
		// This is NOT a "reasoning.type=singleWithLevels" request; it's a compatibility fail-safe.
		// Use non-adaptive enabled thinking.
		// This may need to change when messages api explicitly switches to adaptive as default.
		effectiveAdaptive = false
		effectiveBudget = anthropicDefaultThinkingBudget
	}

	// Apply "effort" mapping ONLY when we are actually going to run adaptive thinking,
	//    and ONLY when caller did not explicitly provide outputParam.verbosity.
	//    (If caller provided verbosity, applyAnthropicOutputParam will set OutputConfig.Effort later.)
	if effectiveEnabled && effectiveAdaptive {
		maybeApplyAnthropicEffortFromReasoning(params, mp)
	}

	// Materialize final SDK params.
	if effectiveEnabled {
		if effectiveBudget <= 0 {
			effectiveBudget = anthropicDefaultThinkingBudget
		}
		if effectiveAdaptive {
			// Reasoning.type=singleWithLevels -> adaptive thinking.
			a := anthropic.NewThinkingConfigAdaptiveParam()
			params.Thinking = anthropic.ThinkingConfigParamUnion{
				OfAdaptive: &a,
			}
		} else {
			// Existing behavior: fixed-budget enabled thinking.
			params.Thinking = anthropic.ThinkingConfigParamOfEnabled(effectiveBudget)
		}
		// Do not set temperature when thinking is enabled (Anthropic requirement).
		return
	}

	// Thinking disabled => temperature is allowed.
	if mp.Temperature != nil {
		params.Temperature = anthropic.Float(*mp.Temperature)
	}
}

func requestedAnthropicThinking(mp *spec.ModelParam) (enabled, adaptive bool, budget int64) {
	if mp == nil || mp.Reasoning == nil {
		return false, false, 0
	}
	rp := mp.Reasoning
	switch rp.Type {
	case spec.ReasoningTypeHybridWithTokens:
		// Enforce minimum budget if requested.
		return true, false, int64(max(rp.Tokens, int(anthropicDefaultThinkingBudget)))

	case spec.ReasoningTypeSingleWithLevels:
		// Map qualitative levels to token budgets; ignore rp.Tokens.
		switch rp.Level {
		case spec.ReasoningLevelNone:
			return false, false, 0
		case spec.ReasoningLevelMinimal, spec.ReasoningLevelLow:
			return true, true, 1024
		case spec.ReasoningLevelMedium:
			return true, true, 2048
		case spec.ReasoningLevelHigh:
			return true, true, 8192
		case spec.ReasoningLevelXHigh:
			return true, true, 16384
		default:
			// Unknown => treat as not requested.
			return false, false, 0
		}
	default:
		return false, false, 0
	}
}

// If reasoning.type=singleWithLevels and caller did NOT explicitly set outputParam.verbosity,
// derive output_config.effort from the reasoning level (Anthropic only supports: low|medium|high|max).
func maybeApplyAnthropicEffortFromReasoning(params *anthropic.MessageNewParams, mp *spec.ModelParam) {
	if params == nil || mp == nil || mp.Reasoning == nil {
		return
	}
	rp := mp.Reasoning
	if rp.Type != spec.ReasoningTypeSingleWithLevels {
		return
	}

	// Do not override explicitly provided effort/verbosity.
	if mp.OutputParam != nil && mp.OutputParam.Verbosity != nil {
		return
	}

	effort, ok := anthropicEffortFromReasoningLevel(rp.Level)
	if !ok {
		return
	}
	params.OutputConfig.Effort = effort
}

func anthropicEffortFromReasoningLevel(level spec.ReasoningLevel) (anthropic.OutputConfigEffort, bool) {
	switch level {
	case spec.ReasoningLevelMinimal, spec.ReasoningLevelLow:
		return anthropic.OutputConfigEffort("low"), true
	case spec.ReasoningLevelMedium:
		return anthropic.OutputConfigEffort("medium"), true
	case spec.ReasoningLevelHigh:
		return anthropic.OutputConfigEffort("high"), true
	case spec.ReasoningLevelXHigh:
		return anthropic.OutputConfigEffort("max"), true
	default:
		// None / Unknown.
		return "", false
	}
}
