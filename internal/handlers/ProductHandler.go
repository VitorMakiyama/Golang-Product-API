package handlers

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"api-produtos/internal/handlers/dtos"
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

type handler struct {
	service ports.ProductService
}

func NewProductHandler(productService ports.ProductService) *handler {
	return &handler{service: productService}
}

func (h *handler) CreateProduct(req *restful.Request, res *restful.Response) {
	newProduct := new(domain.Product)
	if err := req.ReadEntity(newProduct); err != nil {
		res.WriteError(400, err)
		return
	}

	ps, err := h.service.CreateProduct(*newProduct)
	if err != nil {
		res.WriteError(500, err)
		return
	}

	var createdDTOs dtos.ProductDTOList
	createdDTOs.FromDomain(ps)

	res.WriteHeader(http.StatusCreated)
	res.WriteAsJson(createdDTOs)
}

func (h handler) GetAllProducts(req *restful.Request, res *restful.Response) {
	ps, err := h.service.GetAllProducts()
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
		return
	}

	var list dtos.ProductDTOList
	list.FromDomain(ps)

	res.WriteHeader(http.StatusOK)
	res.WriteAsJson(list)
}
