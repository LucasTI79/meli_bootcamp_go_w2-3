package employee

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("employee not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	Save(ctx context.Context, e domain.Employee) (int, error)
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, id int) error
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Delete implements Service.
func (*service) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// Get implements Service.
func (*service) Get(ctx context.Context, id int) (domain.Employee, error) {
	panic("unimplemented")
}

// GetAll implements Service.
func (*service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	panic("unimplemented")
}

// Save implements Service.
func (*service) Save(ctx context.Context, e domain.Employee) (int, error) {
	panic("unimplemented")
}

// Update implements Service.
func (*service) Update(ctx context.Context, e domain.Employee) error {
	panic("unimplemented")
}