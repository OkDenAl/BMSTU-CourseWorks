package redis

import (
	"github.com/redis/go-redis/v9"
)

type Repo struct {
	client *redis.Client
}

func New(client *redis.Client) *Repo {
	return &Repo{
		client: client,
	}
}
