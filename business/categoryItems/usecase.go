package categoryItems

import (
	"context"
	"errors"
	"time"

	"github.com/daffashafwan/pointcuan/app/middlewares"
)

type CategoryItemsUsecase struct {
	Repo           	Repository
	contextTimeout 	time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewCategoryUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &CategoryItemsUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (pc *CategoryItemsUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	category, err := pc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return category, nil
}

func (uc *CategoryItemsUsecase) Delete(ctx context.Context, id int) ( error) {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CategoryItemsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	categoryitem, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return categoryitem, nil
}

func (uc *CategoryItemsUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	categoryitem, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if categoryitem.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return categoryitem, nil
}




func (uc *CategoryItemsUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	categoryItem, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return categoryItem, nil
}


