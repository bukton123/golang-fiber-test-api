package main

import (
	"api/pkg/api"
	"api/pkg/database"
	"api/pkg/env"
	"api/pkg/logging"
	"api/pkg/validator"
	"api/services/users/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

const (
	ServiceName = "Users"
	Version     = "0.0.1"
)

func main() {
	logging.New(ServiceName, Version)
	defer logging.Close()

	connect := database.NewMongoDB(env.GetEnvString("SERVICE_USER_MONGODB", "mongodb://localhost:27017"))
	defer connect.Close()

	app := fiber.New(fiber.Config{
		AppName:      api.AppName(ServiceName, Version),
		ErrorHandler: api.ErrorHandlerFiber(ServiceName, Version, env.GetEnvBool("SERVICE_USER_DEV", false)),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
	})

	handler := Handler{
		validate:    validator.New(),
		userService: service.NewUserService(connect),
	}

	handler.register(app)
	if err := app.Listen(":4000"); err != nil {
		panic(err.Error())
	}
}
