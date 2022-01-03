package categoryItems

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	Name string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, id int) ( error)
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, id int) (error)
}