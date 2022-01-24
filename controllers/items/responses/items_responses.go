package responses

import (
	"github.com/daffashafwan/pointcuan/business/items"
	"strconv"
)

type ItemsResponse struct {
	Id          int         `json:"id"`
	CategoryId  int         `json:"categoryId"`
	Category    interface{} `json:"category"`
	Name        string      `json:"name"`
	PointRedeem string         `json:"pointRedeem"`
	Stock       string      `json:"stock"`
}

func FromDomain(domain items.Domain) ItemsResponse {
	pointR := strconv.Itoa(domain.PointRedeem)
	stockR := strconv.Itoa(domain.Stock)
	return ItemsResponse{
		Id:          domain.Id,
		CategoryId:  domain.CategoryId,
		Category:    domain.Category,
		Name:        domain.Name,
		PointRedeem: pointR,
		Stock:       stockR,
	}
}

func FromListDomain(domain []items.Domain) []ItemsResponse {
	var list []ItemsResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
