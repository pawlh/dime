package api

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	apiGroup := e.Group("/api")
	apiGroup.POST("/login", Login)
	apiGroup.POST("/register", Register)

	apiGroup.Use(echojwt.WithConfig(echojwt.Config{
		TokenLookup: "cookie:token",
		SigningKey:  secret,
	}))
	apiGroup.Use(validateToken)

	apiGroup.POST("/upload", Upload)
	apiGroup.GET("/archive", GetArchives)
	apiGroup.GET("/archive/:id", GetArchive)

	e.Static("/", "frontend/dist")

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		defaultPage := "frontend/dist/index.html"
		if err := c.File(defaultPage); err != nil {
			c.Logger().Error(err)
		}
	}

}
