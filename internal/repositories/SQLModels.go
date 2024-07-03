package repositories

import (
	"api-produtos/internal/core/domain"
)

type sqlProduct struct {
	id          int
	name        string
	description string
	price       float32
	typeId      int
}

func (s *sqlProduct) ToDomain() *domain.Product {
	p := new(domain.Product)
	p.Id = s.id
	p.Name = s.name
	p.Description = s.description
	p.Price = s.price
	p.Type.Id = s.typeId
	return p
}

func (s *sqlProduct) FromDomain(p *domain.Product) {
	s.id = p.Id
	s.name = p.Name
	s.description = p.Description
	s.price = p.Price
	s.typeId = p.Type.Id
}

type sqlProductType struct {
	id     int
	name   string
	active bool
}

func (s *sqlProductType) ToDomain() *domain.ProductType {
	t := new(domain.ProductType)
	t.Id = s.id
	t.Name = s.name
	t.Active = s.active
	return t
}

func (s *sqlProductType) FromDomain(t *domain.ProductType) {
	s.id = t.Id
	s.name = t.Name
	s.active = t.Active
}
