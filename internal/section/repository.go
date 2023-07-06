package section

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

const (
	SectionExists                   = "SELECT id FROM sections WHERE id=?"
	SectionProductsReports          = "SELECT count(pb.id) as `products_count`, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id GROUP BY pb.section_id"
	SectionProductsReportsBySection = "SELECT count(pb.id) as `products_count`, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id WHERE pb.section_id = ? GROUP BY pb.section_id"
)

// Repository encapsulates the storage of a section.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, sectionNumber int) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section) error
	Delete(ctx context.Context, id int) error
	ExistsById(sectionID int) bool
	SectionProductsReportsBySection(id int) (domain.ProductBySection, error)
	SectionProductsReports() ([]domain.ProductBySection, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Section, error) {
	query := "SELECT * FROM sections;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var sections []domain.Section

	for rows.Next() {
		s := domain.Section{}
		_ = rows.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
		sections = append(sections, s)
	}

	return sections, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Section, error) {
	query := "SELECT * FROM sections WHERE id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Section{}
	err := row.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		return domain.Section{}, err
	}

	return s, nil
}

func (r *repository) Exists(ctx context.Context, sectionNumber int) bool {
	query := "SELECT section_number FROM sections WHERE section_number=?;"
	row := r.db.QueryRow(query, sectionNumber)
	err := row.Scan(&sectionNumber)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Section) (int, error) {
	query := "INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Section) error {
	query := "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID, &s.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM sections WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) ExistsById(sectionID int) bool {
	row := r.db.QueryRow(SectionExists, sectionID)
	err := row.Scan(&sectionID)
	return err == nil
}

func (r *repository) SectionProductsReportsBySection(id int) (domain.ProductBySection, error) {
	row := r.db.QueryRow(SectionProductsReportsBySection, id)
	var productBySection domain.ProductBySection
	err := row.Scan(&productBySection.ProductsCount, &productBySection.SectionID, &productBySection.SectionNumber)
	if err != nil {
		return domain.ProductBySection{}, err
	}
	return productBySection, nil
}

func (r *repository) SectionProductsReports() ([]domain.ProductBySection, error) {
	rows, err := r.db.Query(SectionProductsReports)
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
