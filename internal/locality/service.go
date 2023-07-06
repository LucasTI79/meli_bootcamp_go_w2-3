package locality

import (
	"errors"
)

var (
	ErrNotFound = errors.New("seller not found")
)

type Service interface {
}

type LocalityService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &LocalityService{
		repository: r,
	}
}
