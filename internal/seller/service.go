package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Save(ctx context.Context, d domain.Seller) (int, error)
}

type sellerService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &sellerService{
		repository: r,
	}
}

func (s *sellerService) GetAll(ctx context.Context) ([]domain.Seller, error) {
	sellers, err := s.repository.GetAll(ctx)
	return sellers, err
}

func (s *sellerService) Save(ctx context.Context, d domain.Seller) (int, error) {
	userExist := s.repository.Exists(ctx, d.CID)
	if userExist {
		return 0, errors.New("user already exists")
	}
	sellerId, err := s.repository.Save(ctx, d)
	return sellerId, err
}
