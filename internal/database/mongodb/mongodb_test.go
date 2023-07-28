package mongodb

import "testing"

func beforeEach(t *testing.T) {
	db := Init(mongoUri)
	defer db.Disconnect()

	userDao, err := db.UserDAO()
	if err != nil {
		t.Errorf("Error getting user dao: %v", err)
	}

	if userDao.Clear() != nil {
		t.Errorf("Error clearing users: %v", err)
	}

	transactionDao, err := db.TransactionDAO()
	if err != nil {
		t.Errorf("Error getting transaction dao: %v", err)
	}

	if transactionDao.Clear() != nil {
		t.Errorf("Error clearing transactions: %v", err)
	}
}
