package cassandra

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) GetStoryViewStatByID(ctx context.Context, id string) (domain.StoryStat, error) {
	var stat domain.StoryStat
	allColumns := r.storiesViewStatTable.Metadata().Columns
	err := r.storiesViewStatTable.
		GetQueryContext(ctx, r.session, allColumns...).
		Bind(id).
		Consistency(gocql.One).
		GetRelease(&stat)
	if err != nil {
		return domain.StoryStat{}, errors.Wrapf(err, "failed to get story stat with %s", id)
	}

	return stat, nil
}
