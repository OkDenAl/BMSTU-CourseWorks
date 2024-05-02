package cassandra

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2/qb"
)

func (r Repo) UpdateStat(ctx context.Context, storyID string) error {
	err := r.session.
		ContextQuery(ctx,
			`UPDATE story_stat.story_views_stat SET count=count+1 WHERE story_id=?`,
			[]string{"story_id"}).
		BindMap(qb.M{"story_id": storyID}).
		Consistency(gocql.One).
		ExecRelease()
	if err != nil {
		return errors.Wrap(err, "failed to execute update counter query")
	}

	return nil
}
