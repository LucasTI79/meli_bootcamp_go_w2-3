package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SellerServiceMock struct {
    mock.Mock
}

type SellerControllerMock struct {
    mock.Mock
}

func (s *SellerServiceMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
	args := s.Called()
    return args.Get(0).([]domain.Seller), args.Error(1)
}

// func (s *SellerControllerMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
// 	args := s.Called()
//     return args.Get(0).([]domain.Seller), args.Error(1)
// }

func (s *SellerServiceMock) Save(ctx context.Context, d domain.Seller) (domain.Seller, error) {
	args := s.Called(ctx, s)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *SellerServiceMock) Delete(ctx context.Context, id int) error {
	args := s.Called(ctx, id)
	return args.Error(0)
}

func (s *SellerServiceMock) Get(ctx context.Context, id int) (domain.Seller, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *SellerServiceMock) Update(ctx context.Context, d domain.Seller) error {
	return nil
}
