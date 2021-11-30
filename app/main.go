package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID             uint64 `gorm:"primaryKey"`
	Username       string `json:"username" form:"username"`
	Password       string `json:"password" form:"password"`
	Email          string `json:"email" form:"email"`
	Pin            uint32 `json:"pin" form:"pin"`
	Previllage     uint32 `json:"previllage" form:"previllage"`
	PrevillageUser string `json:"previllageuser" form:"previllageuser"`
	FingerId       string `json:"fingerid" form:"fingerid"`
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
	e.GET("/users", GetUserController)
	e.POST("/users", CreateUserController)

	e.Start(":8000")
}

func GetUserController(c echo.Context) error {
	var users []User

	err := DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	// Binding Data
	user := User{}
	c.Bind(&user)

	err := DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Membuat Akun",
		"user":    user,
	})
}
