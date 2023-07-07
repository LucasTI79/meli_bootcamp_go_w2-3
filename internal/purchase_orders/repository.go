package purchase_orders

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Repository encapsulates the storage of a purchased order.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.PurchaseOrdersGetAll, error)
	ExistsOrder(ctx context.Context, orderNumber string) bool
	Save(ctx context.Context, o domain.PurchaseOrders) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.PurchaseOrdersGetAll, error) {
	rows, err := r.db.Query("SELECT * FROM purchase_orders")
	if err != nil {
		return nil, err
	}

	var purchaseOrders []domain.PurchaseOrdersGetAll
	for rows.Next() {
		po := domain.PurchaseOrdersGetAll{}
		_ = rows.Scan(&po.ID, &po.OrderNumber, &po.OrderDate, &po.TrackingCode, &po.BuyerID, &po.CarrierID, &po.ProductRecordID, &po.WarehouseID, &po.OrderStatusID)
		purchaseOrders = append(purchaseOrders, po)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return purchaseOrders, nil
}

func (r *repository) ExistsOrder(ctx context.Context, orderNumber string) bool {
	query := "SELECT order_number FROM purchase_orders WHERE order_number = ?"
	err := r.db.QueryRow(query, orderNumber).Scan(&orderNumber)
	return err == nil
}

func (r *repository) Save(ctx context.Context, o domain.PurchaseOrders) error {
	query := "INSERT INTO purchase_orders(order_number, order_date, tracking_code, buyer_id, product_record_id, order_status_id) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(o.OrderNumber, o.OrderDate, o.TrackingCode, o.BuyerID, o.ProductRecordID, o.OrderStatusID)
	if err != nil {
		return err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	o.ID = int(insertedID)

	return nil
}
