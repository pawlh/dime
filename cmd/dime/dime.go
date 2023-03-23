package main

import (
	"database/sql"
	"dime/internal/dao"
	"dime/internal/dbs"
	"dime/internal/models"
	"log"
	"time"
)

func main() {
	db, err := dbs.InitMariaDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	sampleTransaction := models.Transaction{
		Date:        time.Now(),
		Description: "Sample transaction",
		Amount:      100.00,
		Category:    "Sample category",
		Account:     "Sample account",
	}

	transactionDao := dao.NewMariaDbTransaction(db)
	err = transactionDao.Insert(&sampleTransaction)
	if err != nil {
		log.Fatal(err)
	}

}
