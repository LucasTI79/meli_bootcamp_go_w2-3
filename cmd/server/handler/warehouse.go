package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/pkg/web"
	"github.com/gin-gonic/gin"
)

type warehouseController struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *warehouseController {
	return &warehouseController{
		warehouseService: w,
	}
}

func (w *warehouseController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, "invalid id")
			return
		}

		warehouse, err := w.warehouseService.Get(c, warehouseId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				web.Error(c, http.StatusNotFound, "")
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, warehouse)
	}
}

func (w *warehouseController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses, err := w.warehouseService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error listing sellers")
			return
		}
		web.Success(c, http.StatusOK, warehouses)
	}
}

func (w *warehouseController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseInput := &domain.Warehouse{}
		err := c.ShouldBindJSON(warehouseInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error, try again %s", err)
			return
		}

		if warehouseInput.Address == "" || warehouseInput.MinimumCapacity == 0 || warehouseInput.MinimumTemperature == 0 || warehouseInput.Telephone == "" || warehouseInput.WarehouseCode == "" {
			web.Error(c, http.StatusUnprocessableEntity, "invalid body")
			return
		}

		warehouseId, err := w.warehouseService.Save(c, *warehouseInput)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}

		web.Success(c, http.StatusCreated, warehouseId)
	}
}

func (w *warehouseController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (w *warehouseController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
