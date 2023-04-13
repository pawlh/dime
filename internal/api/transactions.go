package api

import (
	"dime/internal/dbs"
	"dime/internal/transaction"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type activeConnection struct {
	Username string
	Id       uint32
	Conn     *websocket.Conn
}

var activeConnections []activeConnection

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func GetTransactions(c echo.Context) error {
	transactions, err := dbs.DB.TransactionDao().FindByOwner(c.Get("username").(string))
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "unable to search transactions", err)
	}

	return c.JSON(http.StatusOK, transactions.Transactions)
}

func GetPendingTransactions(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var updateRequest transaction.UpdateRequest
		err = json.Unmarshal(message, &updateRequest)
		if err != nil {
			err := ws.WriteMessage(websocket.TextMessage, []byte("Error unmarshalling data"))
			if err != nil {
				log.Println("Error sending error to client:", err)
				return err
			}
		}

		// Save the pending transactions
		//err = dbs.DB.PendingTransactionsDao().SavePendingTransactions(data, c.Get("username").(string))

	}

	return nil

}
