package requests

import (
	"github.com/daffashafwan/pointcuan/business/redeem"
)

type RedeemRequest struct {
	UserId     int    `json:"userId"`
	ItemId     int    `json:"itemId"`
	DataRedeem string `json:"dataRedeem"`
}

func (tr *RedeemRequest) ToDomain() redeem.Domain {
	return redeem.Domain{
		UserId:     tr.UserId,
		ItemId:     tr.ItemId,
		DataRedeem: tr.DataRedeem,
	}
}
