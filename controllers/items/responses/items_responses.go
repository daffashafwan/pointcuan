package responses

import (
	"github.com/daffashafwan/pointcuan/business/items"
)

type ItemsResponse struct {
	CategoryId int     `json:"categoryId"`
	Name  string `json:"name"`
	PointRedeem int `json:"pointRedeem"`
	Stock       int `json:"stock"`
}

func FromDomain(domain items.Domain) ItemsResponse {
	return ItemsResponse{
		CategoryId: domain.CategoryId,
		Name: domain.Name,
		PointRedeem: domain.PointRedeem,
		Stock: domain.Stock,
	}
}

func FromListDomain(domain []items.Domain) []ItemsResponse {
	var list []ItemsResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
