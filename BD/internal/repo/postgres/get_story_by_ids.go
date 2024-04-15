package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) GetStoryByIDs(ctx context.Context, ids ...string) ([]domain.StoryStat, error) {
	qArgs := make([]interface{}, len(ids))
	for i, v := range ids {
		qArgs[i] = v
	}

	req, args, err := psql.Select("*").
		From(storyStatTableName).
		Where(squirrel.Expr("story_id IN ("+squirrel.Placeholders(len(ids))+")", qArgs...)).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sql query")
	}

	fmt.Println(req)

	var stats []domain.StoryStat
	if err = pgxscan.Select(ctx, r.db, &stats, req, args...); err != nil {
		return nil, errors.Wrap(err, "failed to get story stats")
	}

	return stats, nil
}
