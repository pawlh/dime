package csv

import (
	"reflect"
	"testing"
	"time"
)

func TestRenameColumns(t *testing.T) {
	original := []map[string]any{
		{
			"dAtE": time.Now(),
			"old":  "value",
		},
		{
			"dAtE": time.Now().Add(time.Hour * 24),
			"old":  "value",
		},
	}

	columnMapping := map[string]string{
		"dAtE": "date",
		"old":  "new",
	}

	expected := []map[string]any{
		{
			"date": original[0]["dAtE"],
			"new":  "value",
		},
		{
			"date": original[1]["dAtE"],
			"new":  "value",
		},
	}

	actual := RenameColumns(original, columnMapping)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RenameColumns() = %v, want %v", actual, expected)
	}
}

func TestAdaptRequiredFields(t *testing.T) {
	type args struct {
		data       []map[string]any
		dateFormat string
	}
	tests := []struct {
		name     string
		args     args
		expected []map[string]any
	}{
		{
			name: "All required fields are present",
			args: args{
				data: []map[string]any{
					{
						"Date":        "2020-01-01",
						"Description": "test",
						"Amount":      "1.00",
						"Category":    "test",
						"Account":     "test",
					},
				},
				dateFormat: "2006-01-02",
			},
			expected: []map[string]any{
				{
					"Date":        mustParseTime("2020-01-01", "2006-01-02"),
					"Description": "test",
					"Amount":      1.00,
					"Category":    "test",
					"Account":     "test",
				},
			},
		},
		{
			name: "Some fields are missing",
			args: args{
				data: []map[string]any{
					{
						"Date":     "2020-01-01",
						"Amount":   "1.00",
						"Category": "test",
					},
				},
				dateFormat: "2006-01-02",
			},
			expected: nil,
		},
		{
			name: "Date is malformed",
			args: args{
				data: []map[string]any{
					{
						"Date":        "202f0-c01-01ab",
						"Description": "test",
						"Amount":      "1.00",
						"Category":    "test",
						"Account":     "test",
					},
				},
				dateFormat: "2006-01-02",
			},
			expected: nil,
		},
		{
			name: "Amount is not a number",
			args: args{
				data: []map[string]any{
					{
						"Date":        "2020-01-01",
						"Description": "test",
						"Amount":      "abc",
						"Category":    "test",
						"Account":     "test",
					},
				},
				dateFormat: "2006-01-02",
			},
			expected: nil,
		},
		{
			name: "dateFormat is invalid",
			args: args{
				data: []map[string]any{
					{
						"Date":        "2020-01-01",
						"Description": "test",
						"Amount":      "1.00",
						"Category":    "test",
						"Account":     "test",
					},
				},
				dateFormat: "a2006-0d1-02d",
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AdaptRequiredFields(tt.args.data, tt.args.dateFormat); (err != nil) != (tt.expected == nil) {
				t.Errorf("Error expected, but none was returned")
			} else if tt.expected != nil {

				if !reflect.DeepEqual(tt.args.data, tt.expected) {
					t.Errorf("AdaptRequiredFields() = %v, want %v", tt.args.data, tt.expected)
				}
			}
		})
	}
}

func mustParseTime(s, layout string) time.Time {
	t, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}
	return t
}
