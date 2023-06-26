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
	BaseEndpoint = "/warehouses"
	BaseEndpointWithId = "/warehouses/:id"
	BaseEndpointWithIdNumber = "/warehouses/1"
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

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpoint, "")

		mockService.On("GetAll", mock.Anything, mock.AnythingOfType("string")).Return(expectedWarehouses, nil)

		server.GET(BaseEndpoint, handler.GetAll())
		server.ServeHTTP(response, request)

		responseResult := &domain.WarehouseResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedWarehouses, responseResult.Data)
		assert.True(t, len(responseResult.Data) == 2)
	})

	t.Run("Should return status 204 with no content", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("GetAll", mock.Anything, mock.AnythingOfType("string")).Return(emptyWarehouses, nil)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpoint, "")

		server.GET(BaseEndpoint, handler.GetAll())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("Should return status 500 with no content", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("GetAll", mock.Anything, mock.AnythingOfType("string")).Return(emptyWarehouses, warehouse.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpoint, "")

		server.GET(BaseEndpoint, handler.GetAll())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestGetByIdWarehouses(t *testing.T) {
	// emptyWarehouse := domain.Warehouse{}
	t.Run("Should return status 200 and warehouse with id", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Pedro Dias",
			Telephone:          "3712291281",
			WarehouseCode:      "DAE",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		}
		mockService.On("Get", mock.Anything, 1).Return(expectedWarehouse, nil)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointWithIdNumber, "")

		server.GET(BaseEndpointWithId, handler.Get())
		server.ServeHTTP(response, request)

		responseResult := &domain.WarehouseResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedWarehouse, responseResult.Data)
	})
	t.Run("Should return status 400 when the warehouse id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Get", mock.Anything, "invalid").Return(domain.Warehouse{}, warehouse.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodGet, "/warehouses/invalid", "")

		server.GET(BaseEndpointWithId, handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when the warehouse id does not exist", func(t *testing.T) {

		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.Warehouse{}, warehouse.ErrNotFound)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointWithIdNumber, "")

		server.GET(BaseEndpointWithId, handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {

		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.Warehouse{}, warehouse.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointWithIdNumber, "")

		server.GET(BaseEndpointWithId, handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCreateWarehouses(t *testing.T) {
	t.Run("Should return status 200 and the warehouse created", func(t *testing.T) {
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

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)
		
		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		responseResult := domain.WarehouseResponseId{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, expectedWarehouse, responseResult.Data)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":}`)

		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 400 when Address is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":"","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when MinimumCapacity is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":0,"minimum_temperature":10}`)

		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Telephone is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":"Rua Pedro Dias","telephone":"","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return status 400 when WarehouseCode is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":"Rua Pedro Dias","telephone":"37111029","warehouse_code":"","minimum_capacity":10,"minimum_temperature":10}`)

		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 409 when Warehouse already exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		mockService.On("Save", mock.Anything, mock.AnythingOfType("domain.Warehouse")).Return(domain.Warehouse{}, warehouse.ErrAlredyExists)

		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpoint, `{"address":"Rua Pedro Dias","telephone":"3712291281","warehouse_code":"DAEAQ","minimum_capacity":10,"minimum_temperature":10}`)

		mockService.On("Save", mock.Anything, mock.AnythingOfType("domain.Warehouse")).Return(domain.Warehouse{}, warehouse.ErrTryAgain)

		server.POST(BaseEndpoint, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestDeleteWarehouses(t *testing.T) {
	// emptyWarehouse := domain.Warehouse{}
	t.Run("Should return status 204 and delete a warehouse with id", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Delete", mock.Anything, 1).Return(nil)

		request, response := testutil.MakeRequest(http.MethodDelete, BaseEndpointWithIdNumber, "")

		server.DELETE(BaseEndpointWithId, handler.Delete())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
	t.Run("Should return status 404 when warehouse is not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodDelete, BaseEndpointWithIdNumber, "")

		mockService.On("Delete", mock.Anything, 1).Return(warehouse.ErrNotFound)

		server.DELETE(BaseEndpointWithId, handler.Delete())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 400 when the warehouse id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodDelete, "/warehouses/invalid", "")

		mockService.On("Delete", mock.Anything, "invalid").Return(warehouse.ErrInvalidId)

		server.DELETE(BaseEndpointWithId, handler.Delete())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodDelete, BaseEndpointWithIdNumber, "")

		mockService.On("Delete", mock.Anything, 1).Return(warehouse.ErrTryAgain)

		server.DELETE(BaseEndpointWithId, handler.Delete())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestUpdateWarehouse(t *testing.T) {
	t.Run("Should return status 200 and updated warehouse", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
		updatedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Pedro Dias",
			Telephone:          "3712291281",
			WarehouseCode:      "DAE",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		}

		mockService.On("Update", mock.Anything, mock.Anything, 1).Return(updatedWarehouse, nil)

		request, response := testutil.MakeRequest(http.MethodPatch, BaseEndpointWithIdNumber, `{"address":"Rua Pedro Dias","telephone":"371928"}`)

		server.PATCH(BaseEndpointWithId, handler.Update())
		server.ServeHTTP(response, request)

		responseResult := domain.WarehouseResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, updatedWarehouse, responseResult.Data)
	})
	t.Run("Should return status 400 when the warehouse id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		mockService.On("Update", mock.Anything, mock.Anything, "invalid").Return(domain.Warehouse{}, warehouse.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodPatch, "/warehouses/invalid", `{"address":"Rua Pedro Dias","telephone":"371928"}`)

		server.PATCH(BaseEndpointWithId, handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)
		
		mockService.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(domain.Warehouse{}, warehouse.ErrInvalidBody)
		
		request, response := testutil.MakeRequest(http.MethodPatch, BaseEndpointWithIdNumber, "")

		server.PATCH(BaseEndpointWithId, handler.Update())
		server.ServeHTTP(response, request)
		
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 404 when warehouse is not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPatch, BaseEndpointWithIdNumber, `{"address":"Rua Pedro Dias","telephone":"371928"}`)

		mockService.On("Update", mock.Anything, mock.Anything, 1).Return(domain.Warehouse{}, warehouse.ErrNotFound)

		server.PATCH(BaseEndpointWithId, handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithWarehouses(t)

		request, response := testutil.MakeRequest(http.MethodPatch, BaseEndpointWithIdNumber, `{"telephone":"371928"}`)

		mockService.On("Update", mock.Anything, mock.Anything, 1).Return(domain.Warehouse{}, warehouse.ErrTryAgain)

		server.PATCH(BaseEndpointWithId, handler.Update())
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
