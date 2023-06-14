package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type buyerController struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *buyerController {
	return &buyerController{
		buyerService: b,
	}
}

// @Produce json
// GET /buyers/:id @Summary Returns a buyers per Id
// @Router /api/v1/buyers/:id [get]
// @Param   id     path    int     true        "Buyer ID"
// @Accept json
// @Success 200 {object}  domain.Buyer
func (b *buyerController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, "invalid id")
			return
		}
		buyerObj, errGet := b.buyerService.Get(c, buyerId)
		if errGet != nil {
			buyerNotFound := errors.Is(errGet, buyer.ErrNotFound)
			if buyerNotFound {
				web.Error(c, http.StatusNotFound, "buyer not found")
				return
			}
			web.Error(c, http.StatusInternalServerError, "error listing buyer")
			return
		}
		web.Success(c, http.StatusOK, buyerObj)
	}
}

// @Produce json
// GET /buyers @Summary Returns a list of buyers
// @Router /api/v1/buyers [get]
// @Accept json
// @Success 200 {object}  []domain.Buyer
func (b *buyerController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyer, err := b.buyerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error listing buyers")
			return
		}
		web.Success(c, http.StatusOK, buyer)
	}
}

// @Produce json
// POST /buyers @Summary Create a buyer
// @Router /api/v1/buyers [post]
// @Accept json
// @Success 201 {int}  0
// @Param buyer body domain.BuyerRequest true "Buyer Data"
func (b *buyerController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerInput := &domain.BuyerRequest{}
		err := c.ShouldBindJSON(buyerInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}
		if buyerInput.CardNumberID == "" || buyerInput.FirstName == "" || buyerInput.LastName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}
		buyerId, err := b.buyerService.Save(c, domain.Buyer{

			CardNumberID: buyerInput.CardNumberID,
			FirstName:    buyerInput.FirstName,
			LastName:     buyerInput.LastName,
		})
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, buyerId)
	}
}

// @Produce json
// PATCH /buyers/:id @Summary Modifies an existing buyer
// @Router /api/v1/buyers/:id [patch]
// @Accept json
// @Success 200 {object}  domain.Buyer
// @Param buyer body domain.BuyerRequest true "Buyer Data"
func (b *buyerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, "invalid id")
			return
		}
		buyerInput := &domain.BuyerRequest{}
		errBind := c.ShouldBindJSON(buyerInput)
		if errBind != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", errBind)
			return
		}
		if buyerInput.CardNumberID == "" || buyerInput.FirstName == "" || buyerInput.LastName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}
		err := b.buyerService.Update(c, domain.Buyer{
			ID:           buyerId,
			CardNumberID: buyerInput.CardNumberID,
			FirstName:    buyerInput.FirstName,
			LastName:     buyerInput.LastName,
		})
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		buyerObj, errGet := b.buyerService.Get(c, buyerId)
		if errGet != nil {
			buyerNotFound := errors.Is(errGet, buyer.ErrNotFound)
			if buyerNotFound {
				web.Error(c, http.StatusNotFound, "buyer not found")
				return
			}
			web.Error(c, http.StatusInternalServerError, "error listing buyer")
			return
		}
		web.Success(c, http.StatusOK, buyerObj)
	}
}

// @Produce json
// DELETE /buyers/:id @Summary Delete a specific buyer
// @Router /api/v1/buyers/:id [delete]
// @Param   id     path    int     true        "Buyer ID"
// @Accept json
// @Success 200 {string}  " "
func (b *buyerController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, "invalid id")
			return
		}
		err := b.buyerService.Delete(c, buyerId)
		if err != nil {
			buyerNotFound := errors.Is(err, buyer.ErrNotFound)
			if buyerNotFound {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, "error listing buyer")
			return
		}
		web.Success(c, http.StatusNoContent, "")
	}
}
