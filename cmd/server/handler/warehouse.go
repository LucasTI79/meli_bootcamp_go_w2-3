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

type WarehouseController struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *WarehouseController {
	return &WarehouseController{
		warehouseService: w,
	}
}

// @Summary Get Warehouse by ID
// @Produce json
// GET /warehouses/:id @Summary Returns a warehouse per Id
// @Router /api/v1/warehouses/{id} [get]
// @Param id path int true "Warehouse ID"
// @Tags Warehouses
// @Accept json
// @Success 200 {object} domain.Warehouse
// @Description List one by Warehouse id
func (w *WarehouseController) Get() gin.HandlerFunc {
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

			web.Error(c, http.StatusInternalServerError, warehouse.ErrTryAgain.Error(), err)
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
func (w *WarehouseController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses, err := w.warehouseService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, warehouse.ErrTryAgain.Error(), err)
			return
		}

		if len(warehouses) == 0 {
			web.Error(c, http.StatusNoContent, "There are no warehouses stored")
			return
		}

		web.Success(c, http.StatusOK, warehouses)
	}
}

// @Summary Create Warehouse
// @Produce json
// POST /warehouses/:id @Summary Create a warehouse
// @Router /api/v1/warehouses [post]
// @Tags Warehouses
// @Accept json
// @Param warehouse body domain.Warehouse true "Warehouse Data"
// @Success 201 {object} domain.Warehouse
// @Description Create Warehouses
func (w *WarehouseController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseInput := &domain.Warehouse{}
		err := c.ShouldBindJSON(warehouseInput)
		if err != nil {
			web.Error(c, http.StatusBadRequest, warehouse.ErrTryAgain.Error(), err)
			return
		}

		switch {
		case warehouseInput.Address == "":
			web.Error(c, http.StatusUnprocessableEntity, "invalid address field")
			return
		case warehouseInput.MinimumCapacity < 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid minimum_capacity field")
			return
		case warehouseInput.MinimumTemperature < 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid minimum_temperature field")
			return
		case warehouseInput.Telephone == "":
			web.Error(c, http.StatusUnprocessableEntity, "invalid telephone field")
			return
		case warehouseInput.WarehouseCode == "":
			web.Error(c, http.StatusUnprocessableEntity, "invalid warehouse_code field")
			return
		}

		warehouseId, err := w.warehouseService.Save(c, *warehouseInput)
		if err != nil {
			if errors.Is(err, domain.ErrAlreadyExists) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		warehouseInput.ID = warehouseId
		web.Success(c, http.StatusCreated, warehouseInput)
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
func (w *WarehouseController) Delete() gin.HandlerFunc {
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
func (w *WarehouseController) Update() gin.HandlerFunc {
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

		switch {
		case warehouseInput.Address == "":
			web.Error(c, http.StatusUnprocessableEntity, "invalid address field")
			return
		case warehouseInput.MinimumCapacity < 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid minimum_capacity field")
			return
		case warehouseInput.MinimumTemperature < 0:
			web.Error(c, http.StatusUnprocessableEntity, "invalid minimum_temperature field")
			return
		case warehouseInput.Telephone == "":
			web.Error(c, http.StatusUnprocessableEntity, "invalid telephone field")
			return
		case warehouseInput.WarehouseCode == "":
			web.Error(c, http.StatusUnprocessableEntity, "invalid warehouse_code field")
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
