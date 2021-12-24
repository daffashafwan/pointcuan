package pcrcrud

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	NilaiPcr int
	UpdatedAt time.Time 
}

type Usecase interface {
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}
