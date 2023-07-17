package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type BuyerServiceMock struct {
	mock.Mock
}

type BuyerRepositoryMock struct {
	mock.Mock
}

func (m *BuyerServiceMock) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]domain.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) Get(ctx context.Context, id int) (domain.Buyer, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) Update(ctx context.Context, b domain.Buyer, id int) (domain.Buyer, error) {
	args := m.Called(ctx, b, id)
	return args.Get(0).(domain.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *BuyerServiceMock) Create(ctx context.Context, b domain.Buyer) (domain.Buyer, error) {
	args := m.Called(ctx, b)
	return args.Get(0).(domain.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) ExistsID(ctx context.Context, buyerID int) error {
	args := m.Called(ctx, buyerID)
	return args.Error(0)
}

func (m *BuyerServiceMock) GetBuyerOrders(ctx context.Context, id int) (domain.BuyerOrders, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.BuyerOrders), args.Error(1)
}

func (m *BuyerServiceMock) GetBuyersOrders(ctx context.Context) ([]domain.BuyerOrders, error) {
	args := m.Called()
	return args.Get(0).([]domain.BuyerOrders), args.Error(1)
}

func (m *BuyerRepositoryMock) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]domain.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) GetBuyersOrders(ctx context.Context) ([]domain.BuyerOrders, error) {
	args := m.Called()
	return args.Get(0).([]domain.BuyerOrders), args.Error(1)
}

func (m *BuyerRepositoryMock) Get(ctx context.Context, id int) (domain.Buyer, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) GetBuyerOrders(ctx context.Context, id int) (domain.BuyerOrders, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.BuyerOrders), args.Error(1)
}

func (m *BuyerRepositoryMock) Update(ctx context.Context, b domain.Buyer) error {
	args := m.Called(ctx, b)
	return args.Error(0)
}

func (m *BuyerRepositoryMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *BuyerRepositoryMock) Create(ctx context.Context, b domain.Buyer) (domain.Buyer, error) {
	args := m.Called(ctx, b)
	return args.Get(0).(domain.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) ExistsBuyer(ctx context.Context, cardnumber string) bool {
	args := m.Called(ctx, cardnumber)
	return args.Get(0).(bool)
}

func (m *BuyerRepositoryMock) ExistsID(ctx context.Context, buyerID int) bool {
	args := m.Called(buyerID)
	return args.Get(0).(bool)
}

func (m *BuyerRepositoryMock) Save(ctx context.Context, b domain.Buyer) (int, error) {
	args := m.Called(ctx, b)
	return args.Get(0).(int), args.Error(1)
}
