package user

import (
	"context"
	"github.com/daffashafwan/pointcuan/business/users"
	"github.com/daffashafwan/pointcuan/helpers/encrypt"
	"gorm.io/gorm"
	"errors"
)

type UserRepo struct {
	DB *gorm.DB
}

func CreateUserRepo(conn *gorm.DB) users.Repository {
	return &UserRepo{
		DB: conn,
	}
}


func (rep *UserRepo) Login(ctx context.Context, username string, password string) (users.Domain, error) {
	var user User
	result := rep.DB.Table("users").Where("username = ?", username).First(&user).Error

	if result != nil {
		return users.Domain{}, result
	}
	if !(encrypt.Compare(password,user.Password)) {
		return users.Domain{}, errors.New("username tidak cocok")
	}
	return user.ToDomain(), nil

}

// func (e *UserRepoImpl) Create(user *domain.Domain) (*domain.Domain, error) {
// 	err := e.DB.Save(&user).Error
// 	if err != nil {
// 		fmt.Printf("[UserRepoImpl.Create] error execute query %v \n", err)
// 		return nil, fmt.Errorf("failed insert data")
// 	}
// 	return user, nil
// }

// func (e *UserRepoImpl) ReadAll() (*[]domain.Domain, error) {
// 	var users []domain.Domain
// 	err := e.DB.Find(&users).Error
// 	if err != nil {
// 		fmt.Printf("[UserRepoImpl.ReadAll] error execute query %v \n", err)
// 		return nil, fmt.Errorf("failed view all data")
// 	}
// 	return &users, nil
// }

// func (e *UserRepoImpl) ReadById(id int)(*domain.Domain, error) {
// 	var user = domain.Domain{}
// 	err := e.DB.Table("users").Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		fmt.Printf("[UserRepoImpl.ReadById] error execute query %v \n", err)
// 		return nil, fmt.Errorf("id is not exsis")
// 	}
// 	return &user, nil
// }

// func (e *UserRepoImpl) ReadByUsername(username string)(*domain.Domain, error) {
// 	var user = domain.Domain{}
// 	err := e.DB.Table("users").Where("username = ?", username).Where("status", "1").First(&user).Error
// 	if err != nil {
// 		fmt.Printf("[UserRepoImpl.ReadById] error execute query %v \n", err)
// 		return nil, fmt.Errorf("username is not exist or is not activated")
// 	}
// 	return &user, nil
// }

// func (e *UserRepoImpl) ReadByToken(token string)(*domain.Domain, error) {
// 	var user = domain.Domain{}
// 	err := e.DB.Table("users").Where("token = ?", token).First(&user).Error
// 	if err != nil {
// 		fmt.Printf("[UserRepoImpl.ReadById] error execute query %v \n", err)
// 		return nil, fmt.Errorf("username is not exist or is not activated")
// 	}
// 	return &user, nil
// }

// func (e *UserRepoImpl) Update(id int, user *domain.Domain) (*domain.Domain, error) {
// 	var upUser = domain.Domain{}
// 	err := e.DB.Table("users").Where("id = ?", id).First(&upUser).Update(&user).Error
// 	if err != nil {
// 		fmt.Printf("[UserRepoImpl.Update] error execute query %v \n", err)
// 		return nil, fmt.Errorf("failed update data")
// 	}
// 	return &upUser, nil
// }

// func (e *UserRepoImpl) Delete(id int) error {
// 	var user = domain.Domain{}
// 	err := e.DB.Table("users").Where("id = ?", id).First(&user).Delete(&user).Error
// 	if err != nil {
// 		fmt.Printf("[UserRepoImpl.Delete] error execute query %v \n", err)
// 		return fmt.Errorf("id is not exsis")
// 	}
// 	return nil
// }