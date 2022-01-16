package responses

import (
	"time"

	"github.com/daffashafwan/pointcuan/business/transactions"
)

type TransactionResponse struct {
	UserId                int         `json:"userId"`
	TransactionDate       time.Time   `json:"transactionDate"`
	Transaction           float64     `json:"transaction"`
	TransactionAttachment string      `json:"transactionAttachment"`
	Status                int8        `json:"status"`
	Point                 float64     `json:"point"`
	Description           string      `json:"description"`
	User                  interface{} `json:"user"`
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		UserId:                domain.UserId,
		TransactionDate:       domain.TransactionDate,
		Transaction:           domain.Transaction,
		User:                  domain.User,
		TransactionAttachment: domain.TransactionAttachment,
		Status:                domain.Status,
		Description:           domain.Description,
		Point:                 domain.Point,
	}
}

func FromListDomain(domain []transactions.Domain) []TransactionResponse {
	var list []TransactionResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
