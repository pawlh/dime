package api

import (
	"dime/internal/dbs"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type Transaction struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

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
	// upgrade the http connection to a websocket connection
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	// Add the connection to the active connections
	newConnection := activeConnection{
		Username: c.Get("username").(string),
		Conn:     ws,
		Id:       uuid.New().ID(),
	}
	activeConnections = append(activeConnections, newConnection)

	fmt.Println("New connection: ", newConnection.Username)

	sendTransactions(newConnection)

	// this isn't working cross-origin... need to figure out how to do this
	//go pinger(newConnection)

	return nil
}

func pinger(connection activeConnection) {
	for {
		err := connection.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			return
		}

		err = connection.Conn.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			log.Println("Error sending ping to client:", err)
			return
		}

		_, _, err = connection.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading ping from client:", err)
			return
		}

		err = connection.Conn.SetReadDeadline(time.Time{})
		if err != nil {
			err := connection.Conn.Close()
			if err != nil {
				log.Println("Error closing connection:", err)
				return
			}
			removeActiveConnection(connection.Id)
			return
		}

		time.Sleep(5 * time.Second)
	}
}

func sendTransactions(connection activeConnection) {
	transactions, err := dbs.DB.TransactionDao().FindByOwner(connection.Username)
	if err != nil {
		log.Println("Error finding transactions:", err)
		return
	}

	if transactions != nil {

		// Convert the transaction to JSON
		data, err := json.Marshal(transactions)
		if err != nil {
			log.Println("Error marshalling transaction:", err)
		}

		// Write the JSON data to the websocket connection
		err = connection.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Error sending transaction to client:", err)
			return
		}
	}
}

func BroadcastTransactions(username string) {
	for _, connection := range activeConnections {
		if connection.Username == username {
			sendTransactions(connection)
		}
	}
}

func removeActiveConnection(uuid uint32) {
	for i, connection := range activeConnections {
		if connection.Id == uuid {
			activeConnections = append(activeConnections[:i], activeConnections[i+1:]...)
		}
	}
}
