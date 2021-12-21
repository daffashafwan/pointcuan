package admin

import (
	"context"
	"errors"
	"time"
)

type AdminUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUsecase(repo Repository, timeout time.Duration) Usecase {
	return &AdminUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (usecase *AdminUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.New("username empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	var err error

	if err != nil {
		return Domain{}, err
	}

	admin, err := usecase.Repo.Login(ctx, domain.Username, domain.Password)

	if err != nil {
		return Domain{}, err
	}
	return admin, nil
}

func (usecase *AdminUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	admin, err := usecase.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if admin.Id == 0 && admin.Id > 1 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return admin, nil
}
