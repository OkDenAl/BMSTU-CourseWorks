package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r Repo) CleanData(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return errors.Wrap(err, "failed to begin clean data transaction")
	}
	req := `TRUNCATE TABLE ` + storyTableName
	if _, err = tx.Exec(ctx, req); err != nil {
		_ = tx.Rollback(ctx)
		return errors.Wrapf(err, "failed to exec sql query %s", req)
	}

	req = `TRUNCATE TABLE ` + storyViewStatTableName
	if _, err = tx.Exec(ctx, req); err != nil {
		_ = tx.Rollback(ctx)
		return errors.Wrapf(err, "failed to exec sql query %s", req)
	}

	return tx.Commit(ctx)
}
