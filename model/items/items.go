package items

import (
	"time"

	"github.com/daffashafwan/pointcuan/business/items"
)

type Items struct {
	Id          int `gorm:"primaryKey"`
	CategoryId  int
	Name        string
	PointRedeem int
	Stock       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (item *Items) ToDomain() items.Domain {
	return items.Domain{
		Id:          item.Id,
		CategoryId:  item.CategoryId,
		Name:        item.Name,
		PointRedeem: item.PointRedeem,
		Stock:       item.Stock,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func ToListDomain(data []Items) (result []items.Domain) {
	result = []items.Domain{}
	for _, item := range data {
		result = append(result, item.ToDomain())
	}
	return
}

func FromDomain(domain items.Domain) Items {
	return Items{
		Id:          domain.Id,
		CategoryId:  domain.CategoryId,
		Name:        domain.Name,
		PointRedeem: domain.PointRedeem,
		Stock:       domain.Stock,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
