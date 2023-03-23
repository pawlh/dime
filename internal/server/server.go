package server

import (
	"dime/internal/api"
	"github.com/labstack/echo/v4"
)

func Start() {

	e := echo.New()

	api.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
