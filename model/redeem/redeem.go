package redeem

import (
	"github.com/daffashafwan/pointcuan/business/redeem"
	"time"
)

type Redeem struct {
	Id         int `gorm:"primaryKey"`
	UserId     int
	ItemId     int
	DataRedeem string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (tr *Redeem) ToDomain() redeem.Domain {
	return redeem.Domain{
		Id:         tr.Id,
		UserId:     tr.UserId,
		ItemId:     tr.ItemId,
		DataRedeem: tr.DataRedeem,
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
		DataRedeem: domain.DataRedeem,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
