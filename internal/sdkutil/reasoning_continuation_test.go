package sdkutil

import (
	"reflect"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestBuildReasoningContinuationFingerprint(t *testing.T) {
	base := spec.ProviderParam{
		Name:                     "openai-primary",
		SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
		Origin:                   "https://gateway.example/",
		ChatCompletionPathPrefix: "/v1/responses/",
		APIKey:                   "first-api-key",
		DefaultHeaders: map[string]string{
			"x-tenant": "tenant-a",
		},
	}

	baseFingerprint := BuildReasoningContinuationFingerprint(&base)
	if baseFingerprint == "" {
		t.Fatal("base fingerprint must not be empty")
	}

	tests := []struct {
		name     string
		mutate   func(*spec.ProviderParam)
		wantSame bool
	}{
		{
			name: "same endpoint with equivalent slash formatting",
			mutate: func(p *spec.ProviderParam) {
				p.Origin = " https://gateway.example "
				p.ChatCompletionPathPrefix = "v1/responses"
			},
			wantSame: true,
		},
		{
			name: "provider name does not affect continuation domain",
			mutate: func(p *spec.ProviderParam) {
				p.Name = "renamed-openai-provider"
			},
			wantSame: true,
		},
		{
			name: "API key does not affect continuation domain",
			mutate: func(p *spec.ProviderParam) {
				p.APIKey = "rotated-api-key"
			},
			wantSame: true,
		},
		{
			name: "default headers do not affect continuation domain",
			mutate: func(p *spec.ProviderParam) {
				p.DefaultHeaders = map[string]string{
					"x-tenant": "tenant-b",
				}
			},
			wantSame: true,
		},
		{
			name: "SDK type changes continuation domain",
			mutate: func(p *spec.ProviderParam) {
				p.SDKType = spec.ProviderSDKTypeAnthropic
			},
			wantSame: false,
		},
		{
			name: "origin changes continuation domain",
			mutate: func(p *spec.ProviderParam) {
				p.Origin = "https://other-gateway.example"
			},
			wantSame: false,
		},
		{
			name: "path prefix changes continuation domain",
			mutate: func(p *spec.ProviderParam) {
				p.ChatCompletionPathPrefix = "/v1/chat/completions"
			},
			wantSame: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProvider := base
			tt.mutate(&gotProvider)

			gotFingerprint := BuildReasoningContinuationFingerprint(&gotProvider)
			if gotFingerprint == "" {
				t.Fatal("fingerprint must not be empty")
			}

			gotSame := gotFingerprint == baseFingerprint
			if gotSame != tt.wantSame {
				t.Fatalf(
					"fingerprint equality = %v, want %v; got %q, base %q",
					gotSame,
					tt.wantSame,
					gotFingerprint,
					baseFingerprint,
				)
			}
		})
	}
}

func TestFilterReasoningInputsByContinuationFingerprint(t *testing.T) {
	const currentFingerprint = "current-fingerprint"

	tests := []struct {
		name               string
		currentFingerprint string
		inputs             []spec.InputUnion
		wantIDs            []string
	}{
		{
			name:               "no reasoning is unchanged",
			currentFingerprint: currentFingerprint,
			inputs: []spec.InputUnion{
				userInput("user-1"),
				assistantOutput("assistant-1"),
			},
			wantIDs: []string{"user-1", "assistant-1"},
		},
		{
			name:               "matching reasoning is retained",
			currentFingerprint: currentFingerprint,
			inputs: []spec.InputUnion{
				userInput("user-1"),
				reasoningInput("reasoning-1", currentFingerprint),
				assistantOutput("assistant-1"),
			},
			wantIDs: []string{"user-1", "reasoning-1", "assistant-1"},
		},
		{
			name:               "missing fingerprint is dropped",
			currentFingerprint: currentFingerprint,
			inputs: []spec.InputUnion{
				userInput("user-1"),
				reasoningInput("reasoning-unknown", ""),
				assistantOutput("assistant-1"),
			},
			wantIDs: []string{"user-1", "assistant-1"},
		},
		{
			name:               "foreign fingerprint is dropped",
			currentFingerprint: currentFingerprint,
			inputs: []spec.InputUnion{
				userInput("user-1"),
				reasoningInput("reasoning-foreign", "other-fingerprint"),
				assistantOutput("assistant-1"),
			},
			wantIDs: []string{"user-1", "assistant-1"},
		},
		{
			name:               "matching and foreign reasoning can be interleaved",
			currentFingerprint: currentFingerprint,
			inputs: []spec.InputUnion{
				userInput("user-1"),
				reasoningInput("reasoning-keep-1", currentFingerprint),
				assistantOutput("assistant-1"),
				reasoningInput("reasoning-missing", ""),
				reasoningInput("reasoning-foreign", "other-fingerprint"),
				reasoningInput("reasoning-keep-2", currentFingerprint),
				userInput("user-2"),
			},
			wantIDs: []string{
				"user-1",
				"reasoning-keep-1",
				"assistant-1",
				"reasoning-keep-2",
				"user-2",
			},
		},
		{
			name:               "source fingerprint whitespace is ignored",
			currentFingerprint: currentFingerprint,
			inputs: []spec.InputUnion{
				reasoningInput("reasoning-1", " current-fingerprint "),
			},
			wantIDs: []string{"reasoning-1"},
		},
		{
			name:               "empty current fingerprint fails closed",
			currentFingerprint: "",
			inputs: []spec.InputUnion{
				userInput("user-1"),
				reasoningInput("reasoning-1", ""),
			},
			wantIDs: []string{"user-1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterReasoningInputsByContinuationFingerprint(
				tt.inputs,
				tt.currentFingerprint,
			)

			if gotIDs := inputIDs(got); !reflect.DeepEqual(gotIDs, tt.wantIDs) {
				t.Fatalf("input IDs = %#v, want %#v", gotIDs, tt.wantIDs)
			}
		})
	}
}

