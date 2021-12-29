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

func (uc *PcrUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {
	var pcr Domain
	var err error
	data,_ := uc.GetPCR(ctx)
	if data.Id == 0 {
		pcr, err = uc.Repo.Create(ctx, &domain)
		if err != nil {
			return Domain{}, err
		}
	}else{
		domain.Id = data.Id
		pcr, err = uc.Repo.Update(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
	}

	return pcr, nil
}

func (usecase *PcrUsecase) GetPCR(ctx context.Context) (Domain, error) {
	pcr, err := usecase.Repo.GetPCR(ctx)
	if err != nil {
		return Domain{}, err
	}
	if pcr.Id == 0 && pcr.Id > 1 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return pcr, nil
}
