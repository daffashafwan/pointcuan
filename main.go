package main

import (
	// "net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo"

	// "github.com/labstack/echo/v4/middleware"
	"github.com/daffashafwan/pointcuan/config"
	UserHandler "github.com/daffashafwan/pointcuan/user/handler"
	UserRepo "github.com/daffashafwan/pointcuan/user/repo"
	UserUsecase "github.com/daffashafwan/pointcuan/user/usecase"
)

func main() {
	// Echo instance
	e := echo.New()

	db := config.DbConnect()
	defer db.Close()

	UserRepo := UserRepo.CreateUserRepo(db)
	UserUsecase := UserUsecase.CreateUserUsecase(UserRepo)
	UserHandler.CreateUserHandler(e, UserUsecase)
	
	// Start server
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		fmt.Println("err")
	}
	ioutil.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":1323"))
}