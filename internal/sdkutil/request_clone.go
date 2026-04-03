package sdkutil

import (
	"encoding/json"
	"fmt"

	"github.com/flexigpt/inference-go/spec"
)

func CloneFetchCompletionRequest(req *spec.FetchCompletionRequest) (*spec.FetchCompletionRequest, error) {
	if req == nil {
		//nolint:nilnil // Clone of nil is nil for req.
		return nil, nil
	}

	raw, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("clone fetch completion request: marshal: %w", err)
	}

	var out spec.FetchCompletionRequest
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("clone fetch completion request: unmarshal: %w", err)
	}

	return &out, nil
}
