package api

import (
	"bytes"
	"dime/internal/csv"
	"dime/internal/dbs"
	"dime/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"time"
	_ "time"
)

// Upload parses a CSV file and saves it to the pending transactions table
// The upload is expected to be a form with the field:
// - file: the CSV file
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

	//"MM/DD/YYYY hh:MM:SS AM/PM"
	transactionName := time.Now().Format("01-02-2006 15:04:05 PM")

	transactions := models.PendingTransactions{
		WIPTransactions:   parsedCSV,
		SavedTransactions: nil,
		Owner:             c.Get("username").(string),
		Name:              transactionName,
	}

	err = dbs.DB.PendingTransactionsDao().Create(&transactions)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error saving record", err)
	}

	//BroadcastTransactions(c.Get("username").(string))

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
