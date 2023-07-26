package mongodb

import (
	"context"
	"dime/internal/database/dao"
	"errors"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	uri    string

	userDao *UserDAO
}

// UserDAO returns a dao.UserDAO instance. Init() must be used to set up the MongoDB instance
func (m *MongoDB) UserDAO() (dao.UserDAO, error) {
	if m.userDao != nil {
		return m.userDao, nil
	}

	if m.client == nil {
		err := m.connect()
		if err != nil {
			return nil, err
		}
	}

	m.userDao = NewUserDAO(m.client)
	return m.userDao, nil
}

// connect connects to the database using the uri provided in Init()
func (m *MongoDB) connect() error {
	if m.uri == "" {
		return errors.New("Bad usage, call Init() first")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(m.uri))
	if err != nil {
		log.Error(err)
		return err
	}
	m.client = client

	return nil
}

func (m *MongoDB) Disconnect() {
	err := m.client.Disconnect(context.Background())
	if err != nil {
		log.Error(err)
	}

	m.client = nil
	m.userDao = nil
}

// Init initializes configures the connection to the database. The uri should be formatted as specified in the docs https://www.mongodb.com/docs/manual/reference/connection-string/
func Init(uri string) *MongoDB {
	return &MongoDB{uri: uri}
}
