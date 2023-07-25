package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/inbound_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

const layout = "2006-01-02"

type InboundOrdersController struct {
	InboundOrdersService inbound_order.Service
}

func NewInboundOrders(s inbound_order.Service) *InboundOrdersController {
	return &InboundOrdersController{
		InboundOrdersService: s,
	}
}

// @Summary Get Inbound Order by ID
// @Produce json
// GET /InboundOrders/:id @Summary Returns a inbound order per Id
// @Router /api/v1/inboundOrders/{id} [get]
// @Param id path int true "Inbound Order ID"
// @Tags InboundOrders
// @Accept json
// @Success 200 {object} domain.InboundOrders
// @Description List one by Inbound Order id
func (s *InboundOrdersController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		inboundOrderId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, inbound_order.ErrInvalidId.Error())
			return
		}

		inboundOrderGet, err := s.InboundOrdersService.Get(c, inboundOrderId)
		if err != nil {
			if errors.Is(err, inbound_order.ErrNotFound) {
				web.Error(c, http.StatusNotFound, inbound_order.ErrNotFound.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, inbound_order.ErrTryAgain.Error(), err)
			return
		}
		web.Success(c, http.StatusOK, inboundOrderGet)
	}
}

// @Summary Create Inbound Orders
// @Produce json
// POST /InboundOrders @Summary Create a Inbound Order
// @Router /api/v1/inboundOrders [post]
// @Tags InboundOrders
// @Accept json
// @Param inboundOrder body domain.InboundOrders true "Inbound Order Data"
// @Success 201 {object} domain.InboundOrders
// @Description Create Inbound Orders
func (s *InboundOrdersController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		inboundOrdersInput := &domain.InboundOrders{}
		err := c.ShouldBindJSON(inboundOrdersInput)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, inbound_order.ErrInvalidJSON.Error())
			return
		}
		_, err = time.Parse(layout, inboundOrdersInput.OrderDate)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "invalid date")
			return
		}
		switch {
		case inboundOrdersInput.OrderNumber == "":
			web.Error(c, http.StatusBadRequest, "invalid order_number field")
			return
		case inboundOrdersInput.EmployeeID == 0:
			web.Error(c, http.StatusBadRequest, "invalid employee_id field")
			return
		case inboundOrdersInput.ProductBatchID == 0:
			web.Error(c, http.StatusBadRequest, "invalid product_batch_id field")
			return
		case inboundOrdersInput.WarehouseID == 0:
			web.Error(c, http.StatusBadRequest, "invalid product_batch_id field")
			return
		}

		inboundOrdersDomain, err := s.InboundOrdersService.Create(c, *inboundOrdersInput)
		if err != nil {
			switch err {
			case inbound_order.ErrAlredyExists:
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())

		}
		web.Success(c, http.StatusCreated, inboundOrdersDomain)
	}
}

// @Summary Generate a report for all inbound orders
// @Description Generates a report containing information for all inbound orders
// @Tags InboundOrders
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.InboundOrdersReport
// @Failure 500 {string} ErrTryAgain
// @Router /api/v1/reportInboundOrders [get]
func (p *InboundOrdersController) ReportByAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		report, err := p.InboundOrdersService.ReportByAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, inbound_order.ErrTryAgain.Error())

			return
		}
		web.Success(c, http.StatusOK, report)
	}
}

// @Summary Generate a report for a specific employee's inbound orders
// @Description Generates a report containing information for inbound orders of a specific employee based on the provided employee ID
// @Tags InboundOrders
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the employee"
// @Success 200 {object} domain.InboundOrdersReport
// @Failure 400 {string} ErrInvalidId
// @Failure 404 {string} ErrNotFound
// @Failure 500 {string} ErrTryAgain
// @Router /api/v1/reportInboundOrders{id} [get]
func (p *InboundOrdersController) ReportByOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, employee.ErrInvalidId.Error())
			return
		}

		report, err := p.InboundOrdersService.ReportByOne(c, employeeId)
		if err != nil {
			if errors.Is(err, inbound_order.ErrNotFound) {
				web.Error(c, http.StatusNotFound, inbound_order.ErrNotFound.Error())

				return
			}

			web.Error(c, http.StatusInternalServerError, employee.ErrTryAgain.Error())

			return
		}
		web.Success(c, http.StatusOK, report)
	}
}
