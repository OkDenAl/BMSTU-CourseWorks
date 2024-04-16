package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) CreateStat(ctx context.Context, story domain.Story) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	req, args, err := psql.Insert(storyTableName).
		Columns(storyAllColumns()...).
		Values(story.Vals()...).
		ToSql()
	if err != nil {
		_ = tx.Rollback(ctx)
		return errors.Wrap(err, "failed to create sql query for stories")
	}

	if _, err = tx.Exec(ctx, req, args...); err != nil {
		_ = tx.Rollback(ctx)
		return errors.Wrapf(err, "failed to exec sql query %s", req)
	}

	req, args, err = psql.Insert(storyViewStatTableName).
		Columns(storyViewStatAllColumns()...).
		Values(story.StoryID, 0).
		ToSql()
	if err != nil {
		_ = tx.Rollback(ctx)
		return errors.Wrap(err, "failed to create sql query for stories stat")
	}

	if _, err = tx.Exec(ctx, req, args...); err != nil {
		_ = tx.Rollback(ctx)
		return errors.Wrapf(err, "failed to exec sql query %s", req)
	}

	return tx.Commit(ctx)
}
