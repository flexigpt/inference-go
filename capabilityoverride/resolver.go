package capabilityoverride

import (
	"context"
	"errors"
	"fmt"

	"github.com/flexigpt/inference-go/spec"
)

type CompletionKeyResolver struct {
	completionKey string
	capabilities  *spec.ModelCapabilities
}

func NewCompletionKeyResolver(
	completionKey string,
	caps *spec.ModelCapabilities,
) CompletionKeyResolver {
	if caps == nil {
		return CompletionKeyResolver{
			completionKey: completionKey,
		}
	}

	cloned := CloneModelCapabilities(*caps)

	return CompletionKeyResolver{
		completionKey: completionKey,
		capabilities:  &cloned,
	}
}

func (r CompletionKeyResolver) ResolveModelCapabilities(
	ctx context.Context,
	req spec.ResolveModelCapabilitiesRequest,
) (*spec.ModelCapabilities, error) {
	if r.capabilities == nil {
		return nil, errors.New("no model capabilities configured")
	}

	if r.completionKey != "" && req.CompletionKey != r.completionKey {
		return nil, fmt.Errorf("capabilities not found for completionKey %q", req.CompletionKey)
	}

	out := CloneModelCapabilities(*r.capabilities)
	return &out, nil
}

func (r CompletionKeyResolver) CompletionKey() string {
	return r.completionKey
}

func (r CompletionKeyResolver) Capabilities() *spec.ModelCapabilities {
	if r.capabilities == nil {
		return nil
	}

	out := CloneModelCapabilities(*r.capabilities)
	return &out
}
