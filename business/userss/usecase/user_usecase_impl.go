
  
package usecase

import (
	"github.com/daffashafwan/pointcuan/business/domain"
	"github.com/daffashafwan/pointcuan/business/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) Create(user *domain.Domain) (*domain.Domain, error) {
	return e.userRepo.Create(user)
}

func (e *UserUsecaseImpl) ReadAll() (*[]domain.Domain, error) {
	return e.userRepo.ReadAll()
}

func (e *UserUsecaseImpl) ReadById(id int)(*domain.Domain, error) {
	return e.userRepo.ReadById(id)
}

func (e *UserUsecaseImpl) ReadByUsername(username string)(*domain.Domain, error) {
	return e.userRepo.ReadByUsername(username)
}

func (e *UserUsecaseImpl) ReadByToken(token string)(*domain.Domain, error) {
	return e.userRepo.ReadByUsername(token)
}

func (e *UserUsecaseImpl) Update(id int, user *domain.Domain) (*domain.Domain, error) {
	return e.userRepo.Update(id, user)
}

func (e *UserUsecaseImpl) Delete(id int) error {
	return e.userRepo.Delete(id)
}