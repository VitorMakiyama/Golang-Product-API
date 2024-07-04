package handlers

import (
	"api-produtos/internal/core/ports"
	"api-produtos/internal/handlers/dtos"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"strconv"
)

const typeHLogMessage = "TypeHandler:"

type TypeHandler struct {
	service ports.ProductTypeService
}

func NewProductTypeHandler(typeService ports.ProductTypeService) *TypeHandler {
	return &TypeHandler{service: typeService}
}

func (h *TypeHandler) CreateType(req *restful.Request, res *restful.Response) {
	newType := new(dtos.ProductTypeDTO)
	if err := req.ReadEntity(newType); err != nil {
		_ = res.WriteError(http.StatusBadRequest, err)
		return
	}

	ps, err := h.service.CreateType(*newType.ToDomain())
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	var createdDTOs dtos.ProductTypeListDTO
	createdDTOs.FromDomain(ps)

	logType := fmt.Sprintf("%s created type: %v", typeHLogMessage, newType)
	log.Info(logType)
	_ = res.WriteHeaderAndJson(http.StatusCreated, createdDTOs, restful.MIME_JSON)
}

func (h *TypeHandler) GetType(req *restful.Request, res *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		_ = res.WriteError(http.StatusBadRequest, err)
		return
	}

	t, err := h.service.GetType(id)
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	logType := fmt.Sprintf("%s got type: %v", typeHLogMessage, t)
	log.Info(logType)
	_ = res.WriteAsJson(t)
}

func (h *TypeHandler) GetAllTypes(req *restful.Request, res *restful.Response) {
	ts, err := h.service.GetAllTypes()
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	var list dtos.ProductTypeListDTO
	list.FromDomain(ts)

	logType := fmt.Sprintf("%s gotten types: %v", typeHLogMessage, ts)
	log.Info(logType)
	_ = res.WriteAsJson(list)
}

func (h *TypeHandler) UpdateType(req *restful.Request, res *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	update := new(dtos.ProductTypeDTO)
	err2 := req.ReadEntity(update)
	if err != nil {
		_ = res.WriteError(http.StatusBadRequest, err)
		return
	} else if err2 != nil {
		_ = res.WriteError(http.StatusBadRequest, err2)
		return
	}

	t, err := h.service.UpdateType(id, *update.ToDomain())
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	response := new(dtos.ProductTypeDTO)
	response.FromDomain(*t)
	logType := fmt.Sprintf("%s updated type: %v", typeHLogMessage, t)
	log.Info(logType)
	_ = res.WriteAsJson(response)
}

func (h *TypeHandler) DeleteType(req *restful.Request, res *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	active, err2 := strconv.ParseBool(req.QueryParameter("active"))
	if err != nil {
		_ = res.WriteError(http.StatusBadRequest, err)
		return
	} else if err2 != nil {
		_ = res.WriteError(http.StatusBadRequest, err2)
		return
	}

	err = h.service.DeleteType(id, active)
	if err != nil {
		_ = res.WriteError(http.StatusInternalServerError, err)
		return
	}

	logType := fmt.Sprintf("%s changed type 'active' to : %t", typeHLogMessage, active)
	log.Info(logType)
	res.WriteHeader(http.StatusNoContent)
}
