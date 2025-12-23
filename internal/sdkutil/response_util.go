package sdkutil

import (
	"context"
	"log/slog"
	"runtime/debug"
	"strings"

	"github.com/ppipada/inference-go/internal/debugclient"
	"github.com/ppipada/inference-go/internal/maputil"
	"github.com/ppipada/inference-go/spec"
)

// AttachDebugResp adds HTTP-debug information and error context—without panics.
//
// - ctx may or may not contain debug information.
// - respErr is the transport/SDK error (may be nil).
// - isNilResp tells whether the model returned an empty/invalid response.
// - rawModelJSON is an optional, provider-level JSON representation of the *final* model response (e.g. OpenAI
// responses `resp.RawJSON()` or `json.Marshal(fullResponse)` for other SDKs). If provided and the HTTP debug layer
// did not already set ResponseDetails.Data, we will sanitize and store this JSON there.
func AttachDebugResp(
	ctx context.Context,
	completionResp *spec.FetchCompletionResponse,
	respErr error,
	isNilResp bool,
	fullObj any,
) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("attach debug resp panic",
				"recover", r,
				"stack", string(debug.Stack()))
		}
	}()

	if completionResp == nil {
		return
	}

	debugDetails := map[string]any{
		"requestDetails":  map[string]any{},
		"responseDetails": map[string]any{},
		"errorDetails":    map[string]any{},
	}
	completionResp.DebugDetails = debugDetails

	debugResp, _ := debugclient.GetDebugHTTPResponse(ctx)

	// Always attach request/response debug info from the HTTP layer if available.
	if debugResp != nil {
		if debugResp.RequestDetails != nil {
			if d, err := maputil.StructWithJSONTagsToMap(debugResp.RequestDetails); err == nil {
				debugDetails["requestDetails"] = d
			}
		}
		if debugResp.ResponseDetails != nil {
			if d, err := maputil.StructWithJSONTagsToMap(debugResp.ResponseDetails); err == nil {
				debugDetails["responseDetails"] = d
			}
		}
	}

	// If the HTTP layer didn't populate ResponseDetails.Data (most common in
	// streaming/SSE cases), and we have a provider-level raw JSON for the final
	// model response, sanitize that and use it as the debug body.

	if fullObj != nil {
		// We got a object. Lets replace always.
		if m, err := maputil.StructWithJSONTagsToMap(fullObj); err == nil {
			if d, ok := debugDetails["responseDetails"].(map[string]any); ok {
				d["data"] = debugclient.ScrubAnyForDebug(m, true)
			}
		}
	}

	// Gather error-message fragments.
	var msgParts []string
	if debugResp != nil && debugResp.ErrorDetails != nil {
		if m := strings.TrimSpace(debugResp.ErrorDetails.Message); m != "" {
			msgParts = append(msgParts, m)
		}
	}
	if respErr != nil {
		msgParts = append(msgParts, respErr.Error())
	}
	if isNilResp {
		msgParts = append(msgParts, "got nil response from LLM api")
	}

	if len(msgParts) == 0 {
		// Nothing more to add; request/response details (if any) are already attached.

		return
	}

	// Prepare ErrorDetails without aliasing the debug struct pointer.
	if debugResp != nil && debugResp.ErrorDetails != nil {
		ed := *debugResp.ErrorDetails
		ed.Message = strings.Join(msgParts, "; ")

		if d, err := maputil.StructWithJSONTagsToMap(ed); err == nil {
			debugDetails["errorDetails"] = d
		}

	} else {
		if d, ok := debugDetails["errorDetails"].(map[string]any); ok {
			d["message"] = strings.Join(msgParts, "; ")
		}
	}
}
