package usecase

import (
	"context"

	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

type iRepo interface {
	CreateStat(ctx context.Context, stat domain.StoryStat) (string, error)
	UpdateStat(ctx context.Context, storyID string) error
	GetStoryByIDs(ctx context.Context, IDs ...string) ([]domain.StoryStat, error)
}

type Usecase struct {
	repo iRepo
}

func New(repo iRepo) (Usecase, error) {
	if repo == nil {
		return Usecase{}, errors.Wrap(domain.ErrValidationFailed, "failed to validate repo")
	}

	return Usecase{repo: repo}, nil
}
