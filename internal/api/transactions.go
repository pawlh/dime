package api

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

type Transaction struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

var (
	upgrader = websocket.Upgrader{}
)

func GetTransactions(c echo.Context) error {
	// upgrade the http connection to a websocket connection
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	go sendTransactions(ws)

	return nil
}

func sendTransactions(conn *websocket.Conn) {
	// Send a ping message every 30 seconds to keep the connection alive
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// Generate some dummy transactions
	transactions := []Transaction{
		{
			ID:          1,
			Description: "Transaction 1",
			Amount:      10.0,
			Date:        time.Now(),
		},
		{
			ID:          2,
			Description: "Transaction 2",
			Amount:      20.0,
			Date:        time.Now().Add(-24 * time.Hour),
		},
		{
			ID:          3,
			Description: "Transaction 3",
			Amount:      30.0,
			Date:        time.Now().Add(-48 * time.Hour),
		},
	}

	// Loop over the transactions and send them to the client
	for _, transaction := range transactions {
		// Convert the transaction to JSON
		data, err := json.Marshal(transaction)
		if err != nil {
			log.Println("Error marshalling transaction:", err)
			continue
		}

		// Write the JSON data to the websocket connection
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Error sending transaction to client:", err)
			return
		}

		// Wait for 1 second before sending the next transaction
		time.Sleep(1 * time.Second)

		// Check if a ping message has been received from the client
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		_, _, err = conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from client:", err)
			return
		}
	}

	// Close the websocket connection
	conn.Close()
}
