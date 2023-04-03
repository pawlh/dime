package dbs

import (
	"context"
	"dime/internal/dao"
	"dime/internal/dao/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	client *mongo.Client
}

func (m MongoDB) UserDao() dao.UserDAO {
	return mongodb.NewUser(m.client)
}

func (m MongoDB) TransactionDao() dao.TransactionsDao {
	return mongodb.NewTransactions(m.client)
}

func InitMongoDB(url string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	DB = MongoDB{
		client: client,
	}

	return nil
}
