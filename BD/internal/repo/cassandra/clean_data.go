package cassandra

import (
	"context"
	"github.com/pkg/errors"
)

func (r Repo) CleanData(ctx context.Context) error {
	req := `TRUNCATE ` + r.storiesTable.Name()
	if err := r.session.ContextQuery(ctx, req, nil).Exec(); err != nil {
		return errors.Wrapf(err, "failed to exec sql query %s", req)
	}

	req = `TRUNCATE ` + r.storiesViewStatTable.Name()
	if err := r.session.ContextQuery(ctx, req, nil).Exec(); err != nil {
		return errors.Wrapf(err, "failed to exec sql query %s", req)
	}

	return nil
}
