package admins

import (
	"context"

	"github.com/daffashafwan/pointcuan/business/admin/admins"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func CreateAdminRepo(gormDb *gorm.DB) admins.Repository {
	return &AdminRepo{
		DB: gormDb,
	}
}


func (repo AdminRepo) Login(domain admins.Domain, ctx context.Context) (admins.Domain, error){
	adminDb := FromDomain(domain)

	err := repo.DB.Where("username = ? AND password = ?", adminDb.Username, adminDb.Password).First(&adminDb).Error
	if err != nil {
		return admins.Domain{}, err
	}
	return adminDb.ToDomain(), nil
}

func (repo *AdminRepo) GetById(ctx context.Context, id int) (admins.Domain, error) {
	var data Admin
	err := repo.DB.Table("admins").Find(&data, "id=?", id)
	if err.Error != nil {
		return admins.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}