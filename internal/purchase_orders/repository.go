package purchase_orders

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Repository encapsulates the storage of a purchased order.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.PurchaseOrders, error)
	Get(ctx context.Context, id int) (domain.PurchaseOrders, error)
	ExistsOrder(ctx context.Context, orderNumber string) bool
	ExistsBuyer(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, o domain.PurchaseOrders) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.PurchaseOrders, error) {
	var purchaseorders []domain.PurchaseOrders
	return purchaseorders, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.PurchaseOrders, error) {
	var purchaseorders domain.PurchaseOrders
	return purchaseorders, nil
}

func (r *repository) ExistsOrder(ctx context.Context, orderNumber string) bool {
	return false
}

func (r *repository) ExistsBuyer(ctx context.Context, cardNumberID string) bool {
	return false
}

func (r *repository) Save(ctx context.Context, o domain.PurchaseOrders) (int, error) {
	return 0, nil
}