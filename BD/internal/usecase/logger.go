package usecase

import (
	"context"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/logger"
)

func Logger(ctx context.Context, stories []domain.Story, next BenchFunc) (*domain.BenchResult, error) {
	log := logger.New()
	log.Debug().Msg("starting a benchmark")
	res, err := next(ctx, stories)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("%s successfully ended", res.Name)
	return res, err
}
