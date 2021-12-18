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
	e.POST("users/login", cl.UserController.Login)
	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
