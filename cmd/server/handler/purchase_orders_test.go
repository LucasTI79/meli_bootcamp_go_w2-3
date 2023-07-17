package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocksBuyer "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/buyer"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/purchase_orders"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BuyerServiceMocks struct {
	BuyerServiceMock          *mocksBuyer.BuyerServiceMock
	PurchaseOrdersServiceMock *mocks.PurchaseOrdersServiceMock
}

const (
	CreateOrders = "/purchaseOrders"
)

func TestCreateOrders(t *testing.T) {
	t.Run("Should return status 200 and the order created", func(t *testing.T) {
		server, handler, mocks := InitServerWithGetPurchaseOrders(t)
		server.POST(CreateOrders, handler.CreateOrders())

		purchaseOrders := domain.PurchaseOrders{
			OrderNumber:     "order#1000",
			OrderDate:       "2021-04-04",
			TrackingCode:    "abscf123",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		jsonOrders, _ := json.Marshal(purchaseOrders)
		request, response := testutil.MakeRequest("POST", CreateOrders, string(jsonOrders))

		mocks.PurchaseOrdersServiceMock.On("Create", mock.Anything, mock.Anything).Return(purchaseOrders, nil)
		mocks.BuyerServiceMock.On("ExistsID", mock.Anything, 1).Return(nil)

		server.ServeHTTP(response, request)

		responseResult := domain.PurchaseOrdersResponseID{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, purchaseOrders, responseResult.Data)
	})

	t.Run("Should return err invalid body", func(t *testing.T) {
		server, handler, mocks := InitServerWithGetPurchaseOrders(t)
		server.POST(CreateOrders, handler.CreateOrders())

		jsonOrders, _ := json.Marshal(&domain.PurchaseOrders{})
		request, response := testutil.MakeRequest("POST", CreateOrders, string(jsonOrders))

		mocks.PurchaseOrdersServiceMock.On("Create", mock.Anything, mock.Anything).Return(0, nil)
		mocks.BuyerServiceMock.On("ExistsID", mock.Anything, 0).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("Should return status 400 when JSON is invalid", func(t *testing.T) {
		server, handler, mocks := InitServerWithGetPurchaseOrders(t)
		server.POST(CreateOrders, handler.CreateOrders())

		request, response := testutil.MakeRequest(http.MethodPost, CreateOrders, `{"order_number":2}`)

		mocks.PurchaseOrdersServiceMock.On("Create", mock.Anything, mock.Anything).Return(0, nil)
		mocks.BuyerServiceMock.On("ExistsID", mock.Anything, 0).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return status 409 when buyer not exists", func(t *testing.T) {
		server, handler, mocks := InitServerWithGetPurchaseOrders(t)
		server.POST(CreateOrders, handler.CreateOrders())

		purchaseOrders := domain.PurchaseOrders{
			OrderNumber:     "order#1000",
			OrderDate:       "2021-04-04",
			TrackingCode:    "abscf123",
			BuyerID:         100,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		jsonpurchaseOrders, _ := json.Marshal(purchaseOrders)
		request, response := testutil.MakeRequest(http.MethodPost, CreateOrders, string(jsonpurchaseOrders))

		mocks.BuyerServiceMock.On("ExistsID", mock.Anything, 100).Return(errors.New("error"))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return status 500 when internal error", func(t *testing.T) {
		server, handler, mocks := InitServerWithGetPurchaseOrders(t)
		server.POST(CreateOrders, handler.CreateOrders())

		purchaseOrders := domain.PurchaseOrders{
			OrderNumber:     "order#1000",
			OrderDate:       "2021-04-04",
			TrackingCode:    "abscf123",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		jsonpurchaseOrders, _ := json.Marshal(purchaseOrders)
		request, response := testutil.MakeRequest(http.MethodPost, CreateOrders, string(jsonpurchaseOrders))

		mocks.PurchaseOrdersServiceMock.On("Create", mock.Anything, mock.Anything).Return(domain.PurchaseOrders{}, errors.New("error"))
		mocks.BuyerServiceMock.On("ExistsID", mock.Anything, 1).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func InitServerWithGetPurchaseOrders(t *testing.T) (*gin.Engine, *handler.PurchaseOrdersController, BuyerServiceMocks) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.PurchaseOrdersServiceMock)
	mockServiceBuyer := new(mocksBuyer.BuyerServiceMock)
	handler := handler.NewPurchaseOrders(mockService, mockServiceBuyer)
	return server, handler, BuyerServiceMocks{
		BuyerServiceMock:          mockServiceBuyer,
		PurchaseOrdersServiceMock: mockService,
	}
}
