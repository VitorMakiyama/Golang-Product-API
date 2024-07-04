package ports

import "api-produtos/internal/core/domain"

type ProductService interface {
	CreateProduct(product domain.Product) ([]domain.Product, error)
	GetProduct(id int) (*domain.Product, error)
	GetAllProducts(queryParams ...string) ([]domain.Product, error)
	UpdateProduct(id int, update domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
}
