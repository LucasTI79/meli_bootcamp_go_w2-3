package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type WarehouseServiceMock struct {
	mock.Mock
}

type WarehouseRepositoryMock struct {
	mock.Mock
}

func (m *WarehouseServiceMock) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	args := m.Called()
	return args.Get(0).([]domain.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) GetAll() ([]domain.Warehouse, error) {
	args := m.Called()
	return args.Get(0).([]domain.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) Save(ctx context.Context, s domain.Warehouse) (domain.Warehouse, error) {
	args := m.Called(ctx, s)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Save(ctx context.Context, s domain.Warehouse) (domain.Warehouse, error) {
	args := m.Called(ctx, s)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *WarehouseRepositoryMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *WarehouseServiceMock) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) Update(ctx context.Context, d domain.Warehouse, id int) (domain.Warehouse, error) {
	args := m.Called(ctx, d, id)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Update(ctx context.Context, d domain.Warehouse, id int) (domain.Warehouse, error) {
	args := m.Called(ctx, d, id)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}