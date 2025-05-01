package http_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	infrahttp "github.com/hbttundar/diabuddy-api-infra/http"
)

func TestDefaultHTTPClient(t *testing.T) {
	client := infrahttp.DefaultHTTPClient()
	assert.NotNil(t, client)
	assert.Equal(t, 10*time.Second, client.Timeout)
}

func TestNewHTTPClient_CustomTimeout(t *testing.T) {
	timeout := 3 * time.Second
	client := infrahttp.NewHTTPClient(timeout)
	assert.Equal(t, timeout, client.Timeout)
}
