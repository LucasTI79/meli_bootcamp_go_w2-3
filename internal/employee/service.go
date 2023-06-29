package employee

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound      = errors.New("employee not found")
	ErrAlreadyExists = errors.New("employee already exists")
	ErrTryAgain      = errors.New("error, try again %s")
	ErrInvalidId     = errors.New("invalid id")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	Save(ctx context.Context, e domain.Employee) (domain.Employee, error)
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, id int) error
}

type employeeService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &employeeService{
		repository: r,
	}
}

func (s *employeeService) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	return err
}

func (s *employeeService) Get(ctx context.Context, id int) (domain.Employee, error) {
	employee, err := s.repository.Get(ctx, id)
	return employee, err
}

func (s *employeeService) GetAll(ctx context.Context) ([]domain.Employee, error) {
	employees, err := s.repository.GetAll(ctx)
	return employees, err
}

func (s *employeeService) Save(ctx context.Context, e domain.Employee) (domain.Employee, error) {
	employeeExists := s.repository.Exists(ctx, e.CardNumberID)
	if employeeExists {
		return domain.Employee{}, ErrAlreadyExists
	}
	employeeId, err := s.repository.Save(ctx, e)

	if err != nil {
		return domain.Employee{}, err
	}
	e.ID = employeeId
	return e, nil
}

func (s *employeeService) Update(ctx context.Context, e domain.Employee) error {
	err := s.repository.Update(ctx, e)
	return err
}
