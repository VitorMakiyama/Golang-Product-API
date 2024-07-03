package handlers

import (
	"api-produtos/internal/core/ports"
	"api-produtos/internal/handlers/dtos"
	"encoding/json"
	"github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var log, _ = zap.NewProduction()

const logMessage = "ProductHandler:"

type Handler struct {
	service ports.ProductService
}

func NewProductHandler(productService ports.ProductService) *Handler {
	return &Handler{service: productService}
}

func (h *Handler) CreateProduct(req *restful.Request, res *restful.Response) {
	newProduct := new(dtos.ProductDTO)
	if err := req.ReadEntity(newProduct); err != nil {
		_ = res.WriteError(http.StatusBadRequest, err)
		return
	}

	ps, err := h.service.CreateProduct(*newProduct.ToDomain())
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	var createdDTOs dtos.ProductDTOList
	createdDTOs.FromDomain(ps)

	logProduct, _ := json.Marshal(newProduct)
	log.Info(logMessage + " created product: " + string(logProduct))
	_ = res.WriteHeaderAndJson(http.StatusCreated, createdDTOs, restful.MIME_JSON)
}

func (h *Handler) GetAllProducts(req *restful.Request, res *restful.Response) {
	ps, err := h.service.GetAllProducts()
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	var list dtos.ProductDTOList
	list.FromDomain(ps)

	logProducts, _ := json.Marshal(ps)
	log.Info(logMessage + "gotten products: " + string(logProducts))
	_ = res.WriteAsJson(list)
}

func (h *Handler) GetProduct(req *restful.Request, res *restful.Response) {
	idStr := req.PathParameter("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	p, err := h.service.GetProduct(id)
	if err != nil {
		_ = res.WriteError(http.StatusNotFound, err)
		return
	}

	logProduct, _ := json.Marshal(p)
	log.Info(logMessage + "got product: " + string(logProduct))
	_ = res.WriteAsJson(p)
}

func (h *Handler) UpdateProduct(req *restful.Request, res *restful.Response) {
	idStr := req.PathParameter("id")
	id, err := strconv.Atoi(idStr)
	update := new(dtos.ProductDTO)
	err2 := req.ReadEntity(update)
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	} else if err2 != nil {
		_ = res.WriteError(http.StatusInternalServerError, err2)
		return
	}

	p, err := h.service.UpdateProduct(id, *update.ToDomain())
	if err != nil {
		_ = res.WriteError(http.StatusBadRequest, err)
		return
	}

	response := new(dtos.ProductDTO)
	response.FromDomain(*p)
	logP, _ := json.Marshal(p)
	log.Info(logMessage + "updated product id " + strconv.Itoa(id) + ": " + string(logP))
	_ = res.WriteAsJson(response)
}

func (h *Handler) DeleteProduct(req *restful.Request, res *restful.Response) {
	idStr := req.PathParameter("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		_ = res.WriteError(http.StatusBadRequest, err)
		return
	}

	log.Info(logMessage + "deleted product id " + strconv.Itoa(id))
	res.WriteHeader(http.StatusNoContent)
}
