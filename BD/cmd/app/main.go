package main

import (
	"context"
	"fmt"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/cassandra"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/cassandrinit"
	"github.com/ds248a/closer"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/usecase"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/logger"
)

func main() {
	logger.SetupWriter()
	log := logger.New()
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to setup cfg")
	}

	//pool, pgCloser, err := postgresinit.NewPool(ctx, cfg.Postgres)
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to create postgres pool")
	//}
	//closer.Add(pgCloser)

	sess, csCloser, err := cassandrinit.New(ctx, cfg.Cassandra)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create cassandra session")
	}
	closer.Add(csCloser)
	cassandraRepo := cassandra.New(*sess)

	//mongoClient, mgCloser, err := monginit.Connect(ctx, cfg.Mongo)
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to create mongo connection")
	//}
	//closer.Add(mgCloser)
	//
	//storiesCol := mongoClient.Database(cfg.Mongo.DBName).Collection(mongo.Stories.CollectionName)
	//storyViewStatCol := mongoClient.Database(cfg.Mongo.DBName).Collection(mongo.StoryViewsStat.CollectionName)
	//mongoRepo := mongo.New(mongoClient, storiesCol, storyViewStatCol)
	//
	//postgresRepo := postgres.New(pool)

	preparedData := prepareDataForBenchmark(cfg.Benchmark)

	var (
		cassandraResults []*domain.BenchResult
		postgresResults  []*domain.BenchResult
		mongoResults     []*domain.BenchResult
	)

	u := usecase.New(cassandraRepo, cfg.Benchmark)
	cassandraResults, err = u.StartBenchmarks(ctx, preparedData)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to check cassandra")
	}

	//u := usecase.New(postgresRepo, cfg.Benchmark)
	//postgresResults, err = u.StartBenchmarks(ctx, preparedData)
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to check postgres")
	//}
	//
	//u = usecase.New(mongoRepo, cfg.Benchmark)
	//mongoResults, err = u.StartBenchmarks(ctx, preparedData)
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to check mongo")
	//}

	fmt.Println("POSTGRES RESULTS:")
	printResults(postgresResults)

	fmt.Println("CASSANDRA RESULTS:")
	printResults(cassandraResults)

	fmt.Println("MONGO RESULTS:")
	printResults(mongoResults)
}

func printResults(res []*domain.BenchResult) {
	for _, val := range res {
		fmt.Println(val)
	}
}
