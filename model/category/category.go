package category

import (
	"github.com/daffashafwan/pointcuan/business/categoryItems"
	"time"
)

type Category struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Category) ToDomain() categoryItems.Domain {
	return categoryItems.Domain{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}


func ToListDomain(data []Category) (result []categoryItems.Domain) {
	result = []categoryItems.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain categoryItems.Domain) Category {
	return Category{
		Id:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
