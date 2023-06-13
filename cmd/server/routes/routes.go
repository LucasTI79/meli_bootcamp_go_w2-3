package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/warehouse"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := handler.NewSeller(service)
	r.rg.GET("/sellers", handler.GetAll())
	r.rg.GET("/sellers/:id", handler.Get())
	r.rg.POST("/sellers", handler.Create())
	r.rg.DELETE("/sellers/:id", handler.Delete())
	r.rg.PATCH("/sellers/:id", handler.Update())
}

func (r *router) buildProductRoutes() {}

func (r *router) buildSectionRoutes() {}

func (r *router) buildWarehouseRoutes() {
	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	handler := handler.NewWarehouse(service)
	r.rg.GET("/warehouses", handler.GetAll())
	r.rg.GET("/warehouses/:id", handler.Get())
	r.rg.POST("/warehouses", handler.Create())
	r.rg.DELETE("/warehouses/:id", handler.Delete())
	r.rg.PATCH("/warehouses/:id", handler.Update())
}

func (r *router) buildEmployeeRoutes() {}

func (r *router) buildBuyerRoutes() {}
