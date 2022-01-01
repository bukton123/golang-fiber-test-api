package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Service: Users",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			err = ctx.Status(code).JSON(fiber.Map{"error": fiber.Map{
				"type": "UserExaction",
				"code": 0,
				"message": err.Error(),
			}})
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			return nil
		},
	})
	handler := Handler{}

	handler.register(app)
	app.Listen(":8080")
}
