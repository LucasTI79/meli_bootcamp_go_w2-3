package locality_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/locality"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var expectedReport = domain.LocalityReport{
	IdLocality: 1,
	LocalityName:  "São Paulo",
    SellersCount: 22,
}
var expectedLocality = domain.Locality{
	LocalityName: "Florianopolis",
	ProvinceName: "Santa Catarina",
}

var localityInput = domain.LocalityInput{
	LocalityName: "Florianopolis",
	IdProvince: 2,
}

var expectedAllReport = []domain.LocalityReport{
    {
        IdLocality: 1,
        LocalityName:  "São Paulo",
        SellersCount: 22,
    },
    {
        IdLocality: 2,
        LocalityName:  "Rio de Janeiro",
        SellersCount: 22,
    },
}

func TestSave(t *testing.T) {
	t.Run("Should create the locality if it contains the required fields", func(t *testing.T) {
        repository, service := InitServer(t)
		repository.On("GetProvinceByName", mock.Anything, expectedLocality.ProvinceName).Return(2, nil)
		repository.On("Save", mock.Anything, localityInput).Return(1, nil)

		locality, err := service.Save(context.TODO(), expectedLocality)
		fmt.Println(locality)

		assert.Equal(t, 1, locality.ID)
		assert.Equal(t, "Florianopolis", locality.LocalityName)
		assert.Equal(t, 2, locality.IdProvince)

		assert.NoError(t, err)
	})
	t.Run("Should return err ErrProvinceNotFound not found when province not exists", func(t *testing.T) {
		expectedMessage := "province not found"

		repository, service := InitServer(t)

		repository.On("GetProvinceByName", mock.Anything, mock.Anything).Return(0, errors.New("error"))
		repository.On("Save", mock.Anything,  mock.Anything).Return(domain.LocalityInput{}, locality.ErrProvinceNotFound)

		_, err := service.Save(context.TODO(), domain.Locality{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})

	t.Run("should return an error when it hears problems with the database", func(t *testing.T) {
		expectedMessage := "error, try again %s"

		repository, service := InitServer(t)

		repository.On("GetProvinceByName", mock.Anything, mock.Anything).Return(2, nil)
		repository.On("Save", mock.Anything, mock.Anything).Return(0, locality.ErrTryAgain)

		_, err := service.Save(context.TODO(), domain.Locality{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestReportSellersByLocality(t *testing.T) {
    t.Run("should return the report of sellers by locality", func(t *testing.T) {
        var reportExpected []domain.LocalityReport
        reportExpected = append(reportExpected, expectedReport)

        repository, service := InitServer(t)
        repository.On("ExistsById", mock.Anything, 1).Return(true)
        repository.On("ReportLocalityId", mock.Anything, 1).Return(expectedReport, nil)
        reportSlice, err := service.ReportSellersByLocality(context.TODO(), 1)

        assert.Equal(t, reportExpected, reportSlice)
		assert.NoError(t, err)
    })
	t.Run("Should get all the localities with their carriers when the id is not passed and it exists in database", func(t *testing.T) {
		repository, service := InitServer(t)
		repository.On("ReportLocality", mock.Anything).Return(expectedAllReport, nil)

		reportSlice, err := service.ReportSellersByLocality(context.TODO(), 0)

		assert.Equal(t, expectedAllReport, reportSlice)
		assert.NoError(t, err)
	})

	t.Run("Should return err locality id not found when locality passed not exists", func(t *testing.T) {
		expectedMessage := "locality not found"

		repository, service := InitServer(t)
		repository.On("ExistsById", mock.Anything, 1).Return(false)

		_, err := service.ReportSellersByLocality(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})

	t.Run("Should return err error, try again when ReportLocalityId has an error", func(t *testing.T) {
		expectedMessage := "error, try again %s"

		repository, service := InitServer(t)
		repository.On("ExistsById", mock.Anything, 2).Return(true)
		repository.On("ReportLocalityId", mock.Anything, 2).Return(domain.LocalityReport{}, locality.ErrTryAgain)

		_, err := service.ReportSellersByLocality(context.TODO(), 2)

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})

	t.Run("Should return err error, try again when ReadAllCarriers has an error", func(t *testing.T) {
		expectedMessage := "error, try again %s"

		repository, service := InitServer(t)
		repository.On("ReportLocality", mock.Anything).Return([]domain.LocalityReport{}, locality.ErrTryAgain)

		_, err := service.ReportSellersByLocality(context.TODO(), 0)

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}


func TestExistsById(t *testing.T) {
	t.Run("should return a true value when id exists", func(t *testing.T) {
        repository, service := InitServer(t)
        repository.On("ExistsById", mock.Anything, 1).Return(true)

        err := service.ExistsById(context.TODO(), 1)

		assert.Nil(t, err)
    })
    t.Run("should return error when location does not exist", func(t *testing.T) {
		expectedMessage := "locality does not exists"

        repository, service := InitServer(t)
        repository.On("ExistsById", mock.Anything, 1).Return(false, locality.ErrTryAgain)

        err := service.ExistsById(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
    })
}

func InitServer(t *testing.T) (*mocks.LocalityRepositoryMock, locality.Service) {
	t.Helper()
	mockRepository := &mocks.LocalityRepositoryMock{}
	mockService := locality.NewService(mockRepository)
	return mockRepository, mockService
}