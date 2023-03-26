package mongodb

import (
	"dime/internal/models"
	"reflect"
	"testing"
	"time"
)

func TestArchive_Create(t *testing.T) {
	testArchive := models.Archive{
		UploadDate:   time.Time{},
		OriginalName: "",
		Data:         nil,
	}

	archiveDao := NewArchive(client)
	id, err := archiveDao.Create(&testArchive)
	if err != nil {
		t.Errorf("Error creating a new archive: %v", err)
	}

	if id == "" {
		t.Errorf("ID is empty")
	}
}

func TestArchive_UpdateColumnMapping(t *testing.T) {
	testArchive := models.Archive{
		UploadDate:   time.Time{},
		OriginalName: "",
		Data:         nil,
	}

	archiveDao := NewArchive(client)
	id, err := archiveDao.Create(&testArchive)
	if err != nil {
		t.Errorf("Error creating a new archive: %v", err)
	}

	columnMapping := models.ColumnMapping{
		Date:        "testDate",
		Description: "testDescription",
		Amount:      "testAmount",
		Category:    "testCategory",
		Account:     "testAccount",
	}

	err = archiveDao.UpdateColumnMapping(id, &columnMapping)
	if err != nil {
		t.Errorf("Error updating column mapping: %v", err)
	}

	matchedArchive, err := archiveDao.FindByID(id)
	if err != nil {
		t.Errorf("Error retrieving archive: %v", err)
	}

	if matchedArchive.ColumnMapping != columnMapping {
		t.Errorf("Column mapping was not updated")
	}
}

func TestArchive_FindByID(t *testing.T) {

	testdata := []map[string]string{
		{"name": "Alice", "age": "28", "city": "New York"},
		{"name": "Bob", "age": "35", "city": "San Francisco"},
		{"name": "Charlie", "age": "42", "city": "London"},
	}

	testArchive := models.Archive{
		UploadDate:   time.Time{},
		OriginalName: "",
		Data:         testdata,
	}

	archiveDao := NewArchive(client)
	id, err := archiveDao.Create(&testArchive)
	if err != nil {
		t.Errorf("Error creating a new archive: %v", err)
	}

	matchedArchive, err := archiveDao.FindByID(id)
	if err != nil {
		t.Errorf("Error retrieving archive: %v", err)
	}

	//compared testArchive and matchedArchive using reflect deep equal
	if matchedArchive.UploadDate != testArchive.UploadDate {
		t.Errorf("UploadDate does not match. Expected: %v, Actual: %v", testArchive.UploadDate, matchedArchive.UploadDate)
	}

	if matchedArchive.OriginalName != testArchive.OriginalName {
		t.Errorf("OriginalName does not match. Expected: %v, Actual: %v", testArchive.OriginalName, matchedArchive.OriginalName)
	}

	if reflect.DeepEqual(matchedArchive.Data, testArchive.Data) == false {
		t.Errorf("Data does not match. Expected: %v, Actual: %v", testArchive.Data, matchedArchive.Data)
	}
}
