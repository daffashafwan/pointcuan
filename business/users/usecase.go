package users

import (
	"context"
	"errors"
	"time"

	"github.com/daffashafwan/pointcuan/helpers/encrypt"
	"github.com/daffashafwan/pointcuan/helpers/email"
	"github.com/daffashafwan/pointcuan/helpers/randomizer"
)

type UserUsecase struct {
	// ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisinis login
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.New("username empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	var err error

	if err != nil {
		return Domain{}, err
	}

	user, err := uc.Repo.Login(ctx, domain.Username, domain.Password)

	if err != nil {
		return Domain{}, err
	}

	// user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.New("username empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	var err error

	if err != nil {
		return Domain{}, err
	}
	var hashed string
	hashed,_ = encrypt.Encrypt(domain.Password)
	domain.Password = hashed
	domain.Token = randomizer.Randomize()
	user, err := uc.Repo.Create(ctx, &domain)

	if err != nil {
		return Domain{}, err
	}
	email.SendEmail(ctx, domain.Email, "Verifikasi Email Pointcuan", "<a href=`http://localhost:1323/user/verif/"+domain.Token+"`>Link Verifikasi</a>")

	return user, nil
}

func (uc *UserUsecase) Delete(ctx context.Context, id int) ( error) {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return user, nil
}


func (uc *UserUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Password, _ = encrypt.Encrypt(domain.Password)
	domain.Id = id
	user, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
