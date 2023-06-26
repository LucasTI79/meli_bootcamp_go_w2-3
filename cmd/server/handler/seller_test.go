package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/seller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	BaseRoute  = "/sellers"
	BaseRouteWithID = "/sellers/1"
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
		server.GET(BaseRoute, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(expectedSellers, nil)
		request, response := testutil.MakeRequest(http.MethodGet, BaseRoute, "")
		server.ServeHTTP(response, request)

		responseResult := &domain.SellerResponse{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.True(t, len(responseResult.Data) == 2)
	})
	t.Run("Should return status 204 with empty sellers", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.GET(BaseRoute, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptySellers, nil)
		request, response := testutil.MakeRequest(http.MethodGet, BaseRoute, "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("Should return 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.GET(BaseRoute, handler.GetAll())

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(emptySellers, seller.ErrTryAgain)
		request, response := testutil.MakeRequest(http.MethodGet, BaseRoute, "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestGetSeller(t *testing.T) {
	server, mockService, handler := InitServer(t)
	t.Run("Should return status 200 with all seller data request by id", func(t *testing.T) {
		expectedSeller := domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		mockService.On("Get", mock.Anything, 1).Return(expectedSeller, nil)
		server.GET("/sellers/:id", handler.Get())
		request, response := testutil.MakeRequest(http.MethodGet, BaseRouteWithID, "")

		server.ServeHTTP(response, request)
		responseResult := &domain.SellerResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedSeller, responseResult.Data)
	})

	t.Run("Should return status 400 when the seller id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		mockService.On("Get", mock.Anything, "invalid").Return(domain.Seller{}, seller.ErrInvalidId)

		server.GET("/sellers/:id", handler.Get())
		request, response := testutil.MakeRequest(http.MethodGet, "/sellers/invalid", "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when the seller id does not exist", func(t *testing.T) {
		server, mockService, handler := InitServer(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.Seller{}, seller.ErrNotFound)
		server.GET("/sellers/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, BaseRouteWithID, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServer(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.Seller{}, seller.ErrTryAgain)
		server.GET("/sellers/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, BaseRouteWithID, "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCreateSeller(t *testing.T) {
	t.Run("Should return status 201 with all sellers", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		requestBody := domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(jsonSeller))
		mockService.On("Save", mock.Anything, mock.Anything).Return(requestBody, nil)
		server.ServeHTTP(response, request)

		responseResult := domain.SellerResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, requestBody, responseResult.Data)
	})
	t.Run("Should return status 400 when CID is invalid", func(t *testing.T) {
		server, _, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		requestBody := domain.Seller{
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(jsonSeller))
		server.ServeHTTP(response, request)

		var responseData web.ErrorResponse
		_ = json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, "cid is required", responseData.Message)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when CompanyName is invalid", func(t *testing.T) {
		server, _, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		requestBody := domain.Seller{
			CID:       1,
			Address:   "Address",
			Telephone: "88748585",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(jsonSeller))
		server.ServeHTTP(response, request)

		var responseData web.ErrorResponse
		_ = json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, "company name is required", responseData.Message)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Address is invalid", func(t *testing.T) {
		server, _, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		requestBody := domain.Seller{
			CID:       1,
			CompanyName: "Company Name",
			Telephone: "88748585",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(jsonSeller))
		server.ServeHTTP(response, request)

		var responseData web.ErrorResponse
		_ = json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, "address is required", responseData.Message)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Telephone is invalid", func(t *testing.T) {
		server, _, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		requestBody := domain.Seller{
			CID:       1,
			CompanyName: "Company Name",
			Address:   "Address",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(jsonSeller))
		server.ServeHTTP(response, request)

		var responseData web.ErrorResponse
		_ = json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, "phone is required", responseData.Message)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(`invalid-json`))
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 409 for existing CID", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		requestBody := domain.Seller{
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(jsonSeller))
		mockService.On("Save", mock.Anything, mock.Anything).Return(domain.Seller{}, seller.ErrCidAlreadyExists)
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.POST(BaseRoute, handler.Create())

		requestBody := domain.Seller{
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRoute, string(jsonSeller))
		mockService.On("Save", mock.Anything, mock.Anything).Return(domain.Seller{}, seller.ErrSaveSeller)
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestDelete(t *testing.T){
	t.Run("Should return 204 and delete seller with id", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.DELETE("/sellers/:id", handler.Delete())

		request, response := testutil.MakeRequest(http.MethodDelete, BaseRouteWithID, "")
		mockService.On("Delete", mock.Anything, 1).Return(nil)

		server.ServeHTTP(response, request)
		
		assert.Equal(t, http.StatusNoContent, response.Code)
	})
	t.Run("Should return status 404 when id seller is not found", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		server.DELETE("/sellers/:id", handler.Delete())

		request, response := testutil.MakeRequest(http.MethodDelete, BaseRouteWithID, "")

		mockService.On("Delete", mock.Anything, 1).Return(seller.ErrNotFound)

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 400 when the seller id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServer(t)

		server.DELETE("/sellers/:id", handler.Delete())

		request, response := testutil.MakeRequest(http.MethodDelete, "/sellers/invalid", "")
		mockService.On("Delete", mock.Anything, "invalid").Return(seller.ErrInvalidId)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServer(t)

		request, response := testutil.MakeRequest(http.MethodDelete, BaseRouteWithID, "")

		mockService.On("Delete", mock.Anything, 1).Return(seller.ErrTryAgain)

		server.DELETE("/sellers/:id", handler.Delete())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
	
}

func TestUpdateSeller(t *testing.T) {
	t.Run("Should return status 200 and updated seller", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		updatedSeller := domain.Seller{
			ID: 1,
			CID:         1,
			CompanyName: "Company Name",
			Address:     "Address",
			Telephone:   "88748585",
		}

		mockService.On("Update", mock.Anything, mock.Anything, 1).Return(updatedSeller, nil)

		request, response := testutil.MakeRequest(http.MethodPatch, BaseRouteWithID, `{"address":"Address","telephone":"88748585"}`)

		server.PATCH("/sellers/:id", handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)

		responseResult := domain.SellerResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, updatedSeller, responseResult.Data)
	})

	t.Run("Should return status 400 when the seller id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServer(t)

		mockService.On("Update", mock.Anything, mock.Anything, "invalid").Return(domain.Seller{}, seller.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodPatch, "/sellers/invalid", `{"address":"Address","telephone":"88748585"}`)

		server.PATCH("/sellers/:id", handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, mockService, handler := InitServer(t)
		
		mockService.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(domain.Seller{}, seller.ErrInvalidBody)
		
		request, response := testutil.MakeRequest(http.MethodPatch, BaseRouteWithID, "")

		server.PATCH("/sellers/:id", handler.Update())
		server.ServeHTTP(response, request)
		
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 404 when seller is not found", func(t *testing.T) {
		server, mockService, handler := InitServer(t)

		request, response := testutil.MakeRequest(http.MethodPatch, BaseRouteWithID, `{"address":"Address","telephone":"88748585"}`)

		mockService.On("Update", mock.Anything, mock.Anything, 1).Return(domain.Seller{}, seller.ErrNotFound)

		server.PATCH("/sellers/:id", handler.Update())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServer(t)

		request, response := testutil.MakeRequest(http.MethodPatch, BaseRouteWithID, `{"telephone":"88748585"}`)

		mockService.On("Update", mock.Anything, mock.Anything, 1).Return(domain.Seller{}, seller.ErrTryAgain)

		server.PATCH("/sellers/:id", handler.Update())
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
