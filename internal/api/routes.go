package api

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	apiGroup := e.Group("/api")
	apiGroup.POST("/login", Login)
	apiGroup.POST("/register", Register)

	apiGroup.Use(echojwt.JWT(secret))
	apiGroup.Use(validateToken)

	apiGroup.POST("/upload", Upload)
	apiGroup.GET("/archive", GetArchives)
	apiGroup.GET("/archive/:id", GetArchive)

}
