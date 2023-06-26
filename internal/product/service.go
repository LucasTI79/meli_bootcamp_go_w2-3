package product

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

// Errors
var (
	ErrNotFound    = errors.New("product not found")
	ErrInvalidId   = errors.New("invalid id")
	ErrInvalidBody = errors.New("invalid body")
	ErrTryAgain    = errors.New("error, try again %s")

	ErrProductAlreadyExists = errors.New("product already exists")
)

type Service interface {
	Save(ctx context.Context, p domain.Product) (int, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (domain.Product, error)
	Update(ctx context.Context, p domain.Product) error
}

type productService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &productService{
		repository: r,
	}
}

func (s *productService) Save(ctx context.Context, p domain.Product) (int, error) {
	productExists := s.repository.Exists(ctx, p.ProductCode)
	if productExists {
		return 0, ErrProductAlreadyExists
	}

	productId, err := s.repository.Save(ctx, p)

	return productId, err
}

func (s *productService) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	return products, err

}

func (s *productService) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	return err
}

func (s *productService) Get(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repository.Get(ctx, id)
	return product, err
}

func (s *productService) Update(ctx context.Context, p domain.Product) error {
	err := s.repository.Update(ctx, p)
	return err
}
