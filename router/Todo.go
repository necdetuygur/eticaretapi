package router

import (
	"eticaretapi/service"

	"github.com/labstack/echo/v4"
)

func TodoRouter(e *echo.Echo) {
	e.POST("/todo", service.TodoAdd)
	e.GET("/todo", service.TodoList)
	e.GET("/todo/:id", service.TodoGet)
	e.DELETE("/todo/:id", service.TodoDelete)
	e.PUT("/todo/:id", service.TodoSet)
}
