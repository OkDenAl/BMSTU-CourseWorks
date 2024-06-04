package main

import (
	"context"
	"fmt"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/redis"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/usecase"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/logger"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/redisinit"
	"github.com/ds248a/closer"
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
	//postgresRepo := postgres.New(pool)

	//sess, csCloser, err := cassandrinit.New(ctx, cfg.Cassandra)
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to create cassandra session")
	//}
	//closer.Add(csCloser)
	//cassandraRepo := cassandra.New(*sess)

	//mongoClient, mgCloser, err := monginit.Connect(ctx, cfg.Mongo)
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to create mongo connection")
	//}
	//closer.Add(mgCloser)
	//
	//storiesCol := mongoClient.Database(cfg.Mongo.DBName).Collection(mongo.Stories.CollectionName)
	//storyViewStatCol := mongoClient.Database(cfg.Mongo.DBName).Collection(mongo.StoryViewsStat.CollectionName)
	//mongoRepo := mongo.New(mongoClient, storiesCol, storyViewStatCol)

	redisClient, rdCloser, err := redisinit.New(ctx, cfg.Redis)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create mongo connection")
	}
	closer.Add(rdCloser)
	redisRepo := redis.New(redisClient)

	preparedData := prepareDataForBenchmark(cfg.Benchmark)

	var (
		cassandraResults []*domain.BenchResult
		postgresResults  []*domain.BenchResult
		mongoResults     []*domain.BenchResult
		redisResults     []*domain.BenchResult
	)

	//u := usecase.New(cassandraRepo, cfg.Benchmark)
	//cassandraResults, err = u.StartBenchmarks(ctx, preparedData)
	//if err != nil {
	//	log.Panic().Stack().Err(err).Msg("failed to check cassandra")
	//}

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

	u := usecase.New(redisRepo, cfg.Benchmark)
	redisResults, err = u.StartBenchmarks(ctx, preparedData)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to check redis")
	}

	fmt.Println("POSTGRES RESULTS:")
	printResults(postgresResults)

	fmt.Println("CASSANDRA RESULTS:")
	printResults(cassandraResults)

	fmt.Println("MONGO RESULTS:")
	printResults(mongoResults)

	fmt.Println("REDIS RESULTS:")
	printResults(redisResults)
}

func printResults(res []*domain.BenchResult) {
	for _, val := range res {
		fmt.Println(val)
	}
}
