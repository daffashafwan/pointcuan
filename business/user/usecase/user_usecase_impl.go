
  
package usecase

import (
	"github.com/daffashafwan/pointcuan/model"
	"github.com/daffashafwan/pointcuan/business/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) Create(user *model.User) (*model.User, error) {
	return e.userRepo.Create(user)
}

func (e *UserUsecaseImpl) ReadAll() (*[]model.User, error) {
	return e.userRepo.ReadAll()
}

func (e *UserUsecaseImpl) ReadById(id int)(*model.User, error) {
	return e.userRepo.ReadById(id)
}

func (e *UserUsecaseImpl) ReadByUsername(username string)(*model.User, error) {
	return e.userRepo.ReadByUsername(username)
}

func (e *UserUsecaseImpl) Update(id int, user *model.User) (*model.User, error) {
	return e.userRepo.Update(id, user)
}

func (e *UserUsecaseImpl) Delete(id int) error {
	return e.userRepo.Delete(id)
}