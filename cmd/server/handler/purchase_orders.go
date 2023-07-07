package handler

import (
	//"errors"

	"net/http"

	//"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/purchase_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type PurchaseOrdersController struct {
	purchaseordersService purchase_orders.Service
	buyerService          buyer.Service
}

func NewPurchaseOrders(o purchase_orders.Service, b buyer.Service) *PurchaseOrdersController {
	return &PurchaseOrdersController{
		purchaseordersService: o,
		buyerService:          b,
	}
}

// GetAll gets all purchase orders
// @Summary Get all purchase orders
// @Description Get all purchase orders
// @Tags Purchase Orders
// @Produce json
// @Success 200 {array} PurchaseOrder
// @Failure 500 {string} string "Error listing orders"
// @Failure 204 {string} string "No content"
// @Router api/v1/buyers/purchase_orders [get]
func (po *PurchaseOrdersController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := po.purchaseordersService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error listing orders")
			return
		}
		if len(orders) == 0 {
			web.Success(c, http.StatusNoContent, orders)
			return
		}
		web.Success(c, http.StatusOK, orders)
	}
}

// Create creates a new purchase order
// @Summary Create a new purchase order
// @Description Create a new purchase order
// @Tags Purchase Orders
// @Accept json
// @Produce json
// @Param order body PurchaseOrderRequest true "Purchase Order Request"
// @Success 201 {object} PurchaseOrder
// @Failure 400 {string} string "Bad request"
// @Failure 409 {string} string "Conflict"
// @Failure 422 {string} string "Unprocessable entity"
// @Router api/v1/buyers/purchase_orders [post]
func (po *PurchaseOrdersController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderRequest := &domain.PurchaseOrders{}
		err := c.ShouldBindJSON(orderRequest)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}
		err = po.buyerService.ExistsID(c, orderRequest.BuyerID)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		if orderRequest.OrderNumber == "" || orderRequest.OrderDate == "" || orderRequest.TrackingCode == "" || orderRequest.BuyerID == 0 || orderRequest.ProductRecordID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}
		order, err := po.purchaseordersService.Create(c, domain.PurchaseOrders{

			OrderNumber:     orderRequest.OrderNumber,
			OrderDate:       orderRequest.OrderDate,
			TrackingCode:    orderRequest.TrackingCode,
			BuyerID:         orderRequest.BuyerID,
			ProductRecordID: orderRequest.ProductRecordID,
			OrderStatusID:   orderRequest.OrderStatusID,
		})
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, order)
	}
}
