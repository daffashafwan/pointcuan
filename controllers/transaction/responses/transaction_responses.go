package responses

import (
	"github.com/daffashafwan/pointcuan/business/transactions"
	"time"
)

type TransactionResponse struct {
	UserId                int `json:"userId"`
	TransactionDate       time.Time `json:"transactionDate"`
	Transaction           string `json:"transaction"`
	TransactionAttachment string `json:"transactionAttachment"`
	Status                int8 `json:"status"`
	Point                 float64 `json:"point"`
	Description           string  `json:"description"`
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		UserId: domain.UserId,
		TransactionDate: domain.TransactionDate,
		Transaction: domain.Transaction,
		TransactionAttachment: domain.TransactionAttachment,
		Status: domain.Status,
		Description: domain.Description,
		Point:  domain.Point,
	}
}

func FromListDomain(domain []transactions.Domain) []TransactionResponse {
	var list []TransactionResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

