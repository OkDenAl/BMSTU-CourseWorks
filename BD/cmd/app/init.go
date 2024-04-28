package main

import (
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/logger"
)

func prepareDataForBenchmark(cfg config.BenchmarkConfig) []domain.Story {
	log := logger.New()
	log.Debug().Msg("starting to prepare data for benchmark")

	stories := make([]domain.Story, 0, cfg.ObjectsAmount)
	for i := 0; i < cfg.ObjectsAmount; i++ {
		stories = append(stories, domain.NewStory())
	}

	log.Debug().Msg("data successfully prepared")

	return stories
}
