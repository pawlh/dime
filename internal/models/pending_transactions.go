package models

// PendingTransactions holds transactions that have not yet been submitted by a user
type PendingTransactions struct {
	WIPTransactions    []map[string]any `bson:"wip_transactions" json:"wip_transactions"`
	SavedTransactions  []map[string]any `bson:"transactions" json:"transactions"`
	Owner              string           `bson:"owner" json:"owner"`
	Name               string           `bson:"name" json:"name"`
	TransactionGroupId string           `bson:"_id,omitempty" json:"transaction_group_id"`
}
