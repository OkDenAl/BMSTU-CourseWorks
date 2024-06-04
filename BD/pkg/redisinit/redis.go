package redisinit

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Addr     string `yaml:"addr" validate:"required"`
	PoolSize int    `yaml:"pool_size" validate:"required,gt=0"`
}

type CloserFn func()

func New(ctx context.Context, cfg Config) (*redis.Client, CloserFn, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		PoolSize: cfg.PoolSize,
	})

	const pingTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()
	if status := rdb.Ping(ctx); status.Err() != nil {
		disconnect(rdb)()
		return nil, nil, errors.Wrap(status.Err(), "failed to ping redis")
	}

	return rdb, disconnect(rdb), nil
}

func disconnect(client *redis.Client) CloserFn {
	return func() {
		err := client.Close()
		if err != nil {
			log.Panic().Stack().Err(err).Msg("failed to close redis client")
		}
	}
}
