package buyer

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Repository encapsulates the storage of a buyer.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
	ExistsBuyer(ctx context.Context, cardNumberID string) bool
	ExistsID(ctx context.Context, buyerID int) bool
	Save(ctx context.Context, b domain.Buyer) (int, error)
	Update(ctx context.Context, b domain.Buyer) error
	Delete(ctx context.Context, id int) error
	GetBuyerOrders(ctx context.Context, id int) (domain.BuyerOrders, error)
	GetBuyersOrders(ctx context.Context) ([]domain.BuyerOrders, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetBuyerOrders(ctx context.Context, id int) (domain.BuyerOrders, error) {
	query := "SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) AS `purchase_orders_count` FROM buyers b LEFT JOIN purchase_orders po ON b.id = po.buyer_id WHERE b.id = ? GROUP BY b.id, b.card_number_id, b.first_name, b.last_name"
	row := r.db.QueryRow(query, id)
	b := domain.BuyerOrders{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName, &b.PurchaseOrdersCount)
	if err != nil {
		return domain.BuyerOrders{}, err
	}

	return b, nil
}

func (r *repository) GetBuyersOrders(ctx context.Context) ([]domain.BuyerOrders, error) {
	query := "SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) AS `purchase_orders_count` FROM buyers b LEFT JOIN purchase_orders po ON b.id = po.buyer_id GROUP BY b.id, b.card_number_id, b.first_name, b.last_name"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var buyers []domain.BuyerOrders

	for rows.Next() {
		b := domain.BuyerOrders{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName, &b.PurchaseOrdersCount)
		buyers = append(buyers, b)
	}

	return buyers, nil
}

func (r *repository) ExistsID(ctx context.Context, buyerID int) bool {
	query := "SELECT COUNT(*) FROM buyers WHERE id = ?"
	var count int
	err := r.db.QueryRow(query, buyerID).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	query := "SELECT * FROM buyers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var buyers []domain.Buyer

	for rows.Next() {
		b := domain.Buyer{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
		buyers = append(buyers, b)
	}

	return buyers, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Buyer, error) {
	query := "SELECT * FROM buyers WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := domain.Buyer{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return domain.Buyer{}, err
	}

	return b, nil
}

func (r *repository) ExistsBuyer(ctx context.Context, cardNumberID string) bool {
	query := "SELECT card_number_id FROM buyers WHERE card_number_id=?;"
	row := r.db.QueryRow(query, cardNumberID)
	err := row.Scan(&cardNumberID)
	return err == nil
}

func (r *repository) Save(ctx context.Context, b domain.Buyer) (int, error) {
	query := "INSERT INTO buyers(card_number_id,first_name,last_name) VALUES (?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, b domain.Buyer) error {
	query := "UPDATE buyers SET first_name=?, last_name=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&b.FirstName, &b.LastName, &b.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM buyers WHERE id = ?"
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
