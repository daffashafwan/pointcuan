package routes

import (
	"github.com/daffashafwan/pointcuan/app/middlewares"
	admins "github.com/daffashafwan/pointcuan/controllers/admin"
	pcrcrud "github.com/daffashafwan/pointcuan/controllers/pcr_crud"
	point "github.com/daffashafwan/pointcuan/controllers/point"
	users "github.com/daffashafwan/pointcuan/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
	AdminController admins.AdminController
	PcrController pcrcrud.PcrController
	PointController point.PointController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	//USERS
	e.GET("users", cl.UserController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.GET("users/:id", cl.UserController.GetById)
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.PUT("users/:id", cl.UserController.Update)
	e.DELETE("users/:id", cl.UserController.Delete)
	e.GET("users/verify/:token", cl.UserController.Verify)

	//USERS WITH POINT
	e.PUT("users/:id/point", cl.PointController.Update)
	e.DELETE("users/:id/point", cl.PointController.Delete)

	//ADMIN 
	e.GET("admin/:id", cl.AdminController.GetById)
	e.POST("admin/login", cl.AdminController.Login)

	//POINTS
	e.GET("pcr", cl.PcrController.GetPCR)
	e.PUT("pcr", cl.PcrController.Update)
	//PCR

	//Category
	e.GET("categoryitems", cl.CategoryItemController.GetAll)
	e.GET("categoryitem/:id", cl.CategoryItem.GetById)
	e.POST("categoryitem", cl.CategoryItem.Register)
	e.PUT("categoryitem/:id", cl.CategoryItem.Update)
	e.DELETE("categoryitem/:id", cl.CategoryItem.Delete)

	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
