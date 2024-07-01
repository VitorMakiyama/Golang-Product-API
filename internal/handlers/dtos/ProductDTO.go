package dtos

import (
	"api-produtos/internal/core/domain"
)

type ProductDTO struct {
	Name        string
	Description string
	Price       float32
}
type ProductDTOList []ProductDTO

func (dto *ProductDTO) FromDomain(p domain.Product) {

	dto.Name = p.Name
	dto.Description = p.Description
	dto.Price = p.Price
}

func (dtos ProductDTOList) FromDomain(ps []domain.Product) ProductDTOList {
	for _, p := range ps {
		dto := ProductDTO{}
		dto.FromDomain(p)
		dtos = append(dtos, dto)
	}

	return dtos
}

func (dto ProductDTO) ToDomain() *domain.Product {
	return &domain.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
}
