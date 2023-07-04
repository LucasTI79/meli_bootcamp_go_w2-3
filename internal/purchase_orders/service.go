package purchase_orders

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
	GetAll(ctx context.Context) ([]domain.PurchaseOrders, error)
	Get(ctx context.Context, id int) (domain.PurchaseOrders, error)
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

func (o *purchaseordersService) GetAll(ctx context.Context) ([]domain.PurchaseOrders, error) {
	var purchaseorders []domain.PurchaseOrders
	return purchaseorders, nil
}

func (o *purchaseordersService) Get(ctx context.Context, id int) (domain.PurchaseOrders, error) {
	var purchaseorders domain.PurchaseOrders
	return purchaseorders, nil
}

func (o *purchaseordersService) Create(ctx context.Context, p domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	var purchaseorders domain.PurchaseOrders
	return purchaseorders, nil
}