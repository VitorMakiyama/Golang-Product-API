package dtos

import "api-produtos/internal/core/domain"

type ProductTypeDTO struct {
	Id     int
	Name   string
	Active bool
}

type ProductTypeListDTO []ProductTypeDTO

func (d *ProductTypeDTO) FromDomain(t domain.ProductType) {
	d.Id = t.Id
	d.Name = t.Name
	d.Active = t.Active
}

func (d *ProductTypeListDTO) FromDomain(ts []domain.ProductType) {
	for _, t := range ts {
		dto := ProductTypeDTO{}
		dto.FromDomain(t)
		*d = append(*d, dto)
	}
}

func (d *ProductTypeDTO) ToDomain() *domain.ProductType {
	return &domain.ProductType{
		Id:     d.Id,
		Name:   d.Name,
		Active: d.Active,
	}
}
