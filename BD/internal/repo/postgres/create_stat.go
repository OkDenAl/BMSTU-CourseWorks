package postgres

import (
	"context"
	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) CreateStat(ctx context.Context, stat domain.StoryStat) error {
	req, args, err := psql.Insert(storyStatTableName).
		Columns(storyStatAllColumns()...).
		Values(stat.Vals()...).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to create sql query")
	}

	if _, err = r.db.Exec(ctx, req, args...); err != nil {
		return errors.Wrap(err, "failed to exec sql query")
	}

	return nil
}
