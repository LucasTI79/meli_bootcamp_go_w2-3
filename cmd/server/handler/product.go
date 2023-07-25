package handler

import (
	"errors"
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService product.Service
}

func NewProduct(s product.Service) *ProductController {
	return &ProductController{
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
func (p *ProductController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.productService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, product.ErrTryAgain.Error())

			return
		}
		web.Success(c, http.StatusOK, products)
	}
}

// @Summary Get Product by ID
// @Produce json
// @Router /api/v1/products/{id} [get]
// @Param   id     path    int     true        "Product ID"
// @Tags Products
// @Accept json
// @Success 200 {object}  domain.Product
// @Description List one product by it's Product id
func (p *ProductController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id")

		newProduct, err := p.productService.Get(c, id)
		if err != nil {
			if errors.Is(err, product.ErrNotFound) {
				web.Error(c, http.StatusNotFound, product.ErrNotFound.Error())

				return
			}

			web.Error(c, http.StatusInternalServerError, product.ErrTryAgain.Error())
			return
		}
		web.Success(c, http.StatusOK, newProduct)
	}
}

// @Summary Create Product
// @Produce json
// @Router /api/v1/products [post]
// @Tags Products
// @Accept json
// @Param product body domain.Product true "Product Data"
// @Success 201 {object} domain.Product
// @Description Create Product
func (p *ProductController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		productImput := &domain.ProductRequest{}

		err := c.ShouldBindJSON(productImput)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, product.ErrInvalidJson.Error())
			return
		}

		if productImput.Description == "" || productImput.ExpirationRate == 0 || productImput.FreezingRate == 0 || productImput.Height == 0 || productImput.Length == 0 || productImput.Netweight == 0 || productImput.ProductCode == "" || productImput.RecomFreezTemp == 0 || productImput.SellerID == 0 {
			web.Error(c, http.StatusBadRequest, product.ErrInvalidField.Error())
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
			if errors.Is(err, product.ErrProductAlreadyExists) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, product.ErrTryAgain.Error())
			return
		}
		productItem.ID = productId
		web.Success(c, http.StatusCreated, productItem)
	}
}

// @Summary Update Product
// @Produce json
// @Router /api/v1/products/{id} [patch]
// @Accept json
// @Tags Products
// @Success 200 {object}  domain.Product
// @Param id path int true "Product ID"
// @Param product body domain.Product true "Product Data"
// @Description Update Product
func (p *ProductController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id")

		productImput := &domain.ProductRequest{}
		err := c.ShouldBindJSON(productImput)
		if err != nil {

			web.Error(c, http.StatusUnprocessableEntity, product.ErrInvalidJson.Error(), err)
			return
		}

		if productImput.Description == "" || productImput.ExpirationRate == 0 || productImput.FreezingRate == 0 || productImput.Height == 0 || productImput.Length == 0 || productImput.Netweight == 0 || productImput.ProductCode == "" || productImput.RecomFreezTemp == 0 || productImput.SellerID == 0 {

			web.Error(c, http.StatusBadRequest, product.ErrInvalidField.Error())
			return
		}

		productItem := domain.Product{
			ID:             id,
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

// @Summary Delete Product
// @Produce json
// @Router /api/v1/products/{id} [delete]
// @Param   id     path    int     true        "Product ID"
// @Tags Products
// @Accept json
// @Success 204
// @Description Delete Product
func (p *ProductController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id")

		err := p.productService.Delete(c, id)

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
