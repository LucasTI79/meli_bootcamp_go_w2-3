package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	productbatch "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_batch"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatchController struct {
	productBatchService productbatch.Service
	productService      product.Service
	sectionService      section.Service
}

func NewProductBatch(s productbatch.Service, ps product.Service, ss section.Service) *ProductBatchController {
	return &ProductBatchController{
		productBatchService: s,
		productService:      ps,
		sectionService:      ss,
	}
}

// @Summary Save Product Batch
// @Produce json
// POST /productBatches @Summary Save Product Batch
// @Router /api/v1/productBatches [post]
// @Tags ProductBatch
// @Accept json
// @Param productBatch body domain.ProductBatch true "Product Batch"
// @Success 200 {object} domain.ProductBatch
// @Description Save Product Batch
func (s *ProductBatchController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productBatch domain.ProductBatch
		if err := c.ShouldBindJSON(&productBatch); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := productBatch.Validate(); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		err := s.productService.ExistsById(productBatch.ProductID)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}

		productBatchID, err := s.productBatchService.Save(c, productBatch)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, productBatchID)
	}
}
