package routes

import (
	"github.com/daffashafwan/pointcuan/app/middlewares"
	admins "github.com/daffashafwan/pointcuan/controllers/admin"
	pcrcrud "github.com/daffashafwan/pointcuan/controllers/pcr_crud"
	users "github.com/daffashafwan/pointcuan/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
	AdminController admins.AdminController
	PcrController pcrcrud.PcrController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	//USERS
	e.GET("users", cl.UserController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.GET("users/:id", cl.UserController.GetById)
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.PUT("users/:id", cl.UserController.Update)
	e.DELETE("users/:id", cl.UserController.Delete)
	e.GET("users/verif/:token", cl.UserController.Verif)

	//ADMIN 
	e.GET("admin/:id", cl.AdminController.GetById)
	e.POST("admin/login", cl.AdminController.Login)

	//POINTS
	e.GET("pcr/:id", cl.PcrController.GetById)
	e.PUT("pcr/:id", cl.PcrController.Update)
	//PCR

	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
