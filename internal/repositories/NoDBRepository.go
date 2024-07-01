package repositories

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"log"
)

type noDBRepository struct {
	products []domain.Product
}

func NewNoDBRepository() ports.ProductRepository {
	return &noDBRepository{}
}

func (db *noDBRepository) GetProduct(id int) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (db *noDBRepository) GetAllProducts() ([]domain.Product, error) {
	return db.products, nil
}

func (db *noDBRepository) CreateProduct(product domain.Product) ([]domain.Product, error) {
	db.products = append(db.products, product)
	log.Println("passou no db: ", db.products)
	return db.products, nil
}

func (db *noDBRepository) UpdateProduct(id int) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (db *noDBRepository) DeleteProduct(id int) error {
	//TODO implement me
	panic("implement me")
}
