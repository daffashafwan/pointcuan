package main

import (
	"github.com/daffashafwan/pointcuan/app/routes"
	_userUsecase "github.com/daffashafwan/pointcuan/business/users"
	_userController "github.com/daffashafwan/pointcuan/controllers/user"
	_userdb "github.com/daffashafwan/pointcuan/model/user"
	_mysqlDriver "github.com/daffashafwan/pointcuan/config"
	"time"

	_middleware "github.com/daffashafwan/pointcuan/app/middlewares"
	_userRepository "github.com/daffashafwan/pointcuan/model/user"

	"log"

	"github.com/labstack/echo"
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
	db.AutoMigrate(&_userdb.User{})
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
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	routesInit := routes.ControllerList{
		JwtConfig:      configJWT.Init(),
		UserController: *userController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}