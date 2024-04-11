package handler

import (
	"context"

	"github.com/pkg/errors"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

type iUsecase interface {
	StartBenchmark(ctx context.Context) error
}

type Handler struct {
	uc iUsecase
}

func New(uc iUsecase) (Handler, error) {
	if uc == nil {
		return Handler{}, errors.Wrap(domain.ErrValidationFailed, "failed to validate iUsecase")
	}

	return Handler{uc: uc}, nil
}
