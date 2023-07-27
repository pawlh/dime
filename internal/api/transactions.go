package api

import (
	"dime/internal/database"
	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	transactionDao, err := database.DB.TransactionDAO()
	if err != nil {
		return err
	}
	defer database.DB.Disconnect()

	// /transaction?owner=ownerOfTransactions
	owner := c.QueryParam("owner")

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
