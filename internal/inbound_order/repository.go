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
	ReportByAll(ctx context.Context) ([]domain.InboundOrdersReport, error)
	ReportByOne(ctx context.Context, id int) (domain.InboundOrdersReport, error)
}

const (
	ReportByAll = "SELECT io.employee_id as `id`, employees.card_number_id, employees.first_name, employees.last_name, employees.warehouse_id, count(io.id) as `inbound_order_count`FROM inbound_orders io JOIN employees ON io.employee_id = employees.id Group BY io.employee_id"
	ReportByOne = "SELECT io.employee_id as `id`, employees.card_number_id, employees.first_name, employees.last_name, employees.warehouse_id, count(io.id) as `inbound_order_count`FROM inbound_orders io JOIN employees ON io.employee_id = employees.id WHERE employee_id=?Group BY io.employee_id"
)

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

func (r *repository) ReportByAll(ctx context.Context) ([]domain.InboundOrdersReport, error) {

	rows, err := r.db.Query(ReportByAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var report []domain.InboundOrdersReport
	for rows.Next() {
		p := domain.InboundOrdersReport{}
		err := rows.Scan(&p.ID, &p.CardNumberID, &p.FirstName, &p.LastName, &p.WarehouseID, &p.InboundOrdersCount)
		if err != nil {
			return nil, err
		}
		report = append(report, p)
	}
	return report, nil
}

// get all product_records by one specific product
func (r *repository) ReportByOne(ctx context.Context, id int) (domain.InboundOrdersReport, error) {

	row := r.db.QueryRow(ReportByOne, id)
	p := domain.InboundOrdersReport{}
	err := row.Scan(&p.ID, &p.CardNumberID, &p.FirstName, &p.LastName, &p.WarehouseID, &p.InboundOrdersCount)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			return domain.InboundOrdersReport{}, ErrNotFound
		}
		return domain.InboundOrdersReport{}, err
	}
	return p, nil
}
