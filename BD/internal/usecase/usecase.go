package usecase

import (
	"context"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

type iRepo interface {
	CreateStat(ctx context.Context, stat domain.StoryStat) error
	UpdateStat(ctx context.Context, storyID string) error
	GetStoryStatByID(ctx context.Context, id string) (domain.StoryStat, error)
}

type Usecase struct {
	mongoRepo     iRepo
	postgresRepo  iRepo
	cassandraRepo iRepo
}

func New(mongoRepo, postgresRepo, cassandraRepo iRepo) Usecase {
	return Usecase{mongoRepo: mongoRepo, cassandraRepo: cassandraRepo, postgresRepo: postgresRepo}
}
