package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestPendingTransactions_Create(t *testing.T) {
	tests := []struct {
		name    string
		args    *models.PendingTransactions
		wantErr bool
	}{
		{
			name: "Create a new pending transaction",
			args: &models.PendingTransactions{
				WIPTransactions: []map[string]any{
					{
						"amount":   100,
						"date":     "2021-01-01",
						"category": "Food",
					},
				},
				Owner: "testUser",
				Name:  "test transaction group",
			},
			wantErr: false,
		},
		{
			name: "Create a new pending transaction with no owner",
			args: &models.PendingTransactions{
				WIPTransactions: []map[string]any{
					{
						"amount":   100,
						"date":     "2021-01-01",
						"category": "Food",
					},
				},
				Name: "test transaction group",
			},
			wantErr: true,
		},
		{
			name: "Create a new pending transaction with no name",
			args: &models.PendingTransactions{
				WIPTransactions: []map[string]any{
					{
						"amount":   100,
						"date":     "2021-01-01",
						"category": "Food",
					},
				},
				Owner: "testUser",
			},
			wantErr: true,
		},
		{
			name: "Create a new pending transaction with nil wip transactions",
			args: &models.PendingTransactions{
				Owner: "testUser",
				Name:  "test transaction group",
			},
			wantErr: true,
		},
		{
			name: "Create a new pending transaction with empty wip transactions",
			args: &models.PendingTransactions{
				WIPTransactions: []map[string]any{},
				Owner:           "testUser",
				Name:            "test transaction group",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BeforeEach()

			dao := NewPendingTransactions(client)

			if _, err := dao.Create(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPendingTransactions_FindByOwner(t *testing.T) {
	BeforeEach()
	dao := NewPendingTransactions(client)

	testTransactions := []map[string]any{
		{
			"amount":   100,
			"date":     "2021-01-01",
			"category": "Food",
		},
	}

	_, err := dao.Create(&models.PendingTransactions{
		WIPTransactions: testTransactions,
		Name:            "test transaction group",
		Owner:           "testUserA",
	})
	if err != nil {
		t.Errorf("Error creating test pending transactions: %v", err)
	}
	_, err = dao.Create(&models.PendingTransactions{
		WIPTransactions: testTransactions,
		Name:            "test transaction group",
		Owner:           "testUserB",
	})
	if err != nil {
		t.Errorf("Error creating test pending transactions: %v", err)
	}

	tests := []struct {
		name    string
		owner   string
		wantErr bool
	}{
		{
			name:    "Find pending transactions by owner",
			owner:   "testUserA",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dao := NewPendingTransactions(client)

			if transactions, err := dao.FindByOwner(tt.owner); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				if transactions.Owner != tt.owner {
					t.Errorf("Owner does not match: %v", transactions.Owner)
				}
			}
		})
	}
}

func TestPendingTransactions_FindById(t *testing.T) {
	BeforeEach()
	collection := client.Database("dime").Collection("pending_transactions")

	testPendingTransaction := &models.PendingTransactions{
		WIPTransactions: []map[string]any{
			{
				"amount": int32(100),
				"date":   "2021-01-01",
			},
		},
		Name:  "test transaction group",
		Owner: "testUser",
	}

	objectId, err := collection.InsertOne(nil, testPendingTransaction)
	if err != nil {
		t.Errorf("Error creating test pending transactions: %v", err)
	}

	id := objectId.InsertedID.(primitive.ObjectID).Hex()

	dao := NewPendingTransactions(client)

	if transactions, err := dao.FindById(id); err != nil {
		t.Errorf("Error finding pending transaction: %v", err)
	} else {
		if transactions.Name != testPendingTransaction.Name {
			t.Errorf("Name does not match: %v", transactions.Name)
		}
		if transactions.Owner != testPendingTransaction.Owner {
			t.Errorf("Owner does not match: %v", transactions.Owner)
		}

		if transactions.WIPTransactions[0]["amount"] != testPendingTransaction.WIPTransactions[0]["amount"] {
			t.Errorf("Amount does not match: %v", transactions.WIPTransactions[0]["amount"])
		}
		if transactions.WIPTransactions[0]["date"] != testPendingTransaction.WIPTransactions[0]["date"] {
			t.Errorf("Date does not match: %v", transactions.WIPTransactions[0]["date"])
		}
	}

}
