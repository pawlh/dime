package dbs

import "dime/internal/dao"

type Database interface {
	TransactionDao() dao.TransactionDao
	UserDao() dao.UserDAO
	ArchiveDao() dao.ArchiveDao
}

var DB Database
