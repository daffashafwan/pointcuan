package requests

import (
	"github.com/daffashafwan/pointcuan/business/items"
	"strconv"
)

type ItemRequest struct {
	CategoryId  int    `json:"categoryId"`
	Name        string `json:"name"`
	PointRedeem string    `json:"pointRedeem"`
	Stock       string `json:"stock"`
}

type ItemRequestStock struct {
	Stock string `json:"stock"`
}

func (item *ItemRequestStock) StockToDomain() items.Domain {
	stockR,_ := strconv.Atoi(item.Stock)
	return items.Domain{
		Stock: stockR,
	}
}

func (item *ItemRequest) ToDomain() items.Domain {
	pointR,_ := strconv.Atoi(item.PointRedeem)
	stockR,_ := strconv.Atoi(item.Stock)
	return items.Domain{
		CategoryId:  item.CategoryId,
		Name:        item.Name,
		PointRedeem: pointR,
		Stock:       stockR,
	}
}
