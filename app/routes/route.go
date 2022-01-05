package routes

import (
	"github.com/daffashafwan/pointcuan/app/middlewares"
	item "github.com/daffashafwan/pointcuan/controllers/items"
	admins "github.com/daffashafwan/pointcuan/controllers/admin"
	category "github.com/daffashafwan/pointcuan/controllers/categoryitem"
	pcrcrud "github.com/daffashafwan/pointcuan/controllers/pcr_crud"
	point "github.com/daffashafwan/pointcuan/controllers/point"
	transaction "github.com/daffashafwan/pointcuan/controllers/transaction"
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
	TransactionController transaction.TransactionController
	CategoryController category.CategoryItemController
	ItemsController item.ItemsController
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
	e.POST("users/forgotpassword", cl.UserController.ForgotPassword)
	e.GET("users/forgotpassword/:token", cl.UserController.VerifyTokenPassword)
	e.POST("users/resetpassword/:id", cl.UserController.ResetPassword)

	//USERS WITH POINT
	e.PUT("users/:id/point", cl.PointController.Update)
	e.DELETE("users/:id/point", cl.PointController.Delete)

	//USERS WITH TRANSACTION
	e.GET("users/:id/transaction/:tid", cl.TransactionController.GetById)
	e.GET("users/:id/transaction", cl.TransactionController.GetByUserId)
	e.POST("users/:id/transaction", cl.TransactionController.Create)
	e.PUT("users/:id/transaction/:tid", cl.TransactionController.Update)
	e.DELETE("users/:id/transaction/:tid", cl.TransactionController.Delete)

	//TRANSACTION
	e.GET("transaction", cl.TransactionController.GetAll)

	//ADMIN 
	e.GET("admin/:id", cl.AdminController.GetById)
	e.POST("admin/login", cl.AdminController.Login)

	//POINTS
	e.GET("pcr", cl.PcrController.GetPCR)
	e.PUT("pcr", cl.PcrController.Update)
	//PCR

	//Category
	e.GET("categoryitems", cl.CategoryController.GetAll)
	e.GET("categoryitems/:id", cl.CategoryController.GetById)
	e.POST("categoryitems", cl.CategoryController.Create)
	e.PUT("categoryitems/:id", cl.CategoryController.Update)
	e.DELETE("categoryitems/:id", cl.CategoryController.Delete)

	//Items
	e.GET("items", cl.ItemsController.GetAll)
	e.PUT("items/:id", cl.ItemsController.Update)
	e.DELETE("items/:id", cl.ItemsController.Delete)

	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
