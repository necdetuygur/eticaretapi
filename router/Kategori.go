package router

import (
	"eticaretapi/service"

	"github.com/labstack/echo/v4"
)

func KategoriRouter(e *echo.Echo) {
	e.POST("/Kategori", service.KategoriAdd)
	e.GET("/Kategori", service.KategoriList)
	e.GET("/Kategori/:id", service.KategoriGet)
	e.DELETE("/Kategori/:id", service.KategoriDelete)
	e.PUT("/Kategori/:id", service.KategoriSet)
}
