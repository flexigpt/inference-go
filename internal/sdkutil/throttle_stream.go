package sdkutil

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	FlushInterval  = 32 * time.Millisecond
	FlushChunkSize = 1024
)

// NewBufferedStreamerWithError buffers stream data while preserving callback
// errors from size-based, timer-based, and final flushes.
func NewBufferedStreamerWithError(
	onDataFlush func(string) error,
	flushInterval time.Duration,
	maxSize int,
) (write func(string) error, flush func() error) {
	if flushInterval <= 0 {
		flushInterval = FlushInterval
	}
	if maxSize <= 0 {
		maxSize = FlushChunkSize
	}

	var mu sync.Mutex
	var buf strings.Builder
	var firstErr error
	var closed bool

	ticker := time.NewTicker(flushInterval)
	done := make(chan struct{})
	stopped := make(chan struct{})

	// The caller must hold mu. Keeping mu locked while invoking the callback
	// prevents this streamer's timer and size-based flushes from racing.
	flushBufferLocked := func() error {
		if firstErr != nil {
			return firstErr
		}
		if buf.Len() == 0 {
			return nil
		}

		data := buf.String()
		buf.Reset()
		if err := onDataFlush(data); err != nil {
			firstErr = err
			return err
		}
		return nil
	}

	// Background goroutine time-based flush.
	go func() {
		defer Recover("buffered streamer background flush panic")
		defer close(stopped)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				mu.Lock()
				_ = flushBufferLocked()
				mu.Unlock()
			case <-done:
				return
			}
		}
	}()

	// Returns the wrapped write.
	write = func(chunk string) error {
		mu.Lock()
		defer mu.Unlock()

		if firstErr != nil {
			return firstErr
		}
		if closed {
			return errors.New("buffered streamer is closed")
		}
		if chunk == "" {
			return nil
		}

		buf.WriteString(chunk)
		if buf.Len() >= maxSize {
			return flushBufferLocked()
		}
		return nil
	}

	var once sync.Once
	var flushErr error

	// Flush everything, stop the ticker, and report callback errors.
	flush = func() error {
		once.Do(func() {
			mu.Lock()
			closed = true
			mu.Unlock()

			close(done)
			<-stopped

			mu.Lock()
			flushErr = flushBufferLocked()
			mu.Unlock()
		})
		return flushErr
	}

	return write, flush
}

// SafeCallStreamHandler invokes the provided StreamHandler and converts any
// panic into an error while logging the panic details. This prevents user
// callbacks from crashing the streaming loop.
func SafeCallStreamHandler(handler spec.StreamHandler, event spec.StreamEvent) (err error) {
	if handler == nil {
		return nil
	}

	// We use an inline recover here so we can both log and surface an error.
	defer func() {
		if r := recover(); r != nil {
			logutil.Error("stream handler panic",
				"panic", r,
				"kind", event.Kind,
				"provider", event.Provider,
				"model", event.Model,
				"stack", string(debug.Stack()),
			)
			err = fmt.Errorf("stream handler panic: %v", r)
		}
	}()

	return handler(event)
}

// ResolvedStreamConfig is the fully-specified streaming configuration used by
// providers after applying sensible defaults.
type ResolvedStreamConfig struct {
	FlushInterval  time.Duration
	FlushChunkSize int
}

// ResolveStreamConfig converts optional FetchCompletionOptions into a concrete
// ResolvedStreamConfig, falling back to library defaults as needed.
func ResolveStreamConfig(opts *spec.FetchCompletionOptions) ResolvedStreamConfig {
	cfg := ResolvedStreamConfig{
		FlushInterval:  FlushInterval,
		FlushChunkSize: FlushChunkSize,
	}
	if opts == nil || opts.StreamConfig == nil {
		return cfg
	}

	if opts.StreamConfig.FlushIntervalMillis > 0 {
		cfg.FlushInterval = time.Duration(opts.StreamConfig.FlushIntervalMillis) * time.Millisecond
	}
	if opts.StreamConfig.FlushChunkSize > 0 {
		cfg.FlushChunkSize = opts.StreamConfig.FlushChunkSize
	}
	return cfg
}
