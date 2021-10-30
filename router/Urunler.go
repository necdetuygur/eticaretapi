package router

import (
	"eticaretapi/service"

	"github.com/labstack/echo/v4"
)

func UrunlerRouter(e *echo.Echo) {
	e.POST("/Urunler", service.UrunlerAdd)
	e.GET("/Urunler", service.UrunlerList)
	e.GET("/Urunler/:id", service.UrunlerGet)
	e.DELETE("/Urunler/:id", service.UrunlerDelete)
	e.PUT("/Urunler/:id", service.UrunlerSet)
}
