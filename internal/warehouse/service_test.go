package warehouse_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/warehouse"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/warehouse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllWarehouses(t *testing.T) {
	t.Run("Should return all warehouses when repository is called", func(t *testing.T) {
		expectedWarehouses := []domain.Warehouse{
			{
				ID:                 1,
				Address:            "Rua Pedro Dias",
				Telephone:          "3712291281",
				WarehouseCode:      "DAE",
				MinimumCapacity:    10,
				MinimumTemperature: 10,
			},
			{
				ID:                 2,
				Address:            "Rua Maria das Dores",
				Telephone:          "1722919394",
				WarehouseCode:      "EWQ",
				MinimumCapacity:    10,
				MinimumTemperature: 10,
			},
		}

		repository, service := InitServerWithWarehousesRepository(t)
		repository.On("GetAll", mock.Anything).Return(expectedWarehouses, nil)

		warehouses, err := service.GetAll(context.TODO())

		assert.True(t, len(warehouses) == 2)
		assert.NoError(t, err)
	})
}

func TestCreateWarehouses(t *testing.T) {
	t.Run("Should create the warehouse if it contains the required fields", func(t *testing.T) {
		id := 4
		expectedWarehouse := domain.Warehouse{
			Address:            "Rua Pedro Dias",
			Telephone:          "3712291281",
			WarehouseCode:      "AEX",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		}

		repository, service := InitServerWithWarehousesRepository(t)
		repository.On("Exists", mock.Anything, "AEX").Return(false)
		repository.On("Save", mock.Anything, expectedWarehouse).Return(id, nil)

		warehouse, err := service.Save(context.TODO(), expectedWarehouse)

		assert.Equal(t, "Rua Pedro Dias", warehouse.Address)
		assert.Equal(t, "3712291281", warehouse.Telephone)
		assert.Equal(t, "AEX", warehouse.WarehouseCode)
		assert.Equal(t, 10, warehouse.MinimumCapacity)
		assert.Equal(t, 10, warehouse.MinimumTemperature)
		assert.Equal(t, 4, warehouse.ID)

		assert.NoError(t, err)
	})
	t.Run("Should return err warehouse already exists when warehouse already exists", func(t *testing.T) {
		expectedMessage := "warehouse already exists"

		repository, service := InitServerWithWarehousesRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(true)

		_, err := service.Save(context.TODO(), domain.Warehouse{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is an save repository error", func(t *testing.T) {
		repository, service := InitServerWithWarehousesRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(false)

		expectedError := errors.New("some error")
		repository.On("Save", mock.Anything, domain.Warehouse{}).Return(0, expectedError)

		_, err := service.Save(context.TODO(), domain.Warehouse{})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func InitServerWithWarehousesRepository(t *testing.T) (*mocks.WarehouseRepositoryMock, warehouse.Service) {
	t.Helper()
	mockRepository := &mocks.WarehouseRepositoryMock{}
	mockService := warehouse.NewService(mockRepository)
	return mockRepository, mockService
}
