package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) CreateStat(ctx context.Context, stat domain.Story) error {
	if _, err := r.col.InsertOne(ctx, stat); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.Wrapf(err, "failed to insert %+v", stat)
		}
		return errors.Wrapf(err, "failed to insert %+v", stat)
	}

	return nil
}
