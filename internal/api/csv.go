package api

import (
	"bytes"
	"dime/internal/csv"
	"dime/internal/dbs"
	"dime/internal/models"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	_ "time"
)

//var dateLayout = "2006-01-02 15:04:05"

type meta struct {
	ColumnMapping map[string]string `json:"column_mapping"`
	DateFormat    string            `json:"date_format"`
}

// Upload parses a CSV file and saves it to the database
// The upload is expected to be a form with two fields:
// - file: the CSV file
// - meta: information about the CSV file, such as the column mapping
func Upload(c echo.Context) error {

	formFile, err := c.FormFile("file")
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "missing file", err)
	}

	file, err := formFile.Open()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error opening file", err)
	}

	buf, err := fileToBuffer(file)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error reading file", err)
	}

	parsedCSV, err := csv.Parse(buf)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error parsing file", err)
	}

	var meta meta
	//marshal c.FormValue("meta") into meta struct
	if err = json.Unmarshal([]byte(c.FormValue("meta")), &meta); err != nil {
		return mustSendError(c, http.StatusBadRequest, "error parsing meta", err)
	}

	renamedCSV, err := csv.RenameColumns(parsedCSV, meta.ColumnMapping)
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "error renaming columns", err)
	}
	err = csv.AdaptRequiredFields(renamedCSV, meta.DateFormat)
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "error adapting csv", err)
	}

	transactions := models.Transactions{
		Transactions: renamedCSV,
		Owner:        c.Get("username").(string),
		Columns:      csv.GetColumns(renamedCSV),
	}

	err = dbs.DB.TransactionDao().Insert(&transactions)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error saving record", err)
	}

	BroadcastTransactions(c.Get("username").(string))

	return nil
}

func fileToBuffer(file multipart.File) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return nil, fmt.Errorf("error copying file: %w", err)
	}

	return buf, nil
}
