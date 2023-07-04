package purchase_orders

import (
	//"context"
	"errors"
	//"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound  = errors.New("buyer not found")
	ErrExists    = errors.New("buyer already exists")
	ErrInvalidID = errors.New("invalid ID")
)

type Service interface {
	//GetAll(ctx context.Context) ([]domain.Buyer, error)
	//Get(ctx context.Context, id int) (domain.Buyer, error)
	//Create(ctx context.Context, b domain.Buyer) (domain.Buyer, error)
	//Update(ctx context.Context, b domain.Buyer, id int) (domain.Buyer, error)
	//Delete(ctx context.Context, id int) error
}

type purchaseordersService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &purchaseordersService{
		repository: r,
	}
}
