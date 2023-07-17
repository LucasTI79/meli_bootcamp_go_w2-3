package purchase_orders_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/purchase_orders"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/purchase_orders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	t.Run("Should create the purchased order", func(t *testing.T) {
		id := 10
		expectedOrder := domain.PurchaseOrders{
			OrderNumber:     "9423i",
			OrderDate:       "2021-04-04",
			TrackingCode:    "afijaehn",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		repository, service := InitServerWithPurchaseOrdersRepository(t)
		repository.On("ExistsOrder", mock.Anything, "9423i").Return(false)
		repository.On("Save", mock.Anything, expectedOrder).Return(id, nil)

		order, err := service.Create(context.TODO(), expectedOrder)

		assert.Equal(t, "9423i", order.OrderNumber)
		assert.Equal(t, "2021-04-04", order.OrderDate)
		assert.Equal(t, "afijaehn", order.TrackingCode)

		assert.NoError(t, err)
	})
	t.Run("Should return err order_number already exists", func(t *testing.T) {
		expectedMessage := "order already exists"
		repository, service := InitServerWithPurchaseOrdersRepository(t)
		repository.On("ExistsOrder", mock.Anything, mock.Anything).Return(true)

		_, err := service.Create(context.TODO(), domain.PurchaseOrders{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func InitServerWithPurchaseOrdersRepository(t *testing.T) (*mocks.PurchaseOrdersRepositoryMock, purchase_orders.Service) {
	t.Helper()
	mockRepository := &mocks.PurchaseOrdersRepositoryMock{}
	mockService := purchase_orders.NewService(mockRepository)
	return mockRepository, mockService
}
