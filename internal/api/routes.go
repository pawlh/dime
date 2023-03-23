package api

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	apiGroup := e.Group("/api")
	apiGroup.POST("/login", Login)

	apiGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	apiGroup.GET("/ping", Ping)
}
