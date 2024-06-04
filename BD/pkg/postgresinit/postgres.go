package postgresinit

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Config struct {
	ConnString string `yaml:"conn_string" env:"POSTGRES_CONN_STRING" validate:"required"`

	MaxIdleLifetime time.Duration `yaml:"max_idle_lifetime" env-default:"10s" validate:"required"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime" env-default:"5m" validate:"required"`
	MaxConnAmount   int32         `yaml:"max_conn_amount" env-default:"30" validate:"required"`
	MinConnAmount   int32         `yaml:"min_conn_amount" env-default:"5" validate:"required"`
}

type CloserFn func()

func NewPool(ctx context.Context, cfg Config) (*pgxpool.Pool, CloserFn, error) {
	pgxConfig, err := pgxpool.ParseConfig(cfg.ConnString)
	if err != nil {
		return nil, nil, err
	}
	pgxConfig.MaxConnIdleTime = cfg.MaxIdleLifetime
	pgxConfig.MaxConnLifetime = cfg.MaxConnLifetime
	pgxConfig.MaxConns = cfg.MaxConnAmount
	pgxConfig.MinConns = cfg.MinConnAmount

	dbPool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create pgxpool")
	}

	const pingTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()
	if err = dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		return nil, nil, errors.Wrap(err, "failed to ping postgres")
	}

	return dbPool, dbPool.Close, nil
}
