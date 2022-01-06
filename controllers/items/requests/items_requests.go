package requests

import (
	"github.com/daffashafwan/pointcuan/business/items"
)

type ItemRequest struct {
	CategoryId  int    `json:"categoryId"`
	Name        string `json:"name"`
	PointRedeem int    `json:"pointRedeem"`
	Stock       string `json:"stock"`
}

type ItemRequestStock struct {
	Stock string `json:"stock"`
}

func (item *ItemRequestStock) StockToDomain() items.Domain {
	return items.Domain{
		Stock: item.Stock,
	}
}

func (item *ItemRequest) ToDomain() items.Domain {
	return items.Domain{
		CategoryId:  item.CategoryId,
		Name:        item.Name,
		PointRedeem: item.PointRedeem,
		Stock:       item.Stock,
	}
}
