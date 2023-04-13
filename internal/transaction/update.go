package transaction

type columnType string

const (
	UndefinedType columnType = ""
	DateType                 = "date"
	NumberType               = "number"
	TextType                 = "text"
)

type updateType string

const (
	UndefinedUpdate    updateType = ""
	ColumnNameUpdate              = "change_column_name"
	RemoveColumnUpdate            = "remove_column"
	AddColumnUpdate               = "add_column"
	ColumnTypeUpdate              = "change_column_type"
)

type UpdateRequest struct {
	TransactionGroupId int        `json:"transaction_group_id"`
	UpdateType         updateType `json:"update_type"`
	ColumnName         string     `json:"column_name"`
	NewColumnName      string     `json:"new_column_name"`
	NewColumnType      columnType `json:"new_column_type"`
}

func updatePendingTransaction(request UpdateRequest) {
	//// Get the pending transaction
	//pendingTransaction, err := dbs.DB.PendingTransactionsDao().FindByTransactionGroupId(request.TransactionGroupId)
	//if err != nil {
	//	log.Println("Error finding pending transaction:", err)
	//	return
	//}
	//
	//// Update the pending transaction
	//switch request.UpdateType {
	//case ColumnNameUpdate:
	//	pendingTransaction.ColumnName = request.NewColumnName
	//case RemoveColumnUpdate:
	//	pendingTransaction.ColumnName = ""
	//case AddColumnUpdate:
	//	pendingTransaction.ColumnName = request.ColumnName
	//	pendingTransaction.ColumnType = request.NewColumnType
	//case ColumnTypeUpdate:
	//	pendingTransaction.ColumnType = request.NewColumnType
	//case UndefinedUpdate:
	//	log.Println("Undefined update type")
	//}
	//
	//// Save the pending transaction
	//err = dbs.DB.PendingTransactionDao().SavePendingTransaction(pendingTransaction)
	//if err != nil {
	//	log.Println("Error saving pending transaction:", err)
	//	return
	//}
}
