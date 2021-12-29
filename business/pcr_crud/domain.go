package pcrcrud

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	PcrValue float64
	UpdatedAt time.Time 
}

type Usecase interface {
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetPCR(ctx context.Context) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetPCR(ctx context.Context) (Domain, error)
}
