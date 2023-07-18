package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type LocalityController struct {
	localityService locality.Service
}

func NewLocality(l locality.Service) *LocalityController {
	return &LocalityController{
		localityService: l,
	}
}

// Create creates a new locality.
// POST /localities @Summary Create a new locality
// @Description Create a new locality with the provided data
// @Tags Locality
// @Accept json
// @Produce json
// @Param locality body domain.Locality true "Locality object to be created"
// @Success 201 {object} domain.LocalityInput "Locality created successfully"
// @Router /locality [post]
func (l *LocalityController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := &domain.Locality{}
		err := c.ShouldBindJSON(domain)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		switch {
		case domain.LocalityName == "":
			web.Error(c, http.StatusBadRequest, "locality name is required")
			return
		case domain.ProvinceName == "":
			web.Error(c, http.StatusBadRequest, "province name is required")
			return
		}

		localitySaved, err := l.localityService.Save(c, *domain)
		if err != nil {
			if errors.Is(err, locality.ErrProvinceNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
		}
		web.Success(c, http.StatusCreated, localitySaved)
	}
}

// ReportSellersByLocality generates a report of sellers by locality.
// POST /localities/report-sellers @Summary Generate a report of sellers by locality
// @Description Generates a report of sellers based on the provided locality ID
// @Tags Locality
// @Accept json
// @Produce json
// @Param id query integer true "Locality ID"
// @Success 200 {array} domain.LocalityReport "Report of sellers by locality"
// @Router /locality/report [get]
func (l *LocalityController) ReportSellersByLocality() gin.HandlerFunc {
	return func(c *gin.Context) {
		localityIdStr := c.Query("id")

		var localityId int
		var err error

		if localityIdStr != "" {
			localityId, err = strconv.Atoi(localityIdStr)
			if err != nil {
				web.Error(c, http.StatusBadRequest, err.Error())
				return
			}
		}

		report, err := l.localityService.ReportSellersByLocality(c, localityId)
		if err != nil {
			if errors.Is(err, locality.ErrNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
		}
		if len(report) == 0 {
			web.Success(c, http.StatusNoContent, err)
		}
		web.Success(c, http.StatusOK, report)
	}
}
