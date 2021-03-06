package responses

import (
	"time"
	"fmt"
	"github.com/daffashafwan/pointcuan/business/transactions"
)

type TransactionResponse struct {
	Id                    int         `json:"id"`
	UserId                int         `json:"userId"`
	TransactionDate       time.Time   `json:"transactionDate"`
	Transaction           string     `json:"transaction"`
	TransactionAttachment string      `json:"transactionAttachment"`
	Status                int8        `json:"status"`
	Point                 float64     `json:"point"`
	Description           string      `json:"description"`
	User                  interface{} `json:"user"`
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	s := fmt.Sprintf("%f", domain.Transaction)
	return TransactionResponse{
		Id:                    domain.Id,
		UserId:                domain.UserId,
		TransactionDate:       domain.TransactionDate,
		Transaction:           s,
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
