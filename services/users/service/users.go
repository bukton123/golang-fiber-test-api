package service

import "api/services/users/spec"

type UserService struct{}

func NewUserService() spec.UserService {
	return &UserService{}
}

func (u *UserService) Find() string {
	return "test"
}
