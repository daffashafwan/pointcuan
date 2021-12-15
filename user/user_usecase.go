package user

import "github.com/daffashafwan/pointcuan/model"

type UserUsecase interface {
	Create(user *model.User) (*model.User, error)
	ReadAll() (*[]model.User, error)
	ReadById(id int)(*model.User, error)
	ReadByUsername(username string)(*model.User, error)
	Update(id int, user *model.User) (*model.User, error)
	Delete(id int) error
}