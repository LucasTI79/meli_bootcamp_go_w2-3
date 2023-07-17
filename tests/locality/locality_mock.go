package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type LocalityServiceMock struct {
	mock.Mock
}

type LocalityRepositoryMock struct {
	mock.Mock
}

func (m *LocalityServiceMock) ExistsById(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *LocalityRepositoryMock) ExistsById(ctx context.Context, id int) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *LocalityServiceMock) ReportSellersByLocality(ctx context.Context, id int) ([]domain.LocalityReport, error){
	args := m.Called(ctx, id)
	return args.Get(0).([]domain.LocalityReport), args.Error(1)
}

func (m *LocalityServiceMock) Save(ctx context.Context, locality domain.Locality) (domain.LocalityInput, error){
	args := m.Called(ctx,locality)
	return args.Get(0).(domain.LocalityInput), args.Error(1)
}

func (m *LocalityRepositoryMock) Save(ctx context.Context, locality domain.LocalityInput) (int, error){
	args := m.Called(ctx, locality)
	return args.Get(0).(int), args.Error(1)
}

func (m *LocalityRepositoryMock) GetProvinceByName(ctx context.Context, name string) (int, error){
	args := m.Called(ctx, name)
	return args.Get(0).(int), args.Error(1)
}
func (m *LocalityRepositoryMock) ReportLocality(ctx context.Context) ([]domain.LocalityReport, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.LocalityReport), args.Error(1)
}

func (m *LocalityRepositoryMock) ReportLocalityId(ctx context.Context, id int) (domain.LocalityReport, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.LocalityReport), args.Error(1)
}