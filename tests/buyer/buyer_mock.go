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
	args := m.Called()
	return domain.Buyer{}, args.Error(1)
}

func (m *BuyerServiceMock) Update(ctx context.Context, b domain.Buyer) error {
	args := m.Called()
	return args.Error(1)
}

func (m *BuyerServiceMock) Delete(ctx context.Context, id int) error {
	return nil
}

func (m *BuyerServiceMock) Create(ctx context.Context, b domain.Buyer) (int, error) {
	args := m.Called()
	return 0, args.Error(1)
}
