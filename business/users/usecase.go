package users

import (
	"context"
	"fmt"
	"time" 
	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/helpers/email"
	"github.com/daffashafwan/pointcuan/helpers/encrypt"
	"github.com/daffashafwan/pointcuan/helpers/randomizer"
	errors "github.com/daffashafwan/pointcuan/helpers/errors"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisinis login
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.ErrUsernamePasswordNotFound
	}

	if domain.Password == "" {
		return Domain{}, errors.ErrUsernamePasswordNotFound
	}

	user, err := uc.Repo.Login(ctx, domain.Username, domain.Password)

	if err != nil {
		return Domain{}, err
	}
	user.JWTToken, err = uc.ConfigJWT.GenerateTokenJWT(user.Id, 0)

	if err != nil {
		return Domain{}, err
	}

	return user,  nil
}

func (uc *UserUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.ErrPasswordRequired
	}

	if domain.Password == "" {
		return Domain{}, errors.ErrPasswordRequired
	}
	var err error

	if err != nil {
		return Domain{}, err
	}
	var hashed string
	hashed,_ = encrypt.Encrypt(domain.Password)
	domain.Password = hashed
	domain.Token = randomizer.Randomize()
	usern, _ := uc.Repo.GetByUsername(ctx, domain.Username)
	if usern.Id != 0 {
		return Domain{}, errors.ErrUsernameAlreadyExisted
	}
	usere, _ := uc.Repo.GetByEmail(ctx, domain.Email)
	if usere.Id != 0 {
		return Domain{}, errors.ErrEmailHasBeenRegister
	}
	user, err := uc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	email.SendEmail(ctx, domain.Email, "Verifikasi Email Pointcuan", "<a href=`http://localhost:1323/users/verify/"+domain.Token+"`>Link Verifikasi</a>")

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
		return Domain{}, errors.ErrIDNotFound
	}
	return user, nil
}

func (uc *UserUsecase) GetByToken(ctx context.Context, token string) (Domain, error) {
	user, err := uc.Repo.GetByToken(ctx, token)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, errors.ErrIDNotFound
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

func (uc *UserUsecase) Verify(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Status = "1"
	user, err := uc.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) ForgotPassword(ctx context.Context, emails string) (Domain, error) {
	users, errs := uc.Repo.GetByEmail(ctx, emails)
	if errs != nil {
		return Domain{}, errs
	}
	users.Token = randomizer.Randomize()
	fmt.Println(users)
	user, err := uc.Repo.Update(ctx, users)
	if err != nil {
		return Domain{}, err
	}
	email.SendEmail(ctx, users.Email, "Verifikasi Email Pointcuan", "<a href=`http://localhost:1323/users/forgotpassword/"+user.Token+"`>Link Verifikasi</a>")

	return user, nil
}

func (uc *UserUsecase) ResetPassword(ctx context.Context, password string,retypePassword string, id int) (Domain, error) {
	if password != retypePassword {
		return Domain{}, errors.ErrPasswordDidntMatch
	}
	users, errs := uc.Repo.GetById(ctx, id)
	if errs != nil {
		return Domain{}, errs
	}
	users.Token = randomizer.Randomize()
	users.Password,_ = encrypt.Encrypt(password)
	fmt.Println(users)
	user, err := uc.Repo.Update(ctx, users)
	if err != nil {
		return Domain{}, err
	}
	email.SendEmail(ctx, users.Email, "Verifikasi Email Pointcuan", "<a href=`http://localhost:1323/users/forgotpassword/"+user.Token+"`>Link Verifikasi</a>")

	return user, nil
}
