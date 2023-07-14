package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/carry"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	productbatch "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_batch"
	productrecord "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/warehouse"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.buildCarryRoutes()
	r.buildSwagger()
	r.buildProductBatchRoutes()
	r.buildLocalityRoutes()

	r.buildProductRecordRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repoSellers := seller.NewRepository(r.db)
	serviceSellers := seller.NewService(repoSellers)

	repoLocalities := locality.NewRepository(r.db)
	serviceLocalities := locality.NewService(repoLocalities)

	handler := handler.NewSeller(serviceSellers, serviceLocalities)
	r.rg.GET("/sellers", handler.GetAll())
	r.rg.GET("/sellers/:id", handler.Get())
	r.rg.POST("/sellers", handler.Create())
	r.rg.DELETE("/sellers/:id", handler.Delete())
	r.rg.PATCH("/sellers/:id", handler.Update())
}

func (r *router) buildLocalityRoutes() {
	repo := locality.NewRepository(r.db)
	service := locality.NewService(repo)
	handler := handler.NewLocality(service)
	r.rg.POST("/localities", handler.Create())
	r.rg.GET("/localities/report-sellers", handler.ReportSellersByLocality())
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	r.rg.POST("/products", handler.Create())
	r.rg.GET("/products", handler.GetAll())
	r.rg.DELETE("/products/:id", handler.Delete())
	r.rg.GET("/products/:id", handler.Get())
	r.rg.PATCH("/products/:id", handler.Update())
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	handler := handler.NewSection(service)

	r.rg.GET("/sections", handler.GetAll())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.POST("/sections", handler.Create())
	r.rg.DELETE("/sections/:id", handler.Delete())
	r.rg.PATCH("/sections/:id", handler.Update())
	r.rg.GET("/sections/reportProducts", handler.ReportProducts())
}

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

func (r *router) buildEmployeeRoutes() {
	repo := employee.NewRepository(r.db)
	service := employee.NewService(repo)
	handler := handler.NewEmployee(service)

	r.rg.POST("/employees", handler.Create())
	r.rg.GET("/employees", handler.GetAll())
	r.rg.GET("/employees/:id", handler.Get())
	r.rg.DELETE("/employees/:id", handler.Delete())
	r.rg.PATCH("/employees/:id", handler.Update())
}

func (r *router) buildBuyerRoutes() {
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)
	r.rg.GET("/buyers", handler.GetAll())
	r.rg.GET("/buyers/:id", handler.Get())
	r.rg.POST("/buyers", handler.Create())
	r.rg.PATCH("/buyers/:id", handler.Update())
	r.rg.DELETE("/buyers/:id", handler.Delete())
}

func (r *router) buildProductBatchRoutes() {
	productRrepo := product.NewRepository(r.db)
	productService := product.NewService(productRrepo)

	sectionRepo := section.NewRepository(r.db)
	sectionService := section.NewService(sectionRepo)

	repo := productbatch.NewRepository(r.db, productbatch.Querys{})
	service := productbatch.NewService(repo)
	handler := handler.NewProductBatch(service, productService, sectionService)

	r.rg.POST("/productBatches", handler.Create())
}
func (r *router) buildCarryRoutes() {
	repoCarry := carry.NewRepository(r.db)
	repoLocalities := locality.NewRepository(r.db)
	service := carry.NewService(repoCarry, repoLocalities)
	handler := handler.NewCarry(service)
	r.rg.GET("/carriers/:id", handler.Get())
	r.rg.GET("/localities/reportCarriers", handler.Read())
	r.rg.POST("/carriers", handler.Create())
}
func (r *router) buildProductRecordRoutes() {
	productRepo := product.NewRepository(r.db)
	productService := product.NewService(productRepo)

	repo := productrecord.NewRepository(r.db)
	service := productrecord.NewService(repo)
	handler := handler.NewProductRecord(service, productService)

	r.rg.POST("/productRecords", handler.Create())
	r.rg.GET("/products/reportRecords", handler.RecordsByAllProductsReport())
	r.rg.GET("/products/reportRecords/:id", handler.RecordsByOneProductReport())
}

func (r *router) buildSwagger() {
	docs.SwaggerInfo.BasePath = "/"
	r.rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)
	r.rg.GET("/teste", handler.GetAll())
}
