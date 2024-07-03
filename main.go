package main

import (
	"api-produtos/internal/core/services"
	"api-produtos/internal/handlers"
	"api-produtos/internal/repositories"
	"flag"
	"github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
	"log"
	"net/http"
)

var binding string
var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
	flag.StringVar(&binding, "httpbind", ":8000", "address/port to bind listen socket")
}

func main() {
	db := repositories.StartMysqlDb()
	var productRepository = repositories.NewSQLProductRepository(db) //repositories.NewNoDBRepository()
	var productTypeRepository = repositories.NewSQLTypeRepository(db)
	var productService = services.NewProductService(productRepository, productTypeRepository)
	var productTypeService = services.NewProductTypeService(productTypeRepository)

	pws := new(restful.WebService)
	pws = pws.Path("/products")
	productHandler := handlers.NewProductHandler(productService)

	tws := new(restful.WebService)
	tws = tws.Path("/product-types")
	productTypeHandler := handlers.NewProductTypeHandler(productTypeService)

	restful.Add(pws)
	restful.Add(tws)

	// Products routes
	pws.Route(pws.GET("").To(productHandler.GetAllProducts).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	pws.Route(pws.GET("/{id}").To(productHandler.GetProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	pws.Route(pws.POST("").To(productHandler.CreateProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	pws.Route(pws.PATCH("/{id}").To(productHandler.UpdateProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	pws.Route(pws.DELETE("/{id}").To(productHandler.DeleteProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	// Product Types routes
	tws.Route(tws.GET("").To(productTypeHandler.GetAllTypes).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	tws.Route(tws.GET("/{id}").To(productTypeHandler.GetType).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	tws.Route(tws.POST("").To(productTypeHandler.CreateType).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	tws.Route(tws.PATCH("/{id}").To(productTypeHandler.UpdateType).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	tws.Route(tws.DELETE("/{id}").To(productTypeHandler.DeleteType).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	logger.Info("Listening...")
	log.Panicln(http.ListenAndServe(binding, nil))
}
