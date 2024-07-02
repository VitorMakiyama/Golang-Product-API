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
	var repository = repositories.NewNoDBRepository()
	var service = services.NewProductService(repository)

	ws := new(restful.WebService)
	ws = ws.Path("/products")
	handler := handlers.NewProductHandler(service)
	restful.Add(ws)

	// routes
	ws.Route(ws.GET("").To(handler.GetAllProducts).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/{id}").To(handler.GetProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.POST("").To(handler.CreateProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.PATCH("/{id}").To(handler.UpdateProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.DELETE("/{id}").To(handler.DeleteProduct).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	logger.Info("Listening...")
	log.Panicln(http.ListenAndServe(binding, nil))
}
