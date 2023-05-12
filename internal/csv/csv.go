package csv

import (
	"bytes"
	"encoding/csv"
)

// Parse parses a buffer (of a CSV file) into a slice of maps
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
