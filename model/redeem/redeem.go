package redeem

import (
	"time"

	"github.com/daffashafwan/pointcuan/business/redeem"
	"github.com/daffashafwan/pointcuan/model/items"
)

type Redeem struct {
	Id         int `gorm:"primaryKey"`
	UserId     int
	ItemId     int
	Item       items.Items `gorm:"foreignKey:ItemId;association_foreignkey:Id"`
	DataRedeem string
	Point      int
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (tr *Redeem) ToDomain() redeem.Domain {
	return redeem.Domain{
		Id:         tr.Id,
		UserId:     tr.UserId,
		ItemId:     tr.ItemId,
		Item:       tr.Item,
		Point:      tr.Point,
		DataRedeem: tr.DataRedeem,
		Status:     tr.Status,
		CreatedAt:  tr.CreatedAt,
		UpdatedAt:  tr.UpdatedAt,
	}
}

func ToListDomain(data []Redeem) (result []redeem.Domain) {
	result = []redeem.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain redeem.Domain) Redeem {
	return Redeem{
		Id:         domain.Id,
		UserId:     domain.UserId,
		ItemId:     domain.ItemId,
		Point:      domain.Point,
		DataRedeem: domain.DataRedeem,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
