package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) GetStoryViewStatByID(ctx context.Context, id string) (domain.StoryStat, error) {
	req, args, err := psql.Select(storyViewStatAllColumns()...).
		From(storyViewStatTableName).
		Where(squirrel.Eq{"story_id": id}).
		ToSql()
	if err != nil {
		return domain.StoryStat{}, errors.Wrap(err, "failed to create sql query")
	}

	var stat domain.StoryStat
	row := r.db.QueryRow(ctx, req, args...)
	if err = row.Scan(&stat.StoryID, &stat.Count); err != nil {
		return domain.StoryStat{}, errors.Wrap(err, "failed to get story stat")
	}

	return stat, nil
}
