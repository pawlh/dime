package dbs

import "dime/internal/dao"

type Database interface {
	TransactionDao() dao.TransactionsDao
	UserDao() dao.UserDAO
	PendingTransactionsDao() dao.PendingTransactionsDao
}

var DB Database
