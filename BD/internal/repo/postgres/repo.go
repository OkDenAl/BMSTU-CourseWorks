package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

var ErrValidationFailed = errors.New("validation failed")

type Repo struct {
	db *pgxpool.Pool
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewRepo(db *pgxpool.Pool) (Repo, error) {
	if db == nil {
		return Repo{}, errors.Wrap(ErrValidationFailed, "failed to validate postgresinit db")
	}

	return Repo{db: db}, nil
}
