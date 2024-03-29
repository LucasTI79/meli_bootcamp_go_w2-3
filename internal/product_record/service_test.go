package productrecord_test

import (
	"context"
	"errors"

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

		assert.True(t, len(reports) == 2)
		assert.NoError(t, err)
	})
}

func TestRecordsByOneProductReport(t *testing.T) {
	t.Run("Should return the product record when it exists", func(t *testing.T) {

		expectedReport := domain.ProductRecordReport{
			ProductID:    1,
			Description:  "Product 1",
			RecordsCount: 3,
		}
		service, repository := CreateProductRecordService(t)

		repository.On("RecordsByOneProductReport", mock.Anything).Return(expectedReport, nil)

		report, err := service.RecordsByOneProductReport(context.TODO(), 1)

		assert.Equal(t, expectedReport, report)
		assert.NoError(t, err)
	})
	t.Run("Should return an error when the product does not exists", func(t *testing.T) {
		expectedEmpityReport := domain.ProductRecordReport{}
		service, repository := CreateProductRecordService(t)
		expectedError := errors.New("product not found")
		repository.On("RecordsByOneProductReport", mock.Anything).Return(expectedEmpityReport, productrecord.ErrNotFound)
		_, err := service.RecordsByOneProductReport(context.TODO(), 1)
		assert.Equal(t, expectedError, err)
		assert.Error(t, err)
	})
}

func TestCreateProducts(t *testing.T) {

	expectedProductRecord := domain.ProductRecord{
		ID:             1,
		LastUpdateDate: "2021-04-04",
		PurchasePrice:  10,
		SalePrice:      15,
		ProductID:      1,
	}
	id := 1
	t.Run("Should create a product record when it contains the necessary fields", func(t *testing.T) {

		service, repository := CreateProductRecordService(t)
		repository.On("Save", mock.Anything).Return(id, nil)

		productReportId, err := service.Save(context.TODO(), expectedProductRecord)

		assert.Equal(t, expectedProductRecord.ID, productReportId)

		assert.NoError(t, err)
	})
}

func CreateProductRecordService(t *testing.T) (productrecord.Service, *mocks.ProductRecordRepositoryMock) {
	mockRepository := new(mocks.ProductRecordRepositoryMock)
	mockService := productrecord.NewService(mockRepository)
	return mockService, mockRepository
}
