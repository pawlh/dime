package dbs

import (
	"database/sql"
	"dime/internal/dao"
	_ "github.com/go-sql-driver/mysql"
)

var userSchema = `
CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    name TEXT NOT NULL
);
`
var transactionSchema = `
CREATE TABLE IF NOT EXISTS transaction (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    date DATETIME NOT NULL,
    description TEXT NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    category TEXT NOT NULL,
    account TEXT NOT NULL
);
`

type MariaDB struct {
	db *sql.DB
}

func (m MariaDB) TransactionDao() dao.TransactionDao {
	return dao.NewMariaDbTransaction(m.db)
}

func (m MariaDB) UserDao() dao.UserDAO {
	return dao.NewMariaDbUser(m.db)
}

func InitMariaDB() error {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS dime")
	if err != nil {
		return err
	}

	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/dime")
	if err != nil {
		return err
	}

	_, err = db.Exec(userSchema)
	if err != nil {
		return err
	}

	_, err = db.Exec(transactionSchema)
	if err != nil {
		return err
	}

	DB = MariaDB{db: db}
	return nil
}
