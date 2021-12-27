package requests

import (
	"github.com/daffashafwan/pointcuan/business/transactions"
)

type TransactionRequest struct {
	UserId                int       `json:"userId"`
	TransactionDate       string `json:"transactionDate"`
	Transaction           string    `json:"transaction"`
	TransactionAttachment string    `json:"transactionAttachment"`
	Status                int8      `json:"status"`
	Point                 float64   `json:"point"`
	Description           string    `json:"description"`
}

func (tr *TransactionRequest) ToDomain() transactions.Domain {
	return transactions.Domain{
		UserId:                tr.UserId,
		Point:                 tr.Point,
		TransactionDate:       tr.TransactionDate,
		TransactionAttachment: tr.TransactionAttachment,
		Description:           tr.Description,
		Status:                tr.Status,
		Transaction:           tr.Transaction,
	}
}
