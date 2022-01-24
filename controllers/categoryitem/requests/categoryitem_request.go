package requests

import categoryItem "github.com/daffashafwan/pointcuan/business/categoryItems"

type CategoryItem struct {
	Name string `json:"category"`
	Svg  string `json:"svg"`
}

func (cat *CategoryItem) ToDomain() categoryItem.Domain {
	return categoryItem.Domain{
		Name: cat.Name,
		Svg:  cat.Svg,
	}
}
