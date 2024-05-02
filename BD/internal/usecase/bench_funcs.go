package usecase

import (
	"context"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"golang.org/x/sync/errgroup"
	"math/rand"
)

func (b Benchmark) createBench(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error) {
	const benchName = "story_creation"

	for _, story := range stories {
		if err := b.repo.CreateStat(ctx, story); err != nil {
			return nil, err
		}
	}

	return domain.NewBenchResult(benchName, b.cfg.ObjectsAmount), nil
}

func (b Benchmark) createBenchAsync(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error) {
	const benchName = "story_creation_async"

	eg := errgroup.Group{}
	for _, story := range stories {
		story := story
		eg.Go(func() error {
			return b.repo.CreateStat(ctx, story)
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return domain.NewBenchResult(benchName, b.cfg.ObjectsAmount), nil
}

func (b Benchmark) getBench(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error) {
	const benchName = "get_story_stat"

	for range stories {
		if _, err := b.repo.GetStoryViewStatByID(ctx, stories[rand.Intn(b.cfg.ObjectsAmount)].StoryID); err != nil {
			return nil, err
		}
	}

	return domain.NewBenchResult(benchName, b.cfg.ObjectsAmount), nil
}

func (b Benchmark) getBenchAsync(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error) {
	const benchName = "get_story_stat_async"

	eg := errgroup.Group{}
	for range stories {
		eg.Go(func() error {
			if _, err := b.repo.GetStoryViewStatByID(ctx, stories[rand.Intn(b.cfg.ObjectsAmount)].StoryID); err != nil {
				return err
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return domain.NewBenchResult(benchName, b.cfg.ObjectsAmount), nil
}

func (b Benchmark) updateBench(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error) {
	const benchName = "update_story_stat"

	for range stories {
		if err := b.repo.UpdateStat(ctx, stories[rand.Intn(b.cfg.ObjectsAmount)].StoryID); err != nil {
			return nil, err
		}
	}

	return domain.NewBenchResult(benchName, b.cfg.ObjectsAmount), nil
}

func (b Benchmark) updateBenchAsync(ctx context.Context, stories []domain.Story) (*domain.BenchResult, error) {
	const benchName = "update_story_stat_async"

	eg := errgroup.Group{}
	for range stories {
		eg.Go(func() error {
			return b.repo.UpdateStat(ctx, stories[rand.Intn(b.cfg.ObjectsAmount)].StoryID)
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return domain.NewBenchResult(benchName, b.cfg.ObjectsAmount), nil
}
