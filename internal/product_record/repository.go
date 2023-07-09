package productrecord

import (
	"context"
	"database/sql"
	"fmt"

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
	RecordsByOneProductReport(ctx context.Context, id int) (domain.RecordByProduct, error)
	RecordsByAllProductsReport(ctx context.Context) ([]domain.RecordByProduct, error)
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
func (r *repository) RecordsByAllProductsReport(ctx context.Context) ([]domain.RecordByProduct, error) {
	// query := "SELECT * FROM product_records;" //ajustar
	rows, err := r.db.Query(RecordsByAllProductsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recordsByProduct []domain.RecordByProduct
	for rows.Next() {
		p := domain.RecordByProduct{} //product_records by one product
		err := rows.Scan(&p.ProductID, &p.Description, &p.RecordsCount)
		if err != nil {
			return nil, err
		}
		recordsByProduct = append(recordsByProduct, p)
	}
	return recordsByProduct, nil
}

// get the product_records by one specific product
func (r *repository) RecordsByOneProductReport(ctx context.Context, id int) (domain.RecordByProduct, error) {
	// query := "SELECT * FROM products WHERE id=?;"
	row := r.db.QueryRow(RecordsByOneProductQuery, id)
	p := domain.RecordByProduct{}
	err := row.Scan(&p.ProductID, &p.Description, &p.RecordsCount)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return domain.RecordByProduct{}, ErrNotFound
		}
		return domain.RecordByProduct{}, err
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
		fmt.Println("erro da repository", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("id do obj inserido", id)
	return int(id), nil
}
