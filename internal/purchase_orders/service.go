package purchase_orders

import (
	"context"
	"errors"
	"fmt"

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
	GetAll(ctx context.Context) ([]domain.PurchaseOrders, error)
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

func (s *purchaseordersService) GetAll(ctx context.Context) ([]domain.PurchaseOrders, error) {
	fmt.Println("CHEGOU SERVICE")
	orders, err := s.repository.GetAll(ctx)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *purchaseordersService) Create(ctx context.Context, o domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	orderExists := s.repository.ExistsOrder(ctx, o.OrderNumber)
	if !orderExists {
		fmt.Println("order exist")
		return domain.PurchaseOrders{}, ErrExists
	}
	err := s.repository.Save(ctx, o)
	fmt.Println("buyer exist")
	return o, err
}
