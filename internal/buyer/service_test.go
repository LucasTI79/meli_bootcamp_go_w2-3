package buyer_test

/*

UPDATE update_existent Quando a atualização dos dados for bem sucedida, o
comprador será devolvido com as informações
atualizadas

UPDATE update_non_existent Se o comprador a ser atualizado não existir, será retornado null.
*/
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
		repository.On("Exists", mock.Anything, "138935").Return(false)
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
		repository.On("Exists", mock.Anything, mock.Anything).Return(true)

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

func TestDeleteWarehouses(t *testing.T) {
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

func InitServerWithBuyersRepository(t *testing.T) (*mocks.BuyerRepositoryMock, buyer.Service) {
	t.Helper()
	mockRepository := &mocks.BuyerRepositoryMock{}
	mockService := buyer.NewService(mockRepository)
	return mockRepository, mockService
}
