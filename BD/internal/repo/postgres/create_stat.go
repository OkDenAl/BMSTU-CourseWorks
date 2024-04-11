package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/postgres/dbview"
)

func (r Repo) CreateStat(ctx context.Context, stat domain.StoryStat) (string, error) {
	req, args, err := psql.Insert(dbview.StoryStatTableName).
		Columns(dbview.StoryStatAllColumns()...).
		Values(stat.Vals()...).
		Suffix(fmt.Sprintf("RETURNING %s", "story_id")).
		ToSql()
	if err != nil {
		return "", errors.Wrap(err, "failed to create sql query")
	}

	row := r.db.QueryRow(ctx, req, args...)

	var storyID string
	if err = row.Scan(&storyID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.Wrap(domain.ErrNoRows, "failed to insert story stat info")
		}
		return "", errors.Wrapf(err, "%s %q", req, args)
	}

	return storyID, nil
}
