package ports

import "api-produtos/internal/core/domain"

type ProductRepository interface {
	CreateProduct(product domain.Product) error
	GetProduct(id int) (*domain.Product, error)
	GetAllProducts(name string, pTypeId int, minPrice float32, maxPrice float32) ([]domain.Product, error)
	UpdateProduct(id int, update domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
	CheckExistence(name string) bool
}
