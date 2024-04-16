package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	client *mongo.Client
	col    mongo.Collection
}

func New(client *mongo.Client, col mongo.Collection) *Repo {
	return &Repo{client: client, col: col}
}
