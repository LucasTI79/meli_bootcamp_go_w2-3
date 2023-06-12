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
			web.Error(c, http.StatusBadRequest, "Invalid id")
			return
		}
		section, err := s.sectionService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, section)
	}
}

func (s *SectionController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionInput := &domain.SectionRequest{}
		err := c.ShouldBindJSON(sectionInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error creating session, try again %s", err)
			return
		}
		section, err := s.sectionService.Save(c, domain.Section{
			SectionNumber:      sectionInput.SectionNumber,
			CurrentTemperature: sectionInput.CurrentTemperature,
			MinimumTemperature: sectionInput.MinimumTemperature,
			CurrentCapacity:    sectionInput.CurrentCapacity,
			MinimumCapacity:    sectionInput.MinimumCapacity,
			MaximumCapacity:    sectionInput.MaximumCapacity,
			WarehouseID:        sectionInput.WarehouseID,
			ProductTypeID:      sectionInput.ProductTypeID,
		})
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, section)
	}
}

func (s *SectionController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *SectionController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid id")
			return
		}
		err = s.sectionService.Delete(c, id)
		if err != nil {
			if errors.Is(err, section.ErrNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Response(c, http.StatusNoContent, "")
	}
}
