package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/ds248a/closer"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/cassandra"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/postgres"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/usecase"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/cassandrinit"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/logger"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/postgresinit"
)

func main() {
	logger.SetupWriter()
	log := logger.New()
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to setup cfg")
	}

	pool, pgCloser, err := postgresinit.NewPool(ctx, cfg.Postgres)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create postgres pool")
	}
	closer.Add(pgCloser)

	sess, csCloser, err := cassandrinit.New(ctx, cfg.Cassandra)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create cassandra session")
	}
	closer.Add(csCloser)

	cassandraRepo := cassandra.New(*sess)
	postgresRepo := postgres.New(pool)

	log.Debug().Msg("starting to prepare data for benchmark")
	preparedData := prepareDataForBenchmark(cfg.Benchmark)
	log.Debug().Msg("data successfully prepared")

	wg := sync.WaitGroup{}
	var (
		cassandraResults []string
		postgresResults  []string
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		cassandraResults, err = usecase.New(cassandraRepo, cfg.Benchmark, preparedData).StartBenchmarks(cfg.Benchmark)
		if err != nil {
			log.Panic().Stack().Err(err).Msg("failed to check cassandra")
		}
	}()
	go func() {
		defer wg.Done()
		postgresResults, err = usecase.New(postgresRepo, cfg.Benchmark, preparedData).StartBenchmarks(cfg.Benchmark)
		if err != nil {
			log.Panic().Stack().Err(err).Msg("failed to check postgres")
		}
	}()

	wg.Wait()

	fmt.Println(cassandraResults)
	fmt.Println(postgresResults)
}
