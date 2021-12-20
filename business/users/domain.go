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
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	//Delete(ctx context.Context, domain Domain) (Domain, error)
	//GetAll(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, username string, password string) (Domain, error)
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	//Delete(ctx context.Context, id int) (Domain, error)
	//GetAll(ctx context.Context) (Domain, error)
	//GetById(ctx context.Context, id int) (Domain, error)
	//Update(ctx context.Context, id int) (Domain, error)
}
