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

// @Summary Get All Sections
// @Produce json
// GET /sections @Summary Returns a list of Sections
// @Router /api/v1/sections [get]
// @Tags Section
// @Accept json
// @Success 200 {object} []domain.Section
// @Description List All Sections
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

// @Summary Get Section by ID
// @Produce json
// GET /section/:id @Summary Returns a section per Id
// @Router /api/v1/sections/{id} [get]
// @Param id path int true "Section ID"
// @Tags Section
// @Accept json
// @Success 200 {object} domain.Section
// @Description Describe by Section id
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

// @Summary Create Section
// @Produce json
// POST /section/:id @Summary Create a Section
// @Router /api/v1/sections [post]
// @Tags Section
// @Accept json
// @Param section body domain.Section true "Section Data"
// @Success 201 {object} domain.Section
// @Description Create Section
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

// @Summary Update Section
// @Produce json
// PATCH /sections/:id @Summary Update an existing Section
// @Router /api/v1/sections/{id} [patch]
// @Accept json
// @Tags Section
// @Success 200 {object} domain.Section
// @Param id path int true "Section ID"
// @Param section body domain.Section true "Section Data"
// @Description Update Section
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

// @Summary Delete Section
// @Produce json
// DELETE /sections/:id @Summary Delete a specific Section
// @Router /api/v1/sections/{id} [delete]
// @Param   id     path    int     true        "Section ID"
// @Tags Section
// @Accept json
// @Success 204
// @Description Delete Section
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
