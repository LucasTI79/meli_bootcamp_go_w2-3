package buyer

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound  = errors.New("buyer not found")
	ErrExists    = errors.New("buyer already exists")
	ErrInvalidID = errors.New("invalid ID")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
	ExistsID(ctx context.Context, id int) error
	Create(ctx context.Context, b domain.Buyer) (domain.Buyer, error)
	Update(ctx context.Context, b domain.Buyer, id int) (domain.Buyer, error)
	Delete(ctx context.Context, id int) error
	GetBuyerOrders(ctx context.Context, id int) (domain.BuyerOrders, error)
	GetBuyersOrders(ctx context.Context) ([]domain.BuyerOrders, error)
}

type buyerService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &buyerService{
		repository: r,
	}
}

func (b *buyerService) GetBuyerOrders(ctx context.Context, id int) (domain.BuyerOrders, error) {
	buyerOrders, err := b.repository.GetBuyerOrders(ctx, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return domain.BuyerOrders{}, ErrNotFound
		}
		return domain.BuyerOrders{}, err
	}
	return buyerOrders, err
}

func (b *buyerService) GetBuyersOrders(ctx context.Context) ([]domain.BuyerOrders, error) {
	buyersOrders, err := b.repository.GetBuyersOrders(ctx)
	if err != nil {
		return buyersOrders, err
	}
	return buyersOrders, nil
}

func (b *buyerService) ExistsID(ctx context.Context, id int) error {
	buyerExists := b.repository.ExistsID(ctx, id)

	if !buyerExists {
		return errors.New("buyer not exists")
	}

	return nil
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
	userExist := b.repository.ExistsBuyer(ctx, d.CardNumberID)
	if userExist {
		return domain.Buyer{}, ErrExists
	}
	buyerId, err := b.repository.Save(ctx, d)
	d.ID = buyerId
	return d, err
}

func (b *buyerService) Update(ctx context.Context, d domain.Buyer, id int) (domain.Buyer, error) {
	buyer, err := b.repository.Get(ctx, id)
	if err != nil {
		return domain.Buyer{}, errors.New("error getting buyer")
	}
	if d.FirstName != "" {
		buyer.FirstName = d.FirstName
	}
	if d.LastName != "" {
		buyer.LastName = d.LastName
	}
	err = b.repository.Update(ctx, buyer)
	if err != nil {
		return domain.Buyer{}, ErrNotFound
	}

	return buyer, nil

}

func (b *buyerService) Delete(ctx context.Context, id int) error {
	err := b.repository.Delete(ctx, id)
	return err
}
