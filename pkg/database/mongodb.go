package database

import (
	"api/pkg/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type (
	mongodb struct {
		client *mongo.Client
	}
)

func NewMongoDB(connection string) Adapter {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		logging.Fatal(fmt.Sprintf("can't initialize zap logger: %v", err))
	}

	connect := &mongodb{
		client: client,
	}
	return connect
}

func (m *mongodb) init() error {
	return nil
}

func (m *mongodb) AddIndex() error {
	return nil
}

func (m *mongodb) Connect(database, collection string) Connector {
	return &convert{mongoDB: m.client.Database(database).Collection(collection)}
}

func (m *mongodb) Close() {
	if m.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := m.client.Disconnect(ctx); err != nil {
			logging.Error(err.Error())
		}
	}
}
