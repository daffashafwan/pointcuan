package point

import (
	"context"
	"github.com/daffashafwan/pointcuan/business/point"
	"gorm.io/gorm"
	"fmt"
	"errors"
)

type PointRepo struct {
	DB *gorm.DB
}

func CreatePointRepo(conn *gorm.DB) point.Repository {
	return &PointRepo{
		DB: conn,
	}
}


func (rep *PointRepo) Create(ctx context.Context,pointC *point.Domain) (point.Domain, error) {
	points := Point{
		UserId: pointC.UserId,
		Point: pointC.Point,
	}
	err := rep.DB.Create(&points)
	if err.Error != nil {
		fmt.Printf("[PointRepoImpl.Create] error execute query %v \n", err)
		return point.Domain{}, fmt.Errorf("failed insert data")
	}
	return points.ToDomain(), nil
}

func (rep *PointRepo) Update(ctx context.Context, pointU point.Domain) (point.Domain, error) {
	data := FromDomain(pointU)
	err := rep.DB.Table("points").First(&data)
	if err.Error != nil {
		return point.Domain{}, err.Error
	}
	data.UserId = pointU.UserId
	data.Point = pointU.Point
	if rep.DB.Save(&data).Error != nil {
		return point.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *PointRepo) GetAll(ctx context.Context) ([]point.Domain, error) {
	var data []Point
	err := rep.DB.Table("points").Find(&data)
	if err.Error != nil {
		return []point.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *PointRepo) GetByUserId(ctx context.Context, id int) (point.Domain, error) {
	var data Point
	err := rep.DB.Table("point").Find(&data, "user_id=?", id)
	if err.Error != nil {
		return point.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *PointRepo) Delete(ctx context.Context, id int) error {
	user := Point{}
	err := rep.DB.Table("points]").Where("user_id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
