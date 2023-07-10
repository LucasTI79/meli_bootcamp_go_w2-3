package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/inbound_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

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
// @Router /api/v1/InboundOrders/{id} [get]
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
// @Router /api/v1/InboundOrders [post]
// @Tags InboundOrders
// @Accept json
// @Param inbound order body domain.InboundOrders true "Inbound Order Data"
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

		switch {
		case inboundOrdersInput.OrderDate == "":
			web.Error(c, http.StatusBadRequest, "invalid order_date field")
			return
		case inboundOrdersInput.OrderNumber == "":
			web.Error(c, http.StatusBadRequest, "invalid order_number field")
			return
		case inboundOrdersInput.EmployeeID == "":
			web.Error(c, http.StatusBadRequest, "invalid employee_id field")
			return
		case inboundOrdersInput.ProductBatchID == "":
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

			web.Success(c, http.StatusCreated, inboundOrdersDomain)
		}
	}
}
