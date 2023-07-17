package handler_test

import (
	"encoding/json"
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

		purchaseOrders := domain.PurchaseOrders{
			ID:              1,
			OrderNumber:     "order#1000",
			OrderDate:       "2021-04-04",
			TrackingCode:    "abscf123",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		mocks.PurchaseOrdersServiceMock.On("Create", mock.Anything, mock.Anything).Return(purchaseOrders, nil)
		mocks.BuyerServiceMock.On("ExistsID", mock.Anything, 1).Return(nil)

		request, response := testutil.MakeRequest(http.MethodPost, CreateOrders, `{
			"order_number": "order#1000",
			"order_date": "2021-04-04",
			"tracking_code": "abscf123",
			"buyer_id": 1,
			"product_record_id": 1,
			"order_status_id": 1
			}:`)

		server.POST(CreateOrders, handler.CreateOrders())
		server.ServeHTTP(response, request)

		responseResult := domain.PurchaseOrdersResponseID{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, purchaseOrders, responseResult.Data)
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
