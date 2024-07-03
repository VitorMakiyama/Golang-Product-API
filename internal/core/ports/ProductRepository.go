package ports

import "api-produtos/internal/core/domain"

type ProductRepository interface {
	CreateProduct(product domain.Product) error
	GetProduct(id int) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	UpdateProduct(id int, update domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
}
