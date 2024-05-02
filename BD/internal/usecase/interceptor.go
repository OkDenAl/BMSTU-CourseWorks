package usecase

import (
	"context"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

type BenchFuncInterceptor func(
	ctx context.Context,
	stories []domain.Story,
	next BenchFunc,
) (*domain.BenchResult, error)

func applyInterceptors(bFunc BenchFunc, interceptors []BenchFuncInterceptor) BenchFunc {
	finalBenchFunc := bFunc
	for i := len(interceptors) - 1; i >= 0; i-- {
		i := i
		fh := finalBenchFunc
		finalBenchFunc = func(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error) {
			return interceptors[i](ctx, stories, fh)
		}
	}

	return finalBenchFunc
}
