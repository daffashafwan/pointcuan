package admins

import (
	"time"

	admins "github.com/daffashafwan/pointcuan/business/admin"
)

type Admin struct {
	Id        int `gorm:"primaryKey]"`
	Name      string
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (admin *Admin) ToDomain() admins.Domain {
	return admins.Domain {
		Id:        admin.Id,
		Name:      admin.Name,
		Username:  admin.Username,
		Password:  admin.Password,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func FromDomain(domain admins.Domain) Admin {
	return Admin{
		Id:        domain.Id,
		Name:      domain.Name,
		Password:  domain.Password,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
