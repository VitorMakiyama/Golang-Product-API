package services

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"errors"
	"go.uber.org/zap"
	"strconv"
)

var log, _ = zap.NewProduction()

const errorMessage = "Service error: "
const errorStr = " - error: "

// when variable/function starts with lowercase it will be protected, Uppercase will be public
type productServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return productServiceImpl{
		repo: repo,
	}
}

func (s productServiceImpl) GetProduct(id int) (*domain.Product, error) {
	p, err := s.repo.GetProduct(id)
	if err != nil {
		errLog := err.Error()
		log.Error(errorMessage + "repository error getting product id " + strconv.Itoa(id) + errorStr + errLog)
		return nil, err
	}
	return p, nil
}

func (s productServiceImpl) GetAllProducts() ([]domain.Product, error) {
	p, err := s.repo.GetAllProducts()
	if err != nil {
		errLog := err.Error()
		log.Error(errorMessage + "repository error getting all products" + errorStr + errLog)
		return nil, err
	}
	return p, nil
}

func (s productServiceImpl) CreateProduct(product domain.Product) ([]domain.Product, error) {
	if product.Name == "" || product.Price <= 0 {
		return nil, errors.New(errorMessage + " empty name or invalid price")
	}

	p, err := s.repo.CreateProduct(product)

	if err != nil {
		errLog := err.Error()
		log.Error(errorMessage + " repository error creating product" + errorStr + errLog)
		return nil, err
	}
	return p, nil
}

func (s productServiceImpl) UpdateProduct(id int, update domain.Product) (*domain.Product, error) {
	p, err := s.repo.UpdateProduct(id, update)

	if err != nil {
		errLog := err.Error()
		log.Error(errorMessage + " repository error updating product " + strconv.Itoa(id) + errorStr + errLog)
		return nil, err
	}

	return p, nil
}

func (s productServiceImpl) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)

	if err != nil {
		errLog := err.Error()
		log.Error(errorMessage + " repository error deleting product " + strconv.Itoa(id) + errorStr + errLog)
		return err
	}

	return nil
}
