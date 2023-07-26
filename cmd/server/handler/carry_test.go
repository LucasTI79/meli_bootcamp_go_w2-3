package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/middlewares"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/carry"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/carry"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	BaseEndpointCarriers                    = "/carriers"
	BaseEndpointWithIdCarriers              = "/carriers/:id"
	BaseEndpointWithLocalityIdQueryCarriers = "/localities/reportCarriers?id=1"
	BaseEndpointLocalityCarriers            = "/localities/reportCarriers"
)

var expectedCarry = domain.Carry{
	ID:          1,
	Cid:         "1111110",
	CompanyName: "Teste Livre",
	Address:     "Rua Pedro Dias",
	Telephone:   "3712291281",
	LocalityId:  1,
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

func TestGetCarriers(t *testing.T) {
	t.Run("Should return status 200 and carry with id", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Get", mock.Anything, 1).Return(expectedCarry, nil)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointCarriers+"/1", "")

		server.GET(BaseEndpointWithIdCarriers, middlewares.ValidateParams("id"), handler.Get())
		server.ServeHTTP(response, request)

		responseResult := &domain.CarryResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedCarry, responseResult.Data)
	})
	t.Run("Should return status 400 when the carry id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Get", mock.Anything, mock.Anything).Return(domain.Carry{}, carry.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointCarriers+"/invalid", "")

		server.GET(BaseEndpointWithIdCarriers, middlewares.ValidateParams("id"), handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when the carry id does not exist", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.Carry{}, carry.ErrNotFound)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointCarriers+"/1", "")

		server.GET(BaseEndpointWithIdCarriers, middlewares.ValidateParams("id"), handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {

		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Get", mock.Anything, 1).Return(domain.Carry{}, carry.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointCarriers+"/1", "")

		server.GET(BaseEndpointWithIdCarriers, middlewares.ValidateParams("id"), handler.Get())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestReadCarriers(t *testing.T) {
	t.Run("Should return status 200 and locality with his carriers when id is passed", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Read", mock.Anything, 1).Return(expectedLocalitiesCarriersReport, nil)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointWithLocalityIdQueryCarriers, "")

		server.GET(BaseEndpointLocalityCarriers, handler.Read())
		server.ServeHTTP(response, request)

		responseResult := &domain.LocalityCarriersResponse{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedLocalitiesCarriersReport, responseResult.Data)
	})
	t.Run("Should return status 400 when the locality id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Read", mock.Anything, "invalid").Return([]domain.LocalityCarriersReport{}, carry.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointLocalityCarriers+"?id=invalid", "")

		server.GET(BaseEndpointLocalityCarriers, handler.Read())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when the locality id does not exist", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Read", mock.Anything, 1).Return([]domain.LocalityCarriersReport{}, carry.ErrNotFoundLocalityId)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointWithLocalityIdQueryCarriers, "")

		server.GET(BaseEndpointLocalityCarriers, handler.Read())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 200 and all localities with his carriers when id is not passed", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Read", mock.Anything, 0).Return(expectedLocalitiesCarriersReport, nil)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointLocalityCarriers, "")

		server.GET(BaseEndpointLocalityCarriers, handler.Read())
		server.ServeHTTP(response, request)

		responseResult := &domain.LocalityCarriersResponse{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedLocalitiesCarriersReport, responseResult.Data)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Read", mock.Anything, 0).Return([]domain.LocalityCarriersReport{}, nil)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointLocalityCarriers, "")

		server.GET(BaseEndpointLocalityCarriers, handler.Read())
		server.ServeHTTP(response, request)

		responseResult := &domain.LocalityCarriersResponse{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusNoContent, response.Code)
		assert.True(t, len(responseResult.Data) == 0)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Read", mock.Anything, 1).Return([]domain.LocalityCarriersReport{}, carry.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodGet, BaseEndpointWithLocalityIdQueryCarriers, "")

		server.GET(BaseEndpointLocalityCarriers, handler.Read())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCreateCarriers(t *testing.T) {
	t.Run("Should return status 200 and the carry created", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		mockService.On("Create", mock.Anything, mock.Anything).Return(expectedCarry, nil)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"Teste Livre","address":"Rua Pedro Dias","telephone":"3712291281","locality_id":1}`)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		responseResult := domain.CarryResponseId{}

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, expectedCarry, responseResult.Data)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"address":}`)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 400 when Cid is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"","company_name":"Teste Livre","address":"Rua Pedro Dias","telephone":"3712291281","locality_id":1}`)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when CompanyName is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"","address":"Rua Pedro Dias","telephone":"3712291281","locality_id":1}`)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Address is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"Teste Livre","address":"","telephone":"3712291281","locality_id":1}`)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when Telephone is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"Teste Livre","address":"Rua Pedro Dias","telephone":"","locality_id":1}`)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when LocalityID is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"Teste Livre","address":"Rua Pedro Dias","telephone":"3712291281","locality_id":0}`)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 409 when Carry already exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"Teste Livre","address":"Rua Pedro Dias","telephone":"3712291281","locality_id":1}`)

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.Carry")).Return(domain.Carry{}, carry.ErrAlredyExists)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return status 409 when LocalityID is not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"Teste Livre","address":"Rua Pedro Dias","telephone":"3712291281","locality_id":1}`)

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.Carry")).Return(domain.Carry{}, carry.ErrConflictLocalityId)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerWithCarriers(t)

		request, response := testutil.MakeRequest(http.MethodPost, BaseEndpointCarriers, `{"cid":"1111110","company_name":"Teste Livre","address":"Rua Pedro Dias","telephone":"3712291281","locality_id":1}`)

		mockService.On("Create", mock.Anything, mock.AnythingOfType("domain.Carry")).Return(domain.Carry{}, carry.ErrTryAgain)

		server.POST(BaseEndpointCarriers, handler.Create())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func InitServerWithCarriers(t *testing.T) (*gin.Engine, *mocks.CarryServiceMock, *handler.CarryController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.CarryServiceMock)
	handler := handler.NewCarry(mockService)
	return server, mockService, handler
}
