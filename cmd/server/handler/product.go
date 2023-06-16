package handler

import (
	"errors"
	"net/http"
	"strconv"

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

// @Summary Get All Products
// @Produce json
// @Router /api/v1/products [get]
// @Tags Products
// @Accept json
// @Success 200 {object}  []domain.Product
// @Description List all Products
func (p *productController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.productService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error listing products")
			return
		}
		web.Success(c, http.StatusOK, products)
	}
}

func (p *productController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, product.ErrInvalidId.Error())
			return
		}

		product, err := p.productService.Get(c, productId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				web.Error(c, http.StatusNotFound, "")
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, product)
	}
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
			web.Error(c, http.StatusUnprocessableEntity, product.ErrInvalidBody.Error())
			return
		}

		productItem := domain.Product{
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
		}

		productId, err := p.productService.Save(c, productItem)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		productItem.ID = productId
		web.Success(c, http.StatusCreated, productItem)
	}
}

func (p *productController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, errId := strconv.Atoi(c.Param("id"))

		if errId != nil {
			web.Response(c, http.StatusBadRequest, product.ErrInvalidId.Error())
			return
		}

		productImput := &domain.ProductRequest{}
		err := c.ShouldBindJSON(productImput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}

		if productImput.Description == "" || productImput.ExpirationRate == 0 || productImput.FreezingRate == 0 || productImput.Height == 0 || productImput.Length == 0 || productImput.Netweight == 0 || productImput.ProductCode == "" || productImput.RecomFreezTemp == 0 || productImput.SellerID == 0 {
			web.Error(c, http.StatusNotFound, product.ErrInvalidBody.Error())
			return
		}

		productItem := domain.Product{
			ID:             productId,
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
		}

		err = p.productService.Update(c, productItem)

		if err != nil {
			if errors.Is(err, product.ErrNotFound) {
				web.Error(c, http.StatusNotFound, product.ErrNotFound.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, productItem)

	}
}

func (p *productController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			web.Error(c, http.StatusBadRequest, product.ErrInvalidId.Error())
			return
		}
		err = p.productService.Delete(c, productId)

		if err != nil {
			productNotFound := errors.Is(err, product.ErrNotFound)

			if productNotFound {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Response(c, http.StatusNoContent, "")
	}
}
