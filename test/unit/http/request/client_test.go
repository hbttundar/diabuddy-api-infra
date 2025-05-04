package request_test

import (
	infrahttp "github.com/hbttundar/diabuddy-api-infra/http/request"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
