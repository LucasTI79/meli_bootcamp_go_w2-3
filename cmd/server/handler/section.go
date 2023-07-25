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
		id := c.GetInt("id")
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
			ginError := &web.ApiError{}

			errors := ginError.CustomError(err)
			if len(errors) > 0 {
				web.Error(c, http.StatusUnprocessableEntity, errors[0].Message)
				return
			}
			web.Error(c, http.StatusUnprocessableEntity, domain.ErrTryAgain.Error(), err)
			return
		}

		if err := sectionInput.Validate(); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		sectionID, err := s.sectionService.Save(c, *sectionInput)
		if err != nil {
			if errors.Is(err, domain.ErrAlreadyExists) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		sectionInput.ID = sectionID
		web.Success(c, http.StatusCreated, sectionInput)
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
		id := c.GetInt("id")

		sectionInput := &domain.SectionRequest{}
		err := c.ShouldBindJSON(sectionInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, domain.ErrTryAgain.Error(), err)
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
		switch {
		case sectionInput.SectionNumber == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid section_number field")
			return
		case sectionInput.CurrentTemperature == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid current_temperature field")
			return
		case sectionInput.MinimumTemperature == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid minimum_temperature field")
			return
		case sectionInput.CurrentCapacity == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid current_capacity field")
			return
		case sectionInput.MinimumCapacity == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid minimum_capacity field")
			return
		case sectionInput.MaximumCapacity == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid maximum_capacity field")
			return
		case sectionInput.WarehouseID == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid warehouse_id field")
			return
		case sectionInput.ProductTypeID == 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid product_type_id field")
			return
		}

		err = s.sectionService.Update(c, sectionUpdated)
		if err != nil {
			if errors.Is(err, section.ErrNotFound) {
				web.Error(c, http.StatusNotFound, section.ErrNotFound.Error())
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
		id := c.GetInt("id")
		err := s.sectionService.Delete(c, id)
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

// @Summary Report Products by Section or All Sections
// @Produce json
// GET /sections/report @Summary Returns a list of Products by Section or All Sections
// @Router /api/v1/sections/reportProducts [get]
// @Tags Section
// @Accept json
// @Param id query int false "Section ID"
// @Success 200 {object} []domain.Product
// @Description Report Products by Section or All Sections
func (s *SectionController) ReportProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sections []domain.ProductBySection
		var section domain.ProductBySection
		param := c.Query("id")
		id, err := strconv.Atoi(param)
		if id > 0 {
			section, err = s.sectionService.ReportProductsById(c, id)
			sections = append(sections, section)
		}
		if param == "" {
			sections, err = s.sectionService.ReportProducts(c)
		}

		if id <= 0 && param != "" {
			web.Error(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
			return
		}

		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error listing sections")
			return
		}
		web.Success(c, http.StatusOK, sections)
	}
}
