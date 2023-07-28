package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// mustSendError sends an error response to the client. If the response cannot be sent, the original error is logged.
func mustSendError(c echo.Context, code int, error string, originalErr error) error {
	if err := c.JSON(code, echo.Map{
		"error": error,
	}); err != nil {
		log.Error(originalErr)
		return err
	}
	log.Error(originalErr)
	return nil
}
