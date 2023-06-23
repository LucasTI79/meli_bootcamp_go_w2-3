package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/warehouse"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GetAllWarehouses    = "/warehouses"
	GetByIdWarehouses   = "/warehouses/4"
	GetCreateWarehouses = "/warehouses"
)

func TestGetAllWarehouses(t *testing.T) {
	emptyWarehouses := make([]domain.Warehouse, 0)
	t.Run("Should return status 200 with all warehouses", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
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
		server.GET("/warehouses", handler.GetAll())

		request, response := testutil.MakeRequest(http.MethodGet, GetAllWarehouses, "")
		mockService.On("GetAll", mock.AnythingOfType("string")).Return(expectedWarehouses, nil)
		server.ServeHTTP(response, request)

		responseResult := &domain.WarehouseResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusOK, response.Code)

		assert.Equal(t, expectedWarehouses, responseResult.Data)

		assert.True(t, len(responseResult.Data) == 2)
	})

	t.Run("Should return status 204 with no content", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
		server.GET("/warehouses", handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptyWarehouses, nil)

		request, response := testutil.MakeRequest(http.MethodGet, GetAllWarehouses, "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("Should return status 500 with no content", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
		server.GET("/warehouses", handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptyWarehouses, warehouse.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, GetAllWarehouses, "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestGetByIdWarehouses(t *testing.T) {
	// emptyWarehouse := domain.Warehouse{}
	t.Run("Should return a warehouse with id with status 200", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
		expectedWarehouse := domain.Warehouse{
			ID:                 4,
			Address:            "Rua Pedro Dias",
			Telephone:          "3712291281",
			WarehouseCode:      "DAE",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		}
		mockService.On("Get", mock.Anything, 4).Return(expectedWarehouse, nil)

		server.GET("/warehouses/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, GetByIdWarehouses, "")
		server.ServeHTTP(response, request)

		responseResult := &domain.WarehouseResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedWarehouse, responseResult.Data)
	})
	t.Run("Should return status 400 when the warehouse id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Get", mock.Anything, "invalid").Return(domain.Warehouse{}, warehouse.ErrInvalidId)

		server.GET("/warehouses/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, "/warehouses/invalid", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when the warehouse id does not exist", func(t *testing.T) {

		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Get", mock.Anything, 4).Return(domain.Warehouse{}, warehouse.ErrNotFound)

		server.GET("/warehouses/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, GetByIdWarehouses, "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when the warehouse id does not exist", func(t *testing.T) {

		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Get", mock.Anything, 4).Return(domain.Warehouse{}, warehouse.ErrTryAgain)

		server.GET("/warehouses/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, GetByIdWarehouses, "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCreateWarehouses(t *testing.T) {
	t.Run("Should return status 200 with the warehouse created", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Pedro Dias",
			Telephone:          "3712291281",
			WarehouseCode:      "DAE",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		}

		mockService.On("Save", mock.Anything, mock.Anything).Return(expectedWarehouse, nil)

		server.POST("/warehouses", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)
		server.ServeHTTP(response, request)

		responseResult := domain.WarehouseResponseId{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, expectedWarehouse, responseResult.Data)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		server.POST("/warehouses", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 400 when Address is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		server.POST("/warehouses", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":"","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when MinimumCapacity is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		server.POST("/warehouses", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":0,"minimum_temperature":10}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Telephone is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		server.POST("/warehouses", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":"Rua Pedro Dias","telephone":"","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when WarehouseCode is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		server.POST("/warehouses", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"","minimum_capacity":0,"minimum_temperature":10}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 409 when warehouse already exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		mockService.On("Save", mock.Anything, mock.AnythingOfType("domain.Warehouse")).Return(domain.Warehouse{}, warehouse.ErrAlredyExists)

		server.POST("/warehouses", handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, GetCreateWarehouses, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		mockService.On("Save", mock.Anything, mock.AnythingOfType("domain.Warehouse")).Return(domain.Warehouse{}, warehouse.ErrTryAgain)

		server.POST("/warehouses", handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func InitServerWithWarehouses(t *testing.T) (*gin.Engine, *mocks.WarehouseServiceMock, *handler.WarehouseController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.WarehouseServiceMock)
	handler := handler.NewWarehouse(mockService)
	return server, mockService, handler
}
