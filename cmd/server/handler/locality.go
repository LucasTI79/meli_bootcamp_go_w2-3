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

// TODO implementar validações
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

func (l *LocalityController) ReportSellersByLocalities() gin.HandlerFunc {
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
