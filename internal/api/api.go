package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func mustSendError(c echo.Context, code int, error string, originalErr error) error {
	if err := c.JSON(code, echo.Map{
		"error": error,
	}); err != nil {
		log.Error(originalErr)
		return err
	}
	return nil
}
