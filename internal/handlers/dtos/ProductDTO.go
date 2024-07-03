package dtos

import (
	"api-produtos/internal/core/domain"
)

type ProductDTO struct {
	Id          int
	Name        string
	Description string
	Price       float32
	Type        ProductTypeDTO
}
type ProductDTOList []ProductDTO

func (dto *ProductDTO) FromDomain(p domain.Product) {
	t := new(ProductTypeDTO)
	t.FromDomain(p.Type)
	dto.Id = p.Id
	dto.Name = p.Name
	dto.Description = p.Description
	dto.Price = p.Price
	dto.Type = *t
}

func (dtos *ProductDTOList) FromDomain(ps []domain.Product) {
	for _, p := range ps {
		dto := ProductDTO{}
		dto.FromDomain(p)
		*dtos = append(*dtos, dto)
	}
}

func (dto *ProductDTO) ToDomain() *domain.Product {
	return &domain.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Type:        *dto.Type.ToDomain(),
	}
}
