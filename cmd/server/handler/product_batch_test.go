package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocksProduct "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product_batch"
	mocksSection "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/section"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductBatchServiceMocks struct {
	ProductServiceMock      *mocksProduct.ProductServiceMock
	SectionServiceMock      *mocksSection.SectionServiceMock
	ProductBatchServiceMock *mocks.ProductBatchServiceMock
}

func TestCreateProductBatch(t *testing.T) {
	// Should create a new product batch
	newProductBatch := &domain.ProductBatch{
		ID:                 1,
		ProductID:          1,
		SectionID:          1,
		BatchNumber:        1,
		CurrentQuantity:    1,
		InitialQuantity:    1,
		ManufacturingDate:  "2021-01-01",
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		DueDate:            "2021-01-01",
		ManufacturingHour:  "00:00:00",
	}

	t.Run("Should create a new product batch", func(t *testing.T) {
		server, handler, mocks := InitServerWithProductBatch(t)
		server.POST("/productBatches", handler.Create())

		jsonProductBatch, _ := json.Marshal(newProductBatch)
		request, response := testutil.MakeRequest("POST", "/productBatches", string(jsonProductBatch))

		mocks.ProductBatchServiceMock.On("Save", mock.Anything, mock.Anything).Return(1, nil)
		mocks.ProductServiceMock.On("ExistsById", newProductBatch.ProductID).Return(nil)
		mocks.SectionServiceMock.On("ExistsById", newProductBatch.SectionID).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
	// Should not save if product and/or section does not exist
	t.Run("Should not save if product and does not exist", func(t *testing.T) {
		server, handler, mocks := InitServerWithProductBatch(t)
		server.POST("/productBatches", handler.Create())

		jsonProductBatch, _ := json.Marshal(newProductBatch)
		request, response := testutil.MakeRequest("POST", "/productBatches", string(jsonProductBatch))

		mocks.ProductBatchServiceMock.On("Save", mock.Anything, mock.Anything).Return(0, nil)
		mocks.ProductServiceMock.On("ExistsById", newProductBatch.ProductID).Return(domain.ErrNotFound)
		mocks.SectionServiceMock.On("ExistsById", newProductBatch.SectionID).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should not save if section and does not exist", func(t *testing.T) {
		server, handler, mocks := InitServerWithProductBatch(t)
		server.POST("/productBatches", handler.Create())

		jsonProductBatch, _ := json.Marshal(newProductBatch)
		request, response := testutil.MakeRequest("POST", "/productBatches", string(jsonProductBatch))

		mocks.ProductBatchServiceMock.On("Save", mock.Anything, mock.Anything).Return(0, nil)
		mocks.ProductServiceMock.On("ExistsById", newProductBatch.ProductID).Return(nil)
		mocks.SectionServiceMock.On("ExistsById", newProductBatch.SectionID).Return(domain.ErrNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	// Should not save if any error occurs
	t.Run("Should not save if any error occurs", func(t *testing.T) {
		server, handler, mocks := InitServerWithProductBatch(t)
		server.POST("/productBatches", handler.Create())

		jsonProductBatch, _ := json.Marshal(newProductBatch)
		request, response := testutil.MakeRequest("POST", "/productBatches", string(jsonProductBatch))

		mocks.ProductBatchServiceMock.On("Save", mock.Anything, mock.Anything).Return(0, errors.New("error"))
		mocks.ProductServiceMock.On("ExistsById", newProductBatch.ProductID).Return(nil)
		mocks.SectionServiceMock.On("ExistsById", newProductBatch.SectionID).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
	// Should not save if product batch is invalid
	t.Run("Should not save if product batch is invalid", func(t *testing.T) {
		server, handler, mocks := InitServerWithProductBatch(t)
		server.POST("/productBatches", handler.Create())

		jsonProductBatch, _ := json.Marshal(&domain.ProductBatch{})
		request, response := testutil.MakeRequest("POST", "/productBatches", string(jsonProductBatch))

		mocks.ProductBatchServiceMock.On("Save", mock.Anything, mock.Anything).Return(0, nil)
		mocks.ProductServiceMock.On("ExistsById", newProductBatch.ProductID).Return(nil)
		mocks.SectionServiceMock.On("ExistsById", newProductBatch.SectionID).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	// Should not save if product batch fields is invalid
	t.Run("Should not save if product batch fields is invalid", func(t *testing.T) {
		server, handler, mocks := InitServerWithProductBatch(t)
		server.POST("/productBatches", handler.Create())

		jsonProductBatch, _ := json.Marshal(&domain.ProductBatch{
			ID:                 1,
			ProductID:          1,
			SectionID:          1,
			BatchNumber:        1,
			CurrentQuantity:    1,
			InitialQuantity:    1,
			ManufacturingDate:  "2021-01-01",
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			DueDate:            "INVALID DATE",
			ManufacturingHour:  "00:00:00",
		})
		request, response := testutil.MakeRequest("POST", "/productBatches", string(jsonProductBatch))

		mocks.ProductBatchServiceMock.On("Save", mock.Anything, mock.Anything).Return(0, nil)
		mocks.ProductServiceMock.On("ExistsById", newProductBatch.ProductID).Return(nil)
		mocks.SectionServiceMock.On("ExistsById", newProductBatch.SectionID).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

}

func InitServerWithProductBatch(t *testing.T) (*gin.Engine, *handler.ProductBatchController, ProductBatchServiceMocks) {
	t.Helper()
	server := testutil.CreateServer()
	productBatchService := new(mocks.ProductBatchServiceMock)
	productService := new(mocksProduct.ProductServiceMock)
	sectionService := new(mocksSection.SectionServiceMock)
	handler := handler.NewProductBatch(productBatchService, productService, sectionService)

	return server, handler, ProductBatchServiceMocks{
		ProductServiceMock:      productService,
		SectionServiceMock:      sectionService,
		ProductBatchServiceMock: productBatchService,
	}
}
