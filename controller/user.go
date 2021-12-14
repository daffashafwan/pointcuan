package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/daffashafwan/pointcuan/model"
	"github.com/daffashafwan/pointcuan/storage"
)

func GetUser(c echo.Context) error {
	user, _ := GetRepoUser()
	return c.JSON(http.StatusOK, user)
}

func GetRepoUser() ([]model.User, error) {
	db := storage.GetDBInstance()
	user := []model.User{}
	db.AutoMigrate(&user)

	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}