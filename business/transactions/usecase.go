package transactions

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/point"
)

type TransactionUsecase struct {
	PointRepo      point.Repository
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewTransactionUsecase(repo Repository, pointRepo point.Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &TransactionUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		PointRepo:      pointRepo,
		contextTimeout: timeout,
	}
}

func (tc *TransactionUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.TransactionAttachment == "" {
		return Domain{}, errors.New("attachment empty")
	}

	if domain.Transaction == 0 {
		return Domain{}, errors.New("transaction empty")
	}

	if domain.TransactionDate.String() == "" {
		return Domain{}, errors.New("date empty")
	}
	var err error

	if err != nil {
		return Domain{}, err
	}

	transaction, err := tc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tc *TransactionUsecase) Delete(ctx context.Context, id int) error {
	err := tc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (tc *TransactionUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	transaction, err := tc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return transaction, nil
}

func (tc *TransactionUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	transaction, err := tc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if transaction.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return transaction, nil
}

func (tc *TransactionUsecase) GetByUserId(ctx context.Context, id int) ([]Domain, error) {
	transaction, err := tc.Repo.GetByUserId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return transaction, nil
}

// func (tc *TransactionUsecase) GetByUserIdAndStatus(ctx context.Context, id int, status int) ([]Domain, error) {
// 	transaction, err := tc.Repo.GetByUserIdAndStatus(ctx, id, status)
// 	if err != nil {
// 		return []Domain{}, err
// 	}
// 	if transaction.Id == 0 {
// 		return []Domain{}, errors.New("ID NOT FOUND")
// 	}
// 	return transaction, nil
// }

func (tc *TransactionUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	var point point.Domain
	domain.Id = id
	transaction, err := tc.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	if domain.Status == 2 {
		points, _ := tc.PointRepo.GetByUserId(ctx, id)
		point.UserId = domain.UserId
		point.Point = points.Point + domain.Point
		pointU, _ := tc.PointRepo.Update(ctx, point)
		fmt.Println(pointU)
	}

	return transaction, nil
}
