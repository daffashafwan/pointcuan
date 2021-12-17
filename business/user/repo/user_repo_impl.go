package repo

import (
	"fmt"
	"github.com/daffashafwan/pointcuan/model"
	"github.com/daffashafwan/pointcuan/business/user"

	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func CreateUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

func (e *UserRepoImpl) Create(user *model.User) (*model.User, error) {
	err := e.DB.Save(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return user, nil
}

func (e *UserRepoImpl) ReadAll() (*[]model.User, error) {
	var users []model.User
	err := e.DB.Find(&users).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &users, nil
}

func (e *UserRepoImpl) ReadById(id int)(*model.User, error) {
	var user = model.User{}
	err := e.DB.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &user, nil
}

func (e *UserRepoImpl) ReadByUsername(username string)(*model.User, error) {
	var user = model.User{}
	err := e.DB.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("username is not exsis")
	}
	return &user, nil
}

func (e *UserRepoImpl) Update(id int, user *model.User) (*model.User, error) {
	var upUser = model.User{}
	err := e.DB.Table("users").Where("id = ?", id).First(&upUser).Update(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.Update] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &upUser, nil
}

func (e *UserRepoImpl) Delete(id int) error {
	var user = model.User{}
	err := e.DB.Table("users").Where("id = ?", id).First(&user).Delete(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}
	return nil
}