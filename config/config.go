package config

import (
	"fmt"
	"log"

	"github.com/daffashafwan/pointcuan/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DBUser     = "pointcuan"
	DBPassword = "pointcuan"
	DBName     = "pointcuan_db"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBtype     = "mysql"
)

func GetMySQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return dataBase
}

func DbConnect() *gorm.DB {
	consStr := GetMySQLConnectionString()
	db, err := gorm.Open("mysql", consStr)
	if err != nil {
		log.Fatal("Error when connect db" + consStr + " : " + err.Error())
		return nil
	}

	db.Debug().AutoMigrate(
		model.User{}, model.Admin{},
	)
	return db
}