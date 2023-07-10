package handler

import (
	"errors"
	"net/http"
	"strconv"
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

func (p *ProductRecordController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		productRecordImput := &domain.ProductRecordRequest{}

		err := c.ShouldBindJSON(productRecordImput)

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, productrecord.ErrInvalidJson.Error()) //422
			return
		}

		const layout = "2006-01-02"
		_, err = time.Parse(layout, productRecordImput.LastUpdateDate)
		if err != nil {
			web.Error(c, http.StatusBadRequest, productrecord.ErrInvalidField.Error())
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

func (p *ProductRecordController) RecordsByOneProductReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, product.ErrInvalidId.Error())
			return
		}

		productRecord, err := p.productRecordService.RecordsByOneProductReport(c, productId)
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
