package carry

import (
	"errors"
)

var (
	ErrNotFound     = errors.New("carry not found")
	ErrInvalidId    = errors.New("invalid id")
	ErrInvalidBody  = errors.New("invalid body")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrAlredyExists = errors.New("carry already exists")
	ErrInvalidJSON  = errors.New("invalid json")
)
