package api

import (
	"dime/internal/dbs"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTransactions(c echo.Context) error {
	transactions, err := dbs.DB.TransactionDao().FindByOwner(c.Get("username").(string))
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "unable to search transactions", err)
	}

	return c.JSON(http.StatusOK, transactions.Transactions)
}

// GetPendingTransactions returns all pending transactions for the user
func GetPendingTransactions(c echo.Context) error {
	//transactions, err := dbs.DB.TransactionDao().FindPendingByOwner(c.Get("username").(string))
	//if err != nil {
	//	return mustSendError(c, http.StatusInternalServerError, "unable to search pending transactions", err)
	//}
	//
	//return c.JSON(http.StatusOK, transactions.Transactions)
	return nil
}
