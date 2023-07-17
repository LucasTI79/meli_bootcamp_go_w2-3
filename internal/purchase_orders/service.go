package purchase_orders

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound  = errors.New("order not found")
	ErrExists    = errors.New("order already exists")
	ErrInvalidID = errors.New("invalid ID")
	ErrConflict  = errors.New("buyer not found")
)

type Service interface {
	Create(ctx context.Context, o domain.PurchaseOrders) (domain.PurchaseOrders, error)
}

type purchaseordersService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &purchaseordersService{
		repository: r,
	}
}

func (s *purchaseordersService) Create(ctx context.Context, o domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	orderExists := s.repository.ExistsOrder(ctx, o.OrderNumber)
	if orderExists {
		return domain.PurchaseOrders{}, ErrExists
	} else {
		order := s.repository.Save(ctx, o)
		return o, order
	}
}
