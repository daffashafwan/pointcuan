package requests

package requests

import categoryItem "github.com/daffashafwan/pointcuan/business/categoryItem"

type categoryItem struct {
	ctegoryItem    int `json:"pcrValue"`
}

func (categoryItem *categoryItems) ToDomain() categoryItem.Domain {
	return pcrcrud.Domain{
		categoryName:    categoryItem.categoryName,
	}
}
