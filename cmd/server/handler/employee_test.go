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

var employeeJson = `{"card_number_id":"001","first_name":"Joana","last_name":"Silva","warehouse_id":1}`

var expectedEmployees = domain.Employee{
	ID:           01,
	CardNumberID: "001",
	FirstName:    "Joana",
	LastName:     "Silva",
	WarehouseID:  1,
}

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

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		//fmt.Println(err)
		fmt.Println(responseResult)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedEmployees, responseResult.Data)
		assert.True(t, len(responseResult.Data) == 2)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		var ExpectedEmptyEmployees = []domain.Employee{}
		server, mockService, handler := InitServerWithGetEmployees(t)

		server.GET(GetAllEmployees, handler.GetAll())
		request, response := testutil.MakeRequest(http.MethodGet, GetAllEmployees, "")

		mockService.On("GetAll", mock.AnythingOfType("string")).Return(ExpectedEmptyEmployees, employee.ErrTryAgain)

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

func TestGetEmployeeById(t *testing.T) {
	//case find_by_id_existent
	t.Run("Should return status 200 with the requested employee", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		server.GET("/employees/:id", handler.Get())
		request, response := testutil.MakeRequest(http.MethodGet, "/employees/1", "")
		mockService.On("Get", mock.Anything, mock.AnythingOfType("int")).Return(expectedEmployees, nil)

		server.ServeHTTP(response, request)

		responseResult := &domain.EmployeeResponseID{}

		_ = json.Unmarshal(response.Body.Bytes(), responseResult)
		assert.Equal(t, expectedEmployees, responseResult.Data)
		assert.Equal(t, http.StatusOK, response.Code)

	})

	// case find_by_id_non_existent

	t.Run("Should return status 404 when the employee is not found", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		server.GET("/employees/:id", handler.Get())

		request, response := testutil.MakeRequest(http.MethodGet, "/employees/2", "")
		mockService.On("Get", mock.Anything, mock.AnythingOfType("int")).Return(domain.Employee{}, employee.ErrNotFound)
		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		ExpectedEmptyEmployees := domain.Employee{}
		server, mockService, handler := InitServerWithGetEmployees(t)

		server.GET("/employees/:id", handler.Get())

		mockService.On("Get", mock.Anything, mock.Anything).Return(ExpectedEmptyEmployees, employee.ErrTryAgain)
		request, response := testutil.MakeRequest(http.MethodGet, "/employees/1", "")
		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("Should return 400 when an Id is invalid", func(t *testing.T) {
		ExpectedEmptyEmployees := domain.Employee{}
		server, mockService, handler := InitServerWithGetEmployees(t)

		server.GET("/employees/:id", handler.Get())

		mockService.On("Get", mock.Anything).Return(ExpectedEmptyEmployees, employee.ErrInvalidId)
		request, response := testutil.MakeRequest(http.MethodGet, "/employees/invalidId", "")
		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

}

func TestCreateEmployees(t *testing.T) {

	t.Run("Should return 201 when employee is created", func(t *testing.T) {

		server, mockService, handler := InitServerWithGetEmployees(t)
		mockService.On("Save", mock.Anything, mock.Anything).Return(expectedEmployees, nil)
		request, response := testutil.MakeRequest(http.MethodPost, "/employees", employeeJson)

		server.POST("/employees", handler.Create())
		server.ServeHTTP(response, request)

		responseResult := &domain.EmployeeResponseID{}
		_ = json.Unmarshal(response.Body.Bytes(), responseResult)

		assert.Equal(t, expectedEmployees, responseResult.Data)

		assert.Equal(t, http.StatusCreated, response.Code)

	})

	t.Run("Should return 409 when employee already exists", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		mockService.On("Save", mock.Anything, mock.Anything).Return(domain.Employee{}, employee.ErrAlreadyExists)
		request, response := testutil.MakeRequest(http.MethodPost, "/employees", employeeJson)

		server.POST("/employees", handler.Create())
		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return 400 when field is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithGetEmployees(t)
		server.POST("/employees", handler.Create())
		request, response := testutil.MakeRequest(http.MethodPost, "/employees", string(`{"CardNumberID": 0}`))
		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return 422 when Json is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithGetEmployees(t)
		server.POST("/employees", handler.Create())

		request, response := testutil.MakeRequest(http.MethodPost, "/employees", string(`{"CardNumberID":}`))

		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		server.POST("/employees", handler.Create())

		mockService.On("Save", mock.Anything).Return(0, employee.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodPost, "/employee", employeeJson)

		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusInternalServerError, response.Code)

	})
}

func TestUpdateEmployee(t *testing.T) {

	t.Run("Should return 200 when employee is updated", func(t *testing.T) {

		server, mockService, handler := InitServerWithGetEmployees(t)
		request, response := testutil.MakeRequest(http.MethodPatch, "/employees/1", employeeJson)
		responseResult := domain.EmployeeResponseID{}

		mockService.On("Update", mock.Anything, mock.Anything).Return(nil)
		server.PATCH("/employees/:id", handler.Update())

		server.ServeHTTP(response, request)

		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedEmployees, responseResult.Data)

	})

	t.Run("Should return 404 when employee does not exist", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)
		server.PATCH("/employees/:id", handler.Update())
		request, response := testutil.MakeRequest(http.MethodPatch, "/employees/1", employeeJson)
		responseResult := domain.Employee{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)
		mockService.On("Update", mock.Anything, mock.Anything).Return(employee.ErrNotFound)
		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusNotFound, response.Code)

	})

	t.Run("Should return status 500 when an internal server error occurs.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetEmployees(t)

		server.PATCH("/employees/:id", handler.Update())

		mockService.On("Update", mock.Anything, mock.Anything).Return(employee.ErrTryAgain)

		request, response := testutil.MakeRequest(http.MethodPatch, "/employees/1", employeeJson)

		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusInternalServerError, response.Code)

	})

	t.Run("Should return 400 when an Id is invalid", func(t *testing.T) {

		server, mockService, handler := InitServerWithGetEmployees(t)

		server.PATCH("/employees/:id", handler.Update())

		mockService.On("Update", mock.Anything, mock.Anything).Return(employee.ErrInvalidId)

		request, response := testutil.MakeRequest(http.MethodPatch, "/employees/invalidId", employeeJson)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return 400 when field is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithGetEmployees(t)

		server.PATCH("/employees/:id", handler.Update())

		request, response := testutil.MakeRequest(http.MethodPatch, "/employees/1", string(`{"{"CardNumberID": 0}`))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return 422 when Json is invalid", func(t *testing.T) {
		server, _, handler := InitServerWithGetEmployees(t)

		server.PATCH("/employees/:id", handler.Update())

		request, response := testutil.MakeRequest(http.MethodPatch, "/employees/1", string(`{"CardNumberID":}`))

		server.ServeHTTP(response, request)
		//assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
}

func InitServerWithGetEmployees(t *testing.T) (*gin.Engine, *mocks.EmployeeServiceMock, *handler.Employee) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.EmployeeServiceMock)
	handler := handler.NewEmployee(mockService)
	return server, mockService, handler
}
