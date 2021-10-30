package router

import (
	"eticaretapi/service"

	"github.com/labstack/echo/v4"
)

func KategorilerRouter(e *echo.Echo) {
	e.POST("/Kategoriler", service.KategorilerAdd)
	e.GET("/Kategoriler", service.KategorilerList)
	e.GET("/Kategoriler/:id", service.KategorilerGet)
	e.DELETE("/Kategoriler/:id", service.KategorilerDelete)
	e.PUT("/Kategoriler/:id", service.KategorilerSet)
}
