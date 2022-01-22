package FAQ

import (
	"context"
	"errors"
	"github.com/daffashafwan/pointcuan/business/FAQ"
	"gorm.io/gorm"
)

type FAQRepo struct {
	DB *gorm.DB
}

func CreateFAQRepo(conn *gorm.DB) FAQ.Repository {
	return &FAQRepo{
		DB: conn,
	}
}

func (rep *FAQRepo) Create(ctx context.Context, catC *FAQ.Domain) (FAQ.Domain, error) {
	user := Faq{
		Question: catC.Question,
		Answer: catC.Answer,
		Status: catC.Status,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		return FAQ.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (rep *FAQRepo) Update(ctx context.Context, userU FAQ.Domain) (FAQ.Domain, error) {
	data := FromDomain(userU)
	err := rep.DB.Table("faqs").First(&data)
	if err.Error != nil {
		return FAQ.Domain{}, err.Error
	}
	data.Question = userU.Question
	data.Answer = userU.Answer
	data.Status = userU.Status

	if rep.DB.Save(&data).Error != nil {
		return FAQ.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *FAQRepo) GetAll(ctx context.Context) ([]FAQ.Domain, error) {
	var data []Faq
	err := rep.DB.Table("faqs").Find(&data)
	if err.Error != nil {
		return []FAQ.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *FAQRepo) GetActive(ctx context.Context) ([]FAQ.Domain, error) {
	var data []Faq
	err := rep.DB.Table("faqs").Find(&data, "status=?", 1)
	if err.Error != nil {
		return []FAQ.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *FAQRepo) Delete(ctx context.Context, id int) error {
	user := Faq{}
	err := rep.DB.Table("faqs").Where("id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
