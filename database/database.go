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
	config        Config
	Instance      *sql.DB
}

type Config struct {
	MaxConnections    int
	MinConnections    int
	MaxConnectionTime time.Duration
}

func NewDatabase(uri string, logger *logger.Logger, config Config) *Database {
	return &Database{
		connectionURI: uri,
		logger:        logger,
		config: Config{
			MaxConnectionTime: config.MaxConnectionTime,
			MinConnections:    config.MinConnections,
			MaxConnections:    config.MaxConnections,
		},
	}
}

func (db *Database) Connect() error {
	poolConfig, err := pgxpool.ParseConfig(db.connectionURI)
	if err != nil {
		return err
	}
	poolConfig.MaxConns = int32(db.config.MaxConnections)
	poolConfig.MinConns = int32(db.config.MinConnections)
	poolConfig.MaxConnLifetime = db.config.MaxConnectionTime
	pool, err := pgxpool.NewWithConfig(db.ctx, poolConfig)
	if err != nil {
		return err
	}

	dbInstance := stdlib.OpenDBFromPool(pool)

	// drv := entsql.OpenDB(dialect.Postgres, db)
	// client := ent.NewClient(ent.Driver(drv))
	db.Instance = dbInstance

	return nil
}

func (db *Database) Close() error {
	err := db.Instance.Close()
	if err != nil {
		return err
	}
	db.Instance = nil

	return nil
}
