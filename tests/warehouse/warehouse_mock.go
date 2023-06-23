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

func (m *WarehouseServiceMock) Save(ctx context.Context, s domain.Warehouse) (int, error) {
	return 0, nil
}

func (m *WarehouseServiceMock) Delete(ctx context.Context, id int) error {
	return nil
}

func (m *WarehouseServiceMock) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) Update(ctx context.Context, s domain.Warehouse) error {
	return nil
}
