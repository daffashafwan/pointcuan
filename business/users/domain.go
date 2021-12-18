package users

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	Name string 
	Username string 
	Password string 
	Email string 
	Address string 
	Status string 
	Token string 
	JWTToken string
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, username string, password string) (Domain, error)
}
