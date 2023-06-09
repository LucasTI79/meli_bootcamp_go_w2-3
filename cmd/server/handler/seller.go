package handler

import (
	"net/http"

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
