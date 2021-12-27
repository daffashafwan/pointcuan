package transactions

import (
	"context"
	"github.com/daffashafwan/pointcuan/business/transactions"
	"github.com/daffashafwan/pointcuan/helpers/encrypt"
	"gorm.io/gorm"
	"errors"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func CreateTransactionRepo(conn *gorm.DB) transactions.Repository {
	return &TransactionRepo{
		DB: conn,
	}
}


func (rep *TransactionRepo) Login(ctx context.Context, username string, password string) (transactions.Domain, error) {
	var user Transaction
	result := rep.DB.Table("users").Where("username = ?", username).Where("status = ? ", "1").First(&user).Error

	if result != nil {
		return transactions.Domain{}, result
	}
	if !(encrypt.Compare(password,user.Password)) {
		return transactions.Domain{}, errors.New("Password tidak cocok")
	}
	return user.ToDomain(), nil

}

func (rep *TransactionRepo) Create(ctx context.Context,userR *transactions.Domain) (transactions.Domain, error) {
	user := Transaction{
		Name:     userR.Name,
		Email:    userR.Email,
		Transactionname: userR.Transactionname,
		Password: userR.Password,
		Address:  userR.Address,
		Status: userR.Status,
		Token: userR.Token,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (rep *TransactionRepo) Update(ctx context.Context, userU transactions.Domain) (transactions.Domain, error) {
	data := FromDomain(userU)
	err := rep.DB.Table("users").First(&data)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	data.Name = userU.Name
	data.Transaction = userU.Transaction
	data.Password = userU.Password
	data.Status = userU.Status
	data.Email = userU.Email
	data.Address = userU.Address
	

	if rep.DB.Save(&data).Error != nil {
		return transactions.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *TransactionRepo) GetAll(ctx context.Context) ([]transactions.Domain, error) {
	var data []Transaction
	err := rep.DB.Table("users").Find(&data)
	if err.Error != nil {
		return []transactions.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *TransactionRepo) GetById(ctx context.Context, id int) (transactions.Domain, error) {
	var data Transaction
	err := rep.DB.Table("users").Find(&data, "id=?", id)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}


func (rep *TransactionRepo) GetByToken(ctx context.Context, token string) (transactions.Domain, error) {
	var data Transaction
	err := rep.DB.Table("users").Find(&data, "token=?", token)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *TransactionRepo) Delete(ctx context.Context, id int) error {
	user := Transaction{}
	err := rep.DB.Table("users").Where("id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}