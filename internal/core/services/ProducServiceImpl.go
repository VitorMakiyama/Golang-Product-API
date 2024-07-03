package services

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"slices"
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
		return nil, errors.New(fmt.Sprintf("%s type id %d does not exists", errorMessage, product.Type.Id))
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

func (s productServiceImpl) GetAllProducts() ([]domain.Product, error) {
	ps, err := s.pRepo.GetAllProducts()
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
