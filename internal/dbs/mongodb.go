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

func (m MongoDB) TransactionDao() dao.TransactionDao {
	return nil
}

func (m MongoDB) ArchiveDao() dao.ArchiveDao {
	return mongodb.NewArchive(m.client)
}

func InitMongoDB() error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
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
