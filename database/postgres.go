package database

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"github.com/google/uuid"
	"github.com/hbttundar/diabuddy-api-config/config/dbconfig"
	errors "github.com/hbttundar/diabuddy-errors"
	_ "github.com/lib/pq"
)

const (
	SelectPostgresLock   = "SELECT pg_advisory_lock($1)"
	SelectPostgresUnlock = "SELECT pg_advisory_unlock($1)"
)

type PostgresConnection struct {
	db               *sql.DB
	connectionString string
}

type PostgresConnectionOption func(*PostgresConnection) errors.ApiErrors

func WithPostgresConfig(config *dbconfig.DBConfig) PostgresConnectionOption {
	return func(pc *PostgresConnection) errors.ApiErrors {
		connString, err := config.ConnectionString()
		if err != nil {
			return err
		}
		pc.connectionString = connString
		return nil
	}
}

func WithConnectionString(connString string) PostgresConnectionOption {
	return func(pc *PostgresConnection) errors.ApiErrors {
		pc.connectionString = connString
		return nil
	}
}

func NewPostgresConnection(options ...PostgresConnectionOption) (*PostgresConnection, errors.ApiErrors) {
	pc := &PostgresConnection{}
	if len(options) == 0 {
		return nil, errors.NewApiError(errors.InternalServerErrorType, "no configuration options provided for PostgresConnection")
	}
	for _, opt := range options {
		if err := opt(pc); err != nil {
			return nil, err
		}
	}
	if pc.connectionString == "" {
		return nil, errors.NewApiError(errors.InternalServerErrorType, "database string is not configured")
	}
	return pc, nil
}

func (pc *PostgresConnection) Open(ctx context.Context) errors.ApiErrors {
	if pc.connectionString == "" {
		return errors.NewApiError(errors.InternalServerErrorType, "database string is not configured")
	}
	db, err := sql.Open("postgres", pc.connectionString)
	if err != nil {
		return errors.NewApiError(errors.InternalServerErrorType, "could not open connection to the database", errors.WithInternalError(err))
	}
	if err := db.PingContext(ctx); err != nil {
		return errors.NewApiError(errors.InternalServerErrorType, "could not connect to database", errors.WithInternalError(err))
	}
	pc.db = db
	return nil
}

func (pc *PostgresConnection) Close() errors.ApiErrors {
	if pc.db != nil {
		if err := pc.db.Close(); err != nil {
			return errors.NewApiError(errors.InternalServerErrorType, "could not close the database connection", errors.WithInternalError(err))
		}
	}
	return nil
}

func (pc *PostgresConnection) Ping(ctx context.Context) errors.ApiErrors {
	if pc.db == nil {
		return errors.NewApiError(errors.InternalServerErrorType, "database connection not initialized")
	}
	if err := pc.db.PingContext(ctx); err != nil {
		return errors.NewApiError(errors.InternalServerErrorType, "unable to reach the database", errors.WithInternalError(err))
	}
	return nil
}

func (pc *PostgresConnection) IsConnected(ctx context.Context) bool {
	if pc.db == nil {
		return false
	}
	return pc.db.PingContext(ctx) == nil
}

func (pc *PostgresConnection) DB() *sql.DB {
	return pc.db
}

func (pc *PostgresConnection) WrapInTransaction(ctx context.Context, fn func(tx *sql.Tx) errors.ApiErrors) (apiError errors.ApiErrors) {
	if pc.db == nil {
		return errors.NewApiError(errors.InternalServerErrorType, "Database connection is not initialized")
	}
	tx, err := pc.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.NewApiError(errors.InternalServerErrorType, "Failed to begin transaction", errors.WithInternalError(err))
	}
	advisoryLockKey := generateAdvisoryLockKey(uuid.New())
	_, lockErr := tx.ExecContext(ctx, SelectPostgresLock, advisoryLockKey)
	if lockErr != nil {
		_ = tx.Rollback()
		return errors.NewApiError(errors.InternalServerErrorType, "Failed to acquire advisory lock", errors.WithInternalError(lockErr))
	}
	defer func() {
		_, _ = tx.ExecContext(ctx, SelectPostgresUnlock, advisoryLockKey)
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if apiError != nil {
			_ = tx.Rollback()
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				apiError = errors.NewApiError(errors.InternalServerErrorType, "Failed to commit transaction", errors.WithInternalError(commitErr))
			}
		}
	}()
	apiError = fn(tx)
	return apiError
}

func generateAdvisoryLockKey(entityID uuid.UUID) int64 {
	hash := sha256.Sum256(entityID[:])
	return int64(binary.BigEndian.Uint64(hash[:8]))
}
