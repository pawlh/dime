package dbs

import (
	"database/sql"
	"dime/internal/dao"
	"dime/internal/dao/mariadb"
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

var uploadDataSchema = `
CREATE TABLE IF NOT EXISTS upload_data (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    upload_date DATETIME NOT NULL,
    file_name TEXT NOT NULL,
    original_name TEXT NOT NULL,
    owner TEXT NOT NULL
);
`

type MariaDB struct {
	db *sql.DB
}

func (m MariaDB) TransactionDao() dao.TransactionDao {
	return mariadb.NewMariaDbTransaction(m.db)
}

func (m MariaDB) UserDao() dao.UserDAO {
	return mariadb.NewMariaDbUser(m.db)
}

func (m MariaDB) UploadDataDao() dao.UploadDataDAO {
	return mariadb.NewMariaDbUploadData(m.db)
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

	err = loadSchemas(db)
	if err != nil {
		return err
	}

	DB = MariaDB{db: db}
	return nil
}

func loadSchemas(db *sql.DB) error {
	_, err := db.Exec(userSchema)
	if err != nil {
		return err
	}

	_, err = db.Exec(transactionSchema)
	if err != nil {
		return err
	}

	_, err = db.Exec(uploadDataSchema)
	if err != nil {
		return err
	}

	return nil
}
