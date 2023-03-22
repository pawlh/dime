package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Init() error {

	tmpDb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err != nil {
		return err
	}

	_, err = tmpDb.Exec("CREATE DATABASE IF NOT EXISTS dime")
	if err != nil {
		return err
	}

	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/dime")
	if err != nil {
		return err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
