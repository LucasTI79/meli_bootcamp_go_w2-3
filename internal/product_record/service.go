package productrecord

import (
	"context"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

var (
	ErrNotFound     = errors.New("product not found")
	ErrInvalidJson  = errors.New("invalid json")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrInvalidField = errors.New("invalid field")
)

type Service interface {
	Save(ctx context.Context, p domain.ProductRecord) (int, error)
	RecordsByAllProductsReport(ctx context.Context) ([]domain.ProductRecordReport, error)
	RecordsByOneProductReport(ctx context.Context, id int) (domain.ProductRecordReport, error)
}

type ProductRecordService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &ProductRecordService{
		repository: r,
	}
}
func (s *ProductRecordService) Save(ctx context.Context, p domain.ProductRecord) (int, error) {
	productReportId, err := s.repository.Save(ctx, p)
	fmt.Println("productRecordId service", productReportId)
	return productReportId, err
}

func (s *ProductRecordService) RecordsByAllProductsReport(ctx context.Context) ([]domain.ProductRecordReport, error) {
	productRecordsReport, err := s.repository.RecordsByAllProductsReport(ctx)
	fmt.Println("productRecordsReport service", productRecordsReport)
	return productRecordsReport, err
}

func (s *ProductRecordService) RecordsByOneProductReport(ctx context.Context, id int) (domain.ProductRecordReport, error) {
	productRecord, err := s.repository.RecordsByOneProductReport(ctx, id)
	fmt.Println("productRecord service", productRecord)
	return productRecord, err
}
