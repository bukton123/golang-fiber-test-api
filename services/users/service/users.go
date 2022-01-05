package service

import (
	"api/pkg/database"
	"api/services/users/spec"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	mongo *mongo.Collection
}

func NewUserService(connect database.Adapter) spec.UserService {
	return &UserService{
		mongo: connect.Connect("", "").MongoDB(),
	}
}

func (u *UserService) Find() (interface{}, error) {
	ctx, cancel := database.Timeout()
	defer cancel()

	var result map[string]interface{}
	err := u.mongo.FindOne(ctx, bson.M{}).Decode(&result)

	return result, err
}
