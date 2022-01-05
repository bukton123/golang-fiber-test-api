package database

import "go.mongodb.org/mongo-driver/mongo"

func (c *convert) MongoDB() *mongo.Collection {
	return c.mongoDB
}
