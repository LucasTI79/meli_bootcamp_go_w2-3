package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SellerServiceMock struct {
	mock.Mock
}

type LocalityServiceMock struct {
	mock.Mock
}

type SellerRepositoryMock struct {
	mock.Mock
}

func (s *SellerServiceMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
	args := s.Called()
	return args.Get(0).([]domain.Seller), args.Error(1)
}

func (s *SellerRepositoryMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
	args := s.Called()
	return args.Get(0).([]domain.Seller), args.Error(1)
}

func (s *SellerServiceMock) Save(ctx context.Context, d domain.Seller) (domain.Seller, error) {
	args := s.Called(ctx, d)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *SellerRepositoryMock) Save(ctx context.Context, d domain.Seller) (int, error) {
	args := s.Called(ctx, d)
	return args.Get(0).(int), args.Error(1)
}

func (s *SellerServiceMock) Delete(ctx context.Context, id int) error {
	args := s.Called(ctx, id)
	return args.Error(0)
}

func (s *SellerRepositoryMock) Delete(ctx context.Context, id int) error {
	args := s.Called(ctx, id)
	return args.Error(0)
}

func (s *SellerServiceMock) Get(ctx context.Context, id int) (domain.Seller, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *SellerRepositoryMock) Get(ctx context.Context, id int) (domain.Seller, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *SellerRepositoryMock) Update(ctx context.Context, d domain.Seller) error {
	args := s.Called(ctx, d)
	return args.Error(0)
}

func (s *SellerServiceMock) Update(ctx context.Context, id int, d domain.Seller) (domain.Seller, error) {
	args := s.Called(ctx, d, id)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *SellerRepositoryMock) Exists(ctx context.Context, cid int) bool {
	args := s.Called(ctx, cid)
	return args.Get(0).(bool)
}

func (l *LocalityServiceMock) ExistsById(ctx context.Context, id int) error {
	args := l.Called(ctx, id)
	return args.Error(0)
}

func (l *LocalityServiceMock) ReportSellersByLocality(ctx context.Context, id int) ([]domain.LocalityReport, error) {
	args := l.Called(ctx)
	return args.Get(0).([]domain.LocalityReport), args.Error(1)
}

func (l *LocalityServiceMock) Save(ctx context.Context, d domain.Locality) (domain.LocalityInput, error) {
	args := l.Called(ctx, d)
	return args.Get(0).(domain.LocalityInput), args.Error(1)
}