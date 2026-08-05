package debugclient

import (
	"net"
	"net/http"
	"time"
)

const (
	streamDialTimeout  = 30 * time.Second
	streamTCPKeepAlive = 10 * time.Second
)

// defaultStreamingHTTPTransport is shared across provider clients so HTTP/1.1
// and HTTP/2 connections can be pooled. None of these settings imposes an
// inactivity deadline on an active response body.
var defaultStreamingHTTPTransport = newStreamingHTTPTransport()

func newStreamingHTTPTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:   streamDialTimeout,
		KeepAlive: streamTCPKeepAlive,
		KeepAliveConfig: net.KeepAliveConfig{
			Enable:   true,
			Idle:     streamTCPKeepAlive,
			Interval: streamTCPKeepAlive,
			Count:    16,
		},
	}

	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   20,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		// Deliberately zero. Providers can take a long time before returning
		// headers, particularly when reached through compatible gateways.
		ResponseHeaderTimeout: 0,
	}
}

func newStreamingHTTPClient() *http.Client {
	return &http.Client{
		Transport: defaultStreamingHTTPTransport,

		// A non-zero http.Client.Timeout includes reading the complete response
		// body and is therefore unsafe for long-running streams. Provider
		// request options and caller contexts enforce total request deadlines.
		Timeout: 0,
	}
}
