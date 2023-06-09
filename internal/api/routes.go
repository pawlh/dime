package api

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func RegisterRoutes(e *echo.Echo) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:1323",
			"https://dime.pawl.app",
		},
		AllowCredentials: true,
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
	apiGroup.GET("/pending_transactions", GetPendingTransactions)
	apiGroup.GET("/transactions", GetTransactions)

	e.Static("/", "frontend/dist")

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if err.(*echo.HTTPError).Code == 401 {
			_ = mustSendError(c, http.StatusUnauthorized, "missing or invalid JWT", err)
		} else {
			defaultPage := "frontend/dist/index.html"
			if err = c.File(defaultPage); err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
