package database

import (
	"context"
	"database/sql"

	"github.com/hbttundar/diabuddy-api-config/config/dbconfig"
	"github.com/hbttundar/diabuddy-api-config/config/envmanager"
)

type Connection interface {
	DB() *sql.DB
	Open(ctx context.Context) error
	Close() error
}

func NewDefaultTestConnection(ctx context.Context) (Connection, error) {
	envMgr, err := envmanager.NewEnvManager(envmanager.WithEnvironment("test"), envmanager.WithUseDefault(true))
	if err != nil {
		return nil, err
	}

	config, err := dbconfig.NewDBConfig(envMgr, dbconfig.WithType(dbconfig.Postgres), dbconfig.WithDsnParameters(map[string]string{"sslmode": "disable"}))
	if err != nil {
		return nil, err
	}

	conn := NewPostgresConnection(WithPostgresConfig(config))
	err = conn.Open(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
