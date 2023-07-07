package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SectionServiceMock struct {
	mock.Mock
}

type SectionRepositoryMock struct {
	mock.Mock
}

func (m *SectionServiceMock) GetAll(ctx context.Context) ([]domain.Section, error) {
	args := m.Called()

	return args.Get(0).([]domain.Section), args.Error(1)
}

func (m *SectionRepositoryMock) GetAll(ctx context.Context) ([]domain.Section, error) {

	args := m.Called()

	return args.Get(0).([]domain.Section), args.Error(1)
}

func (m *SectionServiceMock) Save(ctx context.Context, s domain.Section) (int, error) {
	args := m.Called(s)
	return args.Int(0), args.Error(1)
}

func (m *SectionRepositoryMock) Save(ctx context.Context, s domain.Section) (int, error) {
	args := m.Called(s)
	return args.Int(0), args.Error(1)
}

func (m *SectionServiceMock) Delete(ctx context.Context, id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *SectionRepositoryMock) Delete(ctx context.Context, id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *SectionServiceMock) Get(ctx context.Context, id int) (domain.Section, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Section), args.Error(1)
}

func (m *SectionRepositoryMock) Get(ctx context.Context, id int) (domain.Section, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Section), args.Error(1)
}

func (m *SectionServiceMock) Update(ctx context.Context, s domain.Section) error {
	args := m.Called(ctx, s)
	return args.Error(0)
}
func (m *SectionRepositoryMock) Update(ctx context.Context, s domain.Section) error {
	args := m.Called(ctx, s)
	return args.Error(0)
}

func (m *SectionRepositoryMock) Exists(ctx context.Context, sectionNumber int) bool {
	args := m.Called(ctx, sectionNumber)
	return args.Get(0).(bool)
}

func (m *SectionServiceMock) ExistsById(productID int) error {
	args := m.Called(productID)
	return args.Error(0)
}

func (m *SectionServiceMock) ReportProducts(ctx context.Context) ([]domain.ProductBySection, error) {
	args := m.Called()
	return args.Get(0).([]domain.ProductBySection), args.Error(1)
}

func (m *SectionRepositoryMock) ExistsById(sectionID int) bool {
	args := m.Called(sectionID)
	return args.Get(0).(bool)
}

func (m *SectionRepositoryMock) SectionProductsReports() ([]domain.ProductBySection, error) {
	args := m.Called()
	return args.Get(0).([]domain.ProductBySection), args.Error(1)
}

func (m *SectionRepositoryMock) SectionProductsReportsBySection(id int) (domain.ProductBySection, error) {
	args := m.Called(id)
	return args.Get(0).(domain.ProductBySection), args.Error(1)
}

func (m *SectionServiceMock) ReportProductsById(ctx context.Context, id int) (domain.ProductBySection, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.ProductBySection), args.Error(1)
}
