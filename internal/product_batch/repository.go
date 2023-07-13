package productbatch

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

const (
	SaveQuery = "INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"
)

type Querys struct {
	SaveQuery string
}
type Repository interface {
	Save(produsctBatch domain.ProductBatch) (int, error)
}

type repository struct {
	db *sql.DB
	Querys
}

func buildQuerys(Querys Querys) Querys {
	if Querys.SaveQuery == "" {
		Querys.SaveQuery = SaveQuery
	}
	return Querys
}

func NewRepository(db *sql.DB, Querys Querys) Repository {
	RepoQuery := buildQuerys(Querys)
	return &repository{
		db:     db,
		Querys: RepoQuery,
	}
}

func (r *repository) Save(produsctBatch domain.ProductBatch) (int, error) {
	stmt, err := r.db.Prepare(r.Querys.SaveQuery)
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
