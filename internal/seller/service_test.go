package seller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/seller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllSelllers(t *testing.T) {
	t.Run("Should return all sellers when repository is called", func(t *testing.T) {
		expectedSellers := []domain.Seller{
			{
				ID:          1,
				CID:         1,
				CompanyName: "Company Name",
				Address:     "Address",
				Telephone:   "88748585",
			},
			{
				ID:          2,
				CID:         2,
				CompanyName: "Company Name2",
				Address:     "Address2",
				Telephone:   "12345698",
			},
		}

		repository, service := InitServerRepository(t)
		repository.On("GetAll", mock.Anything).Return(expectedSellers, nil)

		sellers, err := service.GetAll(context.TODO())

		assert.True(t, len(sellers) == 2)
		assert.NoError(t, err)
	})
}

func TestGetByIdSellers(t *testing.T) {
	t.Run("Should get the seller when it exists in database", func(t *testing.T) {
		expectedSellers := domain.Seller{
            ID: 1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		repository, service := InitServerRepository(t)
		repository.On("Get", mock.Anything, expectedSellers.ID).Return(expectedSellers, nil)

		seller, err := service.Get(context.TODO(), 1)

		assert.Equal(t, expectedSellers, seller)
		assert.NoError(t, err)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		repository, service := InitServerRepository(t)

		expectedError := errors.New("seller not found")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.Seller{}, seller.ErrNotFound)

		_, err := service.Get(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("Should return error when there is an get repository error", func(t *testing.T) {
		repository, service := InitServerRepository(t)

		expectedError := errors.New("some error")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.Seller{}, expectedError)

		_, err := service.Get(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestCreateSellers(t *testing.T) {
	t.Run("Should create the sellers if it contains the required fields", func(t *testing.T) {
		id := 1
		expectedSellers := domain.Seller{
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		repository, service := InitServerRepository(t)
		repository.On("Exists", mock.Anything, 1).Return(false)
		repository.On("Save", mock.Anything, expectedSellers).Return(id, nil)

		seller, err := service.Save(context.TODO(), expectedSellers)

		assert.Equal(t, 1, seller.ID)
		assert.Equal(t, 1, seller.CID)
		assert.Equal(t, "Company Name", seller.CompanyName)
		assert.Equal(t, "Address", seller.Address)
		assert.Equal(t, "88748585", seller.Telephone)

		assert.NoError(t, err)
	})
	t.Run("Should return err when CID seller already exists", func(t *testing.T) {
		expectedMessage := "cid already registered"

		repository, service := InitServerRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(true)

		_, err := service.Save(context.TODO(), domain.Seller{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is an save repository error", func(t *testing.T) {
		repository, service := InitServerRepository(t)

		repository.On("Exists", mock.Anything, mock.Anything).Return(false)

		expectedError := errors.New("error saving seller")
		repository.On("Save", mock.Anything, domain.Seller{}).Return(0, expectedError)

		_, err := service.Save(context.TODO(), domain.Seller{})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestDeleteSellers(t *testing.T) {
	t.Run("Should delete the seller when it exists in database", func(t *testing.T) {
		expectedSellers := domain.Seller{
            ID: 1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}

		repository, service := InitServerRepository(t)
		repository.On("Delete", mock.Anything, expectedSellers.ID).Return(nil)

		err := service.Delete(context.TODO(), 1)

		assert.NoError(t, err)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		repository, service := InitServerRepository(t)

		expectedError := errors.New("seller not found")
		repository.On("Delete", mock.Anything, mock.Anything).Return(seller.ErrNotFound)

		err := service.Delete(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("Should return error when there is an delete repository error", func(t *testing.T) {
		repository, service := InitServerRepository(t)

		expectedError := errors.New("some error")
		repository.On("Delete", mock.Anything, mock.Anything).Return(expectedError)

		err := service.Delete(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestUpdateSellers(t *testing.T) {
	t.Run("Should update the seller when it exists in database", func(t *testing.T) {
		expectedSellers := domain.Seller{
            ID: 1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}

		repository, service := InitServerRepository(t)
		repository.On("Get", mock.Anything, expectedSellers.ID).Return(expectedSellers, nil)
		repository.On("Exists", mock.Anything, expectedSellers.CID).Return(false)
		repository.On("Update", mock.Anything, expectedSellers).Return(nil)

		updatedSeller, err := service.Update(context.TODO(),expectedSellers.ID, expectedSellers )

		assert.NoError(t, err)
		assert.Equal(t, expectedSellers, updatedSeller)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedSeller := domain.Seller{
            ID: 1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}

		repository, service := InitServerRepository(t)

		expectedError := errors.New("seller not found")
		repository.On("Get", mock.Anything, expectedSeller.ID).Return(domain.Seller{}, expectedError)

		updatedSeller, err := service.Update(context.TODO(), expectedSeller.ID, expectedSeller)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, domain.Seller{}, updatedSeller)
	})
	t.Run("Should return err seller already exists when seller already exists", func(t *testing.T) {
		Seller := domain.Seller{
            ID: 1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		updateSeller := domain.Seller{
			CID:      2,
		}

		repository, service := InitServerRepository(t)

		expectedError := errors.New("cid already registered")
		repository.On("Get", mock.Anything, Seller.ID).Return(Seller, nil)
		repository.On("Exists", mock.Anything, updateSeller.CID).Return(true)

		updatedSeller, err := service.Update(context.TODO(), 1, updateSeller)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, domain.Seller{}, updatedSeller)
	})

	t.Run("Should return error when there is an update repository error", func(t *testing.T) {
		expectedSeller := domain.Seller{
            ID: 1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}

		repository, service := InitServerRepository(t)

		expectedError := errors.New("some error")
		repository.On("Get", mock.Anything, expectedSeller.ID).Return(expectedSeller, nil)
		repository.On("Exists", mock.Anything, expectedSeller.CID).Return(false)
		repository.On("Update", mock.Anything, expectedSeller).Return(expectedError)

		_, err := service.Update(context.TODO(), expectedSeller.ID, expectedSeller)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func InitServerRepository(t *testing.T) (*mocks.SellerRepositoryMock, seller.Service) {
	t.Helper()
	mockRepository := &mocks.SellerRepositoryMock{}
	mockService := seller.NewService(mockRepository)
	return mockRepository, mockService
}

