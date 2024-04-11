package domain

import "github.com/pkg/errors"

var (
	ErrNoRows           = errors.New("no rows")
	ErrValidationFailed = errors.New("validation failed")
)
