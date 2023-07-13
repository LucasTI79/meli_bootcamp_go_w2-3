package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductBatchServiceMock struct {
	mock.Mock
}

type ProductBatchRepositoryMock struct {
	mock.Mock
}

func (m *ProductBatchServiceMock) Save(ctx context.Context, p domain.ProductBatch) (int, error) {
	args := m.Called(ctx, p)
	return args.Int(0), args.Error(1)
}

func (m *ProductBatchRepositoryMock) Save(produsctBatch domain.ProductBatch) (int, error) {
	args := m.Called(produsctBatch)
	return args.Int(0), args.Error(1)
}
