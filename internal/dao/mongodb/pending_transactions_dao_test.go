package mongodb

import (
	"dime/internal/models"
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

			if err := dao.Create(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
