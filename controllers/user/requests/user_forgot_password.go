package requests

import "github.com/daffashafwan/pointcuan/business/users"

type UserForgotPassword struct {
	Email    string `json:"email"`
}

func (uf *UserForgotPassword) ToDomain() users.Domain {
	return users.Domain{
		Email:    uf.Email,
	}
}
