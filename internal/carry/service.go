package carry

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
)

var (
	ErrNotFound     = errors.New("carry not found")
	ErrInvalidId    = errors.New("invalid id")
	ErrInvalidBody  = errors.New("invalid body")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrAlredyExists = errors.New("carry already exists")
	ErrInvalidJSON  = errors.New("invalid json")
	ErrConflictLocalityId  = errors.New("locality_id not found")
)

type Service interface {
	Create(ctx context.Context, d domain.Carry) (domain.Carry, error)
	Get(ctx context.Context, id int) (domain.Carry, error)
	Read(ctx context.Context, id int) ([]domain.LocalityCarriersReport, error)
}

type CarryService struct {
	repository Repository
	repositoryLocality locality.Repository
}

func NewService(r Repository, l locality.Repository) Service {
	return &CarryService{
		repository: r,
		repositoryLocality: l,
	}
}

func (c *CarryService) Create(ctx context.Context, d domain.Carry) (domain.Carry, error) {
	if c.repository.ExistsByCidCarry(ctx, d.Cid) {
		return domain.Carry{}, ErrAlredyExists
	}

	if !c.repositoryLocality.ExistsById(ctx, d.LocalityId){
		return domain.Carry{}, ErrConflictLocalityId
	}

	carryId, err := c.repository.Create(ctx, d)
	if err != nil {
		return domain.Carry{}, err
	}
	d.ID = carryId
	return d, nil
}

func (c *CarryService) Read(ctx context.Context, id int) ([]domain.LocalityCarriersReport, error){
	//VALIDAR LOCALITIES
	var readReport []domain.LocalityCarriersReport
	if id != 0 {
		reportWithId, err := c.repository.ReadCarriersWithLocalityId(ctx, id)
		if err != nil {
			return []domain.LocalityCarriersReport{}, ErrTryAgain
		}
		readReport = append(readReport, reportWithId)
	} else {
		reportAll, err := c.repository.ReadAllCarriers(ctx)
		if err != nil {
			return []domain.LocalityCarriersReport{}, ErrTryAgain
		}
		readReport = append(readReport, reportAll...)
	}
	return readReport, nil
}

func (c *CarryService) Get(ctx context.Context, id int) (domain.Carry, error) {
	carry, err := c.repository.Get(ctx, id)
	return carry, err
}
