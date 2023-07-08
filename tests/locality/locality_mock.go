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

func (m *LocalityRepositoryMock) ExistsById(ctx context.Context, id int) bool {
	args := m.Called(ctx, id)
	return args.Get(0).(bool)
}

func (m *LocalityRepositoryMock) ReportLocality(ctx context.Context) ([]domain.LocalityReport, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.LocalityReport), args.Error(1)
}

func (m *LocalityRepositoryMock) ReportLocalityId(ctx context.Context, id int) (domain.LocalityReport, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.LocalityReport), args.Error(1)
}