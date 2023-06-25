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
	args := m.Called(ctx)
	return args.Get(0).([]domain.Employee), args.Error(1)
}
