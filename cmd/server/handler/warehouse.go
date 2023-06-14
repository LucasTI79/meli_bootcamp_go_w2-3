package handler

import (
	"errors"
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

// @Summary Get Warehouse by ID
// @Produce json
// GET /warehouse/:id @Summary Returns a warehouse per Id
// @Router /api/v1/warehouses/{id} [get]
// @Param id path int true "Warehouse ID"
// @Tags Warehouses
// @Accept json
// @Success 200 {object} domain.Warehouse
// @Description List one by Warehouse id
func (w *warehouseController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Response(c, http.StatusBadRequest, warehouse.ErrInvalidId.Error())
			return
		}

		warehouseGet, err := w.warehouseService.Get(c, warehouseId)
		if err != nil {
			if errors.Is(err, warehouse.ErrNotFound) {
				web.Error(c, http.StatusNotFound, warehouse.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, warehouseGet)
	}
}

// @Summary Get all Warehouses
// @Produce json
// GET /warehouses @Summary Returns a list of warehouses
// @Router /api/v1/warehouses [get]
// @Tags Warehouses
// @Accept json
// @Success 200 {object} []domain.Warehouse
// @Description List all Warehouses
func (w *warehouseController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses, err := w.warehouseService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, warehouse.ErrTryAgain.Error(), err)
			return
		}
		web.Success(c, http.StatusOK, warehouses)
	}
}

// @Summary Create Warehouse
// @Produce json
// POST /warehouse/:id @Summary Create a warehouse
// @Router /api/v1/warehouses/ [post]
// @Tags Warehouses
// @Accept json
// @Param warehouse body domain.Warehouse true "Warehouse Data"
// @Success 201 {object} domain.Warehouse
// @Description Create Warehouses
func (w *warehouseController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseInput := &domain.Warehouse{}
		err := c.ShouldBindJSON(warehouseInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, warehouse.ErrTryAgain.Error(), err)
			return
		}

		if warehouseInput.Address == "" || warehouseInput.MinimumCapacity == 0 || warehouseInput.MinimumTemperature == 0 || warehouseInput.Telephone == "" || warehouseInput.WarehouseCode == "" {
			web.Error(c, http.StatusUnprocessableEntity, warehouse.ErrInvalidBody.Error())
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

// @Summary Delete Warehouse
// @Produce json
// DELETE /warehouses/:id @Summary Delete a specific warehouse
// @Router /api/v1/warehouses/{id} [delete]
// @Param  id path  int true  "Warehouse ID"
// @Tags Warehouses
// @Accept json
// @Success 204
// @Description Delete Warehouse
func (w *warehouseController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			web.Error(c, http.StatusBadRequest, warehouse.ErrInvalidId.Error())
			return
		}

		err = w.warehouseService.Delete(c, warehouseId)

		if err != nil {
			if errors.Is(err, warehouse.ErrNotFound) {
				web.Error(c, http.StatusNotFound, warehouse.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Response(c, http.StatusNoContent, "")
	}
}

// @Summary Update Warehouse
// @Produce json
// PATCH /warehouses/:id @Summary Modifies an existing warehouse
// @Router /api/v1/warehouses/{id} [patch]
// @Accept json
// @Tags Warehouses
// @Success 200 {object} domain.Warehouse
// @Param id path int true "Warehouse ID"
// @Param warehouse body domain.Warehouse true "Warehouse Data"
// @Description Update Warehouse
func (w *warehouseController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			web.Response(c, http.StatusBadRequest, warehouse.ErrInvalidId.Error())
			return
		}

		warehouseInput := &domain.Warehouse{}
		err := c.ShouldBindJSON(warehouseInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, warehouse.ErrTryAgain.Error(), err)
			return
		}

		if warehouseInput.Address == "" || warehouseInput.MinimumCapacity == 0 || warehouseInput.MinimumTemperature == 0 || warehouseInput.Telephone == "" || warehouseInput.WarehouseCode == "" {
			web.Error(c, http.StatusUnprocessableEntity, warehouse.ErrInvalidBody.Error())
			return
		}

		warehouseItem := domain.Warehouse{
			ID:                 warehouseId,
			Address:            warehouseInput.Address,
			Telephone:          warehouseInput.Telephone,
			WarehouseCode:      warehouseInput.WarehouseCode,
			MinimumCapacity:    warehouseInput.MinimumCapacity,
			MinimumTemperature: warehouseInput.MinimumTemperature,
		}

		err = w.warehouseService.Update(c, warehouseItem)
		if err != nil {
			if errors.Is(err, warehouse.ErrNotFound) {
				web.Error(c, http.StatusNotFound, warehouse.ErrNotFound.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(c, http.StatusOK, warehouseItem)
	}
}
