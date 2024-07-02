package repositories

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"errors"
	"slices"
)

type noDBRepository struct {
	products []domain.Product
}

var globalId = 0

const errorMessage = "NoDBRepository error: "

func NewNoDBRepository() ports.ProductRepository {
	return &noDBRepository{}
}

func (db *noDBRepository) GetProduct(id int) (*domain.Product, error) {
	index := slices.IndexFunc(db.products, func(p domain.Product) bool { return p.Id == id })
	if index == -1 {
		return nil, errors.New(errorMessage + "id not found")
	}

	return &db.products[index], nil
}

func (db *noDBRepository) GetAllProducts() ([]domain.Product, error) {
	return db.products, nil
}

func (db *noDBRepository) CreateProduct(product domain.Product) ([]domain.Product, error) {
	product.Id = globalId
	globalId++
	db.products = append(db.products, product)
	return db.products, nil
}

func (db *noDBRepository) UpdateProduct(id int, update domain.Product) (*domain.Product, error) {
	index := slices.IndexFunc(db.products, func(p domain.Product) bool { return p.Id == id })
	if index == -1 {
		return nil, errors.New(errorMessage + "id not found")
	}

	db.products[index].Update(update)
	return &db.products[index], nil
}

func (db *noDBRepository) DeleteProduct(id int) error {
	index := slices.IndexFunc(db.products, func(p domain.Product) bool { return p.Id == id })
	if index == -1 {
		return errors.New(errorMessage + "id not found")
	}

	// this deletion method is called "re-slicing", quite expensive, but it's the way to go in Golang
	db.products = append(db.products[:index], db.products[index+1:]...)
	return nil
}
