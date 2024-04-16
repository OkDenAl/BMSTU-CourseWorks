package cassandra

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) CreateStat(ctx context.Context, stat domain.StoryStat) error {
	err := r.storiesTable.
		InsertQueryContext(ctx, r.session).
		BindStruct(stat).
		WithBindTransformer(gocqlx.UnsetEmptyTransformer).
		Consistency(gocql.One).
		ExecRelease()
	if err != nil {
		return errors.Wrap(err, "failed to execute create stat query")
	}

	return nil

}
