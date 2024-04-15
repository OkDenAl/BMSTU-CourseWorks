package main

import (
	"context"
	"fmt"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/postgres"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/usecase"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/logger"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/postgresinit"
	"github.com/ds248a/closer"
	"os"
	"sync"
)

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	logger.SetupWriter()
	log := logger.New()
	ctx := context.Background()

	m := map[string][]string{
		"string":  {"test"},
		"string1": {"test1", "test1"},
	}

	log.Info().Interface("headers", m).Msg("print")

	cfg, err := config.New()
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to setup cfg")
	}

	pool, pgCloser, err := postgresinit.NewPool(ctx, cfg.Postgres)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create postgres pool")
	}
	closer.Add(pgCloser)

	repo := postgres.NewRepo(pool)

	//story, err := repo.CreateStat(ctx, domain.NewStoryStat("userId", "story"))
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to create stat")
	//}

	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			err = repo.UpdateStat(ctx, "5c4e464e-074a-462a-af0e-88cb310a4f60")
			if err != nil {
				log.Panic().Stack().Err(err).Msg("failed to update stat")
			}
		}()
	}

	ds, err := repo.GetStoryByIDs(
		ctx,
		"5c4e464e-074a-462a-af0e-88cb310a4f60",
		"aedeba22-ea42-4f2e-b228-a7b36fea7c65",
	)
	if err != nil {
		fmt.Println(err)
		log.Panic().Stack().Err(err).Msg("failed to create stat")
	}

	wg.Wait()

	fmt.Println(ds)

	usecase.New(repo)
}
