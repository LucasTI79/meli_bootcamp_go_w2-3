package productbatch_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	productbatch "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_batch"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product_batch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductBatch(t *testing.T) {

	// Create new product batch
	t.Run("Should create a new product batch", func(t *testing.T) {
		expectedProductBatch := domain.ProductBatch{
			ID:                 1,
			ProductID:          1,
			SectionID:          1,
			BatchNumber:        1,
			CurrentQuantity:    1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			DueDate:            "2021-01-01",
			InitialQuantity:    1,
			ManufacturingDate:  "2021-01-01",
			ManufacturingHour:  1,
		}
		mockRepository, service := InitProductBatchService(t)
		mockRepository.On("Save", mock.Anything).Return(1, nil)

		productBatchId, err := service.Save(context.TODO(), expectedProductBatch)
		assert.Equal(t, 1, productBatchId)
		assert.NoError(t, err)
	})

	// Should not create a new product batch
	t.Run("Should not create a new product batch", func(t *testing.T) {
		expectedProductBatch := domain.ProductBatch{
			ID:                 1,
			ProductID:          1,
			SectionID:          1,
			BatchNumber:        1,
			CurrentQuantity:    1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			DueDate:            "2021-01-01",
			InitialQuantity:    1,
			ManufacturingDate:  "2021-01-01",
			ManufacturingHour:  1,
		}
		mockRepository, service := InitProductBatchService(t)
		mockRepository.On("Save", mock.Anything).Return(0, errors.New("error"))

		productBatchId, err := service.Save(context.TODO(), expectedProductBatch)
		assert.Equal(t, 0, productBatchId)
		assert.Error(t, err)
	})
}
func InitProductBatchService(t *testing.T) (*mocks.ProductBatchRepositoryMock, productbatch.Service) {
	mockRepository := new(mocks.ProductBatchRepositoryMock)
	mockService := productbatch.NewService(mockRepository)

	return mockRepository, mockService
}
