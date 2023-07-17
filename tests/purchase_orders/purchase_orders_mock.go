package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type PurchaseOrdersServiceMock struct {
	mock.Mock
}

type PurchaseOrdersRepositoryMock struct {
	mock.Mock
}

func (m *PurchaseOrdersServiceMock) Create(ctx context.Context, o domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	args := m.Called(ctx, o)
	return args.Get(0).(domain.PurchaseOrders), args.Error(1)
}

func (m *PurchaseOrdersRepositoryMock) ExistsOrder(ctx context.Context, ordernumber string) bool {
	args := m.Called(ctx, ordernumber)
	return args.Get(0).(bool)
}

func (m *PurchaseOrdersRepositoryMock) Save(ctx context.Context, o domain.PurchaseOrders) error {
	args := m.Called(ctx, o)
	return args.Error(1)
}
