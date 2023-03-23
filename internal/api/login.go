package api

import (
	"dime/internal/dbs"
	"dime/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	tmpUser := &models.User{
		Username: "test",
		Password: "test",
	}
	err := dbs.DB.UserDao().Insert(tmpUser)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "logged in!")
}
