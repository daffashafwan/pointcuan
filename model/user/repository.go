package user

import (
	"context"
	"github.com/daffashafwan/pointcuan/business/users"
	"github.com/daffashafwan/pointcuan/helpers/encrypt"
	"gorm.io/gorm"
	"fmt"
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
	result := rep.DB.Table("users").Where("username = ?", username).Where("status = ? ", "1").First(&user).Error

	if result != nil {
		return users.Domain{}, result
	}
	if !(encrypt.Compare(password,user.Password)) {
		return users.Domain{}, errors.New("username tidak cocok")
	}
	return user.ToDomain(), nil

}

func (rep *UserRepo) Create(ctx context.Context,userR *users.Domain) (users.Domain, error) {
	user := User{
		Name:     userR.Name,
		Email:    userR.Email,
		Username: userR.Username,
		Password: userR.Password,
		Address:  userR.Address,
		Status: userR.Status,
		Token: userR.Token,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		fmt.Printf("[UserRepoImpl.Create] error execute query %v \n", err)
		return users.Domain{}, fmt.Errorf("failed insert data")
	}
	fmt.Println(user.Username)
	fmt.Println(user.ToDomain())
	return user.ToDomain(), nil
}

func (rep *UserRepo) Update(ctx context.Context, userU users.Domain) (users.Domain, error) {
	data := FromDomain(userU)
	err := rep.DB.Table("users").First(&data)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	data.Name = userU.Name
	data.Username = userU.Username
	data.Password = userU.Password
	data.Status = userU.Status
	data.Email = userU.Email
	data.Address = userU.Address
	

	if rep.DB.Save(&data).Error != nil {
		return users.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *UserRepo) GetAll(ctx context.Context) ([]users.Domain, error) {
	var data []User
	err := rep.DB.Table("users").Find(&data)
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *UserRepo) GetById(ctx context.Context, id int) (users.Domain, error) {
	var data User
	err := rep.DB.Table("users").Find(&data, "id=?", id)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *UserRepo) GetByToken(ctx context.Context, token string) (users.Domain, error) {
	var data User
	err := rep.DB.Table("users").Find(&data, "token=?", token)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *UserRepo) Delete(ctx context.Context, id int) error {
	user := User{}
	err := rep.DB.Table("users").Where("id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

func (rep *UserRepo) ForgotPassword(ctx context.Context,userR *users.Domain) (users.Domain, error) {
	user := User{
		Token: userR.Token,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		fmt.Printf("[UserRepoImpl.Create] error execute query %v \n", err)
		return users.Domain{}, fmt.Errorf("failed insert data")
	}
	fmt.Println(user.Username)
	fmt.Println(user.ToDomain())
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