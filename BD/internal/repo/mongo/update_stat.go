package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repo) UpdateStat(ctx context.Context, storyID string) error {
	_, err := r.col.UpdateOne(
		ctx,
		bson.M{
			"story_id": storyID,
		},
		bson.M{
			"$set": "count+1",
		},
		options.Update(),
	)
	if err != nil {
		return errors.Wrapf(err, "failed to update stat for %s", storyID)
	}

	return nil
}
