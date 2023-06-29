package employee_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/employee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var expectedEmployee = domain.Employee{
	ID:           01,
	CardNumberID: "001",
	FirstName:    "Joana",
	LastName:     "Silva",
	WarehouseID:  1,
}

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

func TestCreateEmployees(t *testing.T) {
	t.Run("Should create the employee if it contains the required fields", func(t *testing.T) {
		id := 4
		expectedEmployee := domain.Employee{
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository, service := InitServerWithEmployeesRepository(t)
		repository.On("Exists", mock.Anything, "001").Return(false)
		repository.On("Save", mock.Anything, expectedEmployee).Return(id, nil)

		_, err := service.Save(context.TODO(), domain.Employee{})

		// assert.Equal(t, "Joana", employee.FirstName)
		// assert.Equal(t, "Silva", employee.LastName)
		// assert.Equal(t, "001", employee.CardNumberID)
		// assert.Equal(t, 1, employee.WarehouseID)
		// assert.Equal(t, 4, employee.ID)

		assert.NoError(t, err)
	})
	t.Run("Should return err employee already exists when employee already exists", func(t *testing.T) {
		expectedMessage := "employee already exists"

		repository, service := InitServerWithEmployeesRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(true)

		_, err := service.Save(context.TODO(), domain.Employee{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is an save repository error", func(t *testing.T) {
		repository, service := InitServerWithEmployeesRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(false)

		expectedError := errors.New("some error")
		repository.On("Save", mock.Anything, domain.Employee{}).Return(0, expectedError)

		_, err := service.Save(context.TODO(), domain.Employee{})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetEmployeesById(t *testing.T) {
	t.Run("Should return the employee when it exists", func(t *testing.T) {

		repository, service := InitServerWithEmployeesRepository(t)

		repository.On("Get", mock.Anything).Return(expectedEmployee, nil)

		employee, err := service.Get(context.TODO(), 1)

		assert.Equal(t, expectedEmployee, employee)
		assert.NoError(t, err)
	})
	t.Run("Should return an error when the employee does not exists", func(t *testing.T) {
		repository, service := InitServerWithEmployeesRepository(t)
		expectedError := errors.New("employee not found")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.Employee{}, employee.ErrNotFound)
		_, err := service.Get(context.TODO(), 1)
		assert.Equal(t, expectedError, err)
		assert.Error(t, err)
	})
}

func TestDeleteEmployees(t *testing.T) {
	t.Run("Should delete the employee when it exists in database", func(t *testing.T) {
		expectedEmployee := domain.Employee{
			ID:           4,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository, service := InitServerWithEmployeesRepository(t)
		repository.On("Delete", mock.Anything, expectedEmployee.ID).Return(nil)

		err := service.Delete(context.TODO(), 4)

		assert.NoError(t, err)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		repository, service := InitServerWithEmployeesRepository(t)

		expectedError := errors.New("employee not found")
		repository.On("Delete", mock.Anything, mock.Anything).Return(employee.ErrNotFound)

		err := service.Delete(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("Should return error when there is an delete repository error", func(t *testing.T) {
		repository, service := InitServerWithEmployeesRepository(t)

		expectedError := errors.New("some error")
		repository.On("Delete", mock.Anything, mock.Anything).Return(expectedError)

		err := service.Delete(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestUpdateEmployee(t *testing.T) {

	t.Run("Should update the employee when it exists.", func(t *testing.T) {

		repository, service := InitServerWithEmployeesRepository(t)

		repository.On("Exists", mock.Anything, expectedEmployee.CardNumberID).Return(false)
		repository.On("Update", mock.Anything, expectedEmployee).Return(nil)

		err := service.Update(context.TODO(), expectedEmployee)

		assert.NoError(t, err)
	})

	t.Run("Should return an error when the employee does not exists", func(t *testing.T) {
		repository, service := InitServerWithEmployeesRepository(t)

		expectedError := errors.New("employee not found")
		repository.On("Update", mock.Anything, mock.Anything).Return(employee.ErrNotFound)

		err := service.Update(context.TODO(), expectedEmployee)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

}
