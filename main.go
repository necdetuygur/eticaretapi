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
	model.UrunlerCreateTable()
	model.KategorilerCreateTable()
	model.KullanicilarCreateTable()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// YÖNLENDİRİCİLER
	router.TodoRouter(e)
	router.UrunlerRouter(e)
	router.KategorilerRouter(e)
	router.KullanicilarRouter(e)

	e.Logger.Fatal(e.Start("0.0.0.0:" + config.PORT))
}
