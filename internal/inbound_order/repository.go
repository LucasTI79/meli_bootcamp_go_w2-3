package inbound_order

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, c domain.InboundOrders) (int, error)
	Get(ctx context.Context, id int) (domain.InboundOrders, error)
	Exists(ctx context.Context, order string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, i domain.InboundOrders) (int, error) {
	query := "INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&i.OrderDate, &i.OrderNumber, &i.EmployeeID, &i.ProductBatchID, &i.WarehouseID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.InboundOrders, error) {
	query := "SELECT * FROM inbound_orders WHERE id=?;"
	row := r.db.QueryRow(query, id)
	i := domain.InboundOrders{}
	err := row.Scan(&i.ID, &i.OrderDate, &i.OrderNumber, &i.EmployeeID, &i.ProductBatchID, &i.WarehouseID)
	if err != nil {
		return domain.InboundOrders{}, err
	}

	return i, nil
}

func (r *repository) Exists(ctx context.Context, id string) bool {
	query := "SELECT id FROM inbound_orders WHERE id=?;"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&id)
	return err == nil
}
