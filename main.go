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
	// VERİ TABANINDA TABLO YOKSA OLUŞTUR METODLARI
	model.UrunCreateTable()
	model.KategoriCreateTable()
	model.KullaniciCreateTable()
	model.TodoCreateTable()

	// WEB FRAMEWORK TANIMLAMALARI
	ec := echo.New()
	ec.HideBanner = true
	ec.Use(middleware.Logger())
	ec.Use(middleware.Recover())
	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// WEB FRAMEWORK AUTH
	ec.POST("/Login", Login)
	e := ec.Group("")
	e.Use(middleware.JWT([]byte(config.AUTH_JWT_TOKEN)))

	// WEB FRAMEWORK YÖNLENDİRİCİLERİ
	router.UrunRouter(e)
	router.KategoriRouter(e)
	router.KullaniciRouter(e)
	router.TodoRouter(e)

	// WEB FRAMEWORK BAŞLAT
	host := "localhost"
	port := config.PORT
	if os.Getenv("PORT") != "" {
		host = "0.0.0.0"
		port = os.Getenv("PORT")
	}
	ec.Logger.Fatal(ec.Start(host + ":" + port))
}
