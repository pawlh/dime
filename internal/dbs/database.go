package dbs

import "dime/internal/dao"

type Database interface {
	TransactionDao() dao.TransactionsDao
	UserDao() dao.UserDAO
	ArchiveDao() dao.ArchiveDao
}

var DB Database
