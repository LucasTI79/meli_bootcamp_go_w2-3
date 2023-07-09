package handler

import (
	"fmt"
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

func (p *ProductRecordController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		productRecordImput := &domain.ProductRecordRequest{}

		err := c.ShouldBindJSON(productRecordImput)
		fmt.Println(err)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, productrecord.ErrInvalidJson.Error()) //422
			return
		}
		fmt.Println("productRecordImput handler", productRecordImput)
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
		fmt.Println("productRecordId hendler", productRecordId)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, productrecord.ErrTryAgain.Error())
			return
		}
		productRecordItem.ID = productRecordId
		web.Success(c, http.StatusCreated, productRecordItem) //201
	}
}
