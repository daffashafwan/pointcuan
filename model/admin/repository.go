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


func (repo AdminRepo) Login(domain admin.Domain, ctx context.Context) (admin.Domain, error){
	adminDb := FromDomain(domain)

	err := repo.DB.Where("username = ? AND password = ?", adminDb.Username, adminDb.Password).First(&adminDb).Error
	if err != nil {
		return admin.Domain{}, err
	}
	return adminDb.ToDomain(), nil
}

func (repo *AdminRepo) GetById(ctx context.Context, id int) (admin.Domain, error) {
	var data Admin
	err := repo.DB.Table("admins").Find(&data, "id=?", id)
	if err.Error != nil {
		return admin.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}