package transactions

import (
	"context"
	"errors"

	"github.com/daffashafwan/pointcuan/business/transactions"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func CreateTransactionRepo(conn *gorm.DB) transactions.Repository {
	return &TransactionRepo{
		DB: conn,
	}
}

func (rep *TransactionRepo) Create(ctx context.Context,transR *transactions.Domain) (transactions.Domain, error) {
	user := Transaction{
		UserId: transR.UserId,
		Transaction: transR.Transaction,
		Description: transR.Description,
		Point: transR.Point,
		TransactionAttachment: transR.TransactionAttachment,
		TransactionDate: transR.TransactionDate,
		Status: transR.Status,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (rep *TransactionRepo) Update(ctx context.Context, transU transactions.Domain) (transactions.Domain, error) {
	data := FromDomain(transU)
	err := rep.DB.Table("transactions").First(&data)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	data.Status = transU.Status
	data.Description = transU.Description
	if rep.DB.Save(&data).Error != nil {
		return transactions.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *TransactionRepo) GetAll(ctx context.Context) ([]transactions.Domain, error) {
	var data []Transaction
	err := rep.DB.Table("transactions").Preload("User").Find(&data)
	if err.Error != nil {
		return []transactions.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *TransactionRepo) GetById(ctx context.Context, id int) (transactions.Domain, error) {
	var data Transaction
	err := rep.DB.Table("transactions").Find(&data, "id=?", id)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}


func (rep *TransactionRepo) GetByUserId(ctx context.Context, id int) ([]transactions.Domain, error) {
	var data []Transaction
	err := rep.DB.Table("transactions").Find(&data, "user_id=?", id)
	if err.Error != nil {
		return []transactions.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *TransactionRepo) GetByUserIdAndStatus(ctx context.Context, id int, sid int) ([]transactions.Domain, error) {
	var data []Transaction
	err := rep.DB.Table("transactions").Find(&data, "user_id=?", id).Find(&data, "status=?", sid)
	if err.Error != nil {
		return []transactions.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *TransactionRepo) Delete(ctx context.Context, id int) error {
	user := Transaction{}
	err := rep.DB.Table("transactions").Where("id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}