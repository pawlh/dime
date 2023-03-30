package api

import (
	"bytes"
	"dime/internal/dbs"
	"dime/internal/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"net/http/httptest"
)

var (
	testUser = &models.User{
		Username: "testUsername",
		Password: "testPassword",
		Name:     "testName",
	}
)

var testURI = "mongodb://localhost:27018"

var client *mongo.Client

func newClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(testURI))
	if err != nil {
		return nil, err
	}
	err = client.Connect(nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func BeforeEach() {
	var err error

	if dbs.InitMongoDB("mongodb://localhost:27018") != nil {
		log.Fatal(err)
	}

	client, err = newClient()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	err = client.Database("dime").Drop(nil)
	if err != nil {
		log.Fatalf("Error dropping database: %v", err)
	}
}

func InsertTestUser() {
	if dbs.DB == nil {
		log.Fatal("DB is nil")
	}

	err := dbs.DB.UserDao().Create(testUser)
	if err != nil {
		log.Fatal(err)
	}
}

func ServeAndRequest(method string, target string, handler echo.HandlerFunc, request interface{}) (*httptest.ResponseRecorder, map[string]string) {

	reqBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	switch method {
	case http.MethodGet:
		e.GET(target, handler)
	}
	req := httptest.NewRequest(method, target, bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	var ressponseBody map[string]string
	if json.NewDecoder(res.Body).Decode(&ressponseBody) != nil {
		log.Fatal(err)
	}

	return res, ressponseBody
}
