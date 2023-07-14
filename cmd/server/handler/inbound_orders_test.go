package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
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
	t.Run("Should return status 200 and the inbound order created", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		mockService.On("Create", mock.Anything, mock.Anything).Return(expectedInboundOrder, nil)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

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

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when OrderName is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Employee ID is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when ProductBatchID is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when WarehouseID is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithInboundOrders(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 409 when inbound order already exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.InboundOrders")).Return(domain.InboundOrders{}, inbound_order.ErrAlredyExists)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithInboundOrders(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointInboundOrders, `{"order_date":"01/01/01","order_name":"name one","employee_id":"001","product_batch_id":"01010101","warehouse_id":"01"}`)

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.InboundOrders")).Return(domain.InboundOrders{}, inbound_order.ErrTryAgain)

		server.POST(BaseEndpointInboundOrders, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func InitServerWithInboundOrders(t *testing.T) (*gin.Engine, *mocks.InboundOrderServiceMock, *handler.InboundOrdersController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.InboundOrderServiceMock)
	handler := handler.NewInboundOrders(mockService)
	return server, mockService, handler
}
