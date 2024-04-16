package mongo

import (
	"context"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) GetStoryByID(ctx context.Context, id string) (domain.Story, error) {
	return domain.Story{}, nil
}
