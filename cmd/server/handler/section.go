package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type SectionController struct {
	sectionService section.Service
}

func NewSection(s section.Service) *SectionController {
	return &SectionController{
		sectionService: s,
	}
}

// @Summary List all sections
// @Description List all sections availables
// @Tags ListAllSections
// @Produce json
// @Success 200 array []domain.Section
// @Failure 500 {object} web.Error()
// @Router /api/v1/sections [get]
func (s *SectionController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sections, err := s.sectionService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error listing sections")
			return
		}
		web.Success(c, http.StatusOK, sections)
	}
}

func (s *SectionController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, domain.ErrAlreadyExists.Error())
			return
		}
		section, err := s.sectionService.Get(c, id)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				web.Error(c, http.StatusNotFound, domain.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, section)
	}
}

func (s *SectionController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionInput := &domain.Section{}
		err := c.ShouldBindJSON(sectionInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, domain.ErrTryAgain.Error(), err)
			return
		}
		if sectionInput.SectionNumber == 0 || sectionInput.CurrentTemperature == 0 || sectionInput.MinimumTemperature == 0 || sectionInput.CurrentCapacity == 0 || sectionInput.MinimumCapacity == 0 || sectionInput.MaximumCapacity == 0 || sectionInput.WarehouseID == 0 || sectionInput.ProductTypeID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}
		sectionID, err := s.sectionService.Save(c, *sectionInput)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, sectionID)
	}
}

func (s *SectionController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
			return
		}

		sectionInput := &domain.SectionRequest{}
		err = c.ShouldBindJSON(sectionInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, domain.ErrTryAgain.Error(), err)
			return
		}
		if sectionInput.SectionNumber == 0 || sectionInput.CurrentTemperature == 0 || sectionInput.MinimumTemperature == 0 || sectionInput.CurrentCapacity == 0 || sectionInput.MinimumCapacity == 0 || sectionInput.MaximumCapacity == 0 || sectionInput.WarehouseID == 0 || sectionInput.ProductTypeID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}
		sectionUpdated := domain.Section{
			ID:                 id,
			SectionNumber:      sectionInput.SectionNumber,
			CurrentTemperature: sectionInput.CurrentTemperature,
			MinimumTemperature: sectionInput.MinimumTemperature,
			CurrentCapacity:    sectionInput.CurrentCapacity,
			MinimumCapacity:    sectionInput.MinimumCapacity,
			MaximumCapacity:    sectionInput.MaximumCapacity,
			WarehouseID:        sectionInput.WarehouseID,
			ProductTypeID:      sectionInput.ProductTypeID,
		}
		err = s.sectionService.Update(c, sectionUpdated)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				web.Error(c, http.StatusNotFound, domain.ErrNotFound.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(c, http.StatusOK, sectionUpdated)

	}
}

func (s *SectionController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
			return
		}
		err = s.sectionService.Delete(c, id)
		if err != nil {
			if errors.Is(err, section.ErrNotFound) {
				web.Error(c, http.StatusNotFound, domain.ErrNotFound.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Response(c, http.StatusNoContent, "")
	}
}
