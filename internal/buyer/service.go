package buyer

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type Service interface{	
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
}

type buyerService struct{
	repository Repository
}

func NewService(r Repository) Service {
	return &buyerService{
		repository : r,
	}
}

func (b *buyerService) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	buyers, err := b.repository.GetAll(ctx)
	if err != nil {
		return buyers, err
	}
	return buyers, nil
}

func (b *buyerService) Get(ctx context.Context, id int) (domain.Buyer, error) {
	buyer, err := b.repository.Get(ctx, id)
	if err != nil {
		if err.Error()=="sql: no rows in result set"{
			return domain.Buyer{}, ErrNotFound
		}
		return domain.Buyer{}, err
	}
	return buyer, err
}