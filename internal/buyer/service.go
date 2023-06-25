package buyer

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
	ErrExists   = errors.New("buyer already exists")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
	Create(ctx context.Context, b domain.Buyer) (domain.Buyer, error)
	Update(ctx context.Context, b domain.Buyer) error
	Delete(ctx context.Context, id int) error
}

type buyerService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &buyerService{
		repository: r,
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
		if err.Error() == "sql: no rows in result set" {
			return domain.Buyer{}, ErrNotFound
		}
		return domain.Buyer{}, err
	}
	return buyer, err
}

func (b *buyerService) Create(ctx context.Context, d domain.Buyer) (domain.Buyer, error) {
	userExist := b.repository.Exists(ctx, d.CardNumberID)
	if userExist {
		return domain.Buyer{}, ErrExists
	}
	buyerId, err := b.repository.Save(ctx, d)
	d.ID = buyerId
	return d, err
}

func (b *buyerService) Update(ctx context.Context, d domain.Buyer) error {
	userExist := b.repository.Exists(ctx, d.CardNumberID)
	if !userExist {
		return ErrNotFound
	}
	err := b.repository.Update(ctx, d)
	return err

}

func (b *buyerService) Delete(ctx context.Context, id int) error {
	err := b.repository.Delete(ctx, id)
	return err
}
