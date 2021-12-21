package admin

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	Name string 
	Username string 
	Password string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, username string, password string) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}
