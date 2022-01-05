package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type (
	Adapter interface {
		init() error
		AddIndex() error
		Connect(database, collection string) Connector
		Close()
	}

	Connector interface {
		MongoDB() *mongo.Collection
	}

	convert struct {
		mongoDB *mongo.Collection
	}
)

func Timeout(timeout ...time.Duration) (context.Context, context.CancelFunc) {
	ts := time.Second * 10
	if len(timeout) > 0 {
		ts = timeout[0]
	}

	return context.WithTimeout(context.Background(), ts)
}
