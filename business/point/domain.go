package point

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	UserId int 
	Point float64 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByUserId(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, id int) ( error)
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByUserId(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, id int) ( error)
}
