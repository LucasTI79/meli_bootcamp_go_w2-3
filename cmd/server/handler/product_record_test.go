package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	productrecord "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product"
	mocks1 "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/product_record"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

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

		mockService.On("RecordsByAllProductsReport", mock.AnythingOfType("string")).Return(expectedProductRecords, nil)
		server.ServeHTTP(response, request)

		responseResult := &domain.ProductRecordReports{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, expectedProductRecords, responseResult.Data)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.True(t, len(responseResult.Data) == 2)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		var ExpectedEmpityProductRecords = []domain.ProductRecord{}

		server, mockService, handler := InitServerWithProductRecords(t)
		server.GET("/products/reportRecords", handler.RecordsByAllProductsReport())

		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords", "")

		mockService.On("RecordsByAllProductsReport", mock.AnythingOfType("string")).Return(ExpectedEmpityProductRecords, productrecord.ErrTryAgain)

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

		server.GET("/products/reportRecords/:id", handler.RecordsByOneProductReport())
		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/1", "")

		mockService.On("RecordsByOneProductReport", mock.AnythingOfType("int")).Return(expectedProductRecordReport, nil)

		server.ServeHTTP(response, request)

		responseResult := &domain.ProductRecordReportResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), responseResult)

		assert.Equal(t, expectedProductRecordReport, responseResult.Data)
		assert.Equal(t, http.StatusOK, response.Code)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		ExpectedEmpityProductReport := domain.ProductRecordReport{}

		server, mockService, handler := InitServerWithProductRecords(t)

		server.GET("/products/reportRecords/:id", handler.RecordsByOneProductReport())
		mockService.On("RecordsByOneProductReport", mock.Anything).Return(ExpectedEmpityProductReport, productrecord.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/1", "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("Should return status 404 when the report of a product is not found", func(t *testing.T) {

		server, mockService, handler := InitServerWithProductRecords(t)

		server.GET("/products/reportRecords/:id", handler.RecordsByOneProductReport())
		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/1", "")

		mockService.On("RecordsByOneProductReport", mock.AnythingOfType("int")).Return(domain.ProductRecordReport{}, productrecord.ErrNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return 400 when an product Id is invalid", func(t *testing.T) {

		server, mockService, handler := InitServerWithProductRecords(t)
		server.GET("/products/reportRecords/:id", handler.RecordsByOneProductReport())

		mockService.On("RecordsByOneProductReport", mock.Anything).Return(domain.ProductRecordReport{}, product.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodGet, "/products/reportRecords/invalidId", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func InitServerWithProductRecords(t *testing.T) (*gin.Engine, *mocks1.ProductRecordServiceMock, *handler.ProductRecordController) {
	t.Helper()
	server := testutil.CreateServer()
	mockProductRecordService := new(mocks1.ProductRecordServiceMock)
	mockProductService := new(mocks.ProductServiceMock)
	handler := handler.NewProductRecord(mockProductRecordService, mockProductService)
	return server, mockProductRecordService, handler
}