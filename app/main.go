package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConnMysql() {
	dsn := "root:kita1234@tcp(127.0.0.1:3306)/absencydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func InitMigrateMysql() {
	DB.AutoMigrate(&User)
}

func main() {
	e := echo.Now()

	// Routing with Query Parameter
	e.POST("/users", CreateUser)

	e.Star(":8000")
}
