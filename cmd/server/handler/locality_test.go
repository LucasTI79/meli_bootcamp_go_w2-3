package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/locality"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	BaseRouteLocality     = "/localities"
	RouteLocalityReportId = "/localities/report-sellers?id=1"
	RouteLocalityReport   = "/localities/report-sellers"
)

var localityDomain = domain.Locality{
	ID:           1,
	LocalityName: "Florianopolis",
	ProvinceName: "Santa Catarina",
}

var localityiInput = domain.LocalityInput{
	ID:           1,
	LocalityName: "Florianopolis",
	IdProvince:   1,
}

var listReportExpected = []domain.LocalityReport{
	{
		IdLocality:   1,
		LocalityName: "Florianopolis",
		SellersCount: 33,
	},
	{
		IdLocality:   2,
		LocalityName: "Blumenau",
		SellersCount: 45,
	},
}

func TestCreateLocality(t *testing.T) {
	t.Run("Should return status 201 with all locality", func(t *testing.T) {
		server, mockService, handler := InitServerLocality(t)
		server.POST(BaseRouteLocality, handler.Create())

		jsonLocality, _ := json.Marshal(localityDomain)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRouteLocality, string(jsonLocality))
		mockService.On("Save", mock.Anything, mock.Anything).Return(localityiInput, nil)
		server.ServeHTTP(response, request)

		responseResult := domain.LocalityResponseId{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, localityiInput, responseResult.Data)
	})
	t.Run("Should return status 422 when JSON is invalid", func(t *testing.T) {
		server, _, handler := InitServerLocality(t)
		server.POST(BaseRouteLocality, handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, BaseRouteLocality, string(`invalid-json`))
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
	t.Run("Should return status 400 when locality name is invalid", func(t *testing.T) {
		server, _, handler := InitServerLocality(t)
		server.POST(BaseRouteLocality, handler.Create())

		requestBody := domain.Locality{
			ProvinceName: "Santa Catarina",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRouteLocality, string(jsonSeller))

		server.ServeHTTP(response, request)

		var responseData web.ErrorResponse
		_ = json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, "locality name is required", responseData.Message)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 400 when province name is invalid", func(t *testing.T) {
		server, _, handler := InitServerLocality(t)
		server.POST(BaseRouteLocality, handler.Create())

		requestBody := domain.Locality{
			LocalityName: "Florianopolis",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRouteLocality, string(jsonSeller))

		server.ServeHTTP(response, request)

		var responseData web.ErrorResponse
		_ = json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, "province name is required", responseData.Message)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when province is not found", func(t *testing.T) {
		server, mockService, handler := InitServerLocality(t)
		server.POST(BaseRouteLocality, handler.Create())

		requestBody := domain.Locality{
			ID:           1,
			LocalityName: "Florianopolis",
			ProvinceName: "aa",
		}
		jsonSeller, _ := json.Marshal(requestBody)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRouteLocality, string(jsonSeller))
		mockService.On("Save", mock.Anything, mock.Anything).Return(domain.LocalityInput{}, locality.ErrProvinceNotFound)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerLocality(t)
		server.POST(BaseRouteLocality, handler.Create())

		jsonSeller, _ := json.Marshal(localityDomain)

		request, response := testutil.MakeRequest(http.MethodPost, BaseRouteLocality, string(jsonSeller))
		mockService.On("Save", mock.Anything, mock.Anything).Return(domain.LocalityInput{}, errors.New("error"))
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestReportSellersByLocality(t *testing.T) {
	t.Run("Should return status 200 when report created", func(t *testing.T) {

		server, mockService, handler := InitServerLocality(t)

		mockService.On("ReportSellersByLocality", mock.Anything, 1).Return(listReportExpected, nil)

		request, response := testutil.MakeRequest(http.MethodGet, RouteLocalityReportId, "")

		server.GET(RouteLocalityReport, handler.ReportSellersByLocality())
		server.ServeHTTP(response, request)

		responseResult := &domain.LocalitySellersResponse{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, listReportExpected, responseResult.Data)
	})
	t.Run("Should return status 400 when the locality id is invalid", func(t *testing.T) {
		server, mockService, handler := InitServerLocality(t)

		mockService.On("ReportSellersByLocality", mock.Anything, "invalid").Return([]domain.LocalityReport{}, errors.New("error"))

		request, response := testutil.MakeRequest(http.MethodGet, RouteLocalityReport+"?id=invalid", "")

		server.GET(RouteLocalityReport, handler.ReportSellersByLocality())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return status 404 when the locality id does not exist", func(t *testing.T) {
		server, mockService, handler := InitServerLocality(t)

		mockService.On("ReportSellersByLocality", mock.Anything, 1).Return([]domain.LocalityReport{}, locality.ErrNotFound)

		request, response := testutil.MakeRequest(http.MethodGet, RouteLocalityReportId, "")

		server.GET(RouteLocalityReport, handler.ReportSellersByLocality())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return status 500 when there is an internal error", func(t *testing.T) {
		server, mockService, handler := InitServerLocality(t)

		mockService.On("ReportSellersByLocality", mock.Anything, 1).Return([]domain.LocalityReport{}, errors.New("error"))

		request, response := testutil.MakeRequest(http.MethodGet, RouteLocalityReportId, "")

		server.GET(RouteLocalityReport, handler.ReportSellersByLocality())
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}

func InitServerLocality(t *testing.T) (*gin.Engine, *mocks.LocalityServiceMock, *handler.LocalityController) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.LocalityServiceMock)
	handler := handler.NewLocality(mockService)
	return server, mockService, handler
}
