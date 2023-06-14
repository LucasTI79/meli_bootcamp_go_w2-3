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

// @Produce json
// GET /sellers @Summary Returns a list of sellers
// @Router /api/v1/sellers [get]
// @Tags Sellers
// @Accept json
// @Success 200 {object}  []domain.Seller
// @Description List all Sellers
func (s *sellerController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, seller.ErrTryAgain.Error())
			return
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

// @Produce json
// GET /seller/:id @Summary Returns a seller per Id
// @Router /api/v1/sellers/{id} [get]
// @Param   id     path    int     true        "Seller ID"
// @Tags Sellers
// @Accept json
// @Success 200 {object}  domain.Seller
// @Description List one by Seller id
func (s *sellerController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, seller.ErrInvalidId.Error())
			return
		}
		sellerGet, err := s.sellerService.Get(c, sellerId)
		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, seller.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, sellerGet)
	}
}

// @Produce json
// POST /seller/:id @Summary Create a seller
// @Router /api/v1/sellers/ [post]
// @Tags Sellers
// @Accept json
// @Param seller body domain.Seller true "Seller Data"
// @Success 201 {object} domain.Seller
// @Description Create Sellers
func (s *sellerController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerInput := &domain.Seller{}
		err := c.ShouldBindJSON(sellerInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, seller.ErrTryAgain.Error(), err)
			return
		}

		if sellerInput.Address == "" || sellerInput.CID == 0 || sellerInput.CompanyName == "" || sellerInput.Telephone == "" {
			web.Error(c, http.StatusUnprocessableEntity, seller.ErrInvalidBody.Error())
			return
		}

		sellerItem := domain.Seller{
			CID:         sellerInput.CID,
			CompanyName: sellerInput.CompanyName,
			Address:     sellerInput.Address,
			Telephone:   sellerInput.Telephone,
		}
		sellerId, err := s.sellerService.Save(c, sellerItem)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}

		sellerItem.ID = sellerId
		web.Success(c, http.StatusCreated, sellerItem)
	}
}

// @Produce json
// PATCH /sellers/:id @Summary Modifies an existing seller
// @Router /api/v1/sellers/{id} [patch]
// @Accept json
// @Tags Sellers
// @Success 200 {object}  domain.Seller
// @Param id path int true "Seller ID"
// @Param seller body domain.Seller true "Seller Data"
// @Description Update Seller
func (s *sellerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, seller.ErrInvalidId.Error())
			return
		}

		sellerInput := &domain.Seller{}
		err := c.ShouldBindJSON(sellerInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, seller.ErrTryAgain.Error(), err)
			return
		}

		if sellerInput.Address == "" || sellerInput.CID == 0 || sellerInput.CompanyName == "" || sellerInput.Telephone == "" {
			web.Error(c, http.StatusUnprocessableEntity, seller.ErrInvalidBody.Error())
			return
		}

		sellerItem := domain.Seller{
			ID:          sellerId,
			CID:         sellerInput.CID,
			CompanyName: sellerInput.CompanyName,
			Address:     sellerInput.Address,
			Telephone:   sellerInput.Telephone,
		}

		err = s.sellerService.Update(c, sellerItem)
		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, seller.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(c, http.StatusOK, sellerItem)
	}
}

// @Produce json
// DELETE /sellers/:id @Summary Delete a specific seller
// @Router /api/v1/sellers/{id} [delete]
// @Param   id     path    int     true        "Seller ID"
// @Tags Sellers
// @Accept json
// @Success 204
// @Description Delete Seller
func (s *sellerController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			web.Error(c, http.StatusBadRequest, seller.ErrInvalidId.Error())
			return
		}

		err = s.sellerService.Delete(c, sellerId)

		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, seller.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Response(c, http.StatusNoContent, "")
	}
}
