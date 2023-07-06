package locality

import (
	"context"
	"database/sql"
)

const (
	ExistsById = "SELECT id FROM localities WHERE id = ?"
)

type Repository interface {
	ExistsById(ctx context.Context, id int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ExistsById(ctx context.Context, id int) bool {
	row := r.db.QueryRow(ExistsById, id)
	err := row.Scan(&id)

	return err == nil
}