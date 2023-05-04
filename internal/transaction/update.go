package transaction

import (
	"dime/internal/dbs"
	"errors"
	"log"
	"time"
)

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
	ColumnTypeUpdate              = "change_column_type"
)

type UpdateRequest struct {
	TransactionGroupId string     `json:"transaction_group_id"`
	UpdateType         updateType `json:"update_type"`
	ColumnName         string     `json:"column_name"`
	NewColumnName      string     `json:"new_column_name"`
	NewColumnType      columnType `json:"new_column_type"`
	DateFormat         string     `json:"date_format"`
}

func updatePendingTransaction(request UpdateRequest) error {
	// Get the pending transaction
	pendingTransaction, err := dbs.DB.PendingTransactionsDao().FindById(request.TransactionGroupId)
	if err != nil {
		log.Println("Error finding pending transaction:", err)
		return err
	}

	// Update the pending transaction
	switch request.UpdateType {
	case ColumnNameUpdate:
		for i := range pendingTransaction.WIPTransactions {
			pendingTransaction.WIPTransactions[i][request.NewColumnName] = pendingTransaction.WIPTransactions[i][request.ColumnName]
			delete(pendingTransaction.WIPTransactions[i], request.ColumnName)
		}
	case RemoveColumnUpdate:
		for i := range pendingTransaction.WIPTransactions {
			delete(pendingTransaction.WIPTransactions[i], request.ColumnName)
		}
	case ColumnTypeUpdate:
		// this is a bit more complicated and prone to error
		switch request.NewColumnType {
		case DateType:
			for i := range pendingTransaction.WIPTransactions {
				date, err := time.Parse(request.DateFormat, pendingTransaction.WIPTransactions[i][request.ColumnName].(string))
				if err != nil {
					log.Println("Error parsing date:", err)
					return errors.New("error parsing date: " + err.Error())
				}

				pendingTransaction.WIPTransactions[i][request.ColumnName] = date
			}
		case NumberType:
			for i := range pendingTransaction.WIPTransactions {
				// convert current value to number
				pendingTransaction.WIPTransactions[i][request.ColumnName] = pendingTransaction.WIPTransactions[i][request.ColumnName].(float64)
			}
		}
	case UndefinedUpdate:
		return errors.New("undefined update type")
	}

	// Save the pending transaction
	err = dbs.DB.PendingTransactionsDao().Save(pendingTransaction)
	if err != nil {
		log.Println("Error saving pending transaction:", err)
		return err
	}

	return nil
}
