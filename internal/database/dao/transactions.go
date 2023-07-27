package dao

type TransactionDAO interface {
	// AddTransaction Add a transaction to the database
	AddTransaction(owner string, transaction []map[string]any) (string, error)
	// AddTransactions Add a list of transactions to the database
	AddTransactions(owner string, transactions []map[string]any) (string, error)
	// GetTransactions Fetch all transactions for a user
	GetTransactions(owner string) ([]map[string]any, error)
}
