package user

import "github.com/daffashafwan/pointcuan/business/user/domain"

type UserRepo interface {
	Create(user *domain.Domain) (*domain.Domain, error)
	ReadAll() (*[]domain.Domain, error)
	ReadById(id int)(*domain.Domain, error)
	ReadByUsername(username string)(*domain.Domain, error)
	ReadByToken(token string)(*domain.Domain, error)
	Update(id int, user *domain.Domain) (*domain.Domain, error)
	Delete(id int) error
}