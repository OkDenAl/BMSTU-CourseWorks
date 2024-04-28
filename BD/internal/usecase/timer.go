package usecase

import (
	"context"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"time"
)

func Timer(ctx context.Context, stories []domain.Story, next BenchFunc) (*domain.BenchResult, error) {
	start := time.Now()
	res, err := next(ctx, stories)
	if err != nil {
		return nil, err
	}
	res.UsedTime = time.Since(start)
	return res, err
}
