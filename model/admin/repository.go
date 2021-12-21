package admin

import (
	"context"

	"github.com/daffashafwan/pointcuan/business/admin"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func CreateAdminRepo(gormDb *gorm.DB) admin.Repository {
	return &AdminRepo{
		DB: gormDb,
	}
}

func (repo *AdminRepo) Login(ctx context.Context, username string, password string) (admin.Domain, error) {
	var adm Admin
	result := repo.DB.Table("admins").Where("username = ?", username).First(&adm).Error

	if result != nil {
		return admin.Domain{}, result
	}
	return adm.ToDomain(), nil

}

func (repo *AdminRepo) GetById(ctx context.Context, id int) (admin.Domain, error) {
	var data Admin
	err := repo.DB.Table("admins").Find(&data, "id=?", id)
	if err.Error != nil {
		return admin.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}