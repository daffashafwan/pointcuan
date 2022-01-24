package FAQ

import (
	"context"
	"time"

	"github.com/daffashafwan/pointcuan/app/middlewares"
)

type FAQUsecase struct {
	Repo           	Repository
	contextTimeout 	time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewFAQUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &FAQUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (pc *FAQUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	faq, err := pc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return faq, nil
}

func (uc *FAQUsecase) Delete(ctx context.Context, id int) ( error) {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *FAQUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	faq, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return faq, nil
}

func (uc *FAQUsecase) GetActive(ctx context.Context) ([]Domain, error) {
	faq, err := uc.Repo.GetActive(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return faq, nil
}




func (uc *FAQUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	faq, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return faq, nil
}


