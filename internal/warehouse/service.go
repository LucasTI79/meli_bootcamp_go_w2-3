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
	ErrInvalidJSON  = errors.New("invalid json")
)

type Service interface {
	Save(ctx context.Context, d domain.Warehouse) (domain.Warehouse, error)
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, d domain.Warehouse, id int) (domain.Warehouse, error)
}

type WarehouseService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &WarehouseService{
		repository: r,
	}
}

func (w *WarehouseService) Save(ctx context.Context, d domain.Warehouse) (domain.Warehouse, error) {
	if w.repository.Exists(ctx, d.WarehouseCode) {
		return domain.Warehouse{}, ErrAlredyExists
	}
	warehouseId, err := w.repository.Save(ctx, d)
	if err != nil {
		return domain.Warehouse{}, err
	}
	d.ID = warehouseId
	return d, nil
}

func (w *WarehouseService) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	warehouses, err := w.repository.GetAll(ctx)
	return warehouses, err
}

func (w *WarehouseService) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	warehouse, err := w.repository.Get(ctx, id)
	return warehouse, err
}

func (w *WarehouseService) Delete(ctx context.Context, id int) error {
	err := w.repository.Delete(ctx, id)
	return err
}

func (w *WarehouseService) Update(ctx context.Context, d domain.Warehouse, id int) (domain.Warehouse, error) {
	warehouseDomain, err := w.Get(ctx, id)
	if err != nil {
		return domain.Warehouse{}, ErrNotFound
	}

	if d.Address != "" {
		warehouseDomain.Address = d.Address
	}
	if d.Telephone != "" {
		warehouseDomain.Telephone = d.Telephone
	}
	if d.WarehouseCode != "" {
		exists := w.repository.Exists(ctx, d.WarehouseCode)
		if exists && warehouseDomain.WarehouseCode != d.WarehouseCode {
			return domain.Warehouse{}, ErrAlredyExists
		}
		warehouseDomain.WarehouseCode = d.WarehouseCode
	}
	if d.MinimumCapacity != 0 {
		warehouseDomain.MinimumCapacity = d.MinimumCapacity
	}
	if d.MinimumTemperature != 0 {
		warehouseDomain.MinimumTemperature = d.MinimumTemperature
	}
	
	err = w.repository.Update(ctx, warehouseDomain)
	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouseDomain, nil
}
