package locality

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

const (
	ExistsById = "SELECT id FROM localities WHERE id = ?"
	ReportSellersByLocality = "SELECT l.ID, l.locality_name, COUNT(s.ID) FROM localities l JOIN sellers s ON l.ID = s.locality_id where l.id = ? GROUP BY l.id, l.locality_name"
	ReportLocality = "SELECT l.ID, l.locality_name, COUNT(s.ID) FROM localities l JOIN sellers s GROUP BY l.id, l.locality_name"
)

type Repository interface {
	ExistsById(ctx context.Context, id int) bool
	ReportLocalityId(ctx context.Context, id int) (domain.LocalityReport, error)
	ReportLocality(ctx context.Context) ([]domain.LocalityReport, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ReportLocality(ctx context.Context) ([]domain.LocalityReport, error) {
	rows, err := r.db.Query(ReportLocality)
	if err != nil {
		return nil, err
	}

	var report []domain.LocalityReport
	for rows.Next() {
		l := domain.LocalityReport{}
		_ = rows.Scan(&l.IdLocality, &l.LocalityName, &l.SellersCount)
		report = append(report, l)
	}
	return report, nil
}

func (r *repository) ReportLocalityId(ctx context.Context, id int) (domain.LocalityReport, error) {
	row := r.db.QueryRow(ReportSellersByLocality, id)
	l := domain.LocalityReport{}
	err := row.Scan(&l.IdLocality, &l.LocalityName, &l.SellersCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.LocalityReport{}, err
		}
		return domain.LocalityReport{}, err
	}
	return l, nil
}

func (r *repository) ExistsById(ctx context.Context, id int) bool {
	row := r.db.QueryRow(ExistsById, id)
	err := row.Scan(&id)

	return err == nil
}