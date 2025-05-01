package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hbttundar/diabuddy-api-infra/helpers"
)

func TestToPointer(t *testing.T) {
	str := "hello"
	ptr := helpers.ToPointer(str)
	require.NotNil(t, ptr)
	assert.Equal(t, "hello", *ptr)

	zero := helpers.ToPointer(0)
	assert.Equal(t, 0, *zero)
}

func TestIfNotEmpty(t *testing.T) {
	assert.Equal(t, "fallback", helpers.IfNotEmpty("", "fallback"))
	assert.Equal(t, "given", helpers.IfNotEmpty("given", "fallback"))
}

func TestCoalesce(t *testing.T) {
	assert.Equal(t, "first", helpers.Coalesce("first", "second"))
	assert.Equal(t, "second", helpers.Coalesce("", "second"))
	assert.Equal(t, "", helpers.Coalesce("", ""))
	assert.Equal(t, "val", helpers.Coalesce("", "", "val"))
}
