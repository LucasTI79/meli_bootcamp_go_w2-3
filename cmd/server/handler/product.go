package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productService product.Service
}

func NewProduct(s product.Service) *productController {
	return &productController{
		productService: s,
	}
}

func (p *productController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *productController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *productController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		productImput := &domain.ProductRequest{}

		err := c.ShouldBindJSON(productImput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}

		if productImput.Description == "" || productImput.ExpirationRate == 0 || productImput.FreezingRate == 0 || productImput.Height == 0 || productImput.Length == 0 || productImput.Netweight == 0 || productImput.ProductCode == "" || productImput.RecomFreezTemp == 0 || productImput.SellerID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}

		productId, err := p.productService.Save(c, domain.Product{
			Description:    productImput.Description,
			ExpirationRate: productImput.ExpirationRate,
			FreezingRate:   productImput.FreezingRate,
			Height:         productImput.Height,
			Length:         productImput.Length,
			Netweight:      productImput.Netweight,
			ProductCode:    productImput.ProductCode,
			RecomFreezTemp: productImput.RecomFreezTemp,
			Width:          productImput.Width,
			ProductTypeID:  productImput.ProductTypeID,
			SellerID:       productImput.SellerID,
		})
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusOK, productId)
	}
}

func (p *productController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *productController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
