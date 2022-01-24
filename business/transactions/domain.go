package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id                    int
	UserId                int
	User                  interface{}
	TransactionDate       time.Time
	Transaction           float64
	TransactionAttachment string
	Status                int8
	Point                 float64
	Description           string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type StatisticDomain struct {
	Percentage float64
	PointIn    float64
	PointOut   float64
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByUserId(ctx context.Context, id int) ([]Domain, error)
	GetByUserIdAndStatus(ctx context.Context, id int, sid int) ([]Domain, error)
	//GetByUserIdAndStatus(ctx context.Context, id int, status int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByUserId(ctx context.Context, id int) ([]Domain, error)
	GetByUserIdAndStatus(ctx context.Context, id int, sid int) ([]Domain, error)
	//GetByUserIdAndStatus(ctx context.Context, id int, status int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}
