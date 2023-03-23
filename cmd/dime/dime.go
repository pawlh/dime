package main

import (
	"dime/internal/dbs"
	"dime/internal/server"
	"log"
)

func main() {

	err := dbs.InitMariaDB()
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
