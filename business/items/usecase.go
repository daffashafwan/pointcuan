package items

import (
	"context"
	"errors"
	"time"
)

type ItemsUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewItemsUsecase(repo Repository, timeout time.Duration) Usecase {
	return &ItemsUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}


func (ic *ItemsUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	items, err := ic.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return items, nil
}

func (ic *ItemsUsecase) Delete(ctx context.Context, id int) ( error) {
	err := ic.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ic *ItemsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	items, err := ic.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return items, nil
}

func (ic *ItemsUsecase) GetByItemId(ctx context.Context, id int) (Domain, error) {
	items, err := ic.Repo.GetByItemId(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if items.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return items, nil
}

func (ic *ItemsUsecase) GetByCategoryId(ctx context.Context, id int) ([]Domain, error) {
	items, err := ic.Repo.GetByCategoryId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return items, nil
}

func (ic *ItemsUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	items, err := ic.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return items, nil
}

func (ic *ItemsUsecase) UpdateStock(ctx context.Context, domain Domain, id int) (Domain, error) {
	item, errs := ic.Repo.GetByItemId(ctx, id)
	if errs != nil {
		return Domain{}, errs
	}
	item.Id = id
	item.Stock = domain.Stock
	items, err := ic.Repo.UpdateStock(ctx, item)
	if err != nil {
		return Domain{}, err
	}

	return items, nil
}

