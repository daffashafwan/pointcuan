package requests

import (
	"time"

	"github.com/daffashafwan/pointcuan/business/transactions"
)

type TransactionRequest struct {
	UserId                int    `json:"userId"`
	TransactionDate       string `json:"transactionDate"`
	Transaction           float64 `json:"transaction"`
	TransactionAttachment string `json:"transactionAttachment"`
	Status                int8   `json:"status"`
	Point                 float64
	Description           string `json:"description"`
}

func (tr *TransactionRequest) ToDomain() transactions.Domain {
	layout := "2006-01-02"
	date, _ := time.Parse(layout, tr.TransactionDate)
	return transactions.Domain{
		UserId:                tr.UserId,
		TransactionDate:       date,
		Point:                 tr.Point,
		TransactionAttachment: tr.TransactionAttachment,
		Description:           tr.Description,
		Status:                tr.Status,
		Transaction:           tr.Transaction,
	}
}
