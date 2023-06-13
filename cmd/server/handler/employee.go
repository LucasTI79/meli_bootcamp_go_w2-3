package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type Employee struct {
	employeeService employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees, err := e.employeeService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, employee.ErrTryAgain.Error())
			return
		}
		web.Success(c, http.StatusOK, employees)
	}
}

func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeInput := &domain.Employee{}
		err := c.ShouldBindJSON(employeeInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}
		if employeeInput.CardNumberID == "" || employeeInput.FirstName == "" || employeeInput.LastName == "" || employeeInput.WarehouseID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}
		employeeId, err := e.employeeService.Save(c, *employeeInput)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		employeeInput.ID = employeeId
		web.Success(c, http.StatusCreated, employeeInput)
	}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
