// Package integration contains example-style tests that demonstrate how to
// call the inference-go library against real providers.
//
// The package is internal so it does not become part of the public API
// surface, but you can run the examples locally with:
//
//	go test ./internal/integration -run Example
//
// All examples are best-effort: they only attempt live API calls when the
// relevant environment variables are set, e.g.:
//
//	ANTHROPIC_API_KEY=<your key>
//	OPENAI_API_KEY=<your key>
package integration
