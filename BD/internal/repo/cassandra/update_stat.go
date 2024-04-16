package cassandra

import (
	"context"
)

func (r Repo) UpdateStat(ctx context.Context, storyID string) error {
	//err := r.storiesTable.
	//	UpdateQueryContext(ctx, r.session, "count").
	//	BindStruct(stat).
	//	WithBindTransformer(gocqlx.UnsetEmptyTransformer).
	//	Consistency(gocql.All).
	//	ExecRelease()
	//if err != nil {
	//	return errors.Wrap(err, "failed to execute create stat query")
	//}
	return nil
}
