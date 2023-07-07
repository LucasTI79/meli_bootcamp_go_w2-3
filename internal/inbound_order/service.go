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
)

type Service interface {
	Get(ctx *context.Context, id int) (*domain.InboundOrders, error)
	GetAll(ctx *context.Context) (*[]domain.InboundOrders, error)
	Save(ctx *context.Context, inboundOrders domain.InboundOrders) (*domain.InboundOrders, error)
	Update(ctx *context.Context, id int, reqUpdateInboundOrders *domain.RequestUpdateInboundOrders) (*domain.InboundOrders, error)
	Delete(ctx *context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(ctx *context.Context, id int) (*domain.InboundOrders, error) {
	inboundOrders, err := s.repository.Get(*ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return &inboundOrders, nil
}

func (s *service) GetAll(ctx *context.Context) (*[]domain.InboundOrders, error) {
	inboundOrders := []domain.InboundOrders{}

	inboundOrders, err := s.repository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}
	return &inboundOrders, nil
}

func (s *service) Save(ctx *context.Context, inboundOrders domain.InboundOrders) (*domain.InboundOrders, error) {
	id, err := s.repository.Save(*ctx, inboundOrders)
	if err != nil {
		return nil, err
	}

	inboundOrders.ID = id

	return &inboundOrders, nil
}

func (s *service) Update(ctx *context.Context, id int, reqUpdateInboundOrders *domain.RequestUpdateInboundOrders) (*domain.InboundOrders, error) {
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

func (s *service) Delete(ctx *context.Context, id int) error {
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
