package main

import (
	"api/pkg/api"
	"api/pkg/env"
	"api/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"time"
)

const (
	ServiceName = "Scheduler"
	Version     = "0.0.1"
)

func main() {
	logging.New(ServiceName, Version)
	defer logging.Close()

	app := fiber.New(fiber.Config{
		AppName:      api.AppName(ServiceName, Version),
		ErrorHandler: api.ErrorHandlerFiber(ServiceName, Version, env.GetEnvBool("SERVICE_USER_DEV", false)),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
	})

	h := Handler{}
	h.register(app)

	if err := app.Listen(":4000"); err != nil {
		panic(err.Error())
	}
}
