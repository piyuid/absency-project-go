package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func connectDB() {
	dsn := "root:1234ABC.@tcp(127.0.0.1:8000)/<namaDatabase>?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Tidak Connect")
	}
}

func Migration() {
	db
}

type 	

func main() {
	e := echo.New()

	e.GET("/", HelloController)

	e.Start(":8000")
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