func TestStampReasoningContinuationFingerprint(t *testing.T) {
	tests := []struct {
		name        string
		fingerprint string
		response    *spec.FetchCompletionResponse
		want        []string
	}{
		{
			name:        "nil response is ignored",
			fingerprint: "current-fingerprint",
			response:    nil,
			want:        nil,
		},
		{
			name:        "only reasoning outputs are stamped",
			fingerprint: "current-fingerprint",
			response: &spec.FetchCompletionResponse{
				Outputs: []spec.OutputUnion{
					{
						Kind: spec.OutputKindOutputMessage,
						OutputMessage: &spec.InputOutputContent{
							ID: "assistant-1",
						},
					},
					{
						Kind: spec.OutputKindReasoningMessage,
						ReasoningMessage: &spec.ReasoningContent{
							ID:                      "reasoning-1",
							ContinuationFingerprint: "old-fingerprint",
						},
					},
					{
						Kind: spec.OutputKindReasoningMessage,
						ReasoningMessage: &spec.ReasoningContent{
							ID: "reasoning-2",
						},
					},
				},
			},
			want: []string{
				"reasoning-1=current-fingerprint",
				"reasoning-2=current-fingerprint",
			},
		},
		{
			name:        "empty fingerprint does not overwrite outputs",
			fingerprint: "",
			response: &spec.FetchCompletionResponse{
				Outputs: []spec.OutputUnion{{
					Kind: spec.OutputKindReasoningMessage,
					ReasoningMessage: &spec.ReasoningContent{
						ID:                      "reasoning-1",
						ContinuationFingerprint: "old-fingerprint",
					},
				}},
			},
			want: []string{"reasoning-1=old-fingerprint"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StampReasoningContinuationFingerprint(tt.response, tt.fingerprint)

			if got := outputReasoningFingerprints(tt.response); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("reasoning fingerprints = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func userInput(id string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindInputMessage,
		InputMessage: &spec.InputOutputContent{
			ID:   id,
			Role: spec.RoleUser,
		},
	}
}

func assistantOutput(id string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindOutputMessage,
		OutputMessage: &spec.InputOutputContent{
			ID:   id,
			Role: spec.RoleAssistant,
		},
	}
}

func reasoningInput(id, fingerprint string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindReasoningMessage,
		ReasoningMessage: &spec.ReasoningContent{
			ID:                      id,
			Role:                    spec.RoleAssistant,
			EncryptedContent:        []string{"opaque-reasoning-data"},
			ContinuationFingerprint: fingerprint,
		},
	}
}

func inputIDs(inputs []spec.InputUnion) []string {
	ids := make([]string, 0, len(inputs))

	for _, input := range inputs {
		switch input.Kind {
		case spec.InputKindInputMessage:
			if input.InputMessage != nil {
				ids = append(ids, input.InputMessage.ID)
			}
		case spec.InputKindOutputMessage:
			if input.OutputMessage != nil {
				ids = append(ids, input.OutputMessage.ID)
			}
		case spec.InputKindReasoningMessage:
			if input.ReasoningMessage != nil {
				ids = append(ids, input.ReasoningMessage.ID)
			}
		default:
		}
	}

	return ids
}

func outputReasoningFingerprints(
	response *spec.FetchCompletionResponse,
) []string {
	if response == nil {
		return nil
	}

	var out []string
	for _, output := range response.Outputs {
		if output.Kind != spec.OutputKindReasoningMessage ||
			output.ReasoningMessage == nil {
			continue
		}

		out = append(
			out,
			output.ReasoningMessage.ID+"="+output.ReasoningMessage.ContinuationFingerprint,
		)
	}

	return out
}
