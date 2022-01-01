package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestServiceUsersGet(t *testing.T) {
	app := fiber.New()
	handler := Handler{}
	handler.register(app)

	req := httptest.NewRequest("GET", "/", nil)
	res, _ := app.Test(req, -1)

	bodyByte, err  := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200, "Expected status code equal 200")
	assert.NotEmpty(t, res.Body, "Expected data response")
	assert.Equal(t, string(bodyByte), "{\"test\":123}", "Expected body equal { \"test\": 123 }")
}