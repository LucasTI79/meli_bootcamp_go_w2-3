package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/carry"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type CarryController struct {
	carryService carry.Service
}

func NewCarry(s carry.Service) *CarryController {
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
			web.Error(c, http.StatusBadRequest, carry.ErrInvalidId.Error())
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

// @Summary Read Carriers of a Locality
// @Produce json
// GET /localities/reportCarriers @Summary Returns a list of localities with carriers count
// @Router /api/v1/localities/reportCarries [get]
// @Tags Carriers
// @Accept json
// @Success 200 {object} []domain.LocalityCarriersReport
// @Description List all Carriers of a Locality or All Localities
// @Param id query int false "Locality ID"
func (s *CarryController) Read() gin.HandlerFunc {
	return func(c *gin.Context) {
		localityIDStr := c.Query("id")

		var localityID int
		var err error

		if localityIDStr != "" {
			localityID, err = strconv.Atoi(localityIDStr)
			if err != nil {
				web.Error(c, http.StatusBadRequest, carry.ErrInvalidId.Error())
				return
			}
		}

		report, err := s.carryService.Read(c, localityID)
		if err != nil {
			switch err {
			case carry.ErrNotFoundLocalityId:
				web.Error(c, http.StatusNotFound, err.Error())
				return
			default:
				web.Error(c, http.StatusInternalServerError, carry.ErrTryAgain.Error(), err)
				return
			}
		}

		if len(report) == 0 {
			web.Success(c, http.StatusNoContent, "There are no carriers stored")
			return
		}

		web.Success(c, http.StatusOK, report)
	}
}

// @Summary Create Carry
// @Produce json
// POST /carriers @Summary Create a carry
// @Router /api/v1/carriers [post]
// @Tags Carriers
// @Accept json
// @Param carry body domain.Carry true "Carry Data"
// @Success 201 {object} domain.Carry
// @Description Create Carriers
func (s *CarryController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		carryInput := &domain.Carry{}
		err := c.ShouldBindJSON(carryInput)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, carry.ErrInvalidJSON.Error())
			return
		}

		switch {
		case carryInput.Cid == "":
			web.Error(c, http.StatusBadRequest, "invalid cid field")
			return
		case carryInput.CompanyName == "":
			web.Error(c, http.StatusBadRequest, "invalid company_name field")
			return
		case carryInput.Address == "":
			web.Error(c, http.StatusBadRequest, "invalid address field")
			return
		case carryInput.Telephone == "":
			web.Error(c, http.StatusBadRequest, "invalid telephone field")
			return
		case carryInput.LocalityId == 0:
			web.Error(c, http.StatusBadRequest, "invalid locality_id field")
			return
		}

		carryDomain, err := s.carryService.Create(c, *carryInput)
		if err != nil {
			switch err {
			case carry.ErrAlredyExists:
				web.Error(c, http.StatusConflict, err.Error())
				return
			case carry.ErrConflictLocalityId:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, carry.ErrTryAgain.Error(), err)
				return
			}
		}

		web.Success(c, http.StatusCreated, carryDomain)
	}
}
