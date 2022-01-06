package redeem

import (
	"context"
	"errors"
	"time"
	"github.com/daffashafwan/pointcuan/app/middlewares"
)

type RedeemUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewRedeemUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &RedeemUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (tc *RedeemUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.DataRedeem == "" {
		return Domain{}, errors.New("data empty")
	}
	var err error

	if err != nil {
		return Domain{}, err
	}
	
	redeem, err := tc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return redeem, nil
}

func (tc *RedeemUsecase) Delete(ctx context.Context, id int) error {
	err := tc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (tc *RedeemUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	redeem, err := tc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return redeem, nil
}

func (tc *RedeemUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	redeem, err := tc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if redeem.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return redeem, nil
}

func (tc *RedeemUsecase) GetByUserId(ctx context.Context, id int) ([]Domain, error) {
	redeem, err := tc.Repo.GetByUserId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return redeem, nil
}

func (tc *RedeemUsecase) GetByItemId(ctx context.Context, id int) ([]Domain, error) {
	redeem, err := tc.Repo.GetByItemId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return redeem, nil
}

