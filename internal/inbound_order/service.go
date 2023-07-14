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
	Get(ctx context.Context, id int) (domain.InboundOrders, error)
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

func (c *inboundOrderService) Create(ctx context.Context, d domain.InboundOrders) (domain.InboundOrders, error) {
	if c.repository.Exists(ctx, d.OrderNumber) {
		return domain.InboundOrders{}, ErrAlredyExists
	}

	InboundOrdersId, err := c.repository.Create(ctx, d)
	if err != nil {
		return domain.InboundOrders{}, err
	}
	d.ID = InboundOrdersId
	return d, nil
}

func (s *inboundOrderService) Get(ctx context.Context, id int) (domain.InboundOrders, error) {
	inboundOrders, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.InboundOrders{}, ErrNotFound
	}
	return inboundOrders, nil
}
