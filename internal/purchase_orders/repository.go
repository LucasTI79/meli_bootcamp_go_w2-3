package purchase_orders

import (
	//"context"
	"database/sql"
	//"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Repository encapsulates the storage of a purchased order.
type Repository interface {
	//GetAll(ctx context.Context) ([]domain.Buyer, error)
	//Get(ctx context.Context, id int) (domain.Buyer, error)
	//Exists(ctx context.Context, cardNumberID string) bool
	//Save(ctx context.Context, b domain.Buyer) (int, error)
	//Update(ctx context.Context, b domain.Buyer) error
	//Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
