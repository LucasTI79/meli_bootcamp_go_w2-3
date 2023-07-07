package carry_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/carry"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/carry"
	mocksLocality "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/locality"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var id = 2

var expectedCarry = domain.Carry{
	Cid:         "1111110",
	CompanyName: "Teste Livre",
	Address:     "Rua Pedro Dias",
	Telephone:   "3712291281",
	LocalityId:  1,
}

var expectedLocalityCarriersReport = domain.LocalityCarriersReport{
	LocalityID:    2,
	LocalityName:  "São Paulo",
	CarriersCount: 5,
}

var expectedLocalitiesCarriersReport = []domain.LocalityCarriersReport{
	{
		LocalityID:    2,
		LocalityName:  "São Paulo",
		CarriersCount: 5,
	},
	{
		LocalityID:    3,
		LocalityName:  "Marajó",
		CarriersCount: 10,
	},
}

func TestCreateCarriers(t *testing.T) {
	t.Run("Should create the carry if it contains the required fields", func(t *testing.T) {
		repository, repositoryLocality, service := InitServerWithCarriersRepository(t)
		repository.On("ExistsByCidCarry", mock.Anything, "1111110").Return(false)
		repositoryLocality.On("ExistsById", mock.Anything, 1).Return(true)
		repository.On("Create", mock.Anything, expectedCarry).Return(id, nil)

		carry, err := service.Create(context.TODO(), expectedCarry)

		assert.Equal(t, 2, carry.ID)
		assert.Equal(t, "1111110", carry.Cid)
		assert.Equal(t, "Teste Livre", carry.CompanyName)
		assert.Equal(t, "Rua Pedro Dias", carry.Address)
		assert.Equal(t, "3712291281", carry.Telephone)
		assert.Equal(t, 1, carry.LocalityId)

		assert.NoError(t, err)
	})
	t.Run("Should return err carry already exists when carry already exists", func(t *testing.T) {
		expectedMessage := "carry already exists"

		repository, _, service := InitServerWithCarriersRepository(t)

		repository.On("ExistsByCidCarry", mock.Anything, mock.Anything).Return(true)

		_, err := service.Create(context.TODO(), domain.Carry{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return err locality_id not found when locality not exists", func(t *testing.T) {
		expectedMessage := "locality_id not found"

		repository, repositoryLocality, service := InitServerWithCarriersRepository(t)

		repository.On("ExistsByCidCarry", mock.Anything, mock.Anything).Return(false)
		repositoryLocality.On("ExistsById", mock.Anything, mock.Anything).Return(false)

		_, err := service.Create(context.TODO(), domain.Carry{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is an save repository error", func(t *testing.T) {
		repository, repositoryLocality, service := InitServerWithCarriersRepository(t)

		repository.On("ExistsByCidCarry", mock.Anything, mock.Anything).Return(false)
		repositoryLocality.On("ExistsById", mock.Anything, mock.Anything).Return(true)

		expectedError := errors.New("some error")
		repository.On("Create", mock.Anything, domain.Carry{}).Return(0, expectedError)

		_, err := service.Create(context.TODO(), domain.Carry{})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestReadCarriers(t *testing.T) {
	t.Run("Should get the locality with their carriers when the id is passed and it exists in database", func(t *testing.T) {
		var reportExpected []domain.LocalityCarriersReport
		reportExpected = append(reportExpected, expectedLocalityCarriersReport)

		repository, repositoryLocality, service := InitServerWithCarriersRepository(t)
		repositoryLocality.On("ExistsById", mock.Anything, 2).Return(true)
		repository.On("ReadCarriersWithLocalityId", mock.Anything, 2).Return(expectedLocalityCarriersReport, nil)

		localityCarriersReportSlice, err := service.Read(context.TODO(), 2)

		assert.Equal(t, reportExpected, localityCarriersReportSlice)
		assert.NoError(t, err)
	})
	t.Run("Should get all the localities with their carriers when the id is not passed and it exists in database", func(t *testing.T) {
		repository, _, service := InitServerWithCarriersRepository(t)
		repository.On("ReadAllCarriers", mock.Anything).Return(expectedLocalitiesCarriersReport, nil)

		localityCarriersReportSlice, err := service.Read(context.TODO(), 0)

		assert.Equal(t, expectedLocalitiesCarriersReport, localityCarriersReportSlice)
		assert.NoError(t, err)
	})
	t.Run("Should return err locality_id not found when locality passed not exists", func(t *testing.T) {
		expectedMessage := "locality_id not found"

		_, repositoryLocality, service := InitServerWithCarriersRepository(t)
		repositoryLocality.On("ExistsById", mock.Anything, 2).Return(false)

		_, err := service.Read(context.TODO(), 2)

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return err error, try again when ReadCarriersWithLocalityId has an error", func(t *testing.T) {
		expectedMessage := "error, try again %s"

		repository, repositoryLocality, service := InitServerWithCarriersRepository(t)
		repositoryLocality.On("ExistsById", mock.Anything, 2).Return(true)
		repository.On("ReadCarriersWithLocalityId", mock.Anything, 2).Return(domain.LocalityCarriersReport{}, carry.ErrTryAgain)

		_, err := service.Read(context.TODO(), 2)

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return err error, try again when ReadAllCarriers has an error", func(t *testing.T) {
		expectedMessage := "error, try again %s"

		repository, _, service := InitServerWithCarriersRepository(t)
		repository.On("ReadAllCarriers", mock.Anything).Return([]domain.LocalityCarriersReport{}, carry.ErrTryAgain)

		_, err := service.Read(context.TODO(), 0)

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestGetCarriers(t *testing.T) {
	t.Run("Should get the carry when it exists in database", func(t *testing.T) {
		repository, _, service := InitServerWithCarriersRepository(t)
		repository.On("Get", mock.Anything, id).Return(expectedCarry, nil)

		carry, err := service.Get(context.TODO(), 2)

		assert.Equal(t, expectedCarry, carry)
		assert.NoError(t, err)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		repository, _, service := InitServerWithCarriersRepository(t)

		expectedError := errors.New("carry not found")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.Carry{}, carry.ErrNotFound)

		_, err := service.Get(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("Should return error when there is an get repository error", func(t *testing.T) {
		repository, _, service := InitServerWithCarriersRepository(t)

		expectedError := errors.New("some error")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.Carry{}, expectedError)

		_, err := service.Get(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func InitServerWithCarriersRepository(t *testing.T) (*mocks.CarryRepositoryMock, *mocksLocality.LocalityRepositoryMock, carry.Service) {
	t.Helper()
	mockRepositoryCarriers := &mocks.CarryRepositoryMock{}
	mockRepositoryLocalities := &mocksLocality.LocalityRepositoryMock{}
	mockService := carry.NewService(mockRepositoryCarriers, mockRepositoryLocalities)
	return mockRepositoryCarriers, mockRepositoryLocalities, mockService
}
