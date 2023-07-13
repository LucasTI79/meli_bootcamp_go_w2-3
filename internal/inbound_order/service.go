package inbound_order

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound            = errors.New("inbound orders not found")
	ErrConflict            = errors.New("409 Conflict: inbound orders already exists")
	ErrUnprocessableEntity = errors.New("all fields are required")
	ErrInvalidId           = errors.New("invalid id")
	ErrTryAgain            = errors.New("internal error")
	ErrAlredyExists        = errors.New("already exists")
	ErrInvalidJSON         = errors.New("invalid JSON")
)

type Service interface {
	Get(ctx *context.Context, id int) (*domain.InboundOrders, error)
	GetAll(ctx *context.Context) (*[]domain.InboundOrders, error)
	Save(ctx *context.Context, inboundOrders domain.InboundOrders) (*domain.InboundOrders, error)
	Update(ctx *context.Context, id int, reqUpdateInboundOrders *domain.RequestUpdateInboundOrders) (*domain.InboundOrders, error)
	Delete(ctx *context.Context, id int) error
	Create(ctx context.Context, d domain.InboundOrders) (domain.InboundOrders, error)
}

type inboundOrderService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &inboundOrderService{
		repository: r,
	}
}

func (s *inboundOrderService) Get(ctx *context.Context, id int) (*domain.InboundOrders, error) {
	inboundOrders, err := s.repository.Get(*ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return &inboundOrders, nil
}

func (s *inboundOrderService) GetAll(ctx *context.Context) (*[]domain.InboundOrders, error) {
	inboundOrders := []domain.InboundOrders{}

	inboundOrders, err := s.repository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}
	return &inboundOrders, nil
}

func (s *inboundOrderService) Save(ctx *context.Context, inboundOrders domain.InboundOrders) (*domain.InboundOrders, error) {
	id, err := s.repository.Save(*ctx, inboundOrders)
	if err != nil {
		return nil, err
	}

	inboundOrders.ID = id

	return &inboundOrders, nil
}

func (s *inboundOrderService) Update(ctx *context.Context, id int, reqUpdateInboundOrders *domain.RequestUpdateInboundOrders) (*domain.InboundOrders, error) {
	existingInboundOrders, err := s.repository.Get(*ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	if reqUpdateInboundOrders.OrderDate != nil {
		existingInboundOrders.OrderDate = *reqUpdateInboundOrders.OrderDate
	}
	if reqUpdateInboundOrders.OrderNumber != nil {
		existingInboundOrders.OrderNumber = *reqUpdateInboundOrders.OrderNumber
	}
	if reqUpdateInboundOrders.EmployeeID != nil {
		existingInboundOrders.EmployeeID = *reqUpdateInboundOrders.EmployeeID
	}
	if reqUpdateInboundOrders.ProductBatchID != nil {
		existingInboundOrders.ProductBatchID = *reqUpdateInboundOrders.ProductBatchID
	}
	if reqUpdateInboundOrders.WarehouseID != nil {
		existingInboundOrders.WarehouseID = *reqUpdateInboundOrders.WarehouseID
	}

	err = s.repository.Update(*ctx, existingInboundOrders)
	if err != nil {
		return nil, err
	}

	return &existingInboundOrders, nil
}

func (s *inboundOrderService) Delete(ctx *context.Context, id int) error {
	_, err := s.repository.Get(*ctx, id)
	if err != nil {
		return ErrNotFound
	}

	err = s.repository.Delete(*ctx, id)
	if err != nil {
		return err
	}
	return nil
}
