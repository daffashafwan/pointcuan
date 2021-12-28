package requests

import (
	"github.com/daffashafwan/pointcuan/business/transactions"
	"time"
)

type TransactionRequest struct {
	UserId                int       `json:"userId"`
	TransactionDate       string `json:"transactionDate"`
	Transaction           string    `json:"transaction"`
	TransactionAttachment string    `json:"transactionAttachment"`
	Status                int8      `json:"status"`
	Description           string    `json:"description"`
}

func (tr *TransactionRequest) ToDomain() transactions.Domain {
	layout := "2006-01-02"
	date,_ := time.Parse(layout, tr.TransactionDate)
	return transactions.Domain{
		UserId:                tr.UserId,
		TransactionDate:       date,
		TransactionAttachment: tr.TransactionAttachment,
		Description:           tr.Description,
		Status:                tr.Status,
		Transaction:           tr.Transaction,
	}
}
