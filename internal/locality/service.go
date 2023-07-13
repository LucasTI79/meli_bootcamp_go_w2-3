package locality

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
)

var (
	ErrProvinceNotFound  = errors.New("province not found")
	ErrNotFound          = errors.New("locality not found")
	ErrNoSellersLocality = errors.New("no sellers found in this location")
	ErrTryAgain          = errors.New("error, try again %s")
	ErrLocalityNotExists = errors.New("locality does not exists")
)

type Service interface {
	Save(ctx context.Context, locality domain.Locality) (domain.LocalityInput, error)
	ReportSellersByLocality(ctx context.Context, id int) ([]domain.LocalityReport, error)
	ExistsById(c context.Context, idLocality int) error
}

type LocalityService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &LocalityService{
		repository: r,
	}
}

func (l *LocalityService) Save(c context.Context, locality domain.Locality) (domain.LocalityInput, error) {
	IdProvince, err := l.repository.GetProvinceByName(c, locality.ProvinceName)
	if err != nil {
		return domain.LocalityInput{}, ErrProvinceNotFound
	}

	localityInput := domain.LocalityInput{
		LocalityName: locality.LocalityName,
		IdProvince:   IdProvince,
	}

	localityId, err := l.repository.Save(c, localityInput)
	if err != nil {
		return domain.LocalityInput{}, ErrTryAgain
	}
	localityInput.ID = localityId
	return localityInput, nil
}

func (l *LocalityService) ReportSellersByLocality(c context.Context, id int) ([]domain.LocalityReport, error) {
	var report []domain.LocalityReport

	if id != 0 {
		if !(l.repository.ExistsById(c, id)) {
			return []domain.LocalityReport{}, ErrNotFound
		}

		reportWihId, err := l.repository.ReportLocalityId(c, id)
		if err != nil {
			return []domain.LocalityReport{}, ErrTryAgain
		}
		report = append(report, reportWihId)
	} else {
		reportAll, err := l.repository.ReportLocality(c)
		if err != nil {
			return []domain.LocalityReport{}, ErrTryAgain
		}

		report = append(report, reportAll...)
	}
	return report, nil
}

func (l *LocalityService) ExistsById(c context.Context, idLocality int) error {
	localityExists := l.repository.ExistsById(c, idLocality)

	if !localityExists {
		return ErrLocalityNotExists
	}
	return nil
}
