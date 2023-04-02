package csv

import (
	"bytes"
	"dime/internal/models"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"
)

func Parse(buf *bytes.Buffer) ([]map[string]any, error) {

	reader := csv.NewReader(buf)

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var data []map[string]any

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		rowData := make(map[string]any)

		for i, field := range header {
			rowData[field] = row[i]
		}

		data = append(data, rowData)
	}

	return data, nil
}

func RenameColumns(data []map[string]any, columnMapping map[string]string) []map[string]any {
	var renamedData []map[string]any
	for i, row := range data {
		renamedData = append(renamedData, make(map[string]any))
		for oldName, newName := range columnMapping {
			renamedData[i][newName] = row[oldName]
		}
	}

	return renamedData
}

func AdaptRequiredFields(data []map[string]any, dateFormat string) error {
	for _, row := range data {
		for columnName, columnType := range models.RequiredColumns {
			if _, ok := row[columnName]; !ok {
				return fmt.Errorf("column %s is missing", columnName)
			}

			switch columnType {
			case "string":
				//this should never be false
				if _, ok := row[columnName].(string); !ok {
					return fmt.Errorf("column %s is not a string", columnName)
				}
			case "date":
				date, err := time.Parse(dateFormat, row[columnName].(string))
				if err != nil {
					return fmt.Errorf("could not format %s", row[columnName])
				}
				row[columnName] = date
			case "float":
				// convert string to float
				number, err := strconv.ParseFloat(row[columnName].(string), 64)
				if err != nil {
					return fmt.Errorf("could not format %s", row[columnName])
				}
				row[columnName] = number
			default:
				return fmt.Errorf("unknown column type %s", columnType)
			}
		}
	}

	return nil
}
