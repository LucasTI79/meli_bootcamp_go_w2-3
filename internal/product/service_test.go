package product_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	t.Run("Should return all products when repository is called", func(t *testing.T) {
		expectedProducts := []domain.Product{
			{
				ID:             1,
				Description:    "milk",
				ExpirationRate: 1,
				FreezingRate:   2,
				Height:         6.4,
				Length:         4.5,
				Netweight:      3.4,
				ProductCode:    "PROD01",
				RecomFreezTemp: 1.3,
				Width:          1.2,
				ProductTypeID:  1,
				SellerID:       1,
			},
			{
				ID:             2,
				Description:    "milk",
				ExpirationRate: 1,
				FreezingRate:   2,
				Height:         6.4,
				Length:         4.5,
				Netweight:      3.4,
				ProductCode:    "PROD02",
				RecomFreezTemp: 1.3,
				Width:          1.2,
				ProductTypeID:  2,
				SellerID:       2,
			},
		}

		service, repository := CreareProductService(t)
		repository.On("GetAll", mock.Anything).Return(expectedProducts, nil)

		products, err := service.GetAll(context.TODO())

		assert.True(t, len(products) == 2)
		assert.NoError(t, err)
	})
}

func CreareProductService(t *testing.T) (product.Service, *mocks.ProductRepositoryMock) {
	mockRepository := new(mocks.ProductRepositoryMock)
	mockService := product.NewService(mockRepository)
	return mockService, mockRepository
}
