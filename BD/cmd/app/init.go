package main

import (
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func prepareDataForBenchmark(cfg config.BenchmarkConfig) []domain.Story {
	stories := make([]domain.Story, 0, cfg.ObjectsAmount)
	for i := 0; i < cfg.ObjectsAmount; i++ {
		stories = append(stories, domain.NewStory())
	}

	return stories
}
