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

func (p *ProductServiceMock) ExistsById(productID int) error {
	args := p.Called(productID)
	return args.Error(0)
}

func (p *ProductRepositoryMock) ExistsById(sectionID int) bool {
	args := p.Called(sectionID)
	return args.Get(0).(bool)
}

// mock do controller Product recebe o mock da service
func (p *ProductServiceMock) GetAll(ctx context.Context) ([]domain.Product, error) {
	args := p.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

// Mock da Service Product recebe o mock da repository
func (p *ProductRepositoryMock) GetAll(ctx context.Context) ([]domain.Product, error) {
	args := p.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (p *ProductServiceMock) Save(ctx context.Context, d domain.Product) (int, error) {
	args := p.Called(d)
	return args.Int(0), args.Error(1)
}

func (p *ProductRepositoryMock) Save(ctx context.Context, d domain.Product) (int, error) {
	args := p.Called(d)
	return args.Int(0), args.Error(1)
}

func (p *ProductServiceMock) Delete(ctx context.Context, id int) error {
	args := p.Called(id)
	return args.Error(0)
}

func (p *ProductRepositoryMock) Delete(ctx context.Context, id int) error {
	args := p.Called(id)
	return args.Error(0)
}

func (p *ProductServiceMock) Get(ctx context.Context, id int) (domain.Product, error) {
	args := p.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (p *ProductRepositoryMock) Get(ctx context.Context, id int) (domain.Product, error) {
	args := p.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (p *ProductServiceMock) Update(ctx context.Context, d domain.Product) error {
	args := p.Called(ctx, d)
	return args.Error(0)

}

func (p *ProductRepositoryMock) Update(ctx context.Context, d domain.Product) error {

	args := p.Called(ctx, d)
	return args.Error(0)
}

func (p *ProductRepositoryMock) Exists(ctx context.Context, productCode string) bool {
	args := p.Called(ctx, productCode)
	return args.Get(0).(bool)
}
