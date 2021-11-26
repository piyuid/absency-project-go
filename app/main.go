package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	// Binding Data
	user := User{}
	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Membuat Akun",
		"user":    user,
	})
}

var DB *gorm.DB

type User struct {
	ID             uint64 `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Pin            uint32 `json:"pin"`
	Previllage     uint32 `json:"previllage"`
	PrevillageUser string `json:"previllageuser"`
	FingerId       string `json:"fingerid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func InitConnMysql() {
	dsn := "root:kita1234@tcp(127.0.0.1:3306)/absencydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrateMysql()
}

func InitMigrateMysql() {
	DB.AutoMigrate(&User{})
}

func main() {
	InitConnMysql()

	e := echo.New()

	// Routing with Query Parameter
	e.POST("/users", CreateUser)

	e.Start(":8000")
}
