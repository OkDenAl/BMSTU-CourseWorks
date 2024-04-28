package usecase

import (
	"context"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

type iRepo interface {
	CreateStat(ctx context.Context, stat domain.Story) error
	UpdateStat(ctx context.Context, storyID string) error
	GetStoryViewStatByID(ctx context.Context, id string) (domain.StoryStat, error)
	CleanData(ctx context.Context) error
}

type BenchFunc func(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error)

type Benchmark struct {
	repo iRepo
	cfg  config.BenchmarkConfig
}

func New(repo iRepo, cfg config.BenchmarkConfig) Benchmark {
	return Benchmark{repo: repo, cfg: cfg}
}
