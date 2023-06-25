package handler_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/employee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GetAllEmployees = "/employees"
)

func TestGetAllEmployees(t *testing.T) {
	t.Run("Should return all employeess when repository is called", func(t *testing.T) {
		expectedEmployees := []domain.Employee{
			{
				ID:           01,
				CardNumberID: "001",
				FirstName:    "Joana",
				LastName:     "Silva",
				WarehouseID:  1,
			},
			{
				ID:           02,
				CardNumberID: "002",
				FirstName:    "Beatriz",
				LastName:     "Costa",
				WarehouseID:  1,
			},
		}

		repository, service := InitServerWithEmployeesRepository(t)
		repository.On("GetAll", mock.Anything).Return(expectedEmployees, nil)

		employees, err := service.GetAll(context.TODO())

		assert.True(t, len(employees) == 2)
		assert.NoError(t, err)
	})
}

func InitServerWithEmployeesRepository(t *testing.T) (*mocks.EmployeeRepositoryMock, employee.Service) {
	t.Helper()
	mockRepository := &mocks.EmployeeRepositoryMock{}
	mockService := employee.NewService(mockRepository)
	return mockRepository, mockService
}
