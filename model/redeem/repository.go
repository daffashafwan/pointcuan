package redeem

import (
	"context"
	"errors"
	"github.com/daffashafwan/pointcuan/business/redeem"
	"gorm.io/gorm"
)

type RedeemRepo struct {
	DB *gorm.DB
}

func CreateRedeemRepo(conn *gorm.DB) redeem.Repository {
	return &RedeemRepo{
		DB: conn,
	}
}

func (rep *RedeemRepo) Create(ctx context.Context, redR *redeem.Domain) (redeem.Domain, error) {
	user := Redeem{
		UserId:     redR.UserId,
		ItemId:     redR.ItemId,
		DataRedeem: redR.DataRedeem,
		Point:      redR.Point,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		return redeem.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (rep *RedeemRepo) GetAll(ctx context.Context) ([]redeem.Domain, error) {
	var data []Redeem
	err := rep.DB.Table("redeems").Preload("Item").Find(&data)
	if err.Error != nil {
		return []redeem.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *RedeemRepo) GetById(ctx context.Context, id int) (redeem.Domain, error) {
	var data Redeem
	err := rep.DB.Table("redeems").Preload("Item").Find(&data, "id=?", id)
	if err.Error != nil {
		return redeem.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *RedeemRepo) GetByUserId(ctx context.Context, id int) ([]redeem.Domain, error) {
	var data []Redeem
	err := rep.DB.Table("redeems").Preload("Item").Find(&data, "user_id=?", id)
	if err.Error != nil {
		return []redeem.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *RedeemRepo) GetByItemId(ctx context.Context, id int) ([]redeem.Domain, error) {
	var data []Redeem
	err := rep.DB.Table("redeems").Preload("Item").Find(&data, "item_id=?", id)
	if err.Error != nil {
		return []redeem.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *RedeemRepo) Delete(ctx context.Context, id int) error {
	user := Redeem{}
	err := rep.DB.Table("redeems").Where("id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
