package redis

import (
	"context"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) GetStoryViewStatByID(ctx context.Context, id string) (domain.StoryStat, error) {
	cmd := r.client.Get(ctx, id)
	if err := cmd.Err(); err != nil {
		return domain.StoryStat{}, err
	}

	var storyStat domain.StoryStat
	if err := cmd.Scan(storyStat); err != nil {
		return domain.StoryStat{}, err
	}
	return storyStat, nil
}
