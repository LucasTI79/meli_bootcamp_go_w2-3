package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

type ProductRepositoryMock struct {
	mock.Mock
}

// mock do controller Product recebe o mock da service
func (p *ProductServiceMock) GetAll(ctx context.Context) ([]domain.Product, error) {
	args := p.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

// Mock da Service Product recebe o mock da repository
func (p *ProductRepositoryMock) GetAll() ([]domain.Product, error) {
	args := p.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (p *ProductServiceMock) Update(ctx context.Context, d domain.Product) error {

	return nil
}

func (p *ProductServiceMock) Save(ctx context.Context, d domain.Product) (int, error) {

	return 0, nil
}

func (p *ProductServiceMock) Delete(ctx context.Context, id int) error {

	return nil
}

func (p *ProductServiceMock) Get(ctx context.Context, id int) (domain.Product, error) {
	args := p.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (p *ProductRepositoryMock) Get(ctx context.Context, id int) (domain.Product, error) {
	args := p.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}
