package cassandra

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) GetStoryStatByID(ctx context.Context, id string) (domain.StoryStat, error) {
	var stat domain.StoryStat
	allColumns := r.storiesTable.Metadata().Columns
	err := r.storiesTable.
		GetQueryContext(ctx, r.session, allColumns...).
		Bind(id).
		Consistency(gocql.One).
		GetRelease(&stat)
	if err != nil {
		return domain.StoryStat{}, errors.Wrap(err, "failed to get story stat")
	}

	return stat, nil
}
