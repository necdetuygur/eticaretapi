package router

import (
	"eticaretapi/service"

	"github.com/labstack/echo/v4"
)

func KullaniciRouter(e *echo.Echo) {
	e.POST("/Kullanici", service.KullaniciAdd)
	e.GET("/Kullanici", service.KullaniciList)
	e.GET("/Kullanici/:id", service.KullaniciGet)
	e.DELETE("/Kullanici/:id", service.KullaniciDelete)
	e.PUT("/Kullanici/:id", service.KullaniciSet)
}
