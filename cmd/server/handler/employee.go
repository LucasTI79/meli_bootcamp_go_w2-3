package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
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
	return func(c *gin.Context) {
		employeeId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, seller.ErrInvalidId.Error())
			return
		}

		employeeInput := &domain.Employee{}
		err := c.ShouldBindJSON(employeeInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, employee.ErrTryAgain.Error(), err)
			return
		}
		if employeeInput.CardNumberID == "" || employeeInput.FirstName == "" || employeeInput.LastName == "" || employeeInput.WarehouseID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}

		employeeItem := domain.Employee{
			ID:          employeeId,
			CardNumberID: employeeInput.CardNumberID,
			FirstName:  employeeInput.FirstName,
			LastName:   employeeInput.LastName,
			WarehouseID: employeeInput.WarehouseID,
		}
		err = e.employeeService.Update(c, employeeItem)
		if err != nil {
			if errors.Is(err, employee.ErrNotFound) {
				web.Error(c, http.StatusNotFound, employee.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, employeeItem)
	}
}

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
