package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductRecordServiceMock struct {
	mock.Mock
}

type ProductRecordRepositoryMock struct {
	mock.Mock
}

func (p *ProductRecordServiceMock) Save(ctx context.Context, d domain.ProductRecord) (int, error) {
	args := p.Called(d)
	return args.Int(0), args.Error(1)
}

func (p *ProductRecordRepositoryMock) Save(ctx context.Context, d domain.ProductRecord) (int, error) {
	args := p.Called(d)
	return args.Int(0), args.Error(1)
}

func (p *ProductRecordServiceMock) RecordsByAllProductsReport(ctx context.Context) ([]domain.ProductRecordReport, error) {
	args := p.Called()
	return args.Get(0).([]domain.ProductRecordReport), args.Error(1)
}

func (p *ProductRecordRepositoryMock) RecordsByAllProductsReport(ctx context.Context) ([]domain.ProductRecordReport, error) {
	args := p.Called()
	return args.Get(0).([]domain.ProductRecordReport), args.Error(1)
}

func (p *ProductRecordServiceMock) RecordsByOneProductReport(ctx context.Context, id int) (domain.ProductRecordReport, error) {
	args := p.Called(id)
	return args.Get(0).(domain.ProductRecordReport), args.Error(1)
}

func (p *ProductRecordRepositoryMock) RecordsByOneProductReport(ctx context.Context, id int) (domain.ProductRecordReport, error) {
	args := p.Called(id)
	return args.Get(0).(domain.ProductRecordReport), args.Error(1)
}
