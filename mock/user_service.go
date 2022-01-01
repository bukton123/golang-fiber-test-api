package mock

import "api/spec"

type (
	userServiceMock struct {}
)

func NewUserServiceMock () spec.UserService {
	return &userServiceMock{}
}

func (u *userServiceMock) Find () string {
	return "mock"
}