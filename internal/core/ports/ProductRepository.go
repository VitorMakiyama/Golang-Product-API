package ports

import "api-produtos/internal/core/domain"

type ProductRepository interface {
	GetProduct(id int) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	CreateProduct(product domain.Product) ([]domain.Product, error)
	UpdateProduct(id int) (*domain.Product, error)
	DeleteProduct(id int) error
}
