package handler

import (
	"errors"
	"net/http"
	"strconv"

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
		sellerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, "invalid id")
			return
		}
		seller, err := s.sellerService.Get(c, sellerId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				web.Error(c, http.StatusNotFound, "")
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, seller)
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
		if sellerInput.Address == "" || sellerInput.CID == 0 || sellerInput.CompanyName == "" || sellerInput.Telephone == "" {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}
		sellerId, err := s.sellerService.Save(c, domain.Seller{
			CID:         sellerInput.CID,
			CompanyName: sellerInput.CompanyName,
			Address:     sellerInput.Address,
			Telephone:   sellerInput.Telephone,
		})
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, sellerId)
	}
}

func (s *sellerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, "Invalid id")
			return
		}
		sellerInput := &domain.SellerRequest{}
		err := c.ShouldBindJSON(sellerInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}
		errUpdate := s.sellerService.Update(c, domain.Seller{
			ID:          sellerId,
			CID:         sellerInput.CID,
			CompanyName: sellerInput.CompanyName,
			Address:     sellerInput.Address,
			Telephone:   sellerInput.Telephone,
		})
		_, errGet := s.sellerService.Get(c, sellerId)
		if errGet != nil {
			web.Error(c, http.StatusNotFound, "")
			return
		}
		if errUpdate != nil {
			web.Error(c, http.StatusInternalServerError, errUpdate.Error())
			return
		}
		web.Success(c, http.StatusNoContent, "")
	}
}

func (s *sellerController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid id")
			return
		}

		err = s.sellerService.Delete(c, sellerId)

		if err != nil {
			sellerNotFound := errors.Is(err, seller.ErrNotFound)
			if sellerNotFound {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Response(c, http.StatusNoContent, "")
	}
}
