package dtos

import (
	"api-produtos/internal/core/domain"
)

type ProductDTO struct {
	Id          int
	Name        string
	Description string
	Price       float32
	TypeId      int
	TypeName    string
}
type ProductDTOList []ProductDTO

func (dto *ProductDTO) FromDomain(p domain.Product) {
	dto.Id = p.Id
	dto.Name = p.Name
	dto.Description = p.Description
	dto.Price = p.Price
	dto.TypeId = p.Type.Id
	dto.TypeName = p.Type.Name
}

func (dtos *ProductDTOList) FromDomain(ps []domain.Product) {
	for _, p := range ps {
		dto := ProductDTO{}
		dto.FromDomain(p)
		*dtos = append(*dtos, dto)
	}
}

func (dto *ProductDTO) ToDomain() *domain.Product {
	t := domain.ProductType{
		Id:   dto.TypeId,
		Name: dto.TypeName,
	}
	return &domain.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Type:        t,
	}
}
