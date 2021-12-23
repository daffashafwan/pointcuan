package point

import (
	"context"
	"errors"
	"time"

	"github.com/daffashafwan/pointcuan/app/middlewares"
)

type PointUsecase struct {
	// ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewPointUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &PointUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}


func (pc *PointUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	point, err := pc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return point, nil
}

func (pc *PointUsecase) Delete(ctx context.Context, id int) ( error) {
	err := pc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (pc *PointUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	point, err := pc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return point, nil
}

func (pc *PointUsecase) GetByUserId(ctx context.Context, id int) (Domain, error) {
	point, err := pc.Repo.GetByUserId(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if point.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return point, nil
}

func (pc *PointUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	point, err := pc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return point, nil
}

