package csv

import (
	"bytes"
	"encoding/csv"
)

func Parse(buf *bytes.Buffer) ([]map[string]string, error) {

	reader := csv.NewReader(buf)

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var data []map[string]string

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		rowData := make(map[string]string)

		for i, field := range header {
			rowData[field] = row[i]
		}

		data = append(data, rowData)
	}

	return data, nil
}
