package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"errors"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
)

type buyerController struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *buyerController {
	return &buyerController{
		buyerService: b,
	}
}

func (b *buyerController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, "invalid id")
			return
		}
		buyerObj, errGet := b.buyerService.Get(c, buyerId)
		fmt.Println("teste3")
		if errGet != nil {
			buyerNotFound := errors.Is(errGet, buyer.ErrNotFound)
			if buyerNotFound {
				web.Error(c, http.StatusNotFound, "buyer not found")
				return
			}
			web.Error(c, http.StatusInternalServerError, "error listing buyer")
			return
		}
		web.Success(c, http.StatusOK, buyerObj)}
}

func (b *buyerController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyer, err := b.buyerService.GetAll(c)
		if err != nil { 
			web.Error(c, http.StatusInternalServerError, "error listing buyers")
			return
		}
		web.Success(c, http.StatusOK, buyer)}
}

func (b *buyerController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *buyerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *buyerController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
