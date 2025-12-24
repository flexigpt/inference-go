//go:build !integration

package integration

import (
	"log/slog"
	"os"

	"github.com/ppipada/inference-go"
	"github.com/ppipada/inference-go/debugclient"
	"github.com/ppipada/inference-go/spec"
)

// newProviderSetWithDebug constructs a ProviderSetAPI with:
//
//   - a text slog.Logger writing to stdout at debug level
//   - an HTTPCompletionDebugger that logs HTTP request/response metadata
//
// The examples reuse this helper to keep them short.
func newProviderSetWithDebug() (*inference.ProviderSetAPI, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return inference.NewProviderSetAPI(
		inference.WithLogger(logger),
		inference.WithDebugClientBuilder(func(p spec.ProviderParam) spec.CompletionDebugger {
			cfg := &debugclient.DebugConfig{
				// Capture and log request/response metadata (headers, URLs,
				// status codes, etc.). Bodies are scrubbed by default.
				LogToSlog: true,
			}
			return debugclient.NewHTTPCompletionDebugger(cfg)
		}),
	)
}
