package csv

import (
	"reflect"
	"testing"
)

func TestRenameColumns(t *testing.T) {
	original := []map[string]string{
		{
			"dAtE": "2020-01-01",
			"old":  "value",
		},
		{
			"dAtE": "2020-01-02",
			"old":  "value",
		},
	}

	columnMapping := map[string]string{
		"dAtE": "date",
		"old":  "new",
	}

	expected := []map[string]string{
		{
			"date": "2020-01-01",
			"new":  "value",
		},
		{
			"date": "2020-01-02",
			"new":  "value",
		},
	}

	actual := RenameColumns(original, columnMapping)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RenameColumns() = %v, want %v", actual, expected)
	}
}
