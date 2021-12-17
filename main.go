package main

import (
	// "net/http

	// "github.com/labstack/echo/v4/middleware"
	"github.com/daffashafwan/pointcuan/config"
	"github.com/daffashafwan/pointcuan/route"
	UserHandler "github.com/daffashafwan/pointcuan/business/user/handler"
	UserRepo "github.com/daffashafwan/pointcuan/business/user/repo"
	UserUsecase "github.com/daffashafwan/pointcuan/business/user/usecase"
)

func main() {
	// Echo instance
	routes := route.Init()

	db := config.DbConnect()
	defer db.Close()

	UserRepo := UserRepo.CreateUserRepo(db)
	UserUsecase := UserUsecase.CreateUserUsecase(UserRepo)
	UserHandler.CreateUserHandler(routes, UserUsecase)
	
	// // Start server
	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	fmt.Println("err")
	// }
	// ioutil.WriteFile("routes.json", data, 0644)
	routes.Logger.Fatal(routes.Start(":1323"))
}