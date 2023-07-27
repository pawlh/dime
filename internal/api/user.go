package api

import (
	"dime/internal/models"
	"github.com/labstack/echo/v4"
)

type getMeResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetMe(c echo.Context) error {
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	return c.JSON(200, getMeResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
}
