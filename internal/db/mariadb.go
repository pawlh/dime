package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var schema string = `
CREATE TABLE IF NOT EXISTS transaction (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    date DATETIME NOT NULL,
    description TEXT NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    category TEXT NOT NULL,
    account TEXT NOT NULL
);
`

func InitMariaDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS dime")
	if err != nil {
		return nil, err
	}

	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/dime")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
