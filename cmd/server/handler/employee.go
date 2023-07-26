package handler

import (
	"errors"
	"net/http"
	"strconv"

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

// @Summary Get Employee by ID
// @Produce json
// GET /employee/:id @Summary Returns a employee per Id
// @Router /api/v1/employees/{id} [get]
// @Param id path int true "Employee ID"
// @Tags Employees
// @Accept json
// @Success 200 {object} domain.Employee
// @Description List one by Employee id
func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, employee.ErrInvalidId.Error())
			return
		}
		employeeGet, err := e.employeeService.Get(c, employeeId)
		if err != nil {
			if errors.Is(err, employee.ErrNotFound) {
				web.Error(c, http.StatusNotFound, employee.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, employeeGet)
	}
}

// @Summary Get all Employees
// @Produce json
// GET /employees @Summary Returns a list of employees
// @Router /api/v1/employees [get]
// @Tags Employees
// @Accept json
// @Success 200 {object} []domain.Employee
// @Description List all Employees
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

// @Summary Create Employee
// @Produce json
// POST /employees/:id @Summary Create a employee
// @Router /api/v1/employees [post]
// @Tags Employees
// @Accept json
// @Param employee body domain.Employee true "Employee Data"
// @Success 201 {object} domain.Employee
// @Description Create Employee
func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeInput := &domain.Employee{}
		err := c.ShouldBindJSON(employeeInput)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error, try again %s", err)
			return
		}
		if employeeInput.CardNumberID == "" || employeeInput.FirstName == "" || employeeInput.LastName == "" || employeeInput.WarehouseID == 0 {
			web.Error(c, http.StatusBadRequest, "invalid body")
			return
		}
		employeeResult, err := e.employeeService.Save(c, *employeeInput)
		if err != nil {
			switch err {
			case employee.ErrAlreadyExists:
				web.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web.Error(c, http.StatusInternalServerError, employee.ErrTryAgain.Error(), err)
				return
			}
		}
		web.Success(c, http.StatusCreated, employeeResult)
	}
}

// @Summary Update Employee
// @Produce json
// PATCH /employees/:id @Summary Modifies an existing employee
// @Router /api/v1/employees/{id} [patch]
// @Accept json
// @Tags Employees
// @Success 200 {object} domain.Employee
// @Param id path int true "Employee ID"
// @Param employee body domain.Employee true "Employee Data"
// @Description Update Employee
func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, employee.ErrInvalidId.Error())
			return
		}
		domain := new(domain.Employee)
		if err := c.ShouldBindJSON(domain); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, employee.ErrInvalidBody.Error())
			return
		}
		result, err := e.employeeService.Update(c, *domain, employeeId)
		if err != nil {
			switch err {
			case employee.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
				return
			default:
				web.Error(c, http.StatusInternalServerError, employee.ErrTryAgain.Error(), err)
				return
			}
		}
		web.Success(c, http.StatusOK, result)
	}
}

// @Summary Delete Employee
// @Produce json
// DELETE /employees/:id @Summary Delete a specific employee
// @Router /api/v1/employees/{id} [delete]
// @Param  id path  int true  "Employee ID"
// @Tags Employees
// @Accept json
// @Success 204
// @Description Delete Employee
func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			web.Error(c, http.StatusBadRequest, employee.ErrInvalidId.Error())
			return
		}

		err = e.employeeService.Delete(c, employeeId)

		if err != nil {
			if errors.Is(err, employee.ErrNotFound) {
				web.Error(c, http.StatusNotFound, employee.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Response(c, http.StatusNoContent, "")
	}
}
