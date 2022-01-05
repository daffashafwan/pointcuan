package requests

import "github.com/daffashafwan/pointcuan/business/items"

type ItemRequest struct {
	CategoryId int
	Name  string
	PointRedeem int
	Stock       int
}

func (item *ItemRequest) ToDomain() items.Domain {
	return items.Domain{
		CategoryId: item.CategoryId,
		Name:  item.Name,
		PointRedeem: item.PointRedeem,
		Stock: item.Stock,
	}
}
