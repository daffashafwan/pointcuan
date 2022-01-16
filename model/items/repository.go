package items

import (
	"context"
	"errors"

	"github.com/daffashafwan/pointcuan/business/items"
	"gorm.io/gorm"
)

type ItemsRepo struct {
	DB *gorm.DB
}

func CreateItemRepo(conn *gorm.DB) items.Repository {
	return &ItemsRepo{
		DB: conn,
	}
}

func (rep *ItemsRepo) Create(ctx context.Context, itemCreate *items.Domain) (items.Domain, error) {
	item := Items{
		CategoryId:  itemCreate.CategoryId,
		Name:        itemCreate.Name,
		PointRedeem: itemCreate.PointRedeem,
		Stock:       itemCreate.Stock,
	}
	err := rep.DB.Create(&item)
	if err.Error != nil {
		return items.Domain{}, err.Error
	}
	return item.ToDomain(), nil
}

func (repo *ItemsRepo) Update(ctx context.Context, itemUpdate items.Domain) (items.Domain, error) {
	data := FromDomain(itemUpdate)
	err := repo.DB.Table("items").First(&data)
	if err.Error != nil {
		return items.Domain{}, err.Error
	}
	data.CategoryId = itemUpdate.CategoryId
	data.PointRedeem = itemUpdate.PointRedeem
	data.Name = itemUpdate.Name
	data.Stock = itemUpdate.Stock
	if repo.DB.Save(&data).Error != nil {
		return items.Domain{}, errors.New("bad requests")
	}

	return data.ToDomain(), nil
}

func (repo *ItemsRepo) UpdateStock(ctx context.Context, itemUpdate items.Domain) (items.Domain, error) {
	data := FromDomain(itemUpdate)
	err := repo.DB.Table("items").First(&data)
	if err.Error != nil {
		return items.Domain{}, err.Error
	}
	data.Stock = itemUpdate.Stock
	if repo.DB.Save(&data).Error != nil {
		return items.Domain{}, errors.New("bad requests")
	}

	return data.ToDomain(), nil
}

func (repo *ItemsRepo) GetAll(ctx context.Context) ([]items.Domain, error) {
	var data []Items
	err := repo.DB.Table("items").Preload("Category").Find(&data)
	if err.Error != nil {
		return []items.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *ItemsRepo) GetByItemId(ctx context.Context, id int) (items.Domain, error) {
	var data Items
	err := repo.DB.Table("items").Find(&data, "id=?", id)
	if err.Error != nil {
		return items.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *ItemsRepo) GetByCategoryId(ctx context.Context, id int) ([]items.Domain, error) {
	var data []Items
	err := repo.DB.Table("items").Find(&data, "category_id=?", id)
	if err.Error != nil {
		return []items.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *ItemsRepo) Delete(ctx context.Context, id int) error {
	item := Items{}
	err := repo.DB.Table("items").Where("id = ?", id).First(&item).Delete(&item)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
