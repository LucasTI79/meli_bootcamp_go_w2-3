package employee

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("employee not found")
	ErrCardNumberAlreadyExists = errors.New("employee already exists")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrInvalidId    = errors.New("invalid id")


)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	Save(ctx context.Context, e domain.Employee) (int, error)
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, id int) error
}

type employeeService struct{
	repository Repository
}

func NewService(r Repository) Service {
	return &employeeService{
		repository: r,
	}
}

// Delete implements Service.
func (s *employeeService) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// Get implements Service.
func (s *employeeService) Get(ctx context.Context, id int) (domain.Employee, error) {
	employee, err := s.repository.Get(ctx, id)
	return employee, err
}

func (s *employeeService) GetAll(ctx context.Context) ([]domain.Employee, error) {
	employees, err := s.repository.GetAll(ctx)
	return employees, err
}

func (s *employeeService) Save(ctx context.Context, e domain.Employee) (int, error) {
	employeeExists := s.repository.Exists(ctx, e.CardNumberID)
	if employeeExists {
		return 0, ErrCardNumberAlreadyExists
	}
	employeeId, err := s.repository.Save(ctx, e)
	return employeeId, err
}

// Update implements Service.
func (s *employeeService) Update(ctx context.Context, e domain.Employee) error {
	panic("unimplemented")
}