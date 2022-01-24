package FAQ

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Question  string
	Answer    string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetActive(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetActive(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}
