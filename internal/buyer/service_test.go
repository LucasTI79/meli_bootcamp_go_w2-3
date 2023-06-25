package buyer_test

/*
CREATE create_ok Se contiver os campos necessários, será criado

CREATE create_conflict Se o card_number_id já existir, ele não pode ser criado

READ find_all Se a lista tiver "n" elementos, retornará o número totalde elementos

READ find_by_id_non_existent Se o elemento procurado por id não existir, retorna null

READ find_by_id_existent Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado

UPDATE update_existent Quando a atualização dos dados for bem sucedida, o
comprador será devolvido com as informações
atualizadas

UPDATE update_non_existent Se o comprador a ser atualizado não existir, será retornado null.

DELETE delete_non_existent Quando o comprador não existir, será devolvido null.

DELETE delete_ok Se a exclusão for bem-sucedida, o item não aparecerá na lista.
*/
import (
	"context"
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

func InitServerWithBuyersRepository(t *testing.T) (*mocks.BuyerRepositoryMock, buyer.Service) {
	t.Helper()
	mockRepository := &mocks.BuyerRepositoryMock{}
	mockService := buyer.NewService(mockRepository)
	return mockRepository, mockService
}
