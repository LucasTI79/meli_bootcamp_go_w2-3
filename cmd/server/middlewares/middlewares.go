package middlewares

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

func ValidateParams(params ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, param := range params {
			validatedParam, err := strconv.Atoi(c.Param(param))
			if err != nil {
				web.Error(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
				return
			}
			c.Set(param, validatedParam)
		}
	}
}
