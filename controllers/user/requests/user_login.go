package requests

import (

	"github.com/daffashafwan/pointcuan/business/users"
)

type UserLogin struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomain() users.Domain {
	return users.Domain{
		Username:    user.Username,
		Password: user.Password,
	}
}
