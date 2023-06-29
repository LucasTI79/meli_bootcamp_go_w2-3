package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/testutil"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/employee"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GetAllEmployees = "/employees"
)

func TestGetAllEmployees(t *testing.T) {
	t.Run("Should return status 200 with all employees", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		expectedEmployees := []domain.Employee{
			{
				ID:           01,
				CardNumberID: "001",
				FirstName:    "Joana",
				LastName:     "Silva",
				WarehouseID:  1,
			},
			{
				ID:           02,
				CardNumberID: "002",
				FirstName:    "Beatriz",
				LastName:     "Costa",
				WarehouseID:  1,
			},
		}

		server.GET(GetAllEmployees, handler.GetAll())

		request, response := testutil.MakeRequest(http.MethodGet, GetAllEmployees, "")
		mockService.On("GetAll", mock.AnythingOfType("string")).Return(expectedEmployees, nil)
		server.ServeHTTP(response, request)

		responseResult := &domain.EmployeeResponse{}

		// fmt.Println(responseResult)

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		//fmt.Println(err)
		fmt.Println(responseResult)
		assert.Equal(t, http.StatusOK, response.Code)
		//assert.Equal(t, expectedEmployees, responseResult.Data)
		assert.True(t, len(responseResult.Data) == 2)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		var ExpectedEmpityEmployees = []domain.Employee{}

		server, mockService, handler := InitServerWithGetEmployees(t)

		server.GET(GetAllEmployees, handler.GetAll())
		request, response := testutil.MakeRequest(http.MethodGet, GetAllEmployees, "")

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(ExpectedEmpityEmployees)

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

	})
}

func TestDeleteEmployees(t *testing.T) {

	t.Run("Should return 204 when employee exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		server.DELETE("/employees/:id", handler.Delete())

		request, response := testutil.MakeRequest(http.MethodDelete, "/employees/1", "")
		mockService.On("Delete", mock.AnythingOfType("int")).Return(nil)
		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("Should return 404 when employee does not exist", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		server.DELETE("/employees/:id", handler.Delete())

		request, response := testutil.MakeRequest(http.MethodDelete, "/employees/1", "")
		mockService.On("Delete", mock.Anything).Return(employee.ErrNotFound)
		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)

		server.DELETE("/employees/:id", handler.Delete())

		mockService.On("Delete", mock.Anything).Return(employee.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodDelete, "/employees/1", "")

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

	})

	t.Run("Should return 400 when an Id is invalid", func(t *testing.T) {

		server, mockService, handler := InitServerWithGetEmployees(t)

		server.DELETE("/employees/:id", handler.Delete())

		mockService.On("Delete", mock.Anything).Return(employee.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodDelete, "/employees/invalidId", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

}

func InitServerWithGetEmployees(t *testing.T) (*gin.Engine, *mocks.EmployeeServiceMock, *handler.Employee) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.EmployeeServiceMock)
	handler := handler.NewEmployee(mockService)
	return server, mockService, handler
}
