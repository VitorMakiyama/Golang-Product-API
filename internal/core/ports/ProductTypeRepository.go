package ports

import "api-produtos/internal/core/domain"

type ProductTypeRepository interface {
	CreateType(newType domain.ProductType) ([]domain.ProductType, error)
	GetType(id int) (*domain.ProductType, error)
	GetAllTypes() ([]domain.ProductType, error)
	UpdateType(id int, update domain.ProductType) (*domain.ProductType, error)
	DeleteType(id int, active bool) error
	CheckExistence(name string) bool
	ValidateType(id int) bool
	GetTypeByName(name string) (*domain.ProductType, error)
}
