package sdkutil

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/flexigpt/inference-go/spec"
)

const reasoningContinuationFingerprintVersion = "v1"

type continuationFingerprintInput struct {
	Version                  string `json:"v"`
	SDKType                  string `json:"s"`
	Origin                   string `json:"o"`
	ChatCompletionPathPrefix string `json:"c"`
}

// BuildReasoningContinuationFingerprint identifies the configured endpoint.
// Provider name, SDK type, model, API key, and CompletionKey are intentionally excluded.
func BuildReasoningContinuationFingerprint(
	providerInfo *spec.ProviderParam,
) string {
	if providerInfo == nil {
		return ""
	}
	pathPrefix := strings.Trim(
		strings.TrimSpace(providerInfo.ChatCompletionPathPrefix),
		"/",
	)
	if pathPrefix != "" {
		pathPrefix = "/" + pathPrefix
	}

	input := continuationFingerprintInput{
		Version:                  reasoningContinuationFingerprintVersion,
		SDKType:                  strings.TrimSpace(string(providerInfo.SDKType)),
		Origin:                   strings.TrimRight(strings.TrimSpace(providerInfo.Origin), "/"),
		ChatCompletionPathPrefix: pathPrefix,
	}
	raw, err := json.Marshal(input)
	if err != nil {
		return ""
	}

	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

// FilterReasoningInputsByContinuationFingerprint preserves ordinary history
// and keeps only reasoning produced by the current provider route and model.
func FilterReasoningInputsByContinuationFingerprint(
	inputs []spec.InputUnion,
	currentFingerprint string,
) []spec.InputUnion {
	if len(inputs) == 0 {
		return inputs
	}

	out := make([]spec.InputUnion, 0, len(inputs))
	for _, input := range inputs {
		if input.Kind != spec.InputKindReasoningMessage {
			out = append(out, input)
			continue
		}

		reasoning := input.ReasoningMessage
		if reasoning == nil ||
			currentFingerprint == "" ||
			strings.TrimSpace(reasoning.ContinuationFingerprint) != currentFingerprint {
			continue
		}

		out = append(out, input)
	}

	return out
}

// StampReasoningContinuationFingerprint records the source route on every
// normalized reasoning output before it is returned to the caller.
func StampReasoningContinuationFingerprint(
	response *spec.FetchCompletionResponse,
	fingerprint string,
) {
	if response == nil || fingerprint == "" {
		return
	}

	for i := range response.Outputs {
		output := &response.Outputs[i]
		if output.Kind != spec.OutputKindReasoningMessage ||
			output.ReasoningMessage == nil {
			continue
		}
		output.ReasoningMessage.ContinuationFingerprint = fingerprint
	}
}
