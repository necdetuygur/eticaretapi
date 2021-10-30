package router

import (
	"eticaretapi/service"

	"github.com/labstack/echo/v4"
)

func KullanicilarRouter(e *echo.Echo) {
	e.POST("/Kullanicilar", service.KullanicilarAdd)
	e.GET("/Kullanicilar", service.KullanicilarList)
	e.GET("/Kullanicilar/:id", service.KullanicilarGet)
	e.DELETE("/Kullanicilar/:id", service.KullanicilarDelete)
	e.PUT("/Kullanicilar/:id", service.KullanicilarSet)
}
