package request

import (
	"net/http"
	"time"
)

// DefaultHTTPClient returns a shared HTTP client with sane defaults.
func DefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       100,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: false,
		},
	}
}

// NewHTTPClient allows customization for advanced use cases.
func NewHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConns:       100,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: false,
		},
	}
}
