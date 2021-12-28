package transactions

import (
	"github.com/daffashafwan/pointcuan/business/transactions"
	"time"
)

type Transaction struct {
	Id                    int `gorm:"primaryKey"`
	UserId                int
	TransactionDate       time.Time
	Transaction           string
	TransactionAttachment string
	Status                int8
	Point                 float64
	Description           string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (tr *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:                    tr.Id,
		UserId:                tr.UserId,
		TransactionDate:       tr.TransactionDate,
		Transaction:           tr.Transaction,
		TransactionAttachment: tr.TransactionAttachment,
		Point:                 tr.Point,
		Description:           tr.Description,
		Status:                tr.Status,
		CreatedAt:             tr.CreatedAt,
		UpdatedAt:             tr.UpdatedAt,
	}
}

func ToListDomain(data []Transaction) (result []transactions.Domain) {
	result = []transactions.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		Id:                    domain.Id,
		UserId:                domain.UserId,
		TransactionDate:       domain.TransactionDate,
		Transaction:           domain.Transaction,
		TransactionAttachment: domain.TransactionAttachment,
		Point:                 domain.Point,
		Description:           domain.Description,
		Status:                domain.Status,
		CreatedAt:             domain.CreatedAt,
		UpdatedAt:             domain.UpdatedAt,
	}
}
