package api

import (
	"dime/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	transactionDao, err := database.DB.TransactionDAO()
	if err != nil {
		return err
	}
	defer database.DB.Disconnect()

	user := c.Get("user").(*jwt.Token)

	owner := user.Claims.(jwt.MapClaims)["userId"].(string)

	if owner == "" {
		return c.JSON(400, "Missing owner parameter")
	}

	transactions, err := transactionDao.GetTransactions(owner)
	if err != nil {
		return err
	}

	if len(transactions) == 0 {
		return c.JSON(200, []string{})
	}

	return c.JSON(200, transactions)
}
