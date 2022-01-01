package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	handler := Handler{}

	handler.register(app)
	app.Listen(":8080")
}