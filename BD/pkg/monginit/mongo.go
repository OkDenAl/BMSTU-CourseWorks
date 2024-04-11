package monginit

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Endpoint string `env:"MONGO_ENDPOINT"`
	DBName   string `yaml:"db_name" env:"DB_NAME" validate:"required"`

	ConnectTimeout time.Duration `yaml:"connect_timeout" env-default:"5s" validate:"required"`
	MinPoolSize    uint64        `yaml:"min_pool_size" env-default:"10" validate:"gte=0"`
	MaxPoolSize    uint64        `yaml:"max_pool_size" env-default:"10" validate:"gt=0"`
}

type CloserFn func(ctx context.Context) error

func Connect(ctx context.Context, cfg Config) (client *mongo.Client, closer CloserFn, err error) {
	opts := options.Client().
		ApplyURI(cfg.Endpoint).
		SetConnectTimeout(cfg.ConnectTimeout).
		SetMinPoolSize(cfg.MinPoolSize).
		SetMaxPoolSize(cfg.MaxPoolSize)

	if client, err = mongo.Connect(ctx, opts); err != nil {
		return nil, nil, errors.Wrap(err, "failed to create mongodb client")
	}

	const pingTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		disconnect(ctx, client)
		return nil, nil, errors.Wrap(err, "failed to ping mongodb")
	}

	return client, client.Disconnect, nil
}

func disconnect(ctx context.Context, client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to close mongo client")
	}
}
