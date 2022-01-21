package redeem

import (
	"context"
	"time"
)

type Domain struct {
	Id         int
	UserId     int
	ItemId     int
	DataRedeem string
	Item       interface{}
	Status     int
	Point      int
	ResponseMidtrans interface{}
	RefId      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByUserId(ctx context.Context, id int) ([]Domain, error)
	GetByItemId(ctx context.Context, id int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByUserId(ctx context.Context, id int) ([]Domain, error)
	GetByItemId(ctx context.Context, id int) ([]Domain, error)
	//GetByUserIdAndStatus(ctx context.Context, id int, status int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}
