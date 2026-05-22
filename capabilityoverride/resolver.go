package capabilityoverride

import (
	"context"
	"errors"
	"fmt"

	"github.com/flexigpt/inference-go/spec"
)

type CompletionKeyResolver struct {
	CompletionKey string
	Capabilities  *spec.ModelCapabilities
}

func NewCompletionKeyResolver(
	completionKey string,
	caps *spec.ModelCapabilities,
) CompletionKeyResolver {
	if caps == nil {
		return CompletionKeyResolver{
			CompletionKey: completionKey,
		}
	}

	cloned := CloneModelCapabilities(*caps)

	return CompletionKeyResolver{
		CompletionKey: completionKey,
		Capabilities:  &cloned,
	}
}

func (r CompletionKeyResolver) ResolveModelCapabilities(
	ctx context.Context,
	req spec.ResolveModelCapabilitiesRequest,
) (*spec.ModelCapabilities, error) {
	if r.Capabilities == nil {
		return nil, errors.New("no model capabilities configured")
	}

	if r.CompletionKey != "" && req.CompletionKey != r.CompletionKey {
		return nil, fmt.Errorf("capabilities not found for completionKey %q", req.CompletionKey)
	}

	out := CloneModelCapabilities(*r.Capabilities)
	return &out, nil
}
