package main

import (
	"eticaretapi/config"
	"eticaretapi/model"
	"eticaretapi/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// TABLO YOKSA OLUŞTUR METODLARI
	model.TodoCreateTable()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// YÖNLENDİRİCİLER
	router.TodoRouter(e)

	e.Logger.Fatal(e.Start("0.0.0.0:" + config.PORT))
}
