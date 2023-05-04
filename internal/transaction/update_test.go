package transaction

import (
	"dime/internal/dbs"
	"dime/internal/models"
	"log"
	"testing"
)

var testURI = "mongodb://localhost:27018"

func BeforeEach() {
	err := dbs.InitMongoDB(testURI)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	//clear
	err = dbs.DB.PendingTransactionsDao().Clear()
	if err != nil {
		log.Fatalf("Error clearing database: %v", err)
	}
}

func Test_updatePendingTransaction1(t *testing.T) {

	testPendingTransactions := models.PendingTransactions{
		WIPTransactions: []map[string]any{
			{
				"amount":   100,
				"date":     "2021-01-01",
				"category": "Food",
			},
			{
				"amount":   12.34,
				"date":     "2021-09-13",
				"category": "Food",
			},
		},
		SavedTransactions: nil,
		Owner:             "testUser",
		Name:              "testName",
	}

	tests := []struct {
		name    string
		request UpdateRequest
		wantErr bool
	}{
		{
			"Change Column Name",
			UpdateRequest{
				TransactionGroupId: testPendingTransactions.TransactionGroupId,
				UpdateType:         ColumnNameUpdate,
				ColumnName:         "category",
				NewColumnName:      "dogegory",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BeforeEach()
			id, err := dbs.DB.PendingTransactionsDao().Create(&testPendingTransactions)
			if err != nil {
				t.Errorf("Error creating test pending transactions: %v", err)
			}
			tt.request.TransactionGroupId = id

			if err := updatePendingTransaction(tt.request); (err != nil) != tt.wantErr {
				t.Errorf("updatePendingTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
