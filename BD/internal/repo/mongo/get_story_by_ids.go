package mongo

import (
	"context"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r Repo) GetStoryViewStatByID(ctx context.Context, id string) (domain.StoryStat, error) {
	res := r.storiesViewStat.FindOne(
		ctx,
		bson.M{
			"story_id": id,
		},
	)
	if res.Err() != nil {
		return domain.StoryStat{}, errors.Wrapf(res.Err(), "failed to get stat for %s", id)
	}

	var storyStat domain.StoryStat
	err := res.Decode(&storyStat)
	if err != nil {
		return domain.StoryStat{}, errors.Wrap(err, "failed to decode result")
	}

	return storyStat, nil
}
