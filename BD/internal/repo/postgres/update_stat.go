package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r Repo) UpdateStat(ctx context.Context, storyID string) error {
	req, args, err := psql.Update(storyStatTableName).
		Set("count", squirrel.Expr("count + 1")).
		Where(squirrel.Eq{"story_id": storyID}).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to create sql query")
	}

	if _, err = r.db.Exec(ctx, req, args...); err != nil {
		return errors.Wrapf(err, "%s %q", req, args)
	}

	return nil
}
