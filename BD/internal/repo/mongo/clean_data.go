package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/pkg/errors"
)

func (r Repo) CleanData(ctx context.Context) error {
	if _, err := r.stories.DeleteMany(ctx, bson.D{}); err != nil {
		return errors.Wrapf(err, "failed to clean data from %s", Stories.CollectionName)
	}

	if _, err := r.storiesViewStat.DeleteMany(ctx, bson.D{}); err != nil {
		return errors.Wrapf(err, "failed to clean data from %s", StoryViewsStat.CollectionName)
	}

	return nil
}
