package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type EmployeeServiceMock struct {
	mock.Mock
}

type EmployeeRepositoryMock struct {
	mock.Mock
}

func (m *EmployeeServiceMock) GetAll(ctx context.Context) ([]domain.Employee, error) {
	args := m.Called()
	return args.Get(0).([]domain.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) GetAll(ctx context.Context) ([]domain.Employee, error) {
	args := m.Called()
	return args.Get(0).([]domain.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Save(ctx context.Context, e domain.Employee) (domain.Employee, error) {
	args := m.Called(ctx, e)
	return args.Get(0).(domain.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Save(ctx context.Context, e domain.Employee) (int, error) {
	args := m.Called(ctx, e)
	return args.Get(0).(int), args.Error(1)
}

func (m *EmployeeServiceMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *EmployeeRepositoryMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *EmployeeServiceMock) Get(ctx context.Context, id int) (domain.Employee, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Get(ctx context.Context, id int) (domain.Employee, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Update(ctx context.Context, e domain.Employee, id int) (domain.Employee, error) {
	args := m.Called(ctx, e, id)
	return args.Get(0).(domain.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Update(ctx context.Context, e domain.Employee) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *EmployeeRepositoryMock) Exists(ctx context.Context, warehouseCode string) bool {
	args := m.Called(ctx, warehouseCode)
	return args.Get(0).(bool)
}
