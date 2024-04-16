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
}

type Benchmark struct {
	repo iRepo
	cfg  config.BenchmarkConfig
	data []domain.Story
}

func New(repo iRepo, cfg config.BenchmarkConfig, data []domain.Story) Benchmark {
	return Benchmark{repo: repo, cfg: cfg, data: data}
}
