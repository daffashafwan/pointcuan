package requests

import "github.com/daffashafwan/pointcuan/business/users"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Status   string `json:"status"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (ur *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		Username: ur.Username,
		Password: ur.Password,
		Email:    ur.Email,
		Address:  ur.Address,
		Token:    ur.Token,
		Name:     ur.Name,
		Status:   ur.Status,
	}
}
