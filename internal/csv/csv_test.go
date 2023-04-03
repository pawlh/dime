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

	actual, err := RenameColumns(original, columnMapping)
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}

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
						"id":          "1",
						"date":        "2020-01-01",
						"description": "test",
						"amount":      "1.00",
						"category":    "test",
						"account":     "test",
					},
				},
				dateFormat: "2006-01-02",
			},
			expected: []map[string]any{
				{
					"id":          "1",
					"date":        mustParseTime("2020-01-01", "2006-01-02"),
					"description": "test",
					"amount":      1.00,
					"category":    "test",
					"account":     "test",
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
						"id":          "1",
						"date":        "202f0-c01-01ab",
						"description": "test",
						"amount":      "1.00",
						"category":    "test",
						"account":     "test",
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
						"id":          "1",
						"date":        "2020-01-01",
						"description": "test",
						"amount":      "abc",
						"category":    "test",
						"account":     "test",
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
						"id":          "1",
						"date":        "2020-01-01",
						"description": "test",
						"amount":      "1.00",
						"category":    "test",
						"account":     "test",
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
				if tt.expected == nil {
					t.Errorf("error expected, but none occurred")
				} else {
					t.Errorf("error with AdaptRequiredFields() = %v", err)
				}
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
