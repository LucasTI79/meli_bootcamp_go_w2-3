package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type InboundOrderServiceMock struct {
	mock.Mock
}

type InboundOrderRepositoryMock struct {
	mock.Mock
}

func (m *InboundOrderServiceMock) Read(ctx context.Context, id int) ([]domain.Employee, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]domain.Employee), args.Error(1)
}

func (m *InboundOrderRepositoryMock) ReadAllInboundOrders(ctx context.Context) ([]domain.Employee, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Employee), args.Error(1)
}

func (m *InboundOrderServiceMock) Create(ctx context.Context, c domain.InboundOrders) (domain.InboundOrders, error) {
	args := m.Called(ctx, c)
	return args.Get(0).(domain.InboundOrders), args.Error(1)
}

func (m *InboundOrderRepositoryMock) Create(ctx context.Context, c domain.InboundOrders) (int, error) {
	args := m.Called(ctx, c)
	return args.Get(0).(int), args.Error(1)
}

func (m *InboundOrderServiceMock) Get(ctx context.Context, id int) (domain.InboundOrders, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.InboundOrders), args.Error(1)
}

func (m *InboundOrderRepositoryMock) Get(ctx context.Context, id int) (domain.InboundOrders, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.InboundOrders), args.Error(1)
}

func (m *InboundOrderRepositoryMock) ExistsByCidInboundOrder(ctx context.Context, cid string) bool {
	args := m.Called(ctx, cid)
	return args.Get(0).(bool)
}
