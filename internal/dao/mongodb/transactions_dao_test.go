package mongodb

import (
	"dime/internal/models"
	"reflect"
	"testing"
)

func TestTransactions_InsertNew(t *testing.T) {
	BeforeEach()

	testTransactions := []map[string]any{
		{
			"testColumn1": "transaction1",
			"testColumn2": int32(100),
		},
		{
			"testColumn1": "transaction2",
			"testColumn2": int32(200),
		},
		{
			"testColumn1": "transaction3",
			"testColumn2": int32(300),
		},
	}
	testTransaction := models.Transactions{
		Transactions: testTransactions,
		Owner:        "testUser",
		Columns:      []string{"testColumn1", "testColumn2"},
	}

	transactionsDao := NewTransactions(client)

	err := transactionsDao.Insert(&testTransaction)
	if err != nil {
		t.Errorf("Error inserting new transactions: %v", err)
	}

	if match, err := transactionsDao.FindByOwner("testUser"); err != nil {
		t.Errorf("Error finding transactions: %v", err)
	} else {
		if !reflect.DeepEqual(*match, testTransaction) {
			t.Errorf("Transactions do not match. Expected %v, got %v", testTransaction, *match)
		}
	}
}

func TestTransactions_FindByOwner(t *testing.T) {
	BeforeEach()

	transactionsDao := NewTransactions(client)

	if match, err := transactionsDao.FindByOwner("fakeUser"); match != nil {
		t.Errorf("Expected nil, got %v", match)
	} else if err != nil {
		t.Errorf("Error finding transactions: %v", err)
	}

	testTransaction := models.Transactions{
		Transactions: nil,
		Owner:        "testUser",
		Columns:      []string{"testColumn1", "testColumn2"},
	}

	err := transactionsDao.Insert(&testTransaction)
	if err != nil {
		t.Errorf("Error inserting new transactions: %v", err)
	}

	if match, err := transactionsDao.FindByOwner("testUser"); err != nil {
		t.Errorf("Error finding transactions: %v", err)
	} else {
		if !reflect.DeepEqual(*match, testTransaction) {
			t.Errorf("Transactions do not match. Expected %v, got %v", testTransaction, *match)
		}
	}
}
