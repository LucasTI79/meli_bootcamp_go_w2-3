package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

var (
	ErrNotFound     = errors.New("seller not found")
	ErrInvalidId    = errors.New("invalid id")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrAlredyExists = errors.New("seller already exists")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Save(ctx context.Context, d domain.Seller) (int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, s domain.Seller) error
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

func (s *sellerService) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	return err

}

func (s *sellerService) Get(ctx context.Context, id int) (domain.Seller, error) {
	seller, err := s.repository.Get(ctx, id)
	return seller, err
}

func (s *sellerService) Update(ctx context.Context, d domain.Seller) error {
	err := s.repository.Update(ctx, d)
	return err
}
