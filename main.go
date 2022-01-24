package main

import (
	"time"

	"github.com/daffashafwan/pointcuan/app/routes"
	_userUsecase "github.com/daffashafwan/pointcuan/business/users"
	_mysqlDriver "github.com/daffashafwan/pointcuan/config"
	_userController "github.com/daffashafwan/pointcuan/controllers/user"
	_userdb "github.com/daffashafwan/pointcuan/model/user"

	_pointUsecase "github.com/daffashafwan/pointcuan/business/point"
	_pointController "github.com/daffashafwan/pointcuan/controllers/point"
	_pointRepository "github.com/daffashafwan/pointcuan/model/point"
	_pointdb "github.com/daffashafwan/pointcuan/model/point"

	_transactionUsecase "github.com/daffashafwan/pointcuan/business/transactions"
	_transactionController "github.com/daffashafwan/pointcuan/controllers/transaction"
	_transactionRepository "github.com/daffashafwan/pointcuan/model/transactions"
	_transactiondb "github.com/daffashafwan/pointcuan/model/transactions"

	_categoryUsecase "github.com/daffashafwan/pointcuan/business/categoryItems"
	_categoryController "github.com/daffashafwan/pointcuan/controllers/categoryitem"
	_categoryRepository "github.com/daffashafwan/pointcuan/model/category"
	_categorydb "github.com/daffashafwan/pointcuan/model/category"

	_faqUsecase "github.com/daffashafwan/pointcuan/business/FAQ"
	_faqController "github.com/daffashafwan/pointcuan/controllers/faq"
	_faqRepository "github.com/daffashafwan/pointcuan/model/FAQ"
	_faqdb "github.com/daffashafwan/pointcuan/model/FAQ"

	_itemsUsecase "github.com/daffashafwan/pointcuan/business/items"
	_itemsController "github.com/daffashafwan/pointcuan/controllers/items"
	_itemsRepository "github.com/daffashafwan/pointcuan/model/items"
	_itemsdb "github.com/daffashafwan/pointcuan/model/items"

	_redeemUsecase "github.com/daffashafwan/pointcuan/business/redeem"
	_redeemController "github.com/daffashafwan/pointcuan/controllers/redeem"
	_redeemRepository "github.com/daffashafwan/pointcuan/model/redeem"
	_redeemdb "github.com/daffashafwan/pointcuan/model/redeem"

	_middleware "github.com/daffashafwan/pointcuan/app/middlewares"
	_userRepository "github.com/daffashafwan/pointcuan/model/user"

	//admin
	_adminUsecase "github.com/daffashafwan/pointcuan/business/admin"
	_adminController "github.com/daffashafwan/pointcuan/controllers/admin"
	_admindb "github.com/daffashafwan/pointcuan/model/admin"

	//pcr
	_pcr "github.com/daffashafwan/pointcuan/business/pcr_crud"
	_pcrController "github.com/daffashafwan/pointcuan/controllers/pcr_crud"
	_pcrdb "github.com/daffashafwan/pointcuan/model/pcr_crud"

	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userdb.User{}, 
		&_admindb.Admin{}, 
		&_pointdb.Point{}, 
		&_pcrdb.Pcrcrud{}, 
		&_transactiondb.Transaction{}, 
		&_categorydb.Category{},
		&_itemsdb.Items{},
		&_redeemdb.Redeem{},
		&_faqdb.Faq{})
}

func main() {
	// init koneksi databse
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database-aws.user`),
		DB_Password: viper.GetString(`database-aws.pass`),
		DB_Host:     viper.GetString(`database-aws.host`),
		DB_Port:     viper.GetString(`database-aws.port`),
		DB_Database: viper.GetString(`database-aws.name`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	//feAddress := viper.GetString(`frontend.address`)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderAuthorization,echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowMethods, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlRequestHeaders, echo.HeaderAccessControlAllowCredentials},
	  }))
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	pointRepository := _pointRepository.CreatePointRepo(Conn)
	pointUseCase := _pointUsecase.NewPointUsecase(pointRepository, timeoutContext, configJWT)
	pointController := _pointController.NewPointController(pointUseCase)

	//admin
	adminRepository := _admindb.CreateAdminRepo(Conn)
	adminUseCase := _adminUsecase.NewUsecase(adminRepository, timeoutContext, configJWT)
	adminController := _adminController.NewAdminController(adminUseCase)

	//pcr
	pcrRepo := _pcrdb.CreatePcrRepo(Conn)
	pcrUseCase := _pcr.NewPcrcase(pcrRepo, timeoutContext)
	pcrController := _pcrController.NewPcrController(pcrUseCase)

	transactionRepository := _transactionRepository.CreateTransactionRepo(Conn)
	transactionUseCase := _transactionUsecase.NewTransactionUsecase(transactionRepository,pointRepository, timeoutContext, configJWT)
	transactionController := _transactionController.NewTransactionController(transactionUseCase, pcrUseCase)

	categoryRepository := _categoryRepository.CreateCategoryRepo(Conn)
	categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryRepository, timeoutContext, configJWT)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	faqRepository := _faqRepository.CreateFAQRepo(Conn)
	faqUseCase := _faqUsecase.NewFAQUsecase(faqRepository, timeoutContext, configJWT)
	faqController := _faqController.NewFAQController(faqUseCase)

	itemsRepository := _itemsRepository.CreateItemRepo(Conn)
	itemsUseCase := _itemsUsecase.NewItemsUsecase(itemsRepository, timeoutContext)
	itemsController := _itemsController.NewItemsController(itemsUseCase)

	redeemRepository := _redeemRepository.CreateRedeemRepo(Conn)
	redeemUseCase := _redeemUsecase.NewRedeemUsecase(pointRepository,itemsRepository,redeemRepository, timeoutContext, configJWT)
	redeemController := _redeemController.NewRedeemController(redeemUseCase, itemsUseCase)

	userRepository := _userRepository.CreateUserRepo(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase, pointUseCase, redeemUseCase, transactionUseCase)

	routesInit := routes.ControllerList{
		JwtConfig:      configJWT.Init(),
		UserController: *userController,
		AdminController: *adminController,
		PcrController: *pcrController,
		PointController: *pointController,
		TransactionController: *transactionController,
		CategoryController: *categoryController,
		ItemsController: *itemsController,
		RedeemController: *redeemController,
		FaqController: *faqController,
	}
	
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}