package category

import (
	"context"
	"errors"
	"github.com/daffashafwan/pointcuan/business/categoryItems"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func CreateCategoryRepo(conn *gorm.DB) categoryItems.Repository {
	return &CategoryRepo{
		DB: conn,
	}
}


func (rep *CategoryRepo) Create(ctx context.Context,catC *categoryItems.Domain) (categoryItems.Domain, error) {
	user := Category{
		Name:     catC.Name,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		return categoryItems.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (rep *CategoryRepo) Update(ctx context.Context, userU categoryItems.Domain) (categoryItems.Domain, error) {
	data := FromDomain(userU)
	err := rep.DB.Table("categories").First(&data)
	if err.Error != nil {
		return categoryItems.Domain{}, err.Error
	}
	data.Name = userU.Name
	

	if rep.DB.Save(&data).Error != nil {
		return categoryItems.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *CategoryRepo) GetAll(ctx context.Context) ([]categoryItems.Domain, error) {
	var data []Category
	err := rep.DB.Table("categories").Find(&data)
	if err.Error != nil {
		return []categoryItems.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *CategoryRepo) GetById(ctx context.Context, id int) (categoryItems.Domain, error) {
	var data Category
	err := rep.DB.Table("categories").Find(&data, "id=?", id)
	if err.Error != nil {
		return categoryItems.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}


func (rep *CategoryRepo) Delete(ctx context.Context, id int) error {
	user := Category{}
	err := rep.DB.Table("categories").Where("id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
