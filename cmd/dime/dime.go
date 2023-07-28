package main

import (
	"dime/internal/server"
	"os"
)

func main() {
	if os.Getenv("MONGODB_URI") == "" {
		panic("MONGODB_URI environment variable not set")
	}

	server.Start()
}
