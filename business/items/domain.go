package items

import (
	"context"
	"time"
)

type Domain struct {
	Id          int
	CategoryId  int
	Category    interface{}
	Name        string
	PointRedeem int
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	UpdateStock(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByItemId(ctx context.Context, id int) (Domain, error)
	GetByCategoryId(ctx context.Context, id int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	UpdateStock(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByItemId(ctx context.Context, id int) (Domain, error)
	GetByCategoryId(ctx context.Context, id int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}
