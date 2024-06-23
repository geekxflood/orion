// internal/httpclient/httpclient.go

package httpclient

import (
	"sync/atomic"
)

// Client is a struct that holds an atomic.Value for the configuration.
type Client struct {
	Conf *atomic.Value
}

// RunClient runs the HTTP client with the provided configuration and endpoints.
// It sets up routes for handling HTTP requests and starts the HTTP server.
