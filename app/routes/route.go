package routes

import (
	"github.com/daffashafwan/pointcuan/app/middlewares"
	users "github.com/daffashafwan/pointcuan/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
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

	//POINTS

	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
