package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) CreateStat(ctx context.Context, stat domain.StoryStat) error {
	if _, err := r.col.InsertOne(ctx, stat); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.Wrapf(err, "failed to insert %+v", stat)
		}
		return errors.Wrapf(err, "failed to insert %+v", stat)
	}

	return nil
}
