package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hbttundar/diabuddy-api-infra/database"
)

func TestNewPostgresConnection_NoOptions(t *testing.T) {
	conn, err := database.NewPostgresConnection()
	assert.Nil(t, conn)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no configuration options provided")
}

func TestNewPostgresConnection_EmptyConnectionString(t *testing.T) {
	conn, err := database.NewPostgresConnection(database.WithConnectionString(""))
	assert.Nil(t, conn)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "database string is not configured")
}
