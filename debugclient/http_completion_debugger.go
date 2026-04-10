package debugclient

import (
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

type httpSpan struct {
	cfg  DebugConfig
	ctx  context.Context
	info *spec.CompletionSpanStart
}

// End implements spec.CompletionSpan.End.
func (s *httpSpan) End(end *spec.CompletionSpanEnd) any {
	defer sdkutil.Recover("debugclient.httpSpan.End panic")

	if s.cfg.Disable {
		return nil
	}

	state, _ := httpDebugStateFromContext(s.ctx)
	if state == nil {
		state = &HTTPDebugState{}
	}

	// Attach scrubbed provider response, if any.
	if end.ProviderResponse != nil {
		if m, err := structToMap(end.ProviderResponse); err == nil {
			strip := !s.cfg.DisableContentStripping
			state.ProviderResponse = scrubAnyForDebug(m, strip)
			// Dont send response details data if actual response was available.
			if state.ResponseDetails != nil {
				state.ResponseDetails.Data = nil
			}
		}
	}

	// Compose error message from HTTP-level error + provider error.
	var msgParts []string
	if state.ErrorDetails != nil {
		if msg := strings.TrimSpace(state.ErrorDetails.Message); msg != "" {
			msgParts = append(msgParts, msg)
		}
	}
	if end.Err != nil {
		msgParts = append(msgParts, end.Err.Error())
	}

	if len(msgParts) > 0 {
		if state.ErrorDetails == nil {
			state.ErrorDetails = &APIErrorDetails{}
		}
		state.ErrorDetails.Message = strings.Join(msgParts, "; ")
	}

	return state
}

// HTTPCompletionDebugger implements spec.CompletionDebugger using the HTTP
// instrumentation in this package. It produces an opaque debug payload
// (HTTPDebugState) suitable for attachment to FetchCompletionResponse.DebugDetails.
type HTTPCompletionDebugger struct {
	config DebugConfig
	mu     sync.RWMutex
}

// NewHTTPCompletionDebugger constructs a CompletionDebugger that instruments
// HTTP traffic using an internal RoundTripper and produces a scrubbed debug
// blob from HTTP-level data and the raw provider response.
//
// Config may be nil; in that case DebugConfig{} (defaults) is used.
func NewHTTPCompletionDebugger(config *DebugConfig) *HTTPCompletionDebugger {
	var c DebugConfig
	if config != nil {
		c = *config
	}
	return &HTTPCompletionDebugger{config: c}
}

// HTTPClient implements spec.CompletionDebugger.HTTPClient.
func (d *HTTPCompletionDebugger) HTTPClient(base *http.Client) *http.Client {
	if base == nil {
		base = &http.Client{Transport: http.DefaultTransport}
	}
	rt := base.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}

	clone := *base
	clone.Transport = &logTransport{
		base: rt,
	}
	return &clone
}

// StartSpan implements spec.CompletionDebugger.StartSpan.
func (d *HTTPCompletionDebugger) StartSpan(
	ctx context.Context,
	info *spec.CompletionSpanStart,
) (context.Context, spec.CompletionSpan) {
	cfg := d.GetConfig()
	if cfg.Disable {
		return ctx, nil
	}

	ctx = withHTTPDebugState(ctx)
	ctx = withHTTPDebugConfig(ctx, cfg)
	span := &httpSpan{
		cfg:  cfg,
		ctx:  ctx,
		info: info,
	}
	return ctx, span
}

func (d *HTTPCompletionDebugger) SetConfig(cfg DebugConfig) {
	d.mu.Lock()
	d.config = cfg
	d.mu.Unlock()
}

func (d *HTTPCompletionDebugger) GetConfig() DebugConfig {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.config
}
