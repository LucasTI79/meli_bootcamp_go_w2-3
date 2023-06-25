package handler_test

/*
CREATE create_ok Quando a entrada de dados for
bem-sucedida, um código 201 será
retornado junto com o objeto inserido.
201

CREATE create_fail Se o objeto JSON não contiver os campos
necessários, um código 422 será
retornado.
422

CREATE create_conflict Se o card_number_id já existir, ele retornará um erro 409 Conflict.
409

UPDATE update_ok Quando a atualização dos dados for bem
sucedida, o comprador será devolvido com
as informações atualizadas juntamente
com um código 200
200

UPDATE update_non_existent Se o comprador a ser atualizado não
existir, um código 404 será devolvido
404

DELETE delete_non_existent Quando o comprador não existir, um código 404 será devolvido
404
*/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/buyer"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GetAll = "/buyers"
	Get    = "/buyers/:id"
	Delete = "/buyers/:id"
)

func TestGetAll(t *testing.T) {
	t.Run("Should return status 200 with all buyers", func(t *testing.T) {
		server, mockBuyer, handler := InitServerWithGetBuyers(t)
		expectedBuyers := []domain.Buyer{
			{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "Giulianna",
				LastName:     "Oliveira",
			},
			{
				ID:           2,
				CardNumberID: "1234",
				FirstName:    "Giu",
				LastName:     "Oli",
			},
		}
		server.GET(GetAll, handler.GetAll())

		request, response := testutil.MakeRequest(http.MethodGet, GetAll, "")
		mockBuyer.On("GetAll", mock.AnythingOfType("string")).Return(expectedBuyers, nil)
		server.ServeHTTP(response, request)

		responseResult := &domain.BuyerResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusOK, response.Code)
		fmt.Println(responseResult)

		assert.Equal(t, expectedBuyers, responseResult.Data)

		assert.True(t, len(responseResult.Data) == 2)

	})

	t.Run("Should return status 204 with no content", func(t *testing.T) {
		emptyBuyers := make([]domain.Buyer, 0)
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(GetAll, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptyBuyers, nil)

		request, response := testutil.MakeRequest(http.MethodGet, GetAll, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func TestGet(t *testing.T) {
	t.Run("Find by ID status 200 with buyer content", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		expectedBuyers := domain.Buyer{

			ID:           7,
			CardNumberID: "1234",
			FirstName:    "Giu",
			LastName:     "Oli",
		}

		server.GET(Get, handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/7", "")

		mockService.On("Get", mock.Anything, 7).Return(expectedBuyers, nil)

		server.ServeHTTP(response, request)

		responseResult := &domain.BuyerResponseID{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.EqualValues(t, expectedBuyers, responseResult.Data)

	})

	t.Run("Find by ID status 404 with buyer not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		server.GET(Get, handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/1", "")

		mockService.On("Get", mock.Anything, 1).Return(domain.Buyer{}, buyer.ErrNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)

	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete status 204 successful with no content", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		mockService.On("Delete", mock.Anything, 7).Return(nil)

		request, response := testutil.MakeRequest(http.MethodDelete, "/buyers/7", "")

		server.DELETE(Delete, handler.Delete())

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)

	})
}

func InitServerWithGetBuyers(t *testing.T) (*gin.Engine, *mocks.BuyerServiceMock, *handler.BuyerController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.BuyerServiceMock)
	handler := handler.NewBuyer(mockService)
	return server, mockService, handler
}
