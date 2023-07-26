package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	productrecord "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductRecordController struct {
	productRecordService productrecord.Service
	productService       product.Service
}

func NewProductRecord(s productrecord.Service, ps product.Service) *ProductRecordController {
	return &ProductRecordController{
		productRecordService: s,
		productService:       ps,
	}
}

// @Summary Create ProductRecord
// @Produce json
// @Router /api/v1/productRecords [post]
// @Tags ProductRecord
// @Accept json
// @Param product body domain.ProductRecordRequest true "ProductRecord Data"
// @Success 201 {object} domain.ProductRecord
// @Description Create ProductRecord
func (p *ProductRecordController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		productRecordImput := &domain.ProductRecordRequest{}

		err := c.ShouldBindJSON(productRecordImput)

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, productrecord.ErrInvalidJson.Error()) //422
			return
		}

		const layout = "2006-01-02" // modelo da data
		currentDate := time.Now()   //obtem a data atual

		//transforma data de tipo string para tipo data
		imputDate, err := time.Parse(layout, productRecordImput.LastUpdateDate)
		if err != nil {
			web.Error(c, http.StatusBadRequest, productrecord.ErrInvalidField.Error())
			return
		}

		// se a lastUpdateDate for menor que a data do sistema, não poderá ser criado
		if imputDate.Before(currentDate) {
			// alterar pra conflict
			web.Error(c, http.StatusConflict, productrecord.ErrInvalidDate.Error())
			return
		}

		err = p.productService.ExistsById(productRecordImput.ProductID)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error()) //409
			return
		}

		productRecordItem := domain.ProductRecord{
			LastUpdateDate: productRecordImput.LastUpdateDate,
			PurchasePrice:  productRecordImput.PurchasePrice,
			SalePrice:      productRecordImput.SalePrice,
			ProductID:      productRecordImput.ProductID,
		}

		productRecordId, err := p.productRecordService.Save(c, productRecordItem)

		if err != nil {
			web.Error(c, http.StatusInternalServerError, productrecord.ErrTryAgain.Error())
			return
		}
		productRecordItem.ID = productRecordId
		web.Success(c, http.StatusCreated, productRecordItem) //201
	}
}

// @Summary Get All product record reports
// @Produce json
// @Router /api/v1/products/reportRecords [get]
// @Tags ProductRecord
// @Accept json
// @Success 200 {object}  []domain.ProductRecordReport
// @Description List product record reports of All Products
func (p *ProductRecordController) RecordsByAllProductsReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		productRecordsReport, err := p.productRecordService.RecordsByAllProductsReport(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, productrecord.ErrTryAgain.Error())

			return
		}
		web.Success(c, http.StatusOK, productRecordsReport)
	}
}

// @Summary Get Product Record Report by Product ID
// @Produce json
// @Router /api/v1/products/reportRecords/{id} [get]
// @Param   id     path    int     true        "Product ID"
// @Tags ProductRecord
// @Accept json
// @Success 200 {object}  domain.ProductRecordReport
// @Description List the product record report of one  product by it's Product id
func (p *ProductRecordController) RecordsByOneProductReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id")

		productRecord, err := p.productRecordService.RecordsByOneProductReport(c, id)
		if err != nil {
			if errors.Is(err, productrecord.ErrNotFound) {
				web.Error(c, http.StatusNotFound, productrecord.ErrNotFound.Error())

				return
			}

			web.Error(c, http.StatusInternalServerError, product.ErrTryAgain.Error())

			return
		}
		web.Success(c, http.StatusOK, productRecord)
	}
}
