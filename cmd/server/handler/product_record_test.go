package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/middlewares"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	productrecord "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	productmocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product_record"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

type ProductRecordServiceMocks struct {
	MockProductService       *productmocks.ProductServiceMock
	MockProductRecordService *mocks.ProductRecordServiceMock
}

func TestRecordsByAllProductsReport(t *testing.T) {
	//case success
	t.Run("Should return status 200 with productsRecords of all products", func(t *testing.T) {
		server, mockService, handler := InitServerWithProductRecords(t)
		expectedProductRecords := []domain.ProductRecordReport{
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

		server.GET("/products/reportRecords", handler.RecordsByAllProductsReport())
		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords", "")

		mockService.MockProductRecordService.On("RecordsByAllProductsReport").Return(expectedProductRecords, nil)
		server.ServeHTTP(response, request)

		responseResult := &domain.ProductRecordReports{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, expectedProductRecords, responseResult.Data)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.True(t, len(responseResult.Data) == 2)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		var ExpectedEmpityReports = []domain.ProductRecordReport{}

		server, mockService, handler := InitServerWithProductRecords(t)

		server.GET("/products/reportRecords", handler.RecordsByAllProductsReport())

		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords", "")

		mockService.MockProductRecordService.On("RecordsByAllProductsReport").Return(ExpectedEmpityReports, productrecord.ErrTryAgain)

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

	})

}

func TestRecordsByOneProductReport(t *testing.T) {

	t.Run("Should return status 200 with the product record report", func(t *testing.T) {
		expectedProductRecordReport := domain.ProductRecordReport{
			ProductID:    1,
			Description:  "Product 1",
			RecordsCount: 3,
		}

		server, mockService, handler := InitServerWithProductRecords(t)

		server.GET("/products/reportRecords/:id", middlewares.ValidateParams("id"), handler.RecordsByOneProductReport())
		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/1", "")

		mockService.MockProductRecordService.On("RecordsByOneProductReport", mock.AnythingOfType("int")).Return(expectedProductRecordReport, nil)

		server.ServeHTTP(response, request)

		responseResult := &domain.ProductRecordReportResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), responseResult)

		assert.Equal(t, expectedProductRecordReport, responseResult.Data)
		assert.Equal(t, http.StatusOK, response.Code)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		ExpectedEmpityProductReport := domain.ProductRecordReport{}

		server, mockService, handler := InitServerWithProductRecords(t)

		server.GET("/products/reportRecords/:id", middlewares.ValidateParams("id"), handler.RecordsByOneProductReport())
		mockService.MockProductRecordService.On("RecordsByOneProductReport", mock.Anything).Return(ExpectedEmpityProductReport, productrecord.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/1", "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("Should return status 404 when the report of a product is not found", func(t *testing.T) {

		server, mockService, handler := InitServerWithProductRecords(t)

		server.GET("/products/reportRecords/:id", middlewares.ValidateParams("id"), handler.RecordsByOneProductReport())
		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/1", "")

		mockService.MockProductRecordService.On("RecordsByOneProductReport", mock.AnythingOfType("int")).Return(domain.ProductRecordReport{}, productrecord.ErrNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return 400 when an product Id is invalid", func(t *testing.T) {

		server, mockService, handler := InitServerWithProductRecords(t)
		server.GET("/products/reportRecords/:id", middlewares.ValidateParams("id"), handler.RecordsByOneProductReport())

		mockService.MockProductRecordService.On("RecordsByOneProductReport", mock.Anything).Return(domain.ProductRecordReport{}, product.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/invalidId", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestSave(t *testing.T) {
	var productRecordJson = `{
		"last_update_date": "5050-08-04",
		"purchase_price": 10,
		"sale_price": 15,
		"product_id": 1
	}`
	expectedProductRecord := domain.ProductRecord{
		ID:             1,
		LastUpdateDate: "2021-04-04",
		PurchasePrice:  10,
		SalePrice:      15,
		ProductID:      1,
	}
	t.Run("Should return 201 when a product record is created", func(t *testing.T) {

		server, mockService, handler := InitServerWithProductRecords(t)

		server.POST("/productRecords", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, "/productRecords", productRecordJson)

		mockService.MockProductRecordService.On("Save", mock.Anything, mock.Anything).Return(1, nil)

		mockService.MockProductService.On("ExistsById", expectedProductRecord.ID).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)

	})

	t.Run("Should return 422 when Json is invalid", func(t *testing.T) {

		server, _, handler := InitServerWithProductRecords(t)
		server.POST("/productRecords", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, "/productRecords", string(`{"sale_price":}`))

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("It Should not save a product record if the  product do not exist", func(t *testing.T) {

		server, mockService, handler := InitServerWithProductRecords(t)
		server.POST("/productRecords", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, "/productRecords", productRecordJson)

		mockService.MockProductRecordService.On("Save", mock.Anything, mock.Anything).Return(0, nil)

		mockService.MockProductService.On("ExistsById", expectedProductRecord.ID).Return(productrecord.ErrNotFound)
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should not save if any error occurs", func(t *testing.T) {

		server, mockService, handler := InitServerWithProductRecords(t)
		server.POST("/productRecords", handler.Create())
		request, response := testutil.MakeRequest(http.MethodPost, "/productRecords", productRecordJson)

		mockService.MockProductRecordService.On("Save", mock.Anything, mock.Anything).Return(0, errors.New("error"))

		mockService.MockProductService.On("ExistsById", expectedProductRecord.ID).Return(nil)
		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

	})

	t.Run("Should not save if product lastUpdateDate is invalid", func(t *testing.T) {

		server, _, handler := InitServerWithProductRecords(t)

		server.POST("/productRecords", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, "/productRecords", string(`{"last_update_date": "x"}`))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("It Should not save a product lastUpdateDate is less than system date", func(t *testing.T) {

		server, _, handler := InitServerWithProductRecords(t)

		server.POST("/productRecords", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, "/productRecords", string(`{"last_update_date": "2000-01-01"}`))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
}

func InitServerWithProductRecords(t *testing.T) (*gin.Engine, ProductRecordServiceMocks, *handler.ProductRecordController) {
	t.Helper()
	server := testutil.CreateServer()
	mockProductRecordService := new(mocks.ProductRecordServiceMock)
	mockProductService := new(productmocks.ProductServiceMock)
	handler := handler.NewProductRecord(mockProductRecordService, mockProductService)
	return server, ProductRecordServiceMocks{
		MockProductRecordService: mockProductRecordService,
		MockProductService:       mockProductService,
	}, handler
}
