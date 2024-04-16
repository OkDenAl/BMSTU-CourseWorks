package main

import (
	"context"
	"fmt"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/cassandra"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/repo/postgres"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/cassandrinit"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/logger"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/postgresinit"
	"github.com/ds248a/closer"
	"os"
	"time"
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

	t := time.Now()
	err = cassandraRepo.CreateStat(ctx, domain.NewStoryStat())
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create stat")
	}

	fmt.Println(time.Since(t))

	postgresRepo := postgres.New(pool)

	t = time.Now()
	err = postgresRepo.CreateStat(ctx, domain.NewStoryStat())
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create stat")
	}

	fmt.Println(time.Since(t))

	//wg := &sync.WaitGroup{}
	//wg.Add(1000)
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		defer wg.Done()
	//		err = repo.UpdateStat(ctx, "5c4e464e-074a-462a-af0e-88cb310a4f60")
	//		if err != nil {
	//			log.Panic().Stack().Err(err).Msg("failed to update stat")
	//		}
	//	}()
	//}
	//
	//ds, err := repo.GetStoryStatByID(
	//	ctx,
	//	"5c4e464e-074a-462a-af0e-88cb310a4f60",
	//)
	//if err != nil {
	//	fmt.Println(err)
	//	log.Panic().Stack().Err(err).Msg("failed to create stat")
	//}

	//wg.Wait()
	//
	//fmt.Println(ds)
	//
	//usecase.New(repo)
}
