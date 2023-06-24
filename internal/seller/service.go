package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

var (
	ErrNotFound     = errors.New("seller not found")
	ErrInvalidId    = errors.New("invalid id")
	ErrInvalidBody  = errors.New("invalid body")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrAlredyExists = errors.New("seller already exists")
	ErrCidAlreadyExists = errors.New("cid already registered")
	ErrSaveSeller       = errors.New("error saving seller")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Save(ctx context.Context, d domain.Seller) (domain.Seller, error)
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
	if err != nil {
		return []domain.Seller{}, err
	}
	return sellers, nil
}

func (s *sellerService) Save(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	if s.repository.Exists(ctx, seller.CID) {
		return domain.Seller{}, ErrCidAlreadyExists
	}
	sellerId, err := s.repository.Save(ctx, seller)
	if err != nil {
		return domain.Seller{}, ErrSaveSeller
	}
	seller.ID = sellerId
	return seller, nil
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
	if !s.repository.Exists(ctx, d.CID) {
		return errors.New("cannot modify CID")
	}
	err := s.repository.Update(ctx, d)
	return err
}
