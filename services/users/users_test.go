package main

import (
	"api/services/users/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestServiceUsersGet(t *testing.T) {
	app := fiber.New()
	handler := Handler{
		userService: mock.NewUserServiceMock(),
	}
	handler.register(app)

	req := httptest.NewRequest("GET", "/users", nil)
	res, _ := app.Test(req, -1)

	bodyByte, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200, "Expected status code equal 200")
	assert.NotEmpty(t, res.Body, "Expected data response")
	assert.Equal(t, string(bodyByte), "{\"data\":\"mock\"}", "Expected body equal { \"data\": \"mock\" }")
}
