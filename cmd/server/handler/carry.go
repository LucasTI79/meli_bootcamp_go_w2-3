package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/carry"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type CarryController struct {
	carryService carry.CarryService
}

func NewCarry(s carry.CarryService) *CarryController {
	return &CarryController{
		carryService: s,
	}
}

// @Summary Get Carry by ID
// @Produce json
// GET /carriers/:id @Summary Returns a carry per Id
// @Router /api/v1/carriers/{id} [get]
// @Param id path int true "Carry ID"
// @Tags Carriers
// @Accept json
// @Success 200 {object} domain.Carry
// @Description List one by Carry id
func (s *CarryController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		carryId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, carry.ErrInvalidId.Error())
			return
		}

		carryGet, err := s.carryService.Get(c, carryId)
		if err != nil {
			if errors.Is(err, carry.ErrNotFound) {
				web.Error(c, http.StatusNotFound, carry.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, carry.ErrTryAgain.Error(), err)
			return
		}
		web.Success(c, http.StatusOK, carryGet)
	}
}
