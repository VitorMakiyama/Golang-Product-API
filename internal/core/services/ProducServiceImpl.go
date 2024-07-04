package services

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"slices"
	"strconv"
)

var log, _ = zap.NewProduction()

const errorMessage = "Service error:"
const errorStr = "- error:"

// when variable/function starts with lowercase it will be protected, Uppercase will be public
type productServiceImpl struct {
	pRepo ports.ProductRepository
	tRepo ports.ProductTypeRepository
}

func NewProductService(pRepo ports.ProductRepository, tRepo ports.ProductTypeRepository) ports.ProductService {
	return &productServiceImpl{
		pRepo: pRepo,
		tRepo: tRepo,
	}
}

func (s productServiceImpl) CreateProduct(product domain.Product) ([]domain.Product, error) {
	if product.Name == "" || product.Price <= 0 {
		return nil, errors.New(errorMessage + " empty name or invalid price ")
	}
	if !s.tRepo.ValidateType(product.Type.Id) {
		return nil, errors.New(fmt.Sprintf("%s type id %d does not exists or is inactive", errorMessage, product.Type.Id))
	}
	if s.pRepo.CheckExistence(product.Name) {
		return nil, errors.New(fmt.Sprintf("%s name '%s' already exists", errorMessage, product.Name))
	}

	err := s.pRepo.CreateProduct(product)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error creating product %s %s", errorMessage, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	p, _ := s.GetAllProducts()
	return p, nil
}

func (s productServiceImpl) GetProduct(id int) (*domain.Product, error) {
	p, err := s.pRepo.GetProduct(id)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error getting product id: %d %s %s", errorMessage, id, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	t, _ := s.tRepo.GetType(p.Type.Id)
	p.Type = *t

	return p, nil
}

func (s productServiceImpl) GetAllProducts(queryParams ...string) ([]domain.Product, error) {
	t := new(domain.ProductType)
	n, tName, minP, maxP := getQueryParamsValues(queryParams)
	if maxP >= 0 && minP >= maxP {
		errLog := fmt.Sprintf("%s minPrice (%.2f) is greater or equal to maxPrice (%.2f)", errorMessage, minP, maxP)
		log.Error(errLog)
		return nil, errors.New(errLog)
	}
	if tName != "" {
		var err error
		t, err = s.tRepo.GetTypeByName(tName)
		if err != nil {
			errLog := fmt.Sprintf("%s type '%s' not found!", errorMessage, tName)
			log.Error(errLog)
			err.Error()
			return nil, errors.New(errLog)
		}
	} else {
		t.Id = -1
	}

	ps, err := s.pRepo.GetAllProducts(n, t.Id, minP, maxP)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error getting all products %s %s", errorMessage, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	ts, _ := s.tRepo.GetAllTypes()
	for i, p := range ps {
		t := ts[slices.IndexFunc(ts, func(t domain.ProductType) bool { return p.Type.Id == t.Id })]
		ps[i].Type = t
	}

	return ps, nil
}

func getQueryParamsValues(params []string) (string, string, float32, float32) {
	if len(params) > 0 {
		n := params[0]
		t := params[1]
		minP, _ := strconv.ParseFloat(params[2], 32)
		maxP, err := strconv.ParseFloat(params[3], 32)
		if err != nil {
			maxP = -1
		}
		return n, t, float32(minP), float32(maxP)
	}
	return "", "", 0, -1
}

func (s productServiceImpl) UpdateProduct(id int, update domain.Product) (*domain.Product, error) {
	p, err := s.pRepo.UpdateProduct(id, update)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error updating product id: %d %s %s", errorMessage, id, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	p, _ = s.GetProduct(id)
	return p, nil
}

func (s productServiceImpl) DeleteProduct(id int) error {
	err := s.pRepo.DeleteProduct(id)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error deleting product id: %d %s %s", errorMessage, id, errorStr, err.Error())
		log.Error(errLog)
		return errors.New(errLog)
	}

	return nil
}
