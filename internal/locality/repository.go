package locality

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

const (
	ExistsById = "SELECT id, locality_name, province_id FROM localities WHERE id = ?"
)

type Repository interface {
	ExistsById(ctx context.Context, id int) (domain.LocalityInput, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ExistsById(ctx context.Context, id int) (domain.LocalityInput, error){
	row := r.db.QueryRow(ExistsById, id)

	l := domain.LocalityInput{}
	err := row.Scan(&l.ID, &l.LocalityName, &l.IdProvince)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return domain.LocalityInput{}, ErrNotFound
		}
		return domain.LocalityInput{}, err
	}

	return l, nil
}