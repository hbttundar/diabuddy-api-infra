package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	config *PostgresConfig
	db     *sql.DB
}

type PostgresConfig struct {
	Dsn string
}

type PostgresOption func(*PostgresConfig)

func WithPostgresConfig(config *PostgresConfig) PostgresOption {
	return func(p *PostgresConfig) {
		p.Dsn = config.Dsn
	}
}

func NewPostgresConnection(opts ...PostgresOption) *PostgresConnection {
	config := &PostgresConfig{}
	for _, opt := range opts {
		opt(config)
	}
	return &PostgresConnection{config: config}
}

func (p *PostgresConnection) Open(ctx context.Context) error {
	db, err := sql.Open("postgres", p.config.Dsn)
	if err != nil {
		return fmt.Errorf("failed to open postgres connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping postgres: %w", err)
	}

	p.db = db
	return nil
}

func (p *PostgresConnection) DB() *sql.DB {
	return p.db
}

func (p *PostgresConnection) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}
