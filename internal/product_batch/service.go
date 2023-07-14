package productbatch

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

type Service interface {
	Save(ctx context.Context, p domain.ProductBatch) (int, error)
}

type serviceProductBatch struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &serviceProductBatch{
		repository: r,
	}
}
func (s *serviceProductBatch) Save(ctx context.Context, p domain.ProductBatch) (int, error) {
	productBatchID, err := s.repository.Save(p)
	return productBatchID, err
}
