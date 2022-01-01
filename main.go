package main

import (
	"api/pkg/api"
	"api/pkg/env"
	"api/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

const (
	ServiceName = "Users"
	Version     = "0.0.1"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      api.AppName(ServiceName, Version),
		ErrorHandler: api.ErrorHandlerFiber(ServiceName, Version, env.GetEnvBool("SERVICE_USER_DEV", false)),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
	})
	handler := Handler{
		userService: service.NewUserService(),
	}

	handler.register(app)
	if err := app.Listen(":8080"); err != nil {
		panic(err.Error())
	}
}
