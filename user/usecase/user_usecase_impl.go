
  
package usecase

import (
	"github.com/daffashafwan/pointcuan/model"
	"github.com/daffashafwan/pointcuan/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) Create(person *model.User) (*model.User, error) {
	return e.userRepo.Create(person)
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

func (e *UserUsecaseImpl) Update(id int, person *model.User) (*model.User, error) {
	return e.userRepo.Update(id, person)
}

func (e *UserUsecaseImpl) Delete(id int) error {
	return e.userRepo.Delete(id)
}