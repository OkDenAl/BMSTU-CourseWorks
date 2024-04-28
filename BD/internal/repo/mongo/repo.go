package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	client          *mongo.Client
	stories         *mongo.Collection
	storiesViewStat *mongo.Collection
}

func New(client *mongo.Client, stories, storiesViewStat *mongo.Collection) *Repo {
	return &Repo{client: client, stories: stories, storiesViewStat: storiesViewStat}
}
