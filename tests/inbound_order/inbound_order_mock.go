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

func (m *InboundOrderRepositoryMock) Exists(ctx context.Context, orderNumber string) bool {
	args := m.Called(ctx, orderNumber)
	return args.Get(0).(bool)
}

func (m *InboundOrderServiceMock) ReportByAll(ctx context.Context) ([]domain.InboundOrdersReport, error) {
	args := m.Called()
	return args.Get(0).([]domain.InboundOrdersReport), args.Error(1)
}

func (m *InboundOrderRepositoryMock) ReportByAll(ctx context.Context) ([]domain.InboundOrdersReport, error) {
	args := m.Called()
	return args.Get(0).([]domain.InboundOrdersReport), args.Error(1)
}

func (m *InboundOrderServiceMock) ReportByOne(ctx context.Context, id int) (domain.InboundOrdersReport, error) {
	args := m.Called(id)
	return args.Get(0).(domain.InboundOrdersReport), args.Error(1)
}

func (m *InboundOrderRepositoryMock) ReportByOne(ctx context.Context, id int) (domain.InboundOrdersReport, error) {
	args := m.Called(id)
	return args.Get(0).(domain.InboundOrdersReport), args.Error(1)
}
