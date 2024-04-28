package mongo

import (
	"context"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/pkg/errors"
)

func (r Repo) CreateStat(ctx context.Context, story domain.Story) error {
	storyView := NewStoryView(story)
	if _, err := r.stories.InsertOne(ctx, storyView); err != nil {
		return errors.Wrapf(err, "failed to insert %+v", storyView)
	}

	storyViewsStatView := NewStoryViewsStatView(story.StoryID)
	if _, err := r.storiesViewStat.InsertOne(ctx, storyViewsStatView); err != nil {
		return errors.Wrapf(err, "failed to insert %+v", storyViewsStatView)
	}

	return nil
}
