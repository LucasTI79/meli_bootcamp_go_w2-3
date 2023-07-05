package carry

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	_ "github.com/go-sql-driver/mysql"
)

// Repository encapsulates the storage of a carry.
type Repository interface {
	Create(ctx context.Context, c domain.Carry) (int, error)
	Get(ctx context.Context, id int) (domain.Carry, error)
	ReadAllCarriers(ctx context.Context) ([]domain.LocalityCarriersReport, error)
	ReadCarriersWithLocalityId(ctx context.Context, localityID int) (domain.LocalityCarriersReport, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	CreateCarry = "INSERT INTO carriers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	GetCarry = "SELECT id, cid, company_name, address, telephone, locality_id FROM carriers WHERE id = ?"
	ReadCarriersWithLocalityId = "SELECT L.id, L.locality_name, COUNT(C.id) AS carriers_count " +
	"FROM localities L LEFT JOIN carriers C ON L.id = C.locality_id " +
	"WHERE L.id = ? " +
	"GROUP BY L.id, L.locality_name"
	ReadAllCarriers = "SELECT L.id, L.locality_name, COUNT(C.id) AS carriers_count " +
	"FROM localities L LEFT JOIN carriers C ON L.id = C.locality_id " +
	"GROUP BY L.id, L.locality_name"
)

func (r *repository) Create(ctx context.Context, c domain.Carry) (int, error) {
	stmt, err := r.db.Prepare(CreateCarry)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&c.Cid, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Carry, error) {
	row := r.db.QueryRow(GetCarry, id)
	c := domain.Carry{}
	err := row.Scan(&c.ID, &c.Cid, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return domain.Carry{}, ErrNotFound
		}
		return domain.Carry{}, err
	}

	return c, nil
}

func (r *repository) ReadAllCarriers(ctx context.Context) ([]domain.LocalityCarriersReport, error) {
	rows, err := r.db.Query(ReadAllCarriers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var report []domain.LocalityCarriersReport

	for rows.Next() {
		var localityID int
		var localityName string
		var carriersCount int
		err := rows.Scan(&localityID, &localityName, &carriersCount)
		if err != nil {
			return nil, err
		}

		carriersReport := domain.LocalityCarriersReport{
			LocalityID: localityID,
			LocalityName: localityName,
			CarriersCount: carriersCount,
		}
		report = append(report, carriersReport)
	}
	return report, nil
}

func (r *repository) ReadCarriersWithLocalityId(ctx context.Context, localityID int) (domain.LocalityCarriersReport, error) {
	row := r.db.QueryRow(ReadCarriersWithLocalityId, localityID)

	l := domain.LocalityCarriersReport{}
	err := row.Scan(&l.LocalityID, &l.LocalityName, &l.CarriersCount)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return domain.LocalityCarriersReport{}, ErrNotFound
		}
		return domain.LocalityCarriersReport{}, err
	}

	return l, nil
}
