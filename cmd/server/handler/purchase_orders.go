package handler

import (
	//"errors"
	"net/http"
	//"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/purchase_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type PurchaseOrdersController struct {
	purchaseordersService purchase_orders.Service
}

func NewPurchaseOrders(o purchase_orders.Service) *PurchaseOrdersController {
	return &PurchaseOrdersController{
		purchaseordersService: o,
	}
}

func (po *PurchaseOrdersController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (po *PurchaseOrdersController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (po *PurchaseOrdersController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderRequest := &domain.PurchaseOrders{}
		err := c.ShouldBindJSON(orderRequest)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
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
