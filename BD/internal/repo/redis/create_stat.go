package redis

import (
	"context"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) CreateStat(ctx context.Context, story domain.Story) error {
	r.client.Set(ctx, story.StoryID, story, 0)

	storyStat := domain.StoryStat{StoryID: story.StoryID, Count: 0}
	return r.client.Set(ctx, story.StoryID, storyStat, 0).Err()
}
