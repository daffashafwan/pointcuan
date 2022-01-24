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
	JWTToken  string    `json:"jwtToken"`
	Token     string    `json:"token"`
	Status    string    `json:"status"`
	Point     float64   `json:"point"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Username:  domain.Username,
		Email:     domain.Email,
		Address:   domain.Address,
		JWTToken:  domain.JWTToken,
		Token:     domain.Token,
		Status:    domain.Status,
		Point:     domain.Point,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []users.Domain) []UserResponse {
	var list []UserResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
