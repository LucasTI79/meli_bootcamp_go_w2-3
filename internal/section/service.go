package section

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
)

type Service interface {
	Save(ctx context.Context, s domain.Section) (int, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
}

type serviceSection struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &serviceSection{
		repository: r,
	}
}
func (s *serviceSection) Save(ctx context.Context, sect domain.Section) (int, error) {
	section, err := s.repository.Save(ctx, sect)
	return section, err
}
func (s *serviceSection) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	return err
}
func (s *serviceSection) GetAll(ctx context.Context) ([]domain.Section, error){
	sections, err := s.repository.GetAll(ctx)
	return sections, err
}
func (s *serviceSection) Get(ctx context.Context, id int) (domain.Section, error){
	section, err := s.repository.Get(ctx, id)
	return section, err
}
