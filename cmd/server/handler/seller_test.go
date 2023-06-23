package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/seller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GetAllSellers = "/sellers"
	GetByIdSellers = "/sellers/1"
)

func TestGetAllSeller(t *testing.T) {
	emptySellers := make([]domain.Seller, 0)
	t.Run("Should return status 200 with all sellers", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
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
		server.GET(GetAllSellers, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(expectedSellers, nil)
		request, response := testutil.MakeRequest(http.MethodGet, GetAllSellers, "")
		server.ServeHTTP(response, request)

		responseResult := &domain.SellerResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.True(t, len(responseResult.Data) == 2)
	})
	t.Run("Should return status 204 with empty sellers", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.GET(GetAllSellers, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptySellers, nil)
		request, response := testutil.MakeRequest(http.MethodGet, GetAllSellers, "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("Should return 500", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.GET(GetAllSellers, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptySellers, seller.ErrTryAgain)
		request, response := testutil.MakeRequest(http.MethodGet, GetAllSellers, "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}	

func TestGetSeller(t *testing.T) {
	server, mockService, handler := InitServer(t)
	t.Run("Should return status 200 with all seller data request by id", func(t *testing.T){
		expectedSeller := domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		mockService.On("Get", mock.Anything, 1).Return(expectedSeller, nil)
		server.GET("/sellers/:id", handler.Get())
		request, response := testutil.MakeRequest(http.MethodGet, GetByIdSellers, "")
	
		server.ServeHTTP(response, request)
		responseResult := &domain.SellerResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedSeller, responseResult.Data)
	})

	t.Run("Should return status 400", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		mockService.On("Get", mock.Anything, "invalid").Return(domain.Seller{}, seller.ErrInvalidId)

		server.GET("/sellers/:id", handler.Get())
		request, response := testutil.MakeRequest(http.MethodGet, "/sellers/invalid", "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404", func(t *testing.T) {
		server, mockService, handler := InitServer(t)   

		mockService.On("Get", mock.Anything, 1).Return(domain.Seller{}, seller.ErrNotFound)
		server.GET("/sellers/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, GetByIdSellers, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500", func(t *testing.T) {

		server, mockService, handler := InitServer(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.Seller{}, seller.ErrTryAgain)
		server.GET("/sellers/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, GetByIdSellers, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func InitServer(t *testing.T) (*gin.Engine, *mocks.SellerServiceMock, *handler.SellerController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.SellerServiceMock)
	handler := handler.NewSeller(mockService)
	return server, mockService, handler
}
