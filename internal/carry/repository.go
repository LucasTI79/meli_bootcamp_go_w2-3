package carry

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	_ "github.com/go-sql-driver/mysql"
)

// Repository encapsulates the storage of a carry.
type Repository interface {
	Read(ctx context.Context, localityId int) ([]domain.LocalityCarriersReport, error)
	Create(ctx context.Context, c domain.Carry) (int, error)
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
	ReadCarriersWithLocalityId = "SELECT L.id, L.locality_name, COUNT(C.id) AS carriers_count " +
	"FROM localities L LEFT JOIN carriers C ON L.id = C.locality_id " +
	"WHERE L.id = ? " +
	"GROUP BY L.id, L.locality_name"
	ReadAllCarriers = "SELECT L.id, L.locality_name, COUNT(C.id) AS carriers_count " +
	"FROM localities L LEFT JOIN carriers C ON L.id = C.locality_id " +
	"GROUP BY L.id, L.locality_name"
)

func (r *repository) Read(ctx context.Context, localityId int) ([]domain.LocalityCarriersReport, error) {
	var query string

	if localityId != 0 {
		query = ReadCarriersWithLocalityId
	} else {
		query = ReadAllCarriers
	}

	rows, err := r.db.Query(query, localityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	var report []domain.LocalityCarriersReport

	for rows.Next() {
		var localityID int
		var localityName string
		var carriersCount int
		err := rows.Scan(&localityId, &localityName, &carriersCount)
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
