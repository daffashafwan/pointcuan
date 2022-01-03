package requests

import categoryItem "github.com/daffashafwan/pointcuan/business/categoryItems"

type CategoryItem struct {
	Name    string `json:"category"`
}

func (cat *CategoryItem) ToDomain() categoryItem.Domain {
	return categoryItem.Domain{
		Name:    cat.Name,
	}
}
