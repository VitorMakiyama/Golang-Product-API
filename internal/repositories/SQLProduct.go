package repositories

import (
	"api-produtos/internal/core/domain"
)

type sqlProduct struct {
	id          int
	name        string
	description string
	price       float32
}

func (s *sqlProduct) ToDomain() *domain.Product {
	p := new(domain.Product)
	p.Id = s.id
	p.Name = s.name
	p.Description = s.description
	p.Price = s.price

	return p
}

func (s *sqlProduct) FromDomain(p *domain.Product) {
	s.id = p.Id
	s.name = p.Name
	s.description = p.Description
	s.price = p.Price
}
