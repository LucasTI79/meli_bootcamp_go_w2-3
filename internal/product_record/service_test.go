package productrecord_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	productrecord "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_record"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product_record"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRecordsByAllProductsReport(t *testing.T) {
	t.Run("Should return the product records report when repository is called", func(t *testing.T) {

		expectedReport := []domain.ProductRecordReport{
			{
				ProductID:    1,
				Description:  "Product 1",
				RecordsCount: 3,
			},
			{
				ProductID:    2,
				Description:  "Product 2",
				RecordsCount: 3,
			},
		}

		service, repository := CreateProductRecordService(t)
		repository.On("RecordsByAllProductsReport", mock.Anything).Return(expectedReport, nil)

		reports, err := service.RecordsByAllProductsReport(context.TODO())
		fmt.Println("reports mock", reports)
		assert.True(t, len(reports) == 2)
		assert.NoError(t, err)
	})
}

func CreateProductRecordService(t *testing.T) (productrecord.Service, *mocks.ProductRecordRepositoryMock) {
	mockRepository := new(mocks.ProductRecordRepositoryMock)
	mockService := productrecord.NewService(mockRepository)
	return mockService, mockRepository
}
