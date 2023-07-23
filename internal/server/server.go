package server

import (
	"dime/internal/api"
	"github.com/labstack/echo/v4"
)

const defaultPort = ":8080"

func Start() {
	e := echo.New()

	api.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(defaultPort))
}
