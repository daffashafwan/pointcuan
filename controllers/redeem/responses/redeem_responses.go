package responses

import (
	"github.com/daffashafwan/pointcuan/business/redeem"
	"time"
)

type RedeemResponse struct {
	Id         int       `json:"id"`
	UserId     int       `json:"userId"`
	ItemId     int       `json:"itemId"`
	RedeemDate time.Time `json:"redeemDate"`
	DataRedeem string    `json:"dataRedeem"`
	Status     int       `json:"status"`
}

func FromDomain(domain redeem.Domain) RedeemResponse {
	return RedeemResponse{
		Id:         domain.Id,
		UserId:     domain.UserId,
		DataRedeem: domain.DataRedeem,
		ItemId:     domain.ItemId,
		RedeemDate: domain.CreatedAt,
		Status:     domain.Status,
	}
}

func FromListDomain(domain []redeem.Domain) []RedeemResponse {
	var list []RedeemResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
