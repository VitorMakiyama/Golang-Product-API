package domain

type ProductType struct {
	Id     int
	Name   string
	Active bool
}

func (t *ProductType) Update(update ProductType) {
	if update.Name != "" {
		t.Name = update.Name
	}
	t.Active = update.Active
}
