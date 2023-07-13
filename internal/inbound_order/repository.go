package inbound_order

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.InboundOrders, error)
	Create(ctx context.Context, c domain.InboundOrders) (int, error)
	Get(ctx context.Context, id int) (domain.InboundOrders, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, e domain.InboundOrders) (int, error)
	Update(ctx context.Context, e domain.InboundOrders) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.InboundOrders, error) {
	query := "SELECT * FROM inbound_orders"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var inboundOrders []domain.InboundOrders

	for rows.Next() {
		i := domain.InboundOrders{}
		_ = rows.Scan(&i.ID, &i.OrderDate, &i.OrderNumber, &i.EmployeeID, &i.ProductBatchID, &i.WarehouseID)
		inboundOrders = append(inboundOrders, i)
	}

	return inboundOrders, nil
}

func (r *repository) Create(ctx context.Context, i domain.InboundOrders) (int, error) {
	query := "SELECT * FROM inbound_orders"
	rows, err := r.db.Query(query)
	if err != nil {
		return 0, er
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

func (r *repository) Save(ctx context.Context, i domain.InboundOrders) (int, error) {
	query := "INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)"
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

func (r *repository) Update(ctx context.Context, i domain.InboundOrders) error {
	query := "UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&i.OrderDate, &i.OrderNumber, &i.EmployeeID, &i.ProductBatchID, &i.WarehouseID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM inbound_orders WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
