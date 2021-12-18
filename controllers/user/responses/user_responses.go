package responses

import (
	"github.com/daffashafwan/pointcuan/business/users"
	"time"
)

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Address   string    `json:"address"`
	JWTToken  string    `json:"jwt_token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		JWTToken:  domain.JWTToken,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
