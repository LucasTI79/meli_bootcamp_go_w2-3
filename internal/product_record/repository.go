package productrecord

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

const (
	ProductExistsQuery        = "SELECT id FROM products WHERE id=?"
	SaveQuery                 = "INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)"
	RecordsByAllProductsQuery = "SELECT  pr.product_id,  products.description, count(pr.id) as `records_count` FROM product_records pr JOIN products ON pr.product_id = products.id Group BY pr.product_id"
	RecordsByOneProductQuery  = "SELECT  pr.product_id,  products.description, count(pr.id) as `records_count` FROM product_records pr JOIN products ON pr.product_id = products.id WHERE product_id=? Group BY pr.product_id;"
)

type Repository interface {
	// ProductExists(ctx context.Context, productID int) bool
	Save(ctx context.Context, p domain.ProductRecord) (int, error)
	RecordsByOneProductReport(ctx context.Context, id int) (domain.ProductRecordReport, error)
	RecordsByAllProductsReport(ctx context.Context) ([]domain.ProductRecordReport, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db,
	}
}

// get all product_records by each product
func (r *repository) RecordsByAllProductsReport(ctx context.Context) ([]domain.ProductRecordReport, error) {

	rows, err := r.db.Query(RecordsByAllProductsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recordsByProduct []domain.ProductRecordReport
	for rows.Next() {
		p := domain.ProductRecordReport{}
		err := rows.Scan(&p.ProductID, &p.Description, &p.RecordsCount)
		if err != nil {
			return nil, err
		}
		recordsByProduct = append(recordsByProduct, p)
	}
	return recordsByProduct, nil
}

// get all product_records by one specific product
func (r *repository) RecordsByOneProductReport(ctx context.Context, id int) (domain.ProductRecordReport, error) {

	row := r.db.QueryRow(RecordsByOneProductQuery, id)
	p := domain.ProductRecordReport{}
	err := row.Scan(&p.ProductID, &p.Description, &p.RecordsCount)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			return domain.ProductRecordReport{}, ErrNotFound
		}
		return domain.ProductRecordReport{}, err
	}
	return p, nil
}

func (r *repository) Save(ctx context.Context, productRecord domain.ProductRecord) (int, error) {
	stmt, err := r.db.Prepare(SaveQuery)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductID)
	if err != nil {

		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
