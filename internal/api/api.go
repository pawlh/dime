package api

import (
	"github.com/labstack/echo/v4"
)

func mustSendError(c echo.Context, code int, error string) error {
	if err := c.JSON(code, echo.Map{
		"error": error,
	}); err != nil {
		return err
	}
	return nil
}
