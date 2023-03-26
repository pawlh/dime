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
)

//var dateLayout = "2006-01-02 15:04:05"

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

	uploadDate := time.Now()

	archive := models.Archive{
		UploadDate:   uploadDate,
		OriginalName: formFile.Filename,
		Data:         parsedCSV,
	}

	_, err = dbs.DB.ArchiveDao().Create(&archive)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error saving record", err)
	}

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
