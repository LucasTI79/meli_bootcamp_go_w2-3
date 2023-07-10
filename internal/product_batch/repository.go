package productbatch

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

const (
	SaveQuery = "INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"
)

type Repository interface {
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
