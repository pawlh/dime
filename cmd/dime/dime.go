package main

import (
	"dime/internal/dbs"
	"dime/internal/server"
	"log"
	"os"
)

func main() {

	dbHost := os.Getenv("MONGO_HOST")
	err := dbs.InitMongoDB("mongodb://" + dbHost)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
