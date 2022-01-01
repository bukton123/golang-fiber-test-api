package main

import (
	"api/pkg/logging"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	ServiceName = "Users"
	Version     = "0.0.1"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("Service: %s v%s", ServiceName, Version),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			logging.Error(ServiceName, Version, err.Error())
			return ctx.Status(code).JSON(logging.HandlerErrorHttp(ServiceName, 0, http.StatusText(http.StatusInternalServerError)))
		},
	})
	handler := Handler{}

	handler.register(app)
	app.Listen(":8080")
}
