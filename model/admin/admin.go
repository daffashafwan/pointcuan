package admin

import (
	"time"

	"github.com/daffashafwan/pointcuan/business/admin"
)

type Admin struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ad *Admin) ToDomain() admin.Domain {
	return admin.Domain {
		Id:        ad.Id,
		Name:      ad.Name,
		Username:  ad.Username,
		Password:  ad.Password,
		CreatedAt: ad.CreatedAt,
		UpdatedAt: ad.UpdatedAt,
	}
}

func FromDomain(domain admin.Domain) Admin {
	return Admin{
		Id:        domain.Id,
		Name:      domain.Name,
		Password:  domain.Password,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
