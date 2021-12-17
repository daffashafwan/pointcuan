package main

import (
	// "net/http

	// "github.com/labstack/echo/v4/middleware"
	AdminHandler "github.com/daffashafwan/pointcuan/business/admin/handler"
	AdminRepo "github.com/daffashafwan/pointcuan/business/admin/repo"
	AdminUsecase "github.com/daffashafwan/pointcuan/business/admin/usecase"
	UserHandler "github.com/daffashafwan/pointcuan/business/user/handler"
	UserRepo "github.com/daffashafwan/pointcuan/business/user/repo"
	UserUsecase "github.com/daffashafwan/pointcuan/business/user/usecase"
	"github.com/daffashafwan/pointcuan/config"
	"github.com/daffashafwan/pointcuan/route"
)

func main() {
	// Echo instance
	routes := route.Init()

	db := config.DbConnect()
	defer db.Close()

	UserRepo := UserRepo.CreateUserRepo(db)
	UserUsecase := UserUsecase.CreateUserUsecase(UserRepo)
	UserHandler.CreateUserHandler(routes, UserUsecase)
	
	AdminRepo := AdminRepo.CreateAdminRepo(db)
	AdminUsecase := AdminUsecase.CreateAdminUsecase(AdminRepo)
	AdminHandler.CreateAdminHandler(routes, AdminUsecase)
	// // Start server
	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	fmt.Println("err")
	// }
	// ioutil.WriteFile("routes.json", data, 0644)
	routes.Logger.Fatal(routes.Start(":1323"))
}