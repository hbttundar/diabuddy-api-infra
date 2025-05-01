package database

import (
	"context"
	"database/sql"

	errors "github.com/hbttundar/diabuddy-errors"
)

// Connection represents a general interface for establishing and managing connections.
type Connection interface {
	Open(ctx context.Context) errors.ApiErrors
	Close() errors.ApiErrors
	Ping(ctx context.Context) errors.ApiErrors
	IsConnected(ctx context.Context) bool
	DB() *sql.DB
	WrapInTransaction(ctx context.Context, fn func(tx *sql.Tx) errors.ApiErrors) errors.ApiErrors
}
