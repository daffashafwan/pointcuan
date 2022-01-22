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
	faq "github.com/daffashafwan/pointcuan/controllers/faq"

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
	FaqController faq.FaqController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//USERS
	e.GET("users", cl.UserController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.GET("users/:id", cl.UserController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("admin/users/:id", cl.UserController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.PUT("users/:id", cl.UserController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin, middlewares.IsUserId)
	e.DELETE("users/:id", cl.UserController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.GET("users/verify/:token", cl.UserController.Verify)
	e.POST("users/forgotpassword", cl.UserController.ForgotPassword)
	e.GET("users/forgotpassword/:token", cl.UserController.VerifyTokenPassword)
	e.POST("users/resetpassword/:id", cl.UserController.ResetPassword)

	//USERS WITH POINT (NOT USED)
	e.PUT("users/:id/point", cl.PointController.Update)
	e.DELETE("users/:id/point", cl.PointController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//USERS WITH TRANSACTION
	e.GET("users/:id/transaction/:tid", cl.TransactionController.GetById, middleware.JWTWithConfig(cl.JwtConfig),middlewares.IsUserId)
	e.GET("/:id/transaction/:tid", cl.TransactionController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.GET("users/:id/transaction", cl.TransactionController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.POST("users/:id/transaction/status", cl.TransactionController.GetByUserIdAndStatus, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("/:id/transaction", cl.TransactionController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.POST("users/:id/transaction", cl.TransactionController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.PUT("transaction/:tid", cl.TransactionController.Update)
	e.DELETE("users/:id/transaction/:tid", cl.TransactionController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//TRANSACTION
	e.GET("transaction", cl.TransactionController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//ADMIN 
	e.GET("admin/:id", cl.AdminController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.POST("admin/login", cl.AdminController.Login)

	//POINTS
	e.GET("/:id/pcr", cl.PcrController.GetPCR, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("pcr", cl.PcrController.GetPCR, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.PUT("pcr", cl.PcrController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	//PCR

	//Category
	e.GET("/:id/categoryitems", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("categoryitems", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.GET("/:id/categoryitems/:cid", cl.CategoryController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("categoryitems/:cid", cl.CategoryController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.POST("categoryitems", cl.CategoryController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.PUT("categoryitems/:cid", cl.CategoryController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("categoryitems/:cid", cl.CategoryController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//FAQ
	e.GET("/:id/faq", cl.FaqController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("faq", cl.FaqController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.GET("/:id/faq/active", cl.FaqController.GetActive, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("faq/active", cl.FaqController.GetActive, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.POST("faq", cl.FaqController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.PUT("faq/:fid", cl.FaqController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("faq/:fid", cl.FaqController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//Items
	e.POST("items", cl.ItemsController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.GET("/:id/items", cl.ItemsController.GetAll, middleware.JWTWithConfig(cl.JwtConfig),middlewares.IsUserId)
	e.GET("items", cl.ItemsController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.GET("/:id/items/:iid", cl.ItemsController.GetById, middleware.JWTWithConfig(cl.JwtConfig),middlewares.IsUserId)
	e.GET("items/:iid", cl.ItemsController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.GET("/:id/items/category/:cid", cl.ItemsController.GetByCategoryId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("items/category/:cid", cl.ItemsController.GetByCategoryId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.PUT("items/:iid", cl.ItemsController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.PUT("items/:iid/stock", cl.ItemsController.UpdateStock, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("items/:iid", cl.ItemsController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//Redeem
	e.GET("/:id/redeem", cl.RedeemController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("redeem", cl.RedeemController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.GET("/:id/redeem/:rid", cl.RedeemController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("redeem/:rid", cl.RedeemController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.GET("/:id/redeem/item/:iid", cl.RedeemController.GetByItemId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("redeem/item/:iid", cl.RedeemController.GetByItemId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	e.GET("users/:id/redeem", cl.RedeemController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET(":id/redeem", cl.RedeemController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.POST("users/:id/redeem", cl.RedeemController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.DELETE("redeem/:id", cl.RedeemController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
}
