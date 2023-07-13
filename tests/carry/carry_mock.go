package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type CarryServiceMock struct {
	mock.Mock
}

type CarryRepositoryMock struct {
	mock.Mock
}

func (m *CarryServiceMock) Read(ctx context.Context, id int) ([]domain.LocalityCarriersReport, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]domain.LocalityCarriersReport), args.Error(1)
}

func (m *CarryRepositoryMock) ReadAllCarriers(ctx context.Context) ([]domain.LocalityCarriersReport, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.LocalityCarriersReport), args.Error(1)
}

func (m *CarryRepositoryMock) ReadCarriersWithLocalityId(ctx context.Context, localityID int) (domain.LocalityCarriersReport, error) {
	args := m.Called(ctx, localityID)
	return args.Get(0).(domain.LocalityCarriersReport), args.Error(1)
}

func (m *CarryServiceMock) Create(ctx context.Context, c domain.Carry) (domain.Carry, error) {
	args := m.Called(ctx, c)
	return args.Get(0).(domain.Carry), args.Error(1)
}

func (m *CarryRepositoryMock) Create(ctx context.Context, c domain.Carry) (int, error) {
	args := m.Called(ctx, c)
	return args.Get(0).(int), args.Error(1)
}

func (m *CarryServiceMock) Get(ctx context.Context, id int) (domain.Carry, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Carry), args.Error(1)
}

func (m *CarryRepositoryMock) Get(ctx context.Context, id int) (domain.Carry, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Carry), args.Error(1)
}

func (m *CarryRepositoryMock) ExistsByCidCarry(ctx context.Context, cid string) bool {
	args := m.Called(ctx, cid)
	return args.Get(0).(bool)
}
