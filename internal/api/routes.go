package api

import (
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	api := e.Group("/api")

	api.POST("/login", Login)
	api.POST("/register", Register)
	api.GET("/users", GetUsers)

	authenticatedApi := api.Group("")
	authenticatedApi.Use(echojwt.WithConfig(echojwt.Config{
		TokenLookup: "cookie:token",
		SigningKey:  secret,
	}))
	authenticatedApi.Use(RenewTokenMiddleware)

	authenticatedApi.POST("/logout", Logout)

	authenticatedApi.GET("/me", GetMe)
	authenticatedApi.GET("/transaction", GetTransactions)

	// Default
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Hello, World!")
		return c.String(200, "Hello, World!")
	})
}
