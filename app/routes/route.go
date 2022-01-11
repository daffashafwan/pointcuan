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
	redeem "github.com/daffashafwan/pointcuan/controllers/redeem"

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
	RedeemController redeem.RedeemController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//USERS
	e.GET("users", cl.UserController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("users/:id", cl.UserController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.PUT("users/:id", cl.UserController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.DELETE("users/:id", cl.UserController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("users/verify/:token", cl.UserController.Verify)
	e.POST("users/forgotpassword", cl.UserController.ForgotPassword)
	e.GET("users/forgotpassword/:token", cl.UserController.VerifyTokenPassword)
	e.POST("users/resetpassword/:id", cl.UserController.ResetPassword)

	//USERS WITH POINT
	e.PUT("users/:id/point", cl.PointController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.DELETE("users/:id/point", cl.PointController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//USERS WITH TRANSACTION
	e.GET("users/:id/transaction/:tid", cl.TransactionController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("users/:id/transaction", cl.TransactionController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.POST("users/:id/transaction", cl.TransactionController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.PUT("users/:id/transaction/:tid", cl.TransactionController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("users/:id/transaction/:tid", cl.TransactionController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//TRANSACTION
	e.GET("transaction", cl.TransactionController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//ADMIN 
	e.GET("admin/:id", cl.AdminController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.POST("admin/login", cl.AdminController.Login)

	//POINTS
	e.GET("pcr", cl.PcrController.GetPCR, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.PUT("pcr", cl.PcrController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	//PCR

	//Category
	e.GET("categoryitems", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("categoryitems/:id", cl.CategoryController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.POST("categoryitems", cl.CategoryController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.PUT("categoryitems/:id", cl.CategoryController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("categoryitems/:id", cl.CategoryController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//Items
	e.POST("items", cl.ItemsController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.GET("items", cl.ItemsController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("items/:id", cl.ItemsController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("items/category/:id", cl.ItemsController.GetByCategoryId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.PUT("items/:id", cl.ItemsController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.PUT("items/:id/stock", cl.ItemsController.UpdateStock, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("items/:id", cl.ItemsController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//Redeem
	e.GET("redeem", cl.RedeemController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("redeem/:id", cl.RedeemController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("redeem/item/:id", cl.RedeemController.GetByItemId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.GET("users/:id/redeem", cl.RedeemController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.POST("users/:id/redeem", cl.RedeemController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("redeem/:id", cl.RedeemController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
