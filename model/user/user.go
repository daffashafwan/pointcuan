package user

import (
	"github.com/daffashafwan/pointcuan/business/users"
	"time"
)

type User struct {
	Id        int `gorm:"primaryKey]"`
	Name      string
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	Address   string
	Status    string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Status:    user.Status,
		Token:     user.Token,
		Address:   user.Address,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListDomain(data []User) (result []users.Domain) {
	result = []users.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Password:  domain.Password,
		Username:  domain.Username,
		Status:    domain.Status,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
