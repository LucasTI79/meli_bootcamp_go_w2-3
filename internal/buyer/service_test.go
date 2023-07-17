package buyer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/buyer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	t.Run("Should return all buyers when repository is called", func(t *testing.T) {
		expectedBuyers := []domain.Buyer{
			{
				ID:           9,
				CardNumberID: "2556",
				FirstName:    "Giulianna",
				LastName:     "Oliveira",
			},
		}

		repository, service := InitServerWithBuyersRepository(t)
		repository.On("GetAll", mock.Anything).Return(expectedBuyers, nil)

		_, err := service.GetAll(context.TODO())

		assert.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Should create the buyer", func(t *testing.T) {
		id := 10
		expectedBuyer := domain.Buyer{
			CardNumberID: "138935",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		repository, service := InitServerWithBuyersRepository(t)
		repository.On("ExistsBuyer", mock.Anything, "138935").Return(false)
		repository.On("Save", mock.Anything, expectedBuyer).Return(id, nil)

		buyer, err := service.Create(context.TODO(), expectedBuyer)

		assert.Equal(t, "138935", buyer.CardNumberID)
		assert.Equal(t, "Giulianna", buyer.FirstName)
		assert.Equal(t, "Oliveira", buyer.LastName)

		assert.NoError(t, err)
	})
	t.Run("Should return err card_id already exists", func(t *testing.T) {
		expectedMessage := "buyer already exists"
		repository, service := InitServerWithBuyersRepository(t)
		repository.On("ExistsBuyer", mock.Anything, mock.Anything).Return(true)

		_, err := service.Create(context.TODO(), domain.Buyer{})

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestGetById(t *testing.T) {
	t.Run("Should get the buyer if exists", func(t *testing.T) {
		expectedBuyer := domain.Buyer{
			ID:           9,
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		repository, service := InitServerWithBuyersRepository(t)
		repository.On("Get", mock.Anything, expectedBuyer.ID).Return(expectedBuyer, nil)

		buyer, err := service.Get(context.TODO(), 9)

		assert.Equal(t, expectedBuyer, buyer)
		assert.NoError(t, err)
	})
	t.Run("Should return null if id not exist", func(t *testing.T) {

		repository, service := InitServerWithBuyersRepository(t)
		expectedError := errors.New("buyer not found")
		repository.On("Get", mock.Anything, mock.Anything).Return(domain.Buyer{}, buyer.ErrNotFound)

		_, err := service.Get(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Should delete the buyer when it exists", func(t *testing.T) {
		expectedBuyer := domain.Buyer{
			ID:           9,
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		repository, service := InitServerWithBuyersRepository(t)
		repository.On("Delete", mock.Anything, expectedBuyer.ID).Return(nil)

		err := service.Delete(context.TODO(), 9)

		assert.NoError(t, err)
	})
	t.Run("Should return nill when buyer dont exists", func(t *testing.T) {
		repository, service := InitServerWithBuyersRepository(t)

		expectedError := errors.New("buyer not found")
		repository.On("Delete", mock.Anything, mock.Anything).Return(buyer.ErrNotFound)

		err := service.Delete(context.TODO(), 19)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Should update the buyer", func(t *testing.T) {
		expectedBuyer := domain.Buyer{
			ID:           9,
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		repository, service := InitServerWithBuyersRepository(t)
		repository.On("Get", mock.Anything, expectedBuyer.ID).Return(expectedBuyer, nil)
		repository.On("ExistsID", mock.Anything, expectedBuyer.ID).Return(false)
		repository.On("Update", mock.Anything, expectedBuyer).Return(nil)

		updatedBuyer, err := service.Update(context.TODO(), expectedBuyer, expectedBuyer.ID)

		assert.NoError(t, err)
		assert.Equal(t, expectedBuyer, updatedBuyer)
	})
	t.Run("should not update buyer if not exists", func(t *testing.T) {
		repository, service := InitServerWithBuyersRepository(t)

		expectedBuyer := domain.Buyer{
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		repository.On("Get", mock.Anything, 100).Return(domain.Buyer{}, errors.New("error"))
		repository.On("Update", mock.Anything, mock.Anything).Return(errors.New("error"))
		_, err := service.Update(context.TODO(), expectedBuyer, 100)
		assert.Error(t, err)
	})
}

func TestExistsID(t *testing.T) {
	t.Run("Should return err if buyer not exists", func(t *testing.T) {
		mockRepository, service := InitServerWithBuyersRepository(t)
		mockRepository.On("ExistsID", 20).Return(false)

		err := service.ExistsID(context.TODO(), 20)

		assert.Error(t, err)
	})

	t.Run("should return nil if section exists", func(t *testing.T) {
		mockRepository, service := InitServerWithBuyersRepository(t)
		mockRepository.On("ExistsID", 1).Return(true)

		err := service.ExistsID(context.TODO(), 1)

		assert.NoError(t, err)
	})
}

func TestGetBuyerOrders(t *testing.T) {
	expectedBuyers := domain.BuyerOrders{

		ID:                  7,
		CardNumberID:        "1234",
		FirstName:           "Giu",
		LastName:            "Oli",
		PurchaseOrdersCount: 4,
	}
	t.Run("Should return buyer orders", func(t *testing.T) {
		mockRepository, service := InitServerWithBuyersRepository(t)

		mockRepository.On("GetBuyerOrders", mock.Anything, 7).Return(expectedBuyers, nil)

		buyerOrder, err := service.GetBuyerOrders(context.Background(), 7)
		assert.Equal(t, expectedBuyers, buyerOrder)
		assert.NoError(t, err)
	})

	t.Run("Should not return buyer orders", func(t *testing.T) {
		mockRepository, service := InitServerWithBuyersRepository(t)

		mockRepository.On("GetBuyerOrders", mock.Anything, 1).Return(domain.BuyerOrders{}, errors.New("error"))

		buyerOrder, err := service.GetBuyerOrders(context.Background(), 1)
		assert.Equal(t, domain.BuyerOrders{}, buyerOrder)
		assert.Error(t, err)
	})
}

func TestGetBuyersOrders(t *testing.T) {
	expectedBuyers := []domain.BuyerOrders{
		{ID: 1, CardNumberID: "12345", FirstName: "Giulianna", LastName: "Oliveira", PurchaseOrdersCount: 2},
		{ID: 2, CardNumberID: "12345", FirstName: "Giulianna", LastName: "Oliveira", PurchaseOrdersCount: 2},
	}
	t.Run("Should return buyers orders", func(t *testing.T) {
		mockRepository, service := InitServerWithBuyersRepository(t)

		mockRepository.On("GetBuyersOrders").Return(expectedBuyers, nil)

		buyersOrders, err := service.GetBuyersOrders(context.Background())
		assert.Equal(t, expectedBuyers, buyersOrders)
		assert.NoError(t, err)
	})

	t.Run("Should not return buyers orders", func(t *testing.T) {
		mockRepository, service := InitServerWithBuyersRepository(t)

		mockRepository.On("GetBuyersOrders", mock.Anything).Return([]domain.BuyerOrders{}, errors.New("error"))

		buyersOrders, err := service.GetBuyersOrders(context.Background())
		assert.Equal(t, []domain.BuyerOrders{}, buyersOrders)
		assert.Error(t, err)
	})
}

func InitServerWithBuyersRepository(t *testing.T) (*mocks.BuyerRepositoryMock, buyer.Service) {
	t.Helper()
	mockRepository := &mocks.BuyerRepositoryMock{}
	mockService := buyer.NewService(mockRepository)
	return mockRepository, mockService
}
