package handler_test

import (
	"encoding/json"
	"errors"
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
	GetAllBuyers    = "/buyers"
	Get             = "/buyers/:id"
	Delete          = "/buyers/:id"
	Create          = "/buyers"
	Update          = "/buyers/:id"
	GetBuyerOrders  = "/buyers/reportPurchaseOrders/:id"
	GetBuyersOrders = "/buyers/reportPurchaseOrders/"
)

func TestGetBuyerOrders(t *testing.T) {
	t.Run("Find by ID status 200 with buyer content", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		expectedBuyers := domain.BuyerOrders{

			ID:                  7,
			CardNumberID:        "1234",
			FirstName:           "Giu",
			LastName:            "Oli",
			PurchaseOrdersCount: 4,
		}

		server.GET(GetBuyerOrders, handler.GetBuyerOrders())

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/reportPurchaseOrders/7", "")

		mockService.On("GetBuyerOrders", mock.Anything, 7).Return(expectedBuyers, nil)

		server.ServeHTTP(response, request)

		responseResult := &domain.BuyerOrdersResponseID{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.EqualValues(t, expectedBuyers, responseResult.Data)

	})

	t.Run("Should return err if id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		mockService.On("GetBuyerOrders", mock.Anything, "invalid").Return(domain.BuyerOrders{}, buyer.ErrInvalidID)

		server.GET(GetBuyerOrders, handler.GetBuyerOrders())

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/reportPurchaseOrders/invalid", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Find by ID status 404 with buyer not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		server.GET(GetBuyerOrders, handler.GetBuyerOrders())

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/reportPurchaseOrders/10", "")

		mockService.On("GetBuyerOrders", mock.Anything, 10).Return(domain.BuyerOrders{}, buyer.ErrNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)

	})

	t.Run("Should return status 500 when any internal error occour", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(GetBuyerOrders, handler.GetBuyerOrders())

		mockService.On("GetBuyerOrders", mock.Anything, 50).Return(domain.BuyerOrders{}, errors.New("error listing buyer"))

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/reportPurchaseOrders/50", "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestGetBuyersOrders(t *testing.T) {
	t.Run("Should return 200 with buyers orders when list is not empty", func(t *testing.T) {
		expectedOrders := []domain.BuyerOrders{
			{ID: 1, CardNumberID: "12345", FirstName: "Giulianna", LastName: "Oliveira", PurchaseOrdersCount: 2},
			{ID: 2, CardNumberID: "12345", FirstName: "Giulianna", LastName: "Oliveira", PurchaseOrdersCount: 2},
		}

		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(GetBuyersOrders, handler.GetBuyersOrders())

		mockService.On("GetBuyersOrders", mock.AnythingOfType("string")).Return(expectedOrders, nil)

		request, response := testutil.MakeRequest(http.MethodGet, GetBuyersOrders, "")
		server.ServeHTTP(response, request)

		responseResult := &domain.BuyerOrdersResponse{}
		err := json.Unmarshal(response.Body.Bytes(), responseResult)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)

		assert.Equal(t, expectedOrders, responseResult.Data)

	})

	t.Run("Should return status 204 with no content", func(t *testing.T) {
		emptyBuyers := make([]domain.BuyerOrders, 0)
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(GetBuyersOrders, handler.GetBuyersOrders())

		mockService.On("GetBuyersOrders", mock.AnythingOfType("string")).Return(emptyBuyers, nil)

		request, response := testutil.MakeRequest(http.MethodGet, GetBuyersOrders, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("Should return status 500 when any internal error occour", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(GetBuyersOrders, handler.GetBuyersOrders())

		mockService.On("GetBuyersOrders", mock.AnythingOfType("string")).Return([]domain.BuyerOrders{}, domain.ErrNotFound)

		request, response := testutil.MakeRequest(http.MethodGet, GetBuyersOrders, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestGetAllBuyers(t *testing.T) {
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
		server.GET(GetAllBuyers, handler.GetAll())

		request, response := testutil.MakeRequest(http.MethodGet, GetAllBuyers, "")
		mockBuyer.On("GetAll", mock.AnythingOfType("string")).Return(expectedBuyers, nil)
		server.ServeHTTP(response, request)

		responseResult := &domain.BuyerResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusOK, response.Code)

		assert.Equal(t, expectedBuyers, responseResult.Data)

		assert.True(t, len(responseResult.Data) == 2)

	})

	t.Run("Should return status 204 with no content", func(t *testing.T) {
		emptyBuyers := make([]domain.Buyer, 0)
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(GetAllBuyers, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptyBuyers, nil)

		request, response := testutil.MakeRequest(http.MethodGet, GetAllBuyers, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("Should return status 500 when any internal error occour", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(GetAllBuyers, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return([]domain.Buyer{}, domain.ErrNotFound)

		request, response := testutil.MakeRequest(http.MethodGet, GetAllBuyers, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestGetBuyers(t *testing.T) {
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

	t.Run("Should return err if id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		mockService.On("Get", mock.Anything, "invalid").Return(domain.Buyer{}, buyer.ErrInvalidID)

		server.GET(Get, handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/invalid", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Find by ID status 404 with buyer not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		server.GET(Get, handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/1", "")

		mockService.On("Get", mock.Anything, 1).Return(domain.Buyer{}, buyer.ErrNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)

	})

	t.Run("Should return status 500 when any internal error occour", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.GET(Get, handler.Get())

		mockService.On("Get", mock.Anything, mock.Anything).Return(domain.Buyer{}, errors.New("error listing buyer"))

		request, response := testutil.MakeRequest(http.MethodGet, "/buyers/10", "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestDeleteBuyers(t *testing.T) {
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

	t.Run("Should return err if id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		mockService.On("Delete", mock.Anything, "invalid").Return(domain.Buyer{}, buyer.ErrInvalidID)

		server.DELETE(Delete, handler.Delete())

		request, response := testutil.MakeRequest(http.MethodDelete, "/buyers/invalid", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return status 500 when any internal error occour", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)
		server.DELETE(Delete, handler.Delete())

		mockService.On("Delete", mock.Anything, 50).Return(errors.New("error listing buyer"))

		request, response := testutil.MakeRequest(http.MethodDelete, "/buyers/50", "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCreateBuyers(t *testing.T) {
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

	t.Run("Should return status 422 when JSON is empty", func(t *testing.T) {
		server, _, handler := InitServerWithGetBuyers(t)

		server.POST(Create, handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, Create, `{"card_number_id":""}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("Should return status 400 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithGetBuyers(t)

		server.POST(Create, handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, Create, `{"card_number_id":2}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
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

func TestUpdateBuyers(t *testing.T) {
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

	t.Run("Should return err 400 if id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		mockService.On("Update", mock.Anything, mock.Anything, "invalid").Return(domain.Buyer{}, buyer.ErrInvalidID)

		server.PATCH(Update, handler.Update())

		request, response := testutil.MakeRequest(http.MethodPatch, "/buyers/invalid", `{
			"card_number_id": "5435",
			"first_name": "Giulianna",
			"last_name": "Goncalves"
		  }`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return status 500 when internal error", func(t *testing.T) {
		server, _, handler := InitServerWithGetBuyers(t)

		server.PATCH(Update, handler.Update())

		request, response := testutil.MakeRequest(http.MethodPatch, "/buyers/10", `{"first_name":""}`)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetBuyers(t)

		mockService.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(domain.Buyer{}, buyer.ErrInvalidBody)

		request, response := testutil.MakeRequest(http.MethodPatch, "/buyers/2", "")

		server.PATCH(Update, handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

}

func InitServerWithGetBuyers(t *testing.T) (*gin.Engine, *mocks.BuyerServiceMock, *handler.BuyerController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.BuyerServiceMock)
	handler := handler.NewBuyer(mockService)
	return server, mockService, handler
}
