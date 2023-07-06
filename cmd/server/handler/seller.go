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

type SellerController struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *SellerController {
	return &SellerController{
		sellerService: s,
	}
}

// @Summary Get All Sellers
// @Produce json
// GET /sellers @Summary Returns a list of sellers
// @Router /api/v1/sellers [get]
// @Tags Sellers
// @Accept json
// @Success 200 {object}  []domain.Seller
// @Description List all Sellers
func (s *SellerController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		if len(sellers) == 0 {
			web.Success(c, http.StatusNoContent, sellers)
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

// @Summary Get Seller by ID
// @Produce json
// GET /sellers/:id @Summary Returns a seller per Id
// @Router /api/v1/sellers/{id} [get]
// @Param   id     path    int     true        "Seller ID"
// @Tags Sellers
// @Accept json
// @Success 200 {object}  domain.Seller
// @Description List one by Seller id
func (s *SellerController) Get() gin.HandlerFunc {
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

// @Summary Create Seller
// @Produce json
// POST /sellers/:id @Summary Create a seller
// @Router /api/v1/sellers [post]
// @Tags Sellers
// @Accept json
// @Param seller body domain.Seller true "Seller Data"
// @Success 201 {object} domain.Seller
// @Description Create Sellers
func (s *SellerController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerInput := &domain.Seller{}
		err := c.ShouldBindJSON(sellerInput)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, seller.ErrTryAgain.Error())
			return
		}

		switch {
		case sellerInput.CID == 0:
			web.Error(c, http.StatusBadRequest, "cid is required")
			return
		case sellerInput.CompanyName == "":
			web.Error(c, http.StatusBadRequest, "company name is required")
			return
		case sellerInput.Address == "":
			web.Error(c, http.StatusBadRequest, "address is required")
			return
		case sellerInput.Telephone == "":
			web.Error(c, http.StatusBadRequest, "phone is required")
			return
		}

		sellerSaved, err := s.sellerService.Save(c, *sellerInput)
		switch err{
		case seller.ErrLocality:
			web.Error(c, http.StatusConflict, err.Error())
			return
		case seller.ErrCidAlreadyExists:
			web.Error(c, http.StatusConflict, err.Error())
			return
		default:
			web.Error(c, http.StatusInternalServerError, err.Error())

		}
		web.Success(c, http.StatusCreated, sellerSaved)
	}
}

// @Summary Update Seller
// @Produce json
// PATCH /sellers/:id @Summary Modifies an existing seller
// @Router /api/v1/sellers/{id} [patch]
// @Accept json
// @Tags Sellers
// @Success 200 {object}  domain.Seller
// @Param id path int true "Seller ID"
// @Param seller body domain.Seller true "Seller Data"
// @Description Update Seller
func (s *SellerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		domain := new(domain.Seller)
		if err := c.ShouldBindJSON(domain); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, seller.ErrInvalidBody.Error())
			return
		}
		sellerUpdated, err := s.sellerService.Update(c, id, *domain)
		if err != nil {
			switch err {
			case seller.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
				return
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
		web.Success(c, http.StatusOK, sellerUpdated)
	}
}

// @Summary Delete Seller
// @Produce json
// DELETE /sellers/:id @Summary Delete a specific seller
// @Router /api/v1/sellers/{id} [delete]
// @Param   id     path    int     true        "Seller ID"
// @Tags Sellers
// @Accept json
// @Success 204
// @Description Delete Seller
func (s *SellerController) Delete() gin.HandlerFunc {
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
