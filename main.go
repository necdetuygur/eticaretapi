package main

import (
	"eticaretapi/config"
	"eticaretapi/model"
	"eticaretapi/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// VERİTABANINDA TABLO YOKSA OLUŞTUR METODLARI
	model.UrunCreateTable()
	model.KategoriCreateTable()
	model.KullaniciCreateTable()

	// WEB FRAMEWORK TANIMLAMALARI
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// WEB FRAMEWORK YÖNLENDİRİCİLERİ
	router.UrunRouter(e)
	router.KategoriRouter(e)
	router.KullaniciRouter(e)

	// WEB FRAMEWORK BAŞLAT
	e.Logger.Fatal(e.Start("0.0.0.0:" + config.PORT))
}
