package pcrcrud

import (
	"context"
	"errors"

	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
	"gorm.io/gorm"
)

type PcrRepo struct {
	DB *gorm.DB
}

func CreatePcrRepo(gormDb *gorm.DB) pcrcrud.Repository {
	return &PcrRepo{
		DB: gormDb,
	}
}

func (rep *PcrRepo) Update(ctx context.Context, pcr pcrcrud.Domain) (pcrcrud.Domain, error) {
	data := FromDomain(pcr)
	err := rep.DB.Table("pcrcruds").First(&data)
	if err.Error != nil {
		return pcrcrud.Domain{}, err.Error
	}
	data.PcrValue = pcr.PcrValue

	if rep.DB.Save(&data).Error != nil {
		return pcrcrud.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (repo *PcrRepo) GetById(ctx context.Context, id int) (pcrcrud.Domain, error) {
	var data Pcrcrud
	err := repo.DB.Table("pcrcruds").Find(&data, "id=?", id)
	if err.Error != nil {
		return pcrcrud.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}