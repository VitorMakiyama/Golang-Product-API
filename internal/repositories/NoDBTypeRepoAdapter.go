package repositories

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"errors"
)

type noDBTypeRepoAdapter struct {
	types    []domain.ProductType
	globalId int
}

func NewNoDBTypeRepo() ports.ProductTypeRepository {
	return &noDBTypeRepoAdapter{}
}

func (n *noDBTypeRepoAdapter) CreateType(newType domain.ProductType) ([]domain.ProductType, error) {
	newType.Id = n.globalId
	n.globalId++
	n.types = append(n.types, newType)

	return n.types, nil
}

func (n *noDBTypeRepoAdapter) GetType(id int) (*domain.ProductType, error) {
	for _, t := range n.types {
		if t.Id == id {
			return &t, nil
		}
	}
	return nil, errors.New("")
}

func (n *noDBTypeRepoAdapter) GetAllTypes() ([]domain.ProductType, error) {
	//TODO implement me
	panic("implement me")
}

func (n *noDBTypeRepoAdapter) UpdateType(id int, update domain.ProductType) (*domain.ProductType, error) {
	//TODO implement me
	panic("implement me")
}

func (n *noDBTypeRepoAdapter) DeleteType(id int, active bool) error {
	//TODO implement me
	panic("implement me")
}

func (n *noDBTypeRepoAdapter) CheckExistence(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (n *noDBTypeRepoAdapter) ValidateType(id int) bool {
	//TODO implement me
	panic("implement me")
}

func (n *noDBTypeRepoAdapter) GetTypeByName(name string) (*domain.ProductType, error) {
	//TODO implement me
	panic("implement me")
}
