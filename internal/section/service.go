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
	Update(ctx context.Context, s domain.Section) error
	ExistsById(productID int) error
	ReportProductsById(ctx context.Context, id int) (domain.ProductBySection, error)
	ReportProducts(ctx context.Context) ([]domain.ProductBySection, error)
}

type serviceSection struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &serviceSection{
		repository: r,
	}
}
func (s *serviceSection) GetAll(ctx context.Context) ([]domain.Section, error) {
	sections, err := s.repository.GetAll(ctx)
	return sections, err
}
func (s *serviceSection) Get(ctx context.Context, id int) (domain.Section, error) {
	section, err := s.repository.Get(ctx, id)
	return section, err
}
func (s *serviceSection) Save(ctx context.Context, sect domain.Section) (int, error) {
	sectionExist := s.repository.Exists(ctx, sect.ID)
	if sectionExist {
		return 0, domain.ErrAlreadyExists
	}
	sectionID, err := s.repository.Save(ctx, sect)
	return sectionID, err
}
func (s *serviceSection) Delete(ctx context.Context, sectionNumber int) error {
	err := s.repository.Delete(ctx, sectionNumber)
	return err
}
func (s *serviceSection) Update(ctx context.Context, sect domain.Section) error {
	err := s.repository.Update(ctx, sect)
	return err
}

func (s *serviceSection) ExistsById(productID int) error {
	sectionExists := s.repository.ExistsById(productID)

	if !sectionExists {
		return errors.New("section not exists")
	}
	return nil
}

func (s *serviceSection) ReportProductsById(ctx context.Context, id int) (domain.ProductBySection, error) {
	return s.repository.SectionProductsReportsBySection(id)
}

func (s *serviceSection) ReportProducts(ctx context.Context) ([]domain.ProductBySection, error) {
	return s.repository.SectionProductsReports()
}
