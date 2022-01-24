package items

import (
	"github.com/daffashafwan/pointcuan/business/items"
	"github.com/daffashafwan/pointcuan/model/category"
	"time"
)

type Items struct {
	Id          int `gorm:"primaryKey"`
	CategoryId  int
	Category    category.Category `gorm:"foreignKey:CategoryId;association_foreignkey:Id"`
	Name        string
	PointRedeem int
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (item *Items) ToDomain() items.Domain {
	return items.Domain{
		Id:          item.Id,
		CategoryId:  item.CategoryId,
		Category:    item.Category,
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
