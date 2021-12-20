package routes

import (
	"github.com/daffashafwan/pointcuan/controllers/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	//USERS
	e.GET("users", cl.UserController.GetAll)
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.PUT("users/:id", cl.UserController.Update)
	
	//POINTS



	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
