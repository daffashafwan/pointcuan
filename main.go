package main

import (
	"time"

	"github.com/daffashafwan/pointcuan/app/routes"
	_userUsecase "github.com/daffashafwan/pointcuan/business/users"
	_mysqlDriver "github.com/daffashafwan/pointcuan/config"
	_userController "github.com/daffashafwan/pointcuan/controllers/user"
	_userdb "github.com/daffashafwan/pointcuan/model/user"

	_middleware "github.com/daffashafwan/pointcuan/app/middlewares"
	_userRepository "github.com/daffashafwan/pointcuan/model/user"

	//admin
	_adminUsecase "github.com/daffashafwan/pointcuan/business/admin"
	_adminController "github.com/daffashafwan/pointcuan/controllers/admin"
	_admindb "github.com/daffashafwan/pointcuan/model/admin"

	"log"

	"github.com/labstack/echo/v4"
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
	db.AutoMigrate(&_userdb.User{}, &_admindb.Admin{})
}

func main() {
	// init koneksi databse
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.CreateUserRepo(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase)

	//admin
	adminRepository := _admindb.CreateAdminRepo(Conn)
	adminUseCase := _adminUsecase.NewUsecase(adminRepository, timeoutContext, configJWT)
	adminController := _adminController.NewAdminController(adminUseCase)


	routesInit := routes.ControllerList{
		JwtConfig:      configJWT.Init(),
		UserController: *userController,
		AdminController: *adminController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}