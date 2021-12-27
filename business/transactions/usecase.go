package transactions

import (
	"fmt"
	"context"
	"errors"
	"time"
	"github.com/daffashafwan/pointcuan/app/middlewares"
)

type TransactionUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewTransactionUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &TransactionUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (tc *TransactionUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	fmt.Println(domain.Transaction)
	fmt.Println(domain.UserId)
	fmt.Println(domain.TransactionAttachment)
	if domain.TransactionAttachment == "" {
		return Domain{}, errors.New("attachment empty")
	}

	if domain.Transaction == "" {
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

func (tc *TransactionUsecase) Delete(ctx context.Context, id int) ( error) {
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
	domain.Id = id
	transaction, err := tc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}
