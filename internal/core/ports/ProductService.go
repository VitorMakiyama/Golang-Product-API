package ports

import "api-produtos/internal/core/domain"

type ProductService interface {
	GetProduct(id int) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	CreateProduct(product domain.Product) ([]domain.Product, error)
	UpdateProduct(id int, update domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
}
