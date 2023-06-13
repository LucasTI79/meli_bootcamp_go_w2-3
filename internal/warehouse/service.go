package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

var (
	ErrNotFound     = errors.New("warehouse not found")
	ErrInvalidId    = errors.New("invalid id")
	ErrInvalidBody  = errors.New("invalid body")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrAlredyExists = errors.New("warehouse already exists")
)

type Service interface {
	Save(ctx context.Context, d domain.Warehouse) (int, error)
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, d domain.Warehouse) error
}

type warehouseService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &warehouseService{
		repository: r,
	}
}

func (w *warehouseService) Save(ctx context.Context, d domain.Warehouse) (int, error) {
	if w.repository.Exists(ctx, d.WarehouseCode) {
		return 0, ErrAlredyExists
	}
	warehouseId, err := w.repository.Save(ctx, d)
	return warehouseId, err
}

func (w *warehouseService) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	warehouses, err := w.repository.GetAll(ctx)
	return warehouses, err
}

func (w *warehouseService) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	warehouse, err := w.repository.Get(ctx, id)
	return warehouse, err
}

func (w *warehouseService) Delete(ctx context.Context, id int) error {
	err := w.repository.Delete(ctx, id)
	return err
}

func (w *warehouseService) Update(ctx context.Context, d domain.Warehouse) error {
	if !w.repository.Exists(ctx, d.WarehouseCode) {
		return errors.New("cannot modify warehouse code")
	}
	err := w.repository.Update(ctx, d)
	return err
}
