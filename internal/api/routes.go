package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	api := e.Group("/api")

	api.GET("/me", GetMe)
	api.GET("/transaction", GetTransactions)

	// Default
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Hello, World!")
		return c.String(200, "Hello, World!")
	})
}
