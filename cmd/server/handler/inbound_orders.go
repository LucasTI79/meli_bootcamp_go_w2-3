package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	inbound_order "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/inbound_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = errors.New("inbound orders not found")
	ErrConflict = errors.New("409 Conflict: Inboud Orders with ID already exists")
)

type InboundOrders struct {
	service inbound_order.Service
}

func NewInboundOrders(i inbound_order.Service) *InboundOrders {
	return &InboundOrders{
		service: i,
	}
}

// Method Get
// GetInboundOrders godoc
//
//	@Summary		Get InboundOrders
//	@Tags			InboundOrders
//	@Description	Get the details of a InboundOrders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of InboundOrders to be searched"
//	@Success		200	{object}	web.response
//	@Router			/api/v1/inboundOrders/{id} [get]
func (i *InboundOrders) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid InboundOrders ID: %s", err.Error())
			return
		}

		ctx := c.Request.Context()
		inboundOrders, err := i.service.Get(&ctx, id)

		if err != nil {
			if errors.Is(err, ErrNotFound) {
				web.Error(c, http.StatusNotFound, "InboundOrders not found: %s", err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, "Failed to get InboundOrders: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, *inboundOrders)
	}
}

// Method GetAll
// ListInboundOrders godoc
//
//	@Summary		List InboundOrders
//	@Tags			InboundOrders
//	@Description	getAll inboundOrders
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Router			/api/v1/inboundOrders [get]
func (i *InboundOrders) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		inboundOrders, err := i.service.GetAll(&ctx)

		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get inbound Orders: %s", err.Error())
			return
		}

		if len(*inboundOrders) == 0 {
			web.Error(c, http.StatusNoContent, "There are no inbound Orders stored: ")
			return
		}

		web.Success(c, http.StatusOK, *inboundOrders)
	}
}

// Method Save
// CreateInboundOrders godoc
//
//	@Summary		Create InboundOrders
//	@Tags			InboundOrders
//	@Description	Create inboundOrders
//	@Accept			json
//	@Produce		json
//	@Param			InboundOrders	body		domain.RequestCreateInboundOrders	true	"InboundOrders to Create"
//	@Success		200			{object}	web.response
//	@Router			/api/v1/inboundOrders [post]
func (i *InboundOrders) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		// createEmployee := domain.RequestCreateEmployee{}
		createInboundOrders := new(domain.RequestCreateInboundOrders)
		if err := c.ShouldBindJSON(&createInboundOrders); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}

		inboundOrdersDomain := &domain.InboundOrders{
			OrderDate:      createInboundOrders.OrderDate,
			OrderNumber:    createInboundOrders.OrderNumber,
			EmployeeID:     createInboundOrders.EmployeeID,
			ProductBatchID: createInboundOrders.ProductBatchID,
			WarehouseID:    createInboundOrders.WarehouseID,
		}

		if inboundOrdersDomain.OrderDate == "" {
			web.Error(c, http.StatusBadRequest, "Field Order Date is required: %s", "")
			return
		}

		if inboundOrdersDomain.OrderNumber == "" {
			web.Error(c, http.StatusBadRequest, "Field Order Number is required: %s", "")
		}

		if inboundOrdersDomain.EmployeeID == "" {
			web.Error(c, http.StatusBadRequest, "Field Employee ID is required: %s", "")
		}

		if inboundOrdersDomain.ProductBatchID == "" {
			web.Error(c, http.StatusBadRequest, "Field Product Batch ID is required: %s", "")
		}

		if inboundOrdersDomain.WarehouseID == "" {
			web.Error(c, http.StatusBadRequest, "Field Warehouse ID is required: %s", "")
		}

		ctx := c.Request.Context()
		inboundOrdersDomain, err := i.service.Save(&ctx, *inboundOrdersDomain)
		if err != nil {
			switch err {
			case inbound_order.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web.Error(c, http.StatusBadRequest, "Error to save request: %s", err.Error())
				return
			}

		}

		web.Success(c, http.StatusCreated, *inboundOrdersDomain)
	}
}

// Method Update
// UpdateInboundOrders godoc
//
//	@Summary		Update InboundOrders
//	@Tags			InboundOrders
//	@Description	Update the details of a InboundOrders
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string							true	"ID of InboundOrders to be updated"
//	@Param			InboundOrders	body		domain.RequestUpdateInboundOrders	true	"Updated InboundOrders details"
//	@Success		200			{object}	web.response
//	@Router			/api/v1/inboundOrders/{id} [patch]
func (i *InboundOrders) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		ReqUpdateInboundOrders := new(domain.RequestUpdateInboundOrders)

		if err := c.ShouldBindJSON(&ReqUpdateInboundOrders); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "Error to read request: %s", err.Error())
			return
		}

		ctx := c.Request.Context()
		inboundOrdersUpdate, err := i.service.Update(&ctx, id, ReqUpdateInboundOrders)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to update: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, inboundOrdersUpdate)
	}
}

// Method Delete
// DeleteInboundOrders godoc
//
//	@Summary		Delete InboundOrders
//	@Tags			InboundOrders
//	@Description	Delete InboundOrders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of a InboundOrders to be excluded"
//	@Success		204	{object}	web.response
//	@Router			/api/v1/inboundOrders/{id} [delete]
func (i *InboundOrders) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		ctx := c.Request.Context()
		err = i.service.Delete(&ctx, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to delete: %s", err.Error())
			return
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
