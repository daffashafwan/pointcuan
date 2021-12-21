package requests

import admins "github.com/daffashafwan/pointcuan/business/admin"

type AdminLogin struct {
	Username    string `json:"username"`
	Password 	string `json:"password"`
}

func (al *AdminLogin) ToDomain() admins.Domain {
	return admins.Domain{
		Username:    al.Username,
		Password: al.Password,
	}
}
