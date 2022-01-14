package main

import (
	"eticaretapi/config"
	"eticaretapi/model"
	"eticaretapi/router"
	"os"

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
	e.HideBanner = true
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
	host := "localhost"
	port := config.PORT
	if os.Getenv("PORT") != "" {
		host = "0.0.0.0"
		port = os.Getenv("PORT")
	}
	e.Logger.Fatal(e.Start(host + ":" + port))
}
