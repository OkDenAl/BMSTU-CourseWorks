package cassandra

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2/qb"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) CreateStat(ctx context.Context, story domain.Story) error {
	err := r.storiesTable.
		InsertQueryContext(ctx, r.session).
		BindStruct(story).
		Consistency(gocql.One).
		ExecRelease()
	if err != nil {
		return errors.Wrap(err, "failed to insert story")
	}

	err = r.session.
		ContextQuery(ctx,
			`UPDATE story_stat.story_views_stat SET count=count+0 WHERE story_id=?`,
			[]string{"story_id"}).
		BindMap(qb.M{"story_id": story.StoryID}).
		Consistency(gocql.One).
		ExecRelease()
	if err != nil {
		return errors.Wrap(err, "failed to insert story view stat")
	}

	return nil
}
