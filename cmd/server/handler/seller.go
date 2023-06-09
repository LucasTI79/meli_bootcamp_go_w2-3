package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type sellerController struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *sellerController {
	return &sellerController{
		sellerService: s,
	}
}

func (s *sellerController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error listing sellers")
			return
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

func (s *sellerController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (s *sellerController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerInput := &domain.SellerRequest{}
		err := c.ShouldBindJSON(sellerInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}
		sellerId, err := s.sellerService.Save(c, domain.Seller{

			CID:         sellerInput.CID,
			CompanyName: sellerInput.CompanyName,
			Address:     sellerInput.Address,
			Telephone:   sellerInput.Address,
		})
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusOK, sellerId)
	}
}

func (s *sellerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (s *sellerController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
