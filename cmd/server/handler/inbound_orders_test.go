package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/middlewares"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/inbound_order"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/inbound_order"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	BaseEndpointInboundOrders       = "/inbound-orders"
	BaseEndpointWithIdInboundOrders = "/inbound-orders/:id"
)

var expectedInboundOrder = domain.InboundOrders{
	OrderDate:      "01/01/01",
	OrderNumber:    "001",
	EmployeeID:     1,
	ProductBatchID: 1,
	WarehouseID:    1,
}

func TestGetInboundOrders(t *testing.T) {
	t.Run("Should return status 200 and inbound order with id", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		mockService.On("Get", mock.Anything, 1).Return(expectedInboundOrder, nil)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointInboundOrders+"/1", "")

		server.GET(BaseEndpointWithIdInboundOrders, handler.Get())
		server.ServeHTTP(response, request)

		responseResult := &domain.InboundOrdersResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedInboundOrder, responseResult.Data)
	})
	t.Run("Should return status 400 when the inbound order id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		mockService.On("Get", mock.Anything, "invalid").Return(domain.InboundOrders{}, inbound_order.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointInboundOrders+"/invalid", "")

		server.GET(BaseEndpointWithIdInboundOrders, handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when the inbound order id does not exist", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.InboundOrders{}, inbound_order.ErrNotFound)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointInboundOrders+"/1", "")

		server.GET(BaseEndpointInboundOrders, handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {

		server, mockService, handler := InitServerWithInboundOrders(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.InboundOrders{}, inbound_order.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointInboundOrders+"/1", "")

		server.GET(BaseEndpointWithIdInboundOrders, handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCreateInboundOrders(t *testing.T) {
	newInboudOrders := domain.InboundOrders{
		OrderDate:      "2021-01-01",
		OrderNumber:    "002",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	t.Run("Should return status 201 and the inbound order created", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		mockService.On("Create", mock.Anything, mock.Anything).Return(expectedInboundOrder, nil)
		jsonProductBatch, _ := json.Marshal(newInboudOrders)
		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatch))

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		responseResult := domain.InboundOrdersResponseId{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, expectedInboundOrder, responseResult.Data)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"address":}`)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 400 when Order Date is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)
		newInboudOrdersInvalid := domain.InboundOrders{
			ID:             1,
			OrderDate:      "1",
			OrderNumber:    "1",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		}
		jsonProductBatchInvalid, _ := json.Marshal(newInboudOrdersInvalid)
		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatchInvalid))

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when OrderNumber is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)
		newInboudOrdersInvalid := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "2020-01-02",
			OrderNumber:    "",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		}
		jsonProductBatchInvalid, _ := json.Marshal(newInboudOrdersInvalid)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatchInvalid))

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Employee ID is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)
		newInboudOrdersInvalid := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "2020-01-02",
			OrderNumber:    "1",
			EmployeeID:     0,
			ProductBatchID: 1,
			WarehouseID:    1,
		}
		jsonProductBatchInvalid, _ := json.Marshal(newInboudOrdersInvalid)
		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatchInvalid))

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when ProductBatchID is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)
		newInboudOrdersInvalid := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "2020-01-02",
			OrderNumber:    "1",
			EmployeeID:     1,
			ProductBatchID: 0,
			WarehouseID:    1,
		}
		jsonProductBatchInvalid, _ := json.Marshal(newInboudOrdersInvalid)
		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatchInvalid))

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when WarehouseID is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)

		newInboudOrdersInvalid := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "2020-01-02",
			OrderNumber:    "1",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    0,
		}
		jsonProductBatchInvalid, _ := json.Marshal(newInboudOrdersInvalid)
		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatchInvalid))
		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 409 when inbound order already exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		jsonProductBatch, _ := json.Marshal(newInboudOrders)
		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatch))

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.InboundOrders")).Return(domain.InboundOrders{}, inbound_order.ErrAlredyExists)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		jsonProductBatch, _ := json.Marshal(newInboudOrders)
		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, string(jsonProductBatch))

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.InboundOrders")).Return(domain.InboundOrders{}, inbound_order.ErrTryAgain)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestReportsByAllInboundOrders(t *testing.T) {
	t.Run("Should return status 200 with inbound orders of all employees", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)
		expectedInboundOrders := []domain.InboundOrdersReport{
			{
				ID:                 1,
				CardNumberID:       "1",
				FirstName:          "Joana",
				LastName:           "Costa",
				WarehouseID:        01,
				InboundOrdersCount: 001,
			},
			{
				ID:                 2,
				CardNumberID:       "2",
				FirstName:          "Luiza",
				LastName:           "Silva",
				WarehouseID:        02,
				InboundOrdersCount: 001,
			},
		}

		server.GET("/employees/reportInboundOrders", handler.ReportByAll())
		request, response := testutil.MakeRequest(http.MethodGet, "/employees/reportInboundOrders", "")

		mockService.On("ReportByAll", mock.Anything).Return(expectedInboundOrders, nil)
		server.ServeHTTP(response, request)

		responseResult := &domain.InboundOrdersReport{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusOK, response.Code)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		var ExpectedEmptyReports = []domain.InboundOrdersReport{}

		server, mockService, handler := InitServerWithInboundOrders(t)

		server.GET("/employees/reportInboundOrders", handler.ReportByAll())

		request, response := testutil.MakeRequest(http.MethodGet, "/employees/reportInboundOrders", "")

		mockService.On("ReportByAll", mock.AnythingOfType(mock.Anything)).Return(ExpectedEmptyReports, inbound_order.ErrTryAgain)

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

	})

}

func TestReportsByOneInboundOrders(t *testing.T) {

	t.Run("Should return status 200 with the inbound order report", func(t *testing.T) {
		expectedInboundOrders := domain.InboundOrdersReport{
			ID:                 1,
			CardNumberID:       "1",
			FirstName:          "Joana",
			LastName:           "Costa",
			WarehouseID:        01,
			InboundOrdersCount: 001,
		}
		server, mockService, handler := InitServerWithInboundOrders(t)

		server.GET("/employees/reportInboundOrders/:id", middlewares.ValidateParams("id"), handler.ReportByOne())
		request, response := testutil.MakeRequest(http.MethodGet, "/employees/reportInboundOrders/1", "")

		mockService.On("ReportByOne", mock.Anything, mock.Anything).Return(expectedInboundOrders, nil)

		server.ServeHTTP(response, request)

		responseResult := &domain.InboundOrdersReportResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), responseResult)

		assert.Equal(t, expectedInboundOrders, responseResult.Data)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		ExpectedEmptyInboundOrder := domain.InboundOrdersReport{}

		server, mockService, handler := InitServerWithInboundOrders(t)

		server.GET("/employees/reportInboundOrders/:id", middlewares.ValidateParams("id"), handler.ReportByOne())
		mockService.On("ReportByOne", mock.Anything, mock.Anything).Return(ExpectedEmptyInboundOrder, inbound_order.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, "/employees/reportInboundOrders/1", "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("Should return status 404 when the report of a employee is not found", func(t *testing.T) {

		server, mockService, handler := InitServerWithInboundOrders(t)

		server.GET("/employees/reportInboundOrders/:id", middlewares.ValidateParams("id"), handler.ReportByOne())
		request, response := testutil.MakeRequest(http.MethodGet, "/employees/reportInboundOrders/1", "")

		mockService.On("ReportByOne", mock.Anything, mock.Anything).Return(domain.InboundOrdersReport{}, inbound_order.ErrNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return 400 when an employee Id is invalid", func(t *testing.T) {

		server, mockService, handler := InitServerWithInboundOrders(t)
		server.GET("/employees/reportInboundOrders/:id", middlewares.ValidateParams("id"), handler.ReportByOne())

		mockService.On("ReportByOne", mock.Anything).Return(domain.InboundOrdersReport{}, inbound_order.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodGet, "/employees/reportInboundOrders/invalidId", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func InitServerWithInboundOrders(t *testing.T) (*gin.Engine, *mocks.InboundOrderServiceMock, *handler.InboundOrdersController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.InboundOrderServiceMock)
	handler := handler.NewInboundOrders(mockService)
	return server, mockService, handler
}
