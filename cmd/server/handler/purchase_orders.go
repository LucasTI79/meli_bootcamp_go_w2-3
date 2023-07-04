package handler

import (
	//"errors"
	//"net/http"
	//"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/purchase_orders"
	//"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	//"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
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

func (b *PurchaseOrdersController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (b *PurchaseOrdersController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (b *PurchaseOrdersController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}