package api

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/ping", Ping)
	e.POST("/login", Login)
}
