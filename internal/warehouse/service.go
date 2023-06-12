package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

var (
	ErrNotFound     = errors.New("warehouse not found")
	ErrAlredyExists = errors.New("warehouse already exists")
)

type Service interface {
	Save(ctx context.Context, d domain.Warehouse) (int, error)
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
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
	userExist := w.repository.Exists(ctx, d.WarehouseCode)
	if userExist {
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
