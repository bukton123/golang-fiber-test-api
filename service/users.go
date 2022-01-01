package service

import "api/spec"

type UserService struct{}

func NewUserService() spec.UserService {
	return &UserService{}
}

func (u *UserService) Find() string {
	return "test"
}
