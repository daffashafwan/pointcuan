package pcrcrud

import (
	"context"
	"errors"
	"time"
)

type PcrUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewPcrcase(repo Repository, timeout time.Duration) Usecase {
	return &PcrUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *PcrUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	pcr, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return pcr, nil
}

func (usecase *PcrUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	pcr, err := usecase.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if pcr.Id == 0 && pcr.Id > 1 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return pcr, nil
}
