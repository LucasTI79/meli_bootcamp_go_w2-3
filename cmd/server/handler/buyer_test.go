package handler_test

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
	Create = "/buyers"
	Update = "/buyers/:id"
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

	t.Run("Should return status 404 when buyer is not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		request, response := testutil.MakeRequest(http.MethodDelete, "/buyers/1", "")

		mockService.On("Delete", mock.Anything, 1).Return(buyer.ErrNotFound)

		server.DELETE(Delete, handler.Delete())

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Should return status 201 with the buyer created", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)
		expectedBuyers := domain.Buyer{
			ID:           9,
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		mockService.On("Create", mock.Anything, mock.Anything).Return(expectedBuyers, nil)

		server.POST(Create, handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, Create, `{
			"card_number_id":"2556",
			"first_name":"Giulianna",
			"last_name":"Oliveira"}`)
		server.ServeHTTP(response, request)

		responseResult := domain.BuyerResponseID{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, expectedBuyers, responseResult.Data)
	})

	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithGetBuyers(t)

		server.POST(Create, handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, Create, `{"card_number_id":""}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("Should return status 409 when buyer already exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		request, response := testutil.MakeRequest(http.MethodPost, Create, `{
			"card_number_id":"1234",
			"first_name":"Giu",
			"last_name":"Oli"}`)

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.Buyer")).Return(domain.Buyer{}, buyer.ErrExists)

		server.POST(Create, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Should return status 200 and updated buyer", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)
		updatedBuyer := domain.Buyer{
			ID:           8,
			CardNumberID: "5435",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}
		mockService.On("Update", mock.Anything, mock.Anything, 8).Return(updatedBuyer, nil)

		request, response := testutil.MakeRequest(http.MethodPatch, "/buyers/8", `{
			"first_name": "Giulianna",
			"last_name": "Goncalves"
		  }`)

		server.PATCH(Update, handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)

		responseResult := domain.BuyerResponseID{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, updatedBuyer, responseResult.Data)
	})
	t.Run("Should return status 404 when buyer not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		request, response := testutil.MakeRequest(http.MethodPatch, "/buyers/10", `{
			"card_number_id": "5435",
			"first_name": "Giulianna",
			"last_name": "Goncalves"
		  }`)

		mockService.On("Update", mock.Anything, mock.Anything, 10).Return(domain.Buyer{}, buyer.ErrNotFound)

		server.PATCH(Update, handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func InitServerWithGetBuyers(t *testing.T) (*gin.Engine, *mocks.BuyerServiceMock, *handler.BuyerController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.BuyerServiceMock)
	handler := handler.NewBuyer(mockService)
	return server, mockService, handler
}
