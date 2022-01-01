package requests

import "github.com/daffashafwan/pointcuan/business/users"

type UserResetPassword struct {
	Password    string `json:"password"`
	RetypePassword    string `json:"retypePassword"`
}

func (ur *UserResetPassword) ToDomain() users.Domain {
	return users.Domain{
		Password:    ur.Password,
	}
}
