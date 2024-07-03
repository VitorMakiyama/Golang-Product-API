package domain

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float32
	Type        ProductType
}

func (p *Product) Update(update Product) {
	if update.Name != "" {
		p.Name = update.Name
	}
	if update.Description != "" {
		p.Description = update.Description
	}
	if update.Price >= 0 {
		p.Price = update.Price
	}
	if update.Type.Id < 0 {
		p.Type = update.Type
	}
}
