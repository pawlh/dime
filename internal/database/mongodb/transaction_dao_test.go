package mongodb

import (
	"dime/internal/models"
	"testing"
	"time"
)

func TestTransactionDAO_AddTransaction(t *testing.T) {
	//TODO: Implement
}

func TestTransactionDAO_AddTransactions(t *testing.T) {
	//TODO: Implement
}

func TestTransactionDAO_GetTransactions(t *testing.T) {
	beforeEach(t)

	db := Init(mongoUri)
	defer db.Disconnect()

	transactionDao, err := db.TransactionDAO()
	if err != nil {
		t.Errorf("Error getting testTransaction dao: %v", err)
	}

	testTransaction := models.Transaction{
		Owner: "test",
		BaseFields: models.BaseFields{
			Date:        time.Time{},
			Description: "testDescription",
			Amount:      123.456,
			Category:    "testCategory",
			Account:     "testAccount",
		},
		ExtraFields: map[string]any{
			"testField": "testValue",
		},
	}

	_, err = transactionDao.AddTransaction(testTransaction)
	if err != nil {
		t.Errorf("Error adding testTransaction: %v", err)
	}

	transactions, err := transactionDao.GetTransactions(testTransaction.Owner)
	if err != nil {
		t.Errorf("Error getting transactions: %v", err)
	}

	if transactions[0].Owner != testTransaction.Owner {
		t.Errorf("Expected owner to be %s, got %s", testTransaction.Owner, transactions[0].Owner)
	}

	if transactions[0].BaseFields.Description != testTransaction.BaseFields.Description {
		t.Errorf("Expected description to be %s, got %s", testTransaction.BaseFields.Description, transactions[0].BaseFields.Description)
	}

	if transactions[0].BaseFields.Amount != testTransaction.BaseFields.Amount {
		t.Errorf("Expected amount to be %f, got %f", testTransaction.BaseFields.Amount, transactions[0].BaseFields.Amount)
	}

	if transactions[0].BaseFields.Category != testTransaction.BaseFields.Category {
		t.Errorf("Expected category to be %s, got %s", testTransaction.BaseFields.Category, transactions[0].BaseFields.Category)
	}

	if transactions[0].BaseFields.Account != testTransaction.BaseFields.Account {
		t.Errorf("Expected account to be %s, got %s", testTransaction.BaseFields.Account, transactions[0].BaseFields.Account)
	}

	if transactions[0].ExtraFields["testField"] != testTransaction.ExtraFields["testField"] {
		t.Errorf("Expected testField to be %s, got %s", testTransaction.ExtraFields["testField"], transactions[0].ExtraFields["testField"])
	}
}
