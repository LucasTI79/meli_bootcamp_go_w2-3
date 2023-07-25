package inbound_order_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/inbound_order"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/inbound_order"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var id = 2

var expectedInboundOrder = domain.InboundOrders{
	OrderDate:      "01/01/01",
	OrderNumber:    "001",
	EmployeeID:     1,
	ProductBatchID: 1,
	WarehouseID:    1,
}

func TestCreateInboundOrders(t *testing.T) {
	t.Run("Should create the inbound order if it contains the required fields", func(t *testing.T) {
		repository, service := InitServerWithInboundOrdersRepository(t)
		repository.On("Exists", mock.Anything, "001").Return(false)
		repository.On("Create", mock.Anything, expectedInboundOrder).Return(id, nil)

		inbound_order, err := service.Create(context.TODO(), expectedInboundOrder)

		assert.Equal(t, 2, inbound_order.ID)
		assert.Equal(t, "01/01/01", inbound_order.OrderDate)
		assert.Equal(t, "001", inbound_order.OrderNumber)
		assert.Equal(t, 1, inbound_order.EmployeeID)
		assert.Equal(t, 1, inbound_order.ProductBatchID)
		assert.Equal(t, 1, inbound_order.WarehouseID)

		assert.NoError(t, err)
	})
	t.Run("Should return err inbound order already exists when inbound order already exists", func(t *testing.T) {
		expectedMessage := "already exists"

		repository, service := InitServerWithInboundOrdersRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(true)

		_, err := service.Create(context.TODO(), domain.InboundOrders{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})

	t.Run("Should return error when there is an save repository error", func(t *testing.T) {
		repository, service := InitServerWithInboundOrdersRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(false)

		expectedError := errors.New("some error")
		repository.On("Create", mock.Anything, domain.InboundOrders{}).Return(0, expectedError)

		_, err := service.Create(context.TODO(), domain.InboundOrders{})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetInboundOrders(t *testing.T) {
	t.Run("Should get the inbound order when it exists in database", func(t *testing.T) {
		repository, service := InitServerWithInboundOrdersRepository(t)
		repository.On("Get", mock.Anything, id).Return(expectedInboundOrder, nil)

		inbound_order, err := service.Get(context.TODO(), 2)

		assert.Equal(t, expectedInboundOrder, inbound_order)
		assert.NoError(t, err)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		repository, service := InitServerWithInboundOrdersRepository(t)

		expectedError := errors.New("inbound orders not found")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.InboundOrders{}, inbound_order.ErrNotFound)

		_, err := service.Get(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("Should return error when there is an get repository error", func(t *testing.T) {
		repository, service := InitServerWithInboundOrdersRepository(t)

		expectedError := errors.New("inbound orders not found")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.InboundOrders{}, expectedError)

		_, err := service.Get(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestReportByAllInboundOrders(t *testing.T) {
	t.Run("Should return the inbound orders report when repository is called", func(t *testing.T) {

		expectedReport := []domain.InboundOrdersReport{
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

		repository, service := InitServerWithInboundOrdersRepository(t)
		repository.On("ReportByAll", mock.Anything).Return(expectedReport, nil)

		reports, err := service.ReportByAll(context.TODO())

		assert.True(t, len(reports) == 2)
		assert.NoError(t, err)
	})
}

func TestReportByOneInboundOrders(t *testing.T) {
	t.Run("Should return the inbound order when it exists", func(t *testing.T) {

		expectedReport := domain.InboundOrdersReport{
			ID:                 1,
			CardNumberID:       "1",
			FirstName:          "Joana",
			LastName:           "Costa",
			WarehouseID:        01,
			InboundOrdersCount: 001,
		}
		repository, service := InitServerWithInboundOrdersRepository(t)

		repository.On("ReportByOne", mock.Anything, mock.Anything).Return(expectedReport, nil)

		report, err := service.ReportByOne(context.TODO(), 1)

		assert.Equal(t, expectedReport, report)
		assert.NoError(t, err)
	})
	t.Run("Should return an error when the inbound orders does not exists", func(t *testing.T) {
		expectedEmpityReport := domain.InboundOrdersReport{}
		repository, service := InitServerWithInboundOrdersRepository(t)

		expectedError := errors.New("inbound orders not found")
		repository.On("ReportByOne", mock.Anything, mock.Anything).Return(expectedEmpityReport, inbound_order.ErrNotFound)
		_, err := service.ReportByOne(context.TODO(), 1)
		assert.Equal(t, expectedError, err)
		assert.Error(t, err)
	})
}

func InitServerWithInboundOrdersRepository(t *testing.T) (*mocks.InboundOrderRepositoryMock, inbound_order.Service) {
	t.Helper()
	mockRepositoryInboundOrders := &mocks.InboundOrderRepositoryMock{}
	mockService := inbound_order.NewService(mockRepositoryInboundOrders)
	return mockRepositoryInboundOrders, mockService
}
