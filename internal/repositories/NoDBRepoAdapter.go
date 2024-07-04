package repositories

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"errors"
	"slices"
)

type noDBRepositoryAdapter struct {
	products []domain.Product
}

var globalId = 0

const errorMessage = "NoDBRepository error: "

func NewNoDBRepository() ports.ProductRepository {
	return &noDBRepositoryAdapter{}
}

func (db *noDBRepositoryAdapter) GetProduct(id int) (*domain.Product, error) {
	index := slices.IndexFunc(db.products, func(p domain.Product) bool { return p.Id == id })
	if index == -1 {
		return nil, errors.New(errorMessage + "id not found")
	}

	return &db.products[index], nil
}

func (db *noDBRepositoryAdapter) GetAllProducts(name string, pTypeId int, minPrice float32, maxPrice float32) ([]domain.Product, error) {
	return db.products, nil
}

func (db *noDBRepositoryAdapter) CreateProduct(product domain.Product) error {
	product.Id = globalId
	globalId++
	db.products = append(db.products, product)
	return nil
}

func (db *noDBRepositoryAdapter) UpdateProduct(id int, update domain.Product) (*domain.Product, error) {
	index := slices.IndexFunc(db.products, func(p domain.Product) bool { return p.Id == id })
	if index == -1 {
		return nil, errors.New(errorMessage + "id not found")
	}

	db.products[index].Update(update)
	return &db.products[index], nil
}

func (db *noDBRepositoryAdapter) DeleteProduct(id int) error {
	index := slices.IndexFunc(db.products, func(p domain.Product) bool { return p.Id == id })
	if index == -1 {
		return errors.New(errorMessage + "id not found")
	}

	// this deletion method is called "re-slicing", quite expensive, but it's the way to go in Golang
	db.products = append(db.products[:index], db.products[index+1:]...)
	return nil
}

func (db *noDBRepositoryAdapter) CheckExistence(name string) bool {
	//TODO implement me
	panic("implement me")
}
