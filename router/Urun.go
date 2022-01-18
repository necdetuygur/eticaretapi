package router

import (
	"eticaretapi/service"

	"github.com/labstack/echo/v4"
)

func UrunRouter(e *echo.Group) {
	e.POST("/Urun", service.UrunAdd)
	e.GET("/Urun", service.UrunList)
	e.GET("/Urun/:id", service.UrunGet)
	e.DELETE("/Urun/:id", service.UrunDelete)
	e.PUT("/Urun/:id", service.UrunSet)
}
