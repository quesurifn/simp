package database

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"quesurifn/simple-auth/pkg/logger"
	"time"
)

type Database struct {
	connectionURI string
	logger        *logger.Logger
	ctx           context.Context
	config        DatabaseConfig
}

type DatabaseConfig struct {
	MaxConnections    int
	MinConnections    int
	MaxConnectionTime time.Duration
}

func NewDatabase(uri string, logger *logger.Logger, config DatabaseConfig) *Database {
	return &Database{
		connectionURI: uri,
		logger:        logger,
		config: DatabaseConfig{
			MaxConnectionTime: config.MaxConnectionTime,
			MinConnections:    config.MinConnections,
			MaxConnections:    config.MaxConnections,
		},
	}
}

func (db Database) Connect() (*sql.DB, error) {
	poolConfig, err := pgxpool.ParseConfig(db.connectionURI)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = int32(db.config.MaxConnections)
	poolConfig.MinConns = int32(db.config.MinConnections)
	poolConfig.MaxConnLifetime = db.config.MaxConnectionTime
	pool, err := pgxpool.NewWithConfig(db.ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	dbInstance := stdlib.OpenDBFromPool(pool)

	// drv := entsql.OpenDB(dialect.Postgres, db)
	// client := ent.NewClient(ent.Driver(drv))

	return dbInstance, nil
}
