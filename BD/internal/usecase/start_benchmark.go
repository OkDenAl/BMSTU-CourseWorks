package usecase

import (
	"context"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (b Benchmark) StartBenchmarks(
	ctx context.Context,
	stories []domain.Story,
) ([]*domain.BenchResult, error) {
	if b.cfg.CleanDataBefore {
		if err := b.repo.CleanData(ctx); err != nil {
			return nil, err
		}
	}

	benchFuncs := b.buildBenchFuncs(b.cfg)
	interceptorsArr := []BenchFuncInterceptor{
		Timer,
		Logger,
	}

	results := make([]*domain.BenchResult, 0)
	for _, benchF := range benchFuncs {
		benchF = applyInterceptors(benchF, interceptorsArr)
		res, err := benchF(ctx, stories)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}
