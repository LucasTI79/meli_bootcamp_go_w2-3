package productbatch

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

const (
	SaveQuery                       = "INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"
	SectionExists                   = "SELECT id FROM sections WHERE id=?"
	SectionProductsReports          = "SELECT count(pb.id) as `products_count`, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id GROUP BY pb.section_id"
	SectionProductsReportsBySection = "SELECT count(pb.id) as `products_count`, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id WHERE pb.section_id = ? GROUP BY pb.section_id"
)

type Repository interface {
	GetProductsBySection() ([]domain.ProductBySection, error)
	Save(produsctBatch domain.ProductBatch) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) GetProductsBySection() ([]domain.ProductBySection, error) {
	rows, err := r.db.Query(ProductsBySection)
	if err != nil {
		return nil, err
	}
	var productsBySection []domain.ProductBySection
	for rows.Next() {
		var productBySection domain.ProductBySection
		err := rows.Scan(&productBySection.ProductsCount, &productBySection.SectionID, &productBySection.SectionNumber)
		if err != nil {
			return nil, err
		}
		productsBySection = append(productsBySection, productBySection)
	}
	return productsBySection, nil
}

func (r *repository) Save(produsctBatch domain.ProductBatch) (int, error) {
	stmt, err := r.db.Prepare(SaveQuery)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(produsctBatch.BatchNumber, produsctBatch.CurrentQuantity, produsctBatch.CurrentTemperature, produsctBatch.DueDate, produsctBatch.InitialQuantity, produsctBatch.ManufacturingDate, produsctBatch.ManufacturingHour, produsctBatch.MinimumTemperature, produsctBatch.ProductID, produsctBatch.SectionID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
