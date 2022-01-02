package categoryItems

import (
	"context"
	"errors"
	"time"

	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/helpers/encrypt"
)

type CategotyItemsUsecase struct {
	Repo           	Repository
	contextTimeout 	time.Duration
	ConfigJWT		middlewares.ConfigJWT
}



func (uc *CategoryItemUsecase) Delete(ctx context.Context, id int) ( error) {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CategoryItemUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	categoryitem, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return categoryItem, nil
}

func (uc *CategoryItemUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	categoryitem, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if categoryitem.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return categoryItem, nil
}




func (uc *CategoryItemUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Password, _ = encrypt.Encrypt(domain.Password)
	domain.Id = id
	categoryItem, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return categoryItem, nil
}


